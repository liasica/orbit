// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package boot

import (
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/configure"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/repository"
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

	// 初始化云效集成配置
	yunxiao.Setup(
		&yunxiao.Config{
			AccessKeyId:     config.Get().Yunxiao.AccessKeyId,
			AccessKeySecret: config.Get().Yunxiao.AccessKeySecret,
			Endpoint:        config.Get().Yunxiao.Endpoint,
			OrganizationId:  config.Get().Yunxiao.OrganizationId,
			ProjectId:       config.Get().Yunxiao.ProjectId,
		},
		func() (*yunxiao.WorkitemConfigure, error) {
			data, err := repository.NewConfigure().GetValue(configure.KeyWorkitemConfigure)
			if err != nil {
				return nil, err
			}
			var cfg yunxiao.WorkitemConfigure
			err = sonic.Unmarshal(data, &cfg)
			if err != nil {
				return nil, err
			}

			return &cfg, nil
		},
	)

	log.Info().Msg("Boot completed")
}
