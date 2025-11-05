// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/ent/configure"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/repository"
)

var (
	workitemGroup = &cobra.Group{
		ID:    "yunxiao-workitem",
		Title: "云效 Workitem 指令",
	}

	workitemCmd = &cobra.Command{
		Use:               "workitem",
		Short:             "云效 workitem 相关指令",
		GroupID:           workitemGroup.ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func init() {
	workitemCmd.AddCommand(workitemUpgradeCmd())
}

func workitemUpgradeCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "upgrade",
		Short:             "更新 workitem 配置",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run:               doWorkitemUpgrade,
	}
	return
}

func doWorkitemUpgrade(_ *cobra.Command, _ []string) {
	now := time.Now()
	log.Info().Msg("开始更新云效 Workitem 配置...")

	data, err := yunxiao.GetConfigure()
	if err != nil {
		log.Fatal().Err(err)
	}

	var b []byte
	b, err = sonic.Marshal(data)
	if err != nil {
		log.Fatal().Err(err)
	}
	err = repository.NewConfigure().SetValue(configure.KeyYunxiao, b)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("云效 Workitem 配置更新完成，耗时: " + time.Since(now).String())
}
