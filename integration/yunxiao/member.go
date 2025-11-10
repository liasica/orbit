// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package yunxiao

import "github.com/liasica/orbit/integration/yunxiao/entity"

// ListProjectMembers 获取项目成员列表
func ListProjectMembers() (data []entity.Memeber, err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", instance.projectId).
		SetResult(&data).
		Get("/oapi/v1/projex/organizations/{organizationId}/projects/{id}/members")
	return
}
