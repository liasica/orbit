// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package service

import (
	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/ent/configure"
	"github.com/liasica/orbit/integration/yunxiao/entity"
	"github.com/liasica/orbit/repository"
)

type YunxiaoService struct {
}

func NewYunxiao() *YunxiaoService {
	return &YunxiaoService{}
}

// GetConfigure 获取云效配置
func (s *YunxiaoService) GetConfigure() entity.ConfigureMap {
	data, err := repository.NewConfigure().GetValue(configure.KeyYunxiao)
	if err != nil {
		log.Error().Err(err).Msg("读取云效配置失败")
		return nil
	}

	cfg := make(entity.ConfigureMap)
	err = sonic.Unmarshal(data, &cfg)
	if err != nil {
		log.Error().Err(err).Msg("解析云效配置失败")
		return nil
	}

	return cfg
}

// GetWorkitemConfigure 获取工作项配置
func (s *YunxiaoService) GetWorkitemConfigure(category entity.WorkitemCategory) (cfg *entity.WorkitemConfigure) {
	m := s.GetConfigure()
	w, ok := m[category]
	if !ok {
		log.Error().Msgf("获取工作项 %s 配置失败", category)
		return
	}

	return w.Workitem
}

func (s *YunxiaoService) StatusWebhookHandler(data *entity.WebhookStatusEvent) {
	// 打开 / 重新打开 → 处理中
	if (data.From == entity.ConfigureWorkflowStatusOpen || data.From == entity.ConfigureWorkflowStatusReopen) && data.To == entity.ConfigureWorkflowStatusInProgress {
		NewCooperation().YunxiaoStatusChanged(data)
	}
}
