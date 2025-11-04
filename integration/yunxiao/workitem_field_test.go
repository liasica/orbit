// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
)

func TestDoListWorkitemAllFieldsRequest(t *testing.T) {
	testSetup()

	data, err := NewListWorkitemAllFieldsRequest("ba102e46bc6a8483d9b7f25c").Do()
	require.NoError(t, err)
	require.NotNil(t, data)

	b, _ := sonic.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
