// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"fmt"

	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
)

type CommentFormatType string

const (
	CommentFormatTypeMarkdown CommentFormatType = "MARKDOWN"
	CommentFormatTypeRichtext CommentFormatType = "RICHTEXT"
)

func (ft CommentFormatType) Pointer() *string {
	return tea.String(string(ft))
}

// CreateWorkitemCommentRequest 创建工作项评论
// https://next.api.aliyun.com/api/devops/2021-06-25/CreateWorkitemComment
type CreateWorkitemCommentRequest struct {
	identifier string
	params     devops.CreateWorkitemCommentRequest
}

type CreateWorkitemCommentOption func(*CreateWorkitemCommentRequest)

func CreateWorkitemCommentWithFormatType(formatType CommentFormatType) CreateWorkitemCommentOption {
	return func(req *CreateWorkitemCommentRequest) {
		req.params.FormatType = formatType.Pointer()
	}
}

var _ = CreateWorkitemCommentWithParentId

func CreateWorkitemCommentWithParentId(parentId string) CreateWorkitemCommentOption {
	return func(req *CreateWorkitemCommentRequest) {
		req.params.ParentId = &parentId
	}
}

func NewCreateWorkitemCommentRequest(identifier, content string, options ...CreateWorkitemCommentOption) (req *CreateWorkitemCommentRequest) {
	req = &CreateWorkitemCommentRequest{
		identifier: identifier,
		params: devops.CreateWorkitemCommentRequest{
			Content:    &content,
			FormatType: CommentFormatTypeMarkdown.Pointer(),
		},
	}

	for _, opt := range options {
		opt(req)
	}
	return
}

func (req *CreateWorkitemCommentRequest) Do() (*devops.CreateWorkitemCommentResponseBody, error) {
	// 判断 workitemIdentifier 是不是 workItemId
	// 阿里云文档有错误, 文档注明 workItemId 和 workitemIdentifier 均可以, 但是实际调用时只能用 workitemIdentifier
	// 因此此处如果是 workItemId, 则需要转换为 workitemIdentifier
	if IsWorkitemIdentifier(req.identifier) {
		req.params.WorkitemIdentifier = &req.identifier
	} else {
		data, err := GetWorkitem(req.identifier)
		if err != nil {
			return nil, fmt.Errorf("转换 workitemIdentifier 失败: %w", err)
		}
		req.params.WorkitemIdentifier = &data.ID
	}

	res, err := instance.sdk.CreateWorkitemComment(&instance.organizationId, &req.params)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
