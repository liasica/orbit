// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package service

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/app/model"
	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/message"
	"github.com/liasica/orbit/integration/feishu"
)

type FeishuService struct {
}

func NewFeishu() *FeishuService {
	return &FeishuService{}
}

// SaveMessage 保存飞书消息记录
func (s *FeishuService) SaveMessage(resp *v1.CreateMessageResp, workitemId *string, req any, typ message.Type) {
	if resp == nil || resp.Err != nil || resp.Data == nil {
		log.Error().Msg("飞书消息响应数据为空，无法保存消息记录")
		return
	}

	b, _ := sonic.Marshal(req)

	err := ent.Database.Message.Create().
		SetMessageID(*resp.Data.MessageId).
		SetNillableWorkitemID(workitemId).
		SetVaraibales(b).
		SetCreatedAt(time.Now()).
		SetType(typ).
		OnConflictColumns(message.FieldMessageID).
		UpdateNewValues().
		Exec(context.Background())
	if err != nil {
		log.Error().Err(err).Str("message_id", *resp.Data.MessageId).Msg("保存飞书消息记录失败")
	}
}

// SendApkMessage 发送 apk 测试发包消息
func (s *FeishuService) SendApkMessage(req *model.FeishuApkMessageRequest) {
	cfg := config.Get().Feishu.Message.ApkRelease

	msg := &feishu.ApkMessage{
		ID:       req.ID,
		AppName:  req.AppName,
		Message:  req.Message,
		Intranet: req.Intranet,
		Extranet: req.Extranet,
		Version:  req.Version,
	}

	data := feishu.CreateInteractiveMessageReq[feishu.ApkMessage](
		cfg.TemplateId,
		feishu.ReceiveIdTypeChatID,
		config.Get().Feishu.ApkReleaseGroupId,
		msg,
	)

	_, err := feishu.SendMessage(context.Background(), data)
	if err != nil {
		log.Error().Err(err).Msg("发送飞书 apk 测试发包消息失败")
	}
	return
}

// SendUnderReviewMessage 发送待审查消息
func (s *FeishuService) SendUnderReviewMessage(req *model.FeishuSendUnderReviewMessageRequest) {
	msg := &feishu.UnderReviewMessage{
		ID:          req.ID,
		Title:       req.Title,
		Category:    req.Category,
		Theme:       req.Theme,
		Description: req.Description,
		ReviewUsers: req.ReviewUsers,
		Url:         req.Url,
	}

	cfg := config.Get().Feishu.Message.UnderReview
	data := feishu.CreateInteractiveMessageReq[feishu.UnderReviewMessage](
		cfg.TemplateId,
		feishu.ReceiveIdTypeChatID,
		config.Get().Feishu.DevopsGroupId,
		msg,
	)

	resp, err := feishu.SendMessage(context.Background(), data)
	if err != nil {
		log.Error().Err(err).Msg("发送飞书待审查消息失败")
		return
	}

	// 保存消息记录
	s.SaveMessage(resp, &req.ID, req, message.TypeUnderReview)
}

// SendJobMessage 发送新工作消息
func (s *FeishuService) SendJobMessage(req *model.FeishuSendJobMessageRequest) {
	msg := &feishu.JobMessage{
		ID:          req.ID,
		Title:       req.Title,
		Category:    req.Category,
		Theme:       req.Theme,
		Description: req.Description,
		Url:         req.Url,
		Icon:        &feishu.MessageImageVaraibale{ImgKey: req.Icon},
		Status:      req.Status,
	}

	cfg := config.Get().Feishu.Message.Job
	data := feishu.CreateInteractiveMessageReq[feishu.JobMessage](
		cfg.TemplateId,
		req.ReceiveIdType,
		req.ReceiveId,
		msg,
	)

	resp, err := feishu.SendMessage(context.Background(), data)
	if err != nil {
		log.Error().Err(err).Msg("发送飞书新工作消息失败")
	}

	// 保存消息记录
	s.SaveMessage(resp, &req.ID, req, message.TypeJob)
}

// HookCardActionTrigger 飞书卡片操作触发事件处理
func (s *FeishuService) HookCardActionTrigger(_ context.Context, event *callback.CardActionTriggerEvent) (*callback.CardActionTriggerResponse, error) {
	var card *callback.Card
	toast := &callback.Toast{
		Type:    "error",
		Content: "未知错误",
	}

	if event.Event != nil && event.Event.Action != nil && event.Event.Action.Value != nil && event.Event.Operator != nil && event.Event.Operator.UserID != nil {
		values := event.Event.Action.Value
		action, _ := values["action"].(string)
		switch action {
		case "reviewed":
			toast, card = s.cardActionReviewed(event, values)
		case "job-inProgress":
			toast, card = s.cardActionJobInProgress(event, values)
		}
	}

	output := &callback.CardActionTriggerResponse{
		Toast: toast,
		Card:  card,
	}

	return output, nil
}

// 更新审查完成卡片
func (s *FeishuService) cardActionReviewed(event *callback.CardActionTriggerEvent, values map[string]any) (toast *callback.Toast, card *callback.Card) {
	id, _ := values["ID"].(string)
	title, _ := values["TITLE"].(string)
	category, _ := values["CATEGORY"].(string)
	url, _ := values["URL"].(string)
	allowedStr, _ := values["REVIEW_USERS"].(string)
	userId := *event.Event.Operator.UserID

	if id != "" && title != "" && category != "" && url != "" && allowedStr != "" {
		// 验证权限
		allowed := strings.Split(allowedStr, ",")
		if !slices.Contains(allowed, userId) {
			toast = &callback.Toast{
				Type:    "error",
				Content: "您没有权限执行此操作",
			}
			return
		}

		toast = &callback.Toast{
			Type:    "info",
			Content: "成功",
		}

		// 更新卡片
		// https://open.feishu.cn/document/feishu-cards/card-callback-communication#3ac8c17d
		card = &callback.Card{
			Type: "template",
			Data: &feishu.InteractiveTemplateMessageData[feishu.ReviewedMessage]{
				TemplateId: config.Get().Feishu.Message.Reviewed.TemplateId,
				TemplateVariable: &feishu.ReviewedMessage{
					ID:       id,
					Title:    title,
					Category: category,
					Url:      url,
					ReviewedUsers: []*feishu.MessageUserVaraibale{
						{ID: userId},
					},
					ReviewedTime: time.Now().Format(time.DateTime),
				},
			},
		}

		go NewYunxiao().UpdateStatusByID(id, yc.WorkflowStatusResolved)
	} else {
		log.Warn().Any("values", values).Msg("飞书卡片操作触发事件，跳过卡片更新：参数不完整")
	}

	return
}

// 更新工作中卡片
func (s *FeishuService) cardActionJobInProgress(event *callback.CardActionTriggerEvent, values map[string]any) (toast *callback.Toast, card *callback.Card) {
	// TODO: 待实现
	return
}
