// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
)

// ListWorkitemAllFieldsRequest 获取工作项字段列表
// https://next.api.aliyun.com/api/devops/2021-06-25/ListWorkItemAllFields
type ListWorkitemAllFieldsRequest struct {
	params devops.ListWorkItemAllFieldsRequest
}

func NewListWorkitemAllFieldsRequest(workitemTypeIdentifier string) *ListWorkitemAllFieldsRequest {
	return &ListWorkitemAllFieldsRequest{
		params: devops.ListWorkItemAllFieldsRequest{
			SpaceIdentifier:        &instance.config.ProjectId,
			SpaceType:              tea.String("Project"),
			WorkitemTypeIdentifier: &workitemTypeIdentifier,
		},
	}
}

func (req *ListWorkitemAllFieldsRequest) Do() (*devops.ListWorkItemAllFieldsResponseBody, error) {
	res, err := instance.client.ListWorkItemAllFields(&instance.config.OrganizationId, &req.params)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// UpdateWorkitemFieldRequest 更新工作项字段，可以支持批量更新多个字段
// https://next.api.aliyun.com/api/devops/2021-06-25/UpdateWorkitemField
type UpdateWorkitemFieldRequest struct {
	params devops.UpdateWorkitemFieldRequest
}

// type UpdateWorkitemFieldRequestOption func(*UpdateWorkitemFieldRequest)
//
// func NewUpdateWorkitemFieldRequest(workItemId string) *UpdateWorkitemFieldRequest {
//
// }
