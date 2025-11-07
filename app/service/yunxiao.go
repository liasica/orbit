// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package service

import (
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

type YunxiaoService struct {
}

func NewYunxiao() *YunxiaoService {
	return &YunxiaoService{}
}

// StatusWebhookHandler 处理状态变更 Webhook
func (s *YunxiaoService) StatusWebhookHandler(data *entity.WebhookStatusEvent) {
	// 打开 / 重新打开 → 处理中
	if (data.From == yc.WorkflowStatusOpen || data.From == yc.WorkflowStatusReopen) && data.To == yc.WorkflowStatusInProgress {
		NewCooperation().YunxiaoStatusChanged(data)
	}
}
