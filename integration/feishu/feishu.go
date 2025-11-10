// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-08, by liasica

package feishu

import (
	"net/http"
	"sync"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"

	"github.com/liasica/orbit/config"
)

var (
	instance *Feishu
	once     sync.Once
)

type Feishu struct {
	client *lark.Client
}

func Setup() {
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
	})
}
