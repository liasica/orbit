// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import "github.com/spf13/cobra"

var (
	rootGroup = &cobra.Group{
		ID:    "yunxiao",
		Title: "云效指令",
	}

	rootCmd = &cobra.Command{
		Use:               "yunxiao",
		Short:             "云效相关指令",
		GroupID:           rootGroup.ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func Command() *cobra.Command {
	rootCmd.AddGroup(configureGroup)
	rootCmd.AddCommand(configureCmd)
	return rootCmd
}

func Group() *cobra.Group {
	return rootGroup
}
