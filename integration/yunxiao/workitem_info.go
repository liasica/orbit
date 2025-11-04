// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import devops "github.com/alibabacloud-go/devops-20210625/v5/client"

// GetWorkitemInfoRequest 获取工作项的基本信息
// https://next.api.aliyun.com/api/devops/2021-06-25/GetWorkItemInfo
type GetWorkitemInfoRequest struct {
	workitemId              string
	parseWorkitemIdentifier bool
}

func NewGetWorkitemInfoRequest(workItemId string) *GetWorkitemInfoRequest {
	return &GetWorkitemInfoRequest{
		workitemId: workItemId,
	}
}

func (req *GetWorkitemInfoRequest) Do() (*devops.GetWorkItemInfoResponseBody, error) {
	res, err := instance.client.GetWorkItemInfo(&instance.config.OrganizationId, &req.workitemId)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
