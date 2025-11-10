// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-08, by liasica

package feishu

import (
	"context"
	"net/http"
	"sync"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/config"
)

var (
	instance *Feishu
	once     sync.Once
)

type Feishu struct {
	client *lark.Client

	hookCardActionTrigger HookCardActionTrigger
}

type Option func(*Feishu)

func WithHookCardActionTrigger(fn HookCardActionTrigger) Option {
	return func(feishu *Feishu) {
		feishu.hookCardActionTrigger = fn
	}
}

func eventDispatcher() *dispatcher.EventDispatcher {
	return dispatcher.NewEventDispatcher("", "").
		OnP2CardActionTrigger(instance.hookCardActionTrigger)
}

func Setup(opts ...Option) {
	once.Do(func() {
		cfg := config.Get().Feishu

		client := lark.NewClient(
			cfg.AppId,
			cfg.AppSecret,
			lark.WithLogLevel(larkcore.LogLevelDebug),
			lark.WithLogger(NewLogger()),
			lark.WithReqTimeout(10*time.Second),
			lark.WithEnableTokenCache(true),
			lark.WithHelpdeskCredential("id", "token"),
			lark.WithHttpClient(http.DefaultClient),
			lark.WithSerialization(NewSerialization()),
			lark.WithLogReqAtDebug(cfg.Debug),
		)

		instance = &Feishu{
			client: client,
		}

		for _, opt := range opts {
			opt(instance)
		}

		ws := larkws.NewClient(
			cfg.AppId,
			cfg.AppSecret,
			larkws.WithLogLevel(larkcore.LogLevelInfo),
			larkws.WithLogger(NewLogger()),
			larkws.WithEventHandler(eventDispatcher()),
		)

		go func() {
			err := ws.Start(context.Background())
			if err != nil {
				log.Fatal().Err(err).Msg("飞书WebSocket连接失败")
			}
		}()
	})
}
