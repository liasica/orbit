// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package feishu

import (
	"context"

	"github.com/google/uuid"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	"github.com/liasica/orbit/config"
)

// ReviewedMessage 已审查消息模板变量
// https://open.feishu.cn/cardkit/editor?cardId=AAqhD4KzLB1l1
type ReviewedMessage struct {
	ID            string                  `json:"ID,omitempty"`
	Title         string                  `json:"TITLE,omitempty"`
	Category      string                  `json:"CATEGORY,omitempty"`
	Url           string                  `json:"URL,omitempty"`
	ReviewedUsers []*MessageUserVaraibale `json:"REVIEWED_USERS,omitempty"`
	ReviewedTime  string                  `json:"REVIEWED_TIME,omitempty"`
}

// UnderReviewMessage 待审查消息模板变量
// https://open.feishu.cn/cardkit/editor?cardId=AAqhCfa37Qjtc
type UnderReviewMessage struct {
	ID          string `json:"ID,omitempty"`
	Title       string `json:"TITLE,omitempty"`
	Category    string `json:"CATEGORY,omitempty"`
	Theme       string `json:"THEME,omitempty"`
	Description string `json:"DESCRIPTION,omitempty"`
	ReviewUsers string `json:"REVIEW_USERS,omitempty"`
	Url         string `json:"URL,omitempty"`
}

// SendUnderReviewMessage 发送待审查消息
func SendUnderReviewMessage(ctx context.Context, message *UnderReviewMessage) (*v1.CreateMessageResp, error) {
	cfg := config.Get().Feishu.Message.UnderReview

	req := v1.NewCreateMessageReqBuilder().
		ReceiveIdType(cfg.ReceiveIdType).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &cfg.ReceiveId,
			MsgType:   &cfg.MsgType,
			Content:   NewInteractiveTemplateMessage[UnderReviewMessage](cfg.TemplateId, message).StringPtr(),
			Uuid:      larkcore.StringPtr(uuid.New().String()),
		})

	return SendMessage(ctx, req.Build())
}
