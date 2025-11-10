// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package service

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/app/model"
	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/integration/feishu"
)

type FeishuService struct {
}

func NewFeishu() *FeishuService {
	return &FeishuService{}
}

// SendApkMessage 发送 apk 测试发包消息
func (s *FeishuService) SendApkMessage(req *model.FeishuApkMessageRequest) (err error) {
	_, err = feishu.SendApkMessage(context.Background(), &feishu.ApkMessage{
		ID:       req.ID,
		AppName:  req.AppName,
		Message:  req.Message,
		Intranet: req.Intranet,
		Extranet: req.Extranet,
		Version:  req.Version,
	})
	return
}

// SendUnderReviewMessage 发送待审查消息
func (s *FeishuService) SendUnderReviewMessage(req *model.FeishuUnderReviewMessageRequest) (err error) {
	_, err = feishu.SendUnderReviewMessage(context.Background(), &feishu.UnderReviewMessage{
		ID:          req.ID,
		Title:       req.Title,
		Category:    req.Category,
		Theme:       req.Theme,
		Description: req.Description,
		ReviewUsers: req.ReviewUsers,
		Url:         req.Url,
	})
	return
}

func (s *FeishuService) HookCardActionTrigger(_ context.Context, event *callback.CardActionTriggerEvent) (*callback.CardActionTriggerResponse, error) {
	// go func() {
	// 	time.AfterFunc(1*time.Second, func() {
	// 		b, _ := sonic.MarshalIndent(event, "", "  ")
	// 		fmt.Println(string(b))
	// 	})
	// }()

	var card *callback.Card
	toast := &callback.Toast{
		Type:    "info",
		Content: "成功",
	}

	if event.Event != nil && event.Event.Action != nil && event.Event.Action.Value != nil && event.Event.Operator != nil && event.Event.Operator.UserID != nil {
		values := event.Event.Action.Value
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
				goto output
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

		} else {
			log.Warn().Any("values", values).Msg("飞书卡片操作触发事件，跳过卡片更新：参数不完整")
		}
	}

output:
	output := &callback.CardActionTriggerResponse{
		Toast: toast,
		Card:  card,
	}

	b, _ := sonic.MarshalIndent(output, "", "  ")
	fmt.Println(string(b))

	return output, nil
}
