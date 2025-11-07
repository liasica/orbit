// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package script

import (
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/app"
)

var serverGroup = &cobra.Group{
	ID:    "server",
	Title: "服务端指令",
}

func serverCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:               "server",
		Short:             "服务端相关指令",
		GroupID:           serverGroup.ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	cmd.AddCommand(
		serverRunCommand(),
	)

	return
}

func serverRunCommand() (cmd *cobra.Command) {
	var address string

	cmd = &cobra.Command{
		Use:               "run",
		Short:             "运行服务端",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run: func(_ *cobra.Command, _ []string) {
			app.Run(address)
		},
	}

	cmd.Flags().StringVar(&address, "address", "0.0.0.0:80", "服务端监听地址, 例如: 0.0.0.0:80")

	return
}
