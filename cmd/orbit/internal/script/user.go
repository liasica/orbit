// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package script

import (
	"context"

	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	git "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/user"
	"github.com/liasica/orbit/integration/feishu"
	"github.com/liasica/orbit/integration/gitlab"
	"github.com/liasica/orbit/integration/yunxiao"
	"github.com/liasica/orbit/integration/yunxiao/entity"
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
			Use:               "update",
			Short:             "更新用户数据库",
			CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
			Run: func(_ *cobra.Command, _ []string) {
				ctx := context.Background()

				// 获取飞书用户列表
				fsUsers, err := feishu.FindUserByDepartment(ctx, config.Get().Feishu.DepartmentId)
				if err != nil {
					log.Fatal().Err(err).Msg("获取飞书用户列表失败")
				}
				fsUsersMap := make(map[string]*larkcontact.User)
				for _, fsUser := range fsUsers {
					if fsUser.Name == nil || fsUser.UserId == nil || fsUser.OpenId == nil || fsUser.UnionId == nil {
						continue
					}
					fsUsersMap[*fsUser.Name] = fsUser
				}

				// 获取云效用户列表
				var yxUsers []entity.Memeber
				yxUsers, err = yunxiao.ListProjectMembers()
				if err != nil {
					log.Fatal().Err(err).Msg("获取云效用户列表失败")
				}
				yxUsersMap := make(map[string]entity.Memeber)
				for _, ycUser := range yxUsers {
					yxUsersMap[ycUser.UserName] = ycUser
				}

				// 获取 gitlab 用户列表
				var glUsers []*git.User
				glUsers, err = gitlab.ListUsers(nil)
				if err != nil {
					log.Fatal().Err(err).Msg("获取 GitLab 用户列表失败")
				}
				glUsersMap := make(map[string]*git.User)
				for _, glUser := range glUsers {
					glUsersMap[glUser.Name] = glUser
				}

				for name, fsUser := range fsUsersMap {
					creator := ent.Database.User.Create().
						SetName(name).
						SetLarkUserID(*fsUser.UserId).
						SetLarkOpenID(*fsUser.OpenId).
						SetLarkUnionID(*fsUser.UnionId)

					ycUser, ok := yxUsersMap[name]
					if !ok {
						log.Warn().Str("name", name).Msg("云效用户不存在，跳过")
						continue
					}
					creator.SetYunxiaoUserID(ycUser.UserId)

					var glUser *git.User
					glUser, ok = glUsersMap[name]
					if !ok {
						log.Warn().Str("name", name).Msg("GitLab 用户不存在，跳过")
						continue
					}

					err = creator.SetGitlabUsername(glUser.Username).
						SetGitlabEmail(glUser.Email).
						OnConflictColumns(user.FieldLarkUserID).
						UpdateNewValues().
						Exec(ctx)

					if err != nil {
						log.Error().Err(err).Str("name", name).Msg("创建或更新用户失败")
					}
				}
			},
		},
	)

	return c
}
