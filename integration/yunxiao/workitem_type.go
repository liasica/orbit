// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
)

// ListProjectWorkitemTypesRequest 获取项目中的工作项类型
// https://next.api.aliyun.com/api/devops/2021-06-25/ListProjectWorkitemTypes
type ListProjectWorkitemTypesRequest struct {
	params devops.ListProjectWorkitemTypesRequest
}

func NewListProjectWorkitemTypesRequest(category WorkitemCategory) *ListProjectWorkitemTypesRequest {
	return &ListProjectWorkitemTypesRequest{
		params: devops.ListProjectWorkitemTypesRequest{
			Category:  tea.String(string(category)),
			SpaceType: tea.String("Project"),
		},
	}
}

func (req *ListProjectWorkitemTypesRequest) Do() (*devops.ListProjectWorkitemTypesResponseBody, error) {
	res, err := instance.client.ListProjectWorkitemTypes(&instance.config.OrganizationId, &instance.config.ProjectId, &req.params)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func ListAllWorkitemTypes() (items map[WorkitemCategory][]*devops.ListProjectWorkitemTypesResponseBodyWorkitemTypes) {
	items = make(map[WorkitemCategory][]*devops.ListProjectWorkitemTypesResponseBodyWorkitemTypes)
	for _, cate := range WorkitemCategories {
		req := NewListProjectWorkitemTypesRequest(cate)
		data, err := req.Do()
		if err != nil {
			continue
		}
		items[cate] = data.WorkitemTypes
	}

	return
}
