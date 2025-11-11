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
)

// ReceiveIdType 接收者类型
// open_id/user_id/union_id/email/chat_id
type ReceiveIdType string

const (
	ReceiveIdTypeChatID ReceiveIdType = "chat_id" // 群
	ReceiveIdTypeUserID ReceiveIdType = "user_id" // 用户
	// ReceiveIdTypeOpenID  ReceiveIdType = "open_id"
	// ReceiveIdTypeUnionID ReceiveIdType = "union_id"
	// ReceiveIdTypeEmail   ReceiveIdType = "email"
)

// MessageUserVaraibale 模板消息用户变量
type MessageUserVaraibale struct {
	ID string `json:"id"`
}

// MessageImageVaraibale 模板消息图片变量
type MessageImageVaraibale struct {
	ImgKey string `json:"img_key,omitempty"`
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

func CreateInteractiveMessageReq[T any](templateId string, receiveIdType ReceiveIdType, receiveId string, msg *T) *v1.CreateMessageReq {
	return v1.NewCreateMessageReqBuilder().
		ReceiveIdType(string(receiveIdType)).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &receiveId,
			MsgType:   larkcore.StringPtr("interactive"),
			Content:   NewInteractiveTemplateMessage[T](templateId, msg).StringPtr(),
			Uuid:      larkcore.StringPtr(uuid.New().String()),
		}).
		Build()
}

// SendMessage 发送消息
func SendMessage(ctx context.Context, req *v1.CreateMessageReq, options ...larkcore.RequestOptionFunc) (*v1.CreateMessageResp, error) {
	resp, err := instance.client.Im.V1.Message.Create(ctx, req, options...)
	return resp, err
}

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
