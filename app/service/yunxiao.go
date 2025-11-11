// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/app/model"
	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/user"
	"github.com/liasica/orbit/integration/feishu"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

type YunxiaoService struct {
}

func NewYunxiao() *YunxiaoService {
	return &YunxiaoService{}
}

// GetUserFromWorkitemCustomFieldValues 从工作项获取用户列表
func (s *YunxiaoService) GetUserFromWorkitemCustomFieldValues(workitem *entity.Workitem, field yc.Field) (users []*ent.User) {
	var ids []string

	cfg := yc.GetWorkitem(workitem.CategoryID)
	for _, v := range workitem.CustomFieldValues {
		if v.FieldID == cfg.Fields[field].Id {
			for _, vv := range v.Values {
				ids = append(ids, vv.Identifier)
			}
		}
	}

	return s.GetUsers(ids...)
}

func (s *YunxiaoService) GetUsers(ids ...string) (users []*ent.User) {
	users, _ = ent.Database.User.Query().Where(user.YunxiaoUserIDIn(ids...)).All(context.Background())
	return
}

// Webhook 处理云效 webhook
func (s *YunxiaoService) Webhook(headers http.Header, body []byte) {
	// 获取 action
	action := headers.Get(yunxiao.HookHeaderAction)

	// 获取secret
	secret := headers.Get(yunxiao.HookHeaderSecret)

	// 校验 secret
	if !yunxiao.VerifyWebhookSecret(secret) {
		log.Warn().Msg("Secret 校验失败")
		return
	}

	// 获取数据
	node := new(yunxiao.WebhookWorkitemIdentifier)
	err := sonic.Unmarshal(body, node)
	if err != nil {
		log.Error().Err(err).Msg("获取工作项 identifier 失败")
		return
	}

	// 获取工作项信息
	var workitem *entity.Workitem
	workitem, err = yunxiao.GetWorkitem(node.Identifier)
	if err != nil {
		log.Error().Err(err).Msgf("获取工作项信息失败: %s", node.Identifier)
		return
	}

	if workitem == nil {
		log.Error().Msg("工作项为空")
		return
	}

	switch action {
	case yunxiao.HookActionWorkitemCreated:
		s.hookActionWorkitemCreated(headers, workitem)
	case yunxiao.HookActionWorkitemStatusChanged:
		s.hookActionWorkitemStatusChanged(headers, workitem)
	case yunxiao.HookActionWorkitemUnderReview:
		s.hookActionWorkitemUnderReview(headers, workitem)
	default:
		log.Warn().Msgf("未知的 action 类型: %s", action)
	}
}

// 处理工作项创建
func (s *YunxiaoService) hookActionWorkitemCreated(_ http.Header, workitem *entity.Workitem) {
	var userIds []string
	// 负责人
	if workitem.AssignedTo == nil {
		log.Warn().Msgf("工作项 %s 未分配负责人，跳过处理", workitem.SerialNumber)
		return
	}
	userIds = append(userIds, workitem.AssignedTo.ID)

	// 参与者
	for _, u := range workitem.Participants {
		userIds = append(userIds, u.ID)
	}

	users := s.GetUsers(userIds...)

	theme := "blue"
	if workitem.CategoryID == yc.WorkitemCategoryBug {
		theme = "red"
	}

	desc := workitem.Description
	if desc == "" {
		desc = "无"
	}

	url := fmt.Sprintf("https://devops.aliyun.com/projex/bug/%s", workitem.SerialNumber)

	icon := config.Get().Feishu.Icons.Task
	if workitem.CategoryID == yc.WorkitemCategoryBug {
		icon = config.Get().Feishu.Icons.Bug
	}

	// 给指定人员发送卡片消息
	for _, u := range users {
		NewFeishu().SendJobMessage(&model.FeishuSendJobMessageRequest{
			ReceiveIdType: feishu.ReceiveIdTypeUserID,
			ReceiveId:     u.LarkUserID,
			ID:            workitem.SerialNumber,
			Title:         workitem.Subject,
			Category:      workitem.CategoryID.Text(),
			Theme:         theme,
			Description:   desc,
			Url:           url,
			Icon:          icon,
			Status:        workitem.Status.Name,
		})
	}
}

// 处理工作项变更
func (s *YunxiaoService) hookActionWorkitemStatusChanged(headers http.Header, workitem *entity.Workitem) {
	// 获取 status from / to
	from := headers.Get(yunxiao.HookHeaderStatusFrom)
	to := headers.Get(yunxiao.HookHeaderStatusTo)
	if from == "" || to == "" {
		log.Warn().Msg("缺少状态变更信息")
		return
	}

	// 打开 / 重新打开 → 处理中
	f := yc.WorkflowStatus(from)
	t := yc.WorkflowStatus(to)
	if (f == yc.WorkflowStatusOpen || f == yc.WorkflowStatusReopen) && t == yc.WorkflowStatusInProgress {
		s.statusOpenToInprogress(f, t, workitem)
	}
}

// 状态由 打开 / 重新打开 → 处理中
func (s *YunxiaoService) statusOpenToInprogress(from, to yc.WorkflowStatus, workitem *entity.Workitem) {
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
		_, err = yunxiao.NewCreateWorkitemCommentRequest(workitem.ID, fmt.Sprintf("状态: %s → %s\n\n%s, 创建命令: \n```\ngit checkout -b %s\n```", from.Text(), to.Text(), created, branch)).Do()
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

// 处理待审查工作项
func (s *YunxiaoService) hookActionWorkitemUnderReview(_ http.Header, workitem *entity.Workitem) {
	theme := "blue"
	if workitem.CategoryID == yc.WorkitemCategoryBug {
		theme = "red"
	}

	desc := workitem.Description
	if desc == "" {
		desc = "无"
	}

	users := s.GetUserFromWorkitemCustomFieldValues(workitem, yc.FieldReviewUser)
	if len(users) == 0 {
		log.Warn().Msgf("工作项 %s 未获取到审查人，跳过发送飞书消息", workitem.SerialNumber)
		return
	}

	userIds := make([]string, len(users))
	for i, u := range users {
		userIds[i] = u.LarkUserID
	}

	// 发送待审查消息
	NewFeishu().SendUnderReviewMessage(&model.FeishuSendUnderReviewMessageRequest{
		ID:          workitem.SerialNumber,
		Title:       workitem.Subject,
		Category:    workitem.CategoryID.Text(),
		Theme:       theme,
		Description: desc,
		ReviewUsers: strings.Join(userIds, ","),
		Url:         fmt.Sprintf("https://devops.aliyun.com/projex/bug/%s", workitem.SerialNumber),
	})
}

// UpdateStatusByID 更新工作项状态
func (s *YunxiaoService) UpdateStatusByID(id string, status yc.WorkflowStatus) {
	// 标记工作项为已解决
	workitem, err := yunxiao.GetWorkitem(id)
	if err != nil {
		log.Error().Err(err).Msg("无法标记工作项状态为已解决：获取工作项失败")
		return
	}
	err = yunxiao.UpdateWorkitem(workitem.ID, map[string]string{
		yc.FieldStatus: yc.GetWorkflowStatus(workitem.CategoryID, status).Id,
	})
	if err != nil {
		log.Error().Err(err).Msg("标记工作项状态为已解决失败")
		return
	}
}
