// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package gitlab

import (
	"sync"

	"github.com/rs/zerolog/log"
	gitlab "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/config"
)

var (
	instance *Gitlab
	once     sync.Once
)

type Gitlab struct {
	token         string
	webhookSecret string
	client        *gitlab.Client
}

func Setup() {
	once.Do(func() {
		cfg := config.Get().Gitlab

		client, err := gitlab.NewClient(cfg.Token, gitlab.WithBaseURL(cfg.BaseUrl))
		if err != nil {
			log.Fatal().Err(err).Msg("gitlab 客户端初始化失败")
		}

		instance = &Gitlab{
			token:         cfg.Token,
			webhookSecret: cfg.WebhookSecret,
			client:        client,
		}
	})
}
