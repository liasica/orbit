// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package boot

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/app/service"
	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/integration/feishu"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
)

func Bootstrap(cfgPath string) {
	// 设置全局时区
	tz := "Asia/Shanghai"
	_ = os.Setenv("TZ", tz)
	loc, _ := time.LoadLocation(tz)
	time.Local = loc

	// 日志初始化
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05.000",
	}).With().CallerWithSkipFrameCount(2).Logger()

	log.Info().Msg("Booting ...")

	// 读取配置文件
	config.Setup(cfgPath)

	// 初始化数据库
	ent.Setup(config.Get().Database.Postgres.Dsn, config.Get().Database.Postgres.Debug)

	// 初始化gitlab
	gitlab.Setup()

	// 初始化飞书
	feishu.Setup(
		feishu.WithHookCardActionTrigger(service.NewFeishu().HookCardActionTrigger),
	)

	// 初始化云效集成配置
	yunxiao.Setup()

	log.Info().Msg("Boot completed")
}
