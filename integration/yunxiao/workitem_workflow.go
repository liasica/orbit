// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
)

type ListWorkitemWorkFlowStatusRequest struct {
	organizationId *string
	params         devops.ListWorkItemWorkFlowStatusRequest
}

func NewListWorkitemWorkFlowStatusRequest(category WorkitemCategory) *ListWorkitemWorkFlowStatusRequest {
	return &ListWorkitemWorkFlowStatusRequest{
		organizationId: &instance.config.OrganizationId,
		params: devops.ListWorkItemWorkFlowStatusRequest{
			SpaceIdentifier:            &instance.config.ProjectId,
			SpaceType:                  tea.String("Project"),
			WorkitemCategoryIdentifier: tea.String(string(category)),
		},
	}
}

func (req *ListWorkitemWorkFlowStatusRequest) Do() (*devops.ListWorkItemWorkFlowStatusResponseBody, error) {
	res, err := instance.client.ListWorkItemWorkFlowStatus(req.organizationId, &req.params)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// ListAllWorkitemWorkFlowStatus 列出所有工作项工作流状态
func ListAllWorkitemWorkFlowStatus() (items map[WorkitemCategory][]*devops.ListWorkItemWorkFlowStatusResponseBodyStatuses) {
	items = make(map[WorkitemCategory][]*devops.ListWorkItemWorkFlowStatusResponseBodyStatuses)
	for _, cate := range WorkitemCategories {
		req := NewListWorkitemWorkFlowStatusRequest(cate)
		data, err := req.Do()
		if err != nil {
			continue
		}
		items[cate] = data.Statuses
	}

	return items
}
