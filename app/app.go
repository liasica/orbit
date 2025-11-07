// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-30, by liasica

package app

import (
	"context"
	"os/signal"
	"strings"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/app/service"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/pkg/utils"
)

func Run(addr string) {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true

	e.Use(
		middleware.Recover(),
	)

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

	// 获取端口号
	port := strings.Split(addr, ":")[1]

	// 获取IP地址
	myIp, err := utils.GetMyIP()
	if err != nil {
		log.Fatal().Err(err).Msg("获取本机IP地址失败")
	}

	ctx := context.Background()
	sig, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// 使用协程启动服务
	go func() {
		if err = e.Start(addr); err != nil {
			log.Fatal().Err(err)
		}
	}()

	log.Info().Msgf("REST API 服务已启动，监听地址: %s, 公网IP地址: %s:%s", addr, myIp, port)

	// 当中断信号发生时，关闭服务器并返回
	<-sig.Done()
	err = e.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("服务器关闭失败")
	}
	log.Info().Msg("服务器已关闭")
}
