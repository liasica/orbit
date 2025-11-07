// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao"
)

var (
	configureGroup = &cobra.Group{
		ID:    "yunxiao-configure",
		Title: "云效配置指令",
	}

	configureCmd = &cobra.Command{
		Use:               "configure",
		Short:             "云效配置相关指令",
		GroupID:           configureGroup.ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func init() {
	configureCmd.AddCommand(configureUpdateCmd())
}

func configureUpdateCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "update",
		Short:             "更新配置",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run:               doConfigureUpdate,
	}
	return
}

func doConfigureUpdate(_ *cobra.Command, _ []string) {
	now := time.Now()
	log.Info().Msg("开始更新云效配置...")

	data, err := yunxiao.GetConfigure()
	if err != nil {
		log.Fatal().Err(err)
	}

	err = yc.Store(data)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("云效配置更新完成，耗时: " + time.Since(now).String())
}
