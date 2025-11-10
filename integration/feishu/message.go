// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-08, by liasica

package feishu

import (
	"context"

	"github.com/bytedance/sonic"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// MessageUserVaraibale 模板消息用户变量
type MessageUserVaraibale struct {
	ID string `json:"id"`
}

// SendMessage 发送消息
func SendMessage(ctx context.Context, req *v1.CreateMessageReq, options ...larkcore.RequestOptionFunc) (*v1.CreateMessageResp, error) {
	resp, err := instance.client.Im.V1.Message.Create(ctx, req, options...)
	return resp, err
}

// InteractiveTemplateMessage 交互式模板消息
type InteractiveTemplateMessage[T any] struct {
	Type string                            `json:"type"`
	Data InteractiveTemplateMessageData[T] `json:"data"`
}

// InteractiveTemplateMessageData 交互式模板消息数据
type InteractiveTemplateMessageData[T any] struct {
	TemplateId          string `json:"template_id,omitempty"`
	TemplateVersionName string `json:"template_version_name,omitempty"`
	TemplateVariable    *T     `json:"template_variable,omitempty"`
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
