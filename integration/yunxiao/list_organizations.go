// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import devops "github.com/alibabacloud-go/devops-20210625/v5/client"

type ListOrganizationsRequest struct {
	params devops.ListOrganizationsRequest
}

func NewListOrganizationsRequest() *ListOrganizationsRequest {
	return &ListOrganizationsRequest{
		params: devops.ListOrganizationsRequest{},
	}
}

func (req *ListOrganizationsRequest) Do() (*devops.ListOrganizationsResponseBody, error) {
	res, err := instance.client.ListOrganizations(&req.params)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
