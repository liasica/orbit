// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package feishu

import (
	"context"

	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	v3 "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// FindUserByDepartmentOption 根据部门ID获取用户列表可选参数
type FindUserByDepartmentOption func(builder *v3.FindByDepartmentUserReqBuilder)

// FindUserByDepartmentWithPageToken 设置分页Token
func FindUserByDepartmentWithPageToken(pageToken string) FindUserByDepartmentOption {
	return func(builder *v3.FindByDepartmentUserReqBuilder) {
		builder.PageToken(pageToken)
	}
}

// FindUserByDepartmentWithPageSize 设置分页大小
func FindUserByDepartmentWithPageSize(pageSize int) FindUserByDepartmentOption {
	return func(builder *v3.FindByDepartmentUserReqBuilder) {
		builder.PageSize(pageSize)
	}
}

// FindUserByDepartment 根据部门ID获取用户列表
func FindUserByDepartment(ctx context.Context, departmentId string, options ...FindUserByDepartmentOption) (users []*v3.User, err error) {
	// 创建请求对象
	req := larkcontact.NewFindByDepartmentUserReqBuilder().
		UserIdType("user_id").
		DepartmentIdType("open_department_id").
		DepartmentId(departmentId).
		PageSize(50)

	for _, opt := range options {
		opt(req)
	}

	var resp *v3.FindByDepartmentUserResp
	resp, err = instance.client.Contact.V3.User.FindByDepartment(ctx, req.Build())
	if err != nil {
		return
	}

	users = resp.Data.Items

	// 如果有更多数据，递归获取
	if resp.Data.HasMore != nil && *resp.Data.HasMore && resp.Data.PageToken != nil {
		var moreUsers []*v3.User
		moreUsers, err = FindUserByDepartment(ctx, departmentId, FindUserByDepartmentWithPageToken(*resp.Data.PageToken))
		if err != nil {
			return
		}

		users = append(users, moreUsers...)
	}

	return
}
