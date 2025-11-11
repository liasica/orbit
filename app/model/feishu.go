// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package model

import "github.com/liasica/orbit/integration/feishu"

// FeishuApkMessageRequest 飞书 APK 测试发包消息请求体
type FeishuApkMessageRequest struct {
	ID       string `json:"id,omitempty" validate:"required"`
	AppName  string `json:"appName,omitempty" validate:"required"`
	Message  string `json:"message,omitempty" validate:"required"`
	Intranet string `json:"intranet,omitempty" validate:"required"`
	Extranet string `json:"extranet,omitempty" validate:"required"`
	Version  string `json:"version,omitempty" validate:"required"`
}

// FeishuSendUnderReviewMessageRequest 飞书待审查消息请求体
type FeishuSendUnderReviewMessageRequest struct {
	ID          string `json:"id,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Category    string `json:"category,omitempty" validate:"required"`
	Theme       string `json:"theme,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	ReviewUsers string `json:"reviewUsers,omitempty" validate:"required"`
	Url         string `json:"url,omitempty" validate:"required"`
}

// FeishuSendJobMessageRequest 飞书新工作消息请求体
type FeishuSendJobMessageRequest struct {
	ReceiveIdType feishu.ReceiveIdType `json:"receiveIdType" validate:"required"`
	ReceiveId     string               `json:"receiveId" validate:"required"`

	ID          string `json:"id,omitempty" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Category    string `json:"category,omitempty" validate:"required"`
	Theme       string `json:"theme,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Url         string `json:"url,omitempty" validate:"required"`
	Icon        string `json:"icon,omitempty" validate:"required"`
	Status      string `json:"status,omitempty" validate:"required"`
}
