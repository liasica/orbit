// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-11, by liasica

package feishu

import (
	"context"

	"github.com/google/uuid"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	"github.com/liasica/orbit/config"
)

// JobMessage 新工作消息模板
// https://open.feishu.cn/cardkit/editor?cardId=AAqhj54SE7sTh
type JobMessage struct {
	ID          string                 `json:"ID,omitempty"`
	Title       string                 `json:"TITLE,omitempty"`
	Category    string                 `json:"CATEGORY,omitempty"`
	Theme       string                 `json:"THEME,omitempty"`
	Description string                 `json:"DESCRIPTION,omitempty"`
	Url         string                 `json:"URL,omitempty"`
	Icon        *MessageImageVaraibale `json:"ICON,omitempty"`
	Status      string                 `json:"STATUS,omitempty"`
}

// SendJobMessage 发送新工作消息
func SendJobMessage(ctx context.Context, message *JobMessage) (*v1.CreateMessageResp, error) {
	cfg := config.Get().Feishu.Message.Job

	req := v1.NewCreateMessageReqBuilder().
		ReceiveIdType(cfg.ReceiveIdType).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &cfg.ReceiveId,
			MsgType:   &cfg.MsgType,
			Content:   NewInteractiveTemplateMessage[JobMessage](cfg.TemplateId, message).StringPtr(),
			Uuid:      larkcore.StringPtr(uuid.New().String()),
		})

	return SendMessage(ctx, req.Build())
}
