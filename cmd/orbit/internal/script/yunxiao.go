// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package script

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/app/service"
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao"
)

type Yunxiao struct {
	cmd *cobra.Command
}

func (c *Yunxiao) Group() *cobra.Group {
	return &cobra.Group{
		ID:    "yunxiao",
		Title: "云效指令",
	}
}

func (c *Yunxiao) Command() *cobra.Command {
	return c.cmd
}

func NewYunxiao() (c *Yunxiao) {
	c = &Yunxiao{}

	c.cmd = &cobra.Command{
		Use:               "yunxiao",
		Short:             "云效相关指令",
		GroupID:           c.Group().ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	c.cmd.AddCommand(
		c.configureUpdateCmd(),
		c.fieldRepositoryCmd(),
	)

	return
}

func (c *Yunxiao) configureUpdateCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "configure",
		Short:             "更新配置",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run: func(_ *cobra.Command, _ []string) {
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
		},
	}

	return
}

func (c *Yunxiao) fieldRepositoryCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "field",
		Short:             "云效更新字段配置",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run: func(_ *cobra.Command, _ []string) {
			service.NewGitlab().StoreProjects()

			service.NewYunxiao().UpdateRepositoryField()
		},
	}

	return
}
