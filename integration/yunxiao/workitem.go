// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	"regexp"

	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

// IsWorkitemIdentifier 判断字符串是否是 workitemIdentifier
// 如果字符串是 hexadecimal string, 则认为是 workitemIdentifier, 否则调用 GetWorkitemInfo 接口获取 workitemIdentifier
func IsWorkitemIdentifier(str string) bool {
	// 允许大写和小写，且长度为偶数
	matched, _ := regexp.MatchString("^[0-9a-fA-F]+$", str)
	return matched && len(str)%2 == 0
}

// ListAllWorkitemTypes - 获取组织下所有工作项类型列表
// https://help.aliyun.com/zh/yunxiao/developer-reference/listallworkitemtypes
func ListAllWorkitemTypes() (data []entity.WorkitemType, err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", instance.projectId).
		SetQueryParam("categories", "Bug,Task").
		SetResult(&data).
		Get("/oapi/v1/projex/organizations/{organizationId}/workitemTypes")
	return
}

// ListWorkitemTypes 获取工作项类型列表
// https://help.aliyun.com/zh/yunxiao/developer-reference/listworkitemtypes
func ListWorkitemTypes(category yc.WorkitemCategory) (data []entity.WorkitemType, err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", instance.projectId).
		SetQueryParam("category", category.String()).
		SetResult(&data).
		Get("/oapi/v1/projex/organizations/{organizationId}/projects/{id}/workitemTypes")
	return
}

// GetWorkitemTypeFieldConfig 获取工作项类型字段配置
// https://help.aliyun.com/zh/yunxiao/developer-reference/getworkitemtypefieldconfig
func GetWorkitemTypeFieldConfig(workitemTypeId string) (data []entity.Field, err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("projectId", instance.projectId).
		SetPathParam("id", workitemTypeId).
		SetResult(&data).
		Get("/oapi/v1/projex/organizations/{organizationId}/projects/{projectId}/workitemTypes/{id}/fields")
	return
}

// GetWorkitemWorkflow 获取工作项工作流信息
func GetWorkitemWorkflow(workitemTypeId string) (data entity.Workflow, err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("projectId", instance.projectId).
		SetPathParam("id", workitemTypeId).
		SetResult(&data).
		Get("/oapi/v1/projex/organizations/{organizationId}/projects/{projectId}/workitemTypes/{id}/workflows")
	return
}

// GetWorkitem 获取工作项
// https://help.aliyun.com/zh/yunxiao/developer-reference/getworkitem
func GetWorkitem(workitemId string) (data *entity.Workitem, err error) {
	data = new(entity.Workitem)
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", workitemId).
		SetResult(data).
		Get("/oapi/v1/projex/organizations/{organizationId}/workitems/{id}")
	return
}

// CreateWorkitemComment 创建工作项评论
// https://help.aliyun.com/zh/yunxiao/developer-reference/createworkitemcomment
func CreateWorkitemComment(workitemId, content string) (err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", workitemId).
		SetBody(map[string]any{
			"content": content,
		}).
		Post("/oapi/v1/projex/organizations/{organizationId}/workitems/{id}/comments")
	return
}

// UpdateWorkitem 更新工作项
func UpdateWorkitem(workitemId string, body map[string]string) (err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", workitemId).
		SetBody(body).
		Put("/oapi/v1/projex/organizations/{organizationId}/workitems/{id}")
	return
}
