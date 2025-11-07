// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package yunxiao

import (
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

var (
	fieldGroup = &cobra.Group{
		ID:    "yunxiao-field",
		Title: "云效字段配置指令",
	}

	fieldCmd = &cobra.Command{
		Use:               "field",
		Short:             "云效字段配置相关指令",
		GroupID:           fieldGroup.ID,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func init() {
	fieldCmd.AddCommand(fieldRepositoryCmd())
}

func fieldRepositoryCmd() (cmd *cobra.Command) {
	var fieldId string

	cmd = &cobra.Command{
		Use:               "repository",
		Short:             "更新仓库字段配置",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run: func(_ *cobra.Command, _ []string) {
			ps, err := gitlab.ListProjects(nil)
			if err != nil {
				log.Fatal().Err(err)
			}
			var options []string
			for _, p := range ps {
				options = append(options, p.PathWithNamespace)
			}

			// 按名称排序
			slices.SortFunc(options, func(a, b string) int {
				return strings.Compare(a, b)
			})

			err = yunxiao.UpdateCustomField(fieldId, &entity.CustomField{
				Name:    "代码仓库",
				Options: options,
			})
			if err != nil {
				log.Fatal().Err(err)
			}
		},
	}

	cmd.Flags().StringVar(&fieldId, "id", "1e216b5770aac61d8778651c86", "字段ID")

	return
}
