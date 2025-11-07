// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package gitlab

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GetProject 获取单个项目
func GetProject(pid any) (p *gitlab.Project, err error) {
	p, _, err = instance.client.Projects.GetProject(pid, &gitlab.GetProjectOptions{})
	return
}

// ListProjects 列出所有项目
func ListProjects(opt *gitlab.ListProjectsOptions) (ps []*gitlab.Project, err error) {
	if opt == nil {
		opt = &gitlab.ListProjectsOptions{ListOptions: gitlab.ListOptions{}}
	}

	var res *gitlab.Response
	ps, res, err = instance.client.Projects.ListProjects(opt)
	if err != nil {
		return
	}

	if res.NextPage > 0 {
		var morePs []*gitlab.Project
		opt.Page = res.NextPage
		morePs, err = ListProjects(opt)
		if err != nil {
			return
		}

		ps = append(ps, morePs...)
	}

	return
}
