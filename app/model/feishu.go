// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package model

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
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Category    string `json:"category,omitempty"`
	Theme       string `json:"theme,omitempty"`
	Description string `json:"description,omitempty"`
	ReviewUsers string `json:"reviewUsers,omitempty"`
	Url         string `json:"url,omitempty"`
}

// FeishuSendJobMessageRequest 飞书新工作消息请求体
type FeishuSendJobMessageRequest struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Category    string `json:"category,omitempty"`
	Theme       string `json:"theme,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Status      string `json:"status,omitempty"`
}
