// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package gitlab

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type WebhookHandler func(eventType gitlab.EventType, data any)

func Webhook(req *http.Request, w http.ResponseWriter, handle WebhookHandler) {
	defer func() {
		// 直接返回 200 OK
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}()

	// 校验key
	// X-Gitlab-Token -> len:1, cap:1
	if gitlab.HookEventToken(req) != instance.webhookSecret {
		log.Error().Msg("Webhook Secret 校验失败")
		return
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		log.Error().Err(err).Msg("读取 GitLab Webhook 请求体失败")
		return
	}

	// 读取事件类型
	eventType := gitlab.HookEventType(req)

	var data any
	data, err = gitlab.ParseHook(eventType, b)
	if err != nil {
		log.Error().Err(err).Msgf("解析 GitLab Webhook 请求体失败, event: %s", eventType)
		return
	}

	go handle(eventType, data)
}
