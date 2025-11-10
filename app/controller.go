// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package app

import (
	"github.com/labstack/echo/v4"

	"github.com/liasica/orbit/app/model"
	"github.com/liasica/orbit/app/service"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
)

func bindRequest[T any](c echo.Context) (req *T) {
	req = new(T)
	err := c.Bind(req)
	if err != nil {
		panic(err)
	}

	err = c.Validate(req)
	if err != nil {
		panic(err)
	}
	return
}

func addControllers(e *echo.Echo) {
	// gitlab webhook
	e.POST("/gitlab/webhook", func(c echo.Context) error {
		gitlab.Webhook(c.Request(), c.Response(), service.NewGitlab().WebhookHandler)
		return nil
	})

	// 云效 webhook
	e.POST("/yunxiao/webhook/status", func(c echo.Context) error {
		yunxiao.StatusWebhook(c.Request(), c.Response(), service.NewYunxiao().StatusWebhookHandler)
		return nil
	})

	// 发送 apk message
	e.POST("/feishu/message/apk", func(c echo.Context) error {
		return service.NewFeishu().SendApkMessage(bindRequest[model.FeishuApkMessageRequest](c))
	})

	// 发送 review message
	e.POST("/feishu/message/review", func(c echo.Context) error {
		return service.NewFeishu().SendUnderReviewMessage(bindRequest[model.FeishuUnderReviewMessageRequest](c))
	})
}
