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

// ApkMessage apk发包消息模板变量
// https://open.feishu.cn/cardkit/editor?cardId=AAqhHKY2IV0xD
type ApkMessage struct {
	ID       string `json:"CI_JOB_ID,omitempty"`
	AppName  string `json:"APP_NAME,omitempty"`
	Message  string `json:"CI_COMMIT_MESSAGE,omitempty"`
	Intranet string `json:"INTRANET,omitempty"`
	Extranet string `json:"EXTRANET,omitempty"`
	Version  string `json:"VERSION,omitempty"`
}

// SendApkMessage 发送apk发包消息
func SendApkMessage(ctx context.Context, message *ApkMessage) (*v1.CreateMessageResp, error) {
	cfg := config.Get().Feishu.Message.ApkRelease

	req := v1.NewCreateMessageReqBuilder().
		ReceiveIdType(cfg.ReceiveIdType).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &cfg.ReceiveId,
			MsgType:   &cfg.MsgType,
			Content:   NewInteractiveTemplateMessage[ApkMessage](cfg.TemplateId, message).StringPtr(),
			Uuid:      larkcore.StringPtr(uuid.New().String()),
		})

	return SendMessage(ctx, req.Build())
}
