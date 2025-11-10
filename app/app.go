// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-30, by liasica

package app

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/pkg/utils"
)

func Run(addr string) {
	e := echo.New()

	e.HidePort = true
	e.HideBanner = true
	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	// 绑定校验器
	e.Validator = NewValidator()

	// 添加中间件
	recoverConfig := middleware.DefaultRecoverConfig
	recoverConfig.LogErrorFunc = func(_ echo.Context, err error, stack []byte) error {
		log.Error().
			Err(err).
			Bytes("stack", stack).
			Msg("[RECOVER] 捕获HTTP未处理崩溃")
		return err
	}
	e.Use(
		middleware.RecoverWithConfig(recoverConfig),
	)

	// 错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		code := http.StatusInternalServerError
		message := err.Error()

		var target validator.ValidationErrors
		if errors.As(err, &target) {
			code = http.StatusBadRequest
		}

		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, map[string]interface{}{"error": message})
		}

		if err != nil {
			log.Error().Err(err).Msg("发送错误响应失败")
		}
	}

	// 添加路由
	addControllers(e)

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

	log.Info().Msgf("REST API 服务已启动，监听地址: %s, 公网IP地址: %s", addr, myIp)

	// 当中断信号发生时，关闭服务器并返回
	<-sig.Done()
	err = e.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("服务器关闭失败")
	}
	log.Info().Msg("服务器已关闭")
}
