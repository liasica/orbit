// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-08, by liasica

package feishu

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

	"github.com/liasica/orbit/config"
)

// SendMessage 发送消息
func SendMessage(ctx context.Context, req *v1.CreateMessageReq, options ...larkcore.RequestOptionFunc) (*v1.CreateMessageResp, error) {
	return instance.client.Im.V1.Message.Create(ctx, req, options...)
}

// InteractiveTemplateMessage 交互式模板消息
type InteractiveTemplateMessage[T any] struct {
	Type string                            `json:"type"`
	Data InteractiveTemplateMessageData[T] `json:"data"`
}

// InteractiveTemplateMessageData 交互式模板消息数据
type InteractiveTemplateMessageData[T any] struct {
	TemplateId       string `json:"template_id"`
	TemplateVariable *T     `json:"template_variable"`
}

// NewInteractiveTemplateMessage 创建交互式模板消息
func NewInteractiveTemplateMessage[T any](templateId string, data *T) *InteractiveTemplateMessage[T] {
	return &InteractiveTemplateMessage[T]{
		Type: "template",
		Data: InteractiveTemplateMessageData[T]{TemplateId: templateId, TemplateVariable: data},
	}
}

// StringPtr 将消息转换为字符串指针
func (m *InteractiveTemplateMessage[T]) StringPtr() *string {
	s, _ := sonic.MarshalString(m)
	return &s
}

// ApkMessageVariables apk发包消息模板变量
type ApkMessageVariables struct {
	ID         string `json:"CI_JOB_ID,omitempty"`
	AppName    string `json:"APP_NAME,omitempty"`
	Message    string `json:"CI_COMMIT_MESSAGE,omitempty"`
	RCIntranet string `json:"RC_INTRANET_DOWNLOAD,omitempty"`
	RCExtranet string `json:"RC_EXTRANET_DOWNLOAD,omitempty"`
	Version    string `json:"VERSION,omitempty"`
}

// SendApkMessage 发送apk发包消息
func SendApkMessage(ctx context.Context, variables *ApkMessageVariables) (*v1.CreateMessageResp, error) {
	cfg := config.Get().Feishu.Message.ApkRelease

	req := v1.NewCreateMessageReqBuilder().
		ReceiveIdType(cfg.ReceiveIdType).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &cfg.ReceiveId,
			MsgType:   &cfg.MsgType,
			Content:   NewInteractiveTemplateMessage[ApkMessageVariables](cfg.TemplateId, variables).StringPtr(),
			Uuid:      larkcore.StringPtr(uuid.New().String()),
		})

	return SendMessage(ctx, req.Build())
}
