// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package app

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/app/model"
	"github.com/liasica/orbit/app/service"
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

func getBody(c echo.Context) (b []byte) {
	var err error
	b, err = io.ReadAll(c.Request().Body)
	if err != nil {
		log.Warn().Err(err).Msg("读取请求体失败")
	}
	if b == nil {
		b = make([]byte, 0)
	}
	return
}

func addControllers(e *echo.Echo) {
	// gitlab webhook
	e.POST("/gitlab/webhook", func(c echo.Context) error {
		secret := git.HookEventToken(c.Request())
		eventType := git.HookEventType(c.Request())
		body := getBody(c)

		go service.NewGitlab().Webhook(secret, eventType, body)

		return c.NoContent(http.StatusOK)
	})

	// 云效 webhook
	e.POST("/yunxiao/webhook", func(c echo.Context) error {
		headers := c.Request().Header

		// 获取 body
		body := getBody(c)

		go service.NewYunxiao().Webhook(headers, body)

		return c.NoContent(http.StatusOK)
	})

	// 发送 apk message
	e.POST("/feishu/message/apk", func(c echo.Context) error {
		req := bindRequest[model.FeishuApkMessageRequest](c)
		go service.NewFeishu().SendApkMessage(req)
		return c.NoContent(http.StatusOK)
	})

	// 发送 review message
	e.POST("/feishu/message/review", func(c echo.Context) error {
		req := bindRequest[model.FeishuSendUnderReviewMessageRequest](c)
		go service.NewFeishu().SendUnderReviewMessage(req)
		return c.NoContent(http.StatusOK)
	})

	// 发送 job message
	e.POST("/feishu/message/job", func(c echo.Context) error {
		req := bindRequest[model.FeishuSendJobMessageRequest](c)
		go service.NewFeishu().SendJobMessage(req)
		return c.NoContent(http.StatusOK)
	})
}
