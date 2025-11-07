// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package service

import (
	"fmt"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

type CooperationService struct {
}

func NewCooperation() *CooperationService {
	return &CooperationService{}
}

// GetBranchWorkitemID 从 branch 获取工作项 id
func (s *CooperationService) GetBranchWorkitemID(branch string) (id string, ok bool) {
	ok = strings.HasPrefix(branch, config.GitlabBranchPrefix)
	if ok {
		id = branch[len(config.GitlabBranchPrefix):]
	}

	return
}

// GitlabMerged gitlab 分支已被合并
func (s *CooperationService) GitlabMerged(source, target string) {
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
		yc.WorkitemStatusKey: underReview.Id,
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

// YunxiaoStatusChanged 云效工作项状态变更处理
func (s *CooperationService) YunxiaoStatusChanged(data *entity.WebhookStatusEvent) {
	workitem := data.Workitem
	if workitem == nil {
		log.Error().Msg("工作项为空，无法处理状态变更")
		return
	}

	cfg := yc.GetWorkitem(workitem.CategoryID)

	var err error

	branch := config.GitlabBranchPrefix + workitem.SerialNumber

	// 添加评论
	created := "未自动创建 gitlab 分支"
	defer func() {
		_, err = yunxiao.NewCreateWorkitemCommentRequest(workitem.ID, fmt.Sprintf("状态: %s → %s\n\n%s, 创建命令: \n```\ngit checkout -b %s\n```", data.From.Text(), data.To.Text(), created, branch)).Do()
		if err != nil {
			log.Error().Err(err).Msgf("创建工作项 %s 评论失败", workitem.SerialNumber)
			return
		}
	}()

	// 获取 代码仓库 字段
	var pid string
	for _, field := range workitem.CustomFieldValues {
		if field.FieldID == cfg.Fields[yc.FieldRepository].Id {
			pid = field.Values[0].Identifier
			break
		}
	}

	if pid == "" {
		log.Info().Msg("未配置代码仓库，跳过创建分支")
		return
	}

	// 创建gitlab分支
	var b *git.Branch
	b, err = gitlab.CreateBranch(pid, branch, "")
	if err != nil {
		log.Error().Err(err).Msgf("创建 Gitlab 分支 %s 失败", branch)
		return
	}

	created = fmt.Sprintf("已创建分支 `%s` (commit: %s)", b.Name, b.Commit.ID)
}
