// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package pingcode

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
)

func testSetup(t *testing.T) {
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

	config.Setup("../../configs/config.yaml")

	// 初始化数据库
	ent.Setup(config.Get().Database.Postgres.Dsn, config.Get().Database.Postgres.Debug)

	Setup(&Config{
		BaseUrl:  config.Get().PingCode.BaseUrl,
		ClientID: config.Get().PingCode.ClientID,
		Secret:   config.Get().PingCode.Secret,
	})
}
