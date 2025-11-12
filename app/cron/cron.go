// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-12, by liasica

package cron

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/liasica/orbit/app/service"
)

func Run() {
	log.Info().Msg("启动定时任务服务")

	ticker := time.NewTicker(10 * time.Minute)
	for ; ; <-ticker.C {
		now := time.Now()
		log.Info().Msg("正在执行定时任务...")

		// 更新仓库列表
		service.NewGitlab().StoreProjects()

		// 更新云效仓库字段
		service.NewYunxiao().UpdateRepositoryField()

		// 更新用户列表
		service.NewUser().Sync()

		log.Info().Msgf("定时任务执行完毕，耗时 %s", time.Since(now).String())
	}
}
