// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package gitlab

import gitlab "gitlab.com/gitlab-org/api/client-go"

// ListUsers 列出所有用户
func ListUsers(opt *gitlab.ListUsersOptions) (users []*gitlab.User, err error) {
	if opt == nil {
		opt = &gitlab.ListUsersOptions{ListOptions: gitlab.ListOptions{}}
	}

	var res *gitlab.Response
	users, res, err = instance.client.Users.ListUsers(opt)
	if err != nil {
		return
	}

	if res.NextPage > 0 {
		var moreUsers []*gitlab.User
		opt.Page = res.NextPage
		moreUsers, err = ListUsers(opt)
		if err != nil {
			return
		}

		users = append(users, moreUsers...)
	}

	return
}
