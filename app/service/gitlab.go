// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package service

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/repository"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
)

type GitlabService struct {
}

func NewGitlab() *GitlabService {
	return &GitlabService{}
}

// StoreProjects 存储 Gitlab 项目到本地数据库
func (s *GitlabService) StoreProjects() {
	projects, err := gitlab.ListProjects(&git.ListProjectsOptions{
		IncludeHidden:        git.Ptr(false),
		IncludePendingDelete: git.Ptr(false),
	})
	if err != nil {
		log.Error().Err(err).Msg("获取 Gitlab 项目列表失败")
		return
	}

	var needDelete []int
	for _, p := range projects {
		// 如果项目 在跳过列表中、已删除、已归档、命名空间不是 group 或在跳过列表中，则跳过
		if slices.Contains(config.Get().Gitlab.SkipRepositories, p.PathWithNamespace) ||
			p.MarkedForDeletionOn != nil ||
			p.Archived ||
			(p.Namespace != nil && (p.Namespace.Kind != "group" || slices.Contains(config.Get().Gitlab.SkipNamespaces, p.Namespace.Path))) {
			needDelete = append(needDelete, p.ID)
		}

		// 创建或更新项目
		err = ent.Database.Repository.Create().
			SetPath(p.PathWithNamespace).
			SetID(p.ID).
			OnConflictColumns(repository.FieldID).
			UpdateNewValues().
			Exec(context.Background())
		if err != nil {
			log.Error().Err(err).Msgf("保存 Gitlab 项目 %s 失败", p.PathWithNamespace)
			continue
		}
	}

	if len(needDelete) > 0 {
		_, _ = ent.Database.Repository.Delete().
			Where(repository.IDIn(needDelete...)).
			Exec(context.Background())
	}
}

// GetProjects 获取所有已存储的 Gitlab 项目路径
func (s *GitlabService) GetProjects() (items []string) {
	projects, _ := ent.Database.Repository.Query().All(context.Background())
	items = make([]string, len(projects))
	for i, p := range projects {
		items[i] = p.Path
	}

	slices.SortFunc(items, func(a, b string) int {
		return strings.Compare(a, b)
	})

	return
}

// Webhook 处理 Gitlab Webhook 请求
func (s *GitlabService) Webhook(secret string, eventType git.EventType, body []byte) {
	// 校验key
	if !gitlab.VerifyWebhookSecret(secret) {
		log.Error().Msg("Webhook Secret 校验失败")
		return
	}

	data, err := git.ParseHook(eventType, body)
	if err != nil {
		log.Error().Err(err).Msgf("解析 GitLab Webhook 请求体失败, event: %s", eventType)
		return
	}

	switch v := data.(type) {
	case *git.MergeEvent:
		s.doMergeEvent(v)
	}
}

// 处理 Gitlab Merge Event
func (s *GitlabService) doMergeEvent(data *git.MergeEvent) {
	state := data.ObjectAttributes.State
	source := data.ObjectAttributes.SourceBranch
	target := data.ObjectAttributes.TargetBranch

	log.Info().Msgf("收到 Gitlab Merge Event (%s): %s => %s", state, source, target)

	switch state {
	case gitlab.MergeStateMerged:
		s.GitlabMerged(source, target)
	}
}

// GetBranchWorkitemID 从 branch 获取工作项 id
func (s *GitlabService) GetBranchWorkitemID(branch string) (id string, ok bool) {
	ok = strings.HasPrefix(branch, config.GitlabBranchPrefix)
	if ok {
		id = branch[len(config.GitlabBranchPrefix):]
	}

	return
}

// GitlabMerged gitlab 分支已被合并
func (s *GitlabService) GitlabMerged(source, target string) {
	id, ok := s.GetBranchWorkitemID(source)
	if !ok {
		log.Info().Msgf("合并请求 %s 非 dev/ 分支", source)
		return
	}

	// 判断合并目标是否在配置项中
	if !slices.Contains(config.Get().Gitlab.MergeTargetBranchs, target) {
		log.Info().Msgf("目标分支 %s 不在配置列表中 (%s)，跳过处理", target, strings.Join(config.Get().Gitlab.MergeTargetBranchs, ","))
		return
	}

	// 获取工作项
	workitem, err := yunxiao.GetWorkitem(id)
	if err != nil {
		log.Error().Err(err).Msg("获取工作项失败")
		return
	}

	// 获取工作项状态
	status := workitem.Status
	if status == nil {
		log.Warn().Msgf("工作项 %s 状态为空，跳过处理", id)
		return
	}

	// 如果工作项状态不是 处理中 状态, 跳过
	inProgress := yc.GetWorkflowStatus(workitem.CategoryID, yc.WorkflowStatusInProgress)
	if status.ID != inProgress.Id {
		log.Info().Msgf("工作项 %s 状态为 %s，跳过处理", id, status.Name)
		return
	}

	underReview := yc.GetWorkflowStatus(workitem.CategoryID, yc.WorkflowStatusUnderReview)
	if underReview.Id == "" {
		log.Warn().Msg("获取待审查状态配置失败，跳过处理")
		return
	}

	log.Info().Msgf("获取到工作项 ID: %s, 状态转换: %s (%s) => 待审查 (%s)", id, status.Name, status.ID, underReview.Id)

	// 修改工作项状态为 待审查
	err = yunxiao.UpdateWorkitem(id, map[string]string{
		yc.FieldStatus: underReview.Id,
	})
	if err != nil {
		log.Error().Err(err).Msg("更新工作项状态失败")
		return
	}

	// 添加评论
	err = yunxiao.CreateWorkitemComment(id, fmt.Sprintf("分支 %s 已合并到 %s，工作项状态自动更新为 待审查", source, target))
	if err != nil {
		log.Error().Err(err).Msg("创建工作项评论失败")
		return
	}
}
