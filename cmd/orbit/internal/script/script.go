// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package script

import (
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/boot"
	"github.com/liasica/orbit/config"
)

type CommandGroup interface {
	Group() *cobra.Group
	Command() *cobra.Command
}

func Run() {
	var (
		configFile string
	)

	cmd := cobra.Command{
		Use:               "orbit",
		Short:             "极光出行协作工具",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Version:           config.GetVersion(),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			boot.Bootstrap(configFile)
		},
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./configs/config.yaml", "配置文件路径, 例如: ./configs/config.yaml")

	// 服务端指令
	serverCmd := NewServer()

	// 云效指令
	yunxiaoCmd := NewYunxiao()

	// 用户指令
	userCmd := NewUser()

	// 添加指令组
	cmd.AddGroup(
		serverCmd.Group(),
		yunxiaoCmd.Group(),
		userCmd.Group(),
	)

	// 添加子命令
	cmd.AddCommand(
		serverCmd.Command(),
		yunxiaoCmd.Command(),
		userCmd.Command(),
	)

	_ = cmd.Execute()
}
