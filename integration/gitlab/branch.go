// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package gitlab

import gitlab "gitlab.com/gitlab-org/api/client-go"

// GetBranch 获取分支
func GetBranch(pid any, branch string) (b *gitlab.Branch, err error) {
	b, _, err = instance.client.Branches.GetBranch(pid, branch)
	return
}

// CreateBranch 创建分支
// 如果 ref 为空，则使用项目的默认分支
func CreateBranch(pid any, branch, ref string) (b *gitlab.Branch, err error) {
	// 先检查分支是否存在
	b, _ = GetBranch(pid, branch)
	if b != nil {
		return
	}

	// ref 为空则使用默认分支
	if ref == "" {
		var p *gitlab.Project
		p, err = GetProject(pid)
		if err != nil {
			return
		}
		ref = p.DefaultBranch
	}

	// 创建分支
	b, _, err = instance.client.Branches.CreateBranch(pid, &gitlab.CreateBranchOptions{
		Branch: &branch,
		Ref:    &ref,
	})
	return
}
