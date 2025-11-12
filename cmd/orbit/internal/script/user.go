// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package script

import (
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/app/service"
)

type User struct {
	cmd *cobra.Command
}

func (c *User) Group() *cobra.Group {
	return &cobra.Group{
		ID:    "user",
		Title: "用户指令",
	}
}

func (c *User) Command() *cobra.Command {
	return c.cmd
}

func NewUser() *User {
	c := &User{}

	c.cmd = &cobra.Command{
		Use:               "user",
		Short:             "用户相关指令",
		GroupID:           c.Group().ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	c.cmd.AddCommand(
		&cobra.Command{
			Use:               "sync",
			Short:             "同步用户数据库",
			CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
			Run: func(_ *cobra.Command, _ []string) {
				service.NewUser().Sync()
			},
		},
	)

	return c
}
