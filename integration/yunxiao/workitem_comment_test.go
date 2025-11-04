// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoCreateWorkitemCommentOption(t *testing.T) {
	testSetup()

	req := NewCreateWorkitemCommentRequest("6041e9a37b671d2724aec0ade2", "这是一个测试评论", CreateWorkitemCommentWithFormatType(CommentFormatTypeRichtext))
	data, err := req.Do()
	require.NoError(t, err)
	require.NotNil(t, data)

	req = NewCreateWorkitemCommentRequest("DAUR-316", "```go\nstr := \"test\"\n```", CreateWorkitemCommentWithFormatType(CommentFormatTypeMarkdown))
	data, err = req.Do()
	require.NoError(t, err)
	require.NotNil(t, data)
}
