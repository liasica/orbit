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

func TestDoListProjectWorkitemTypesRequest(t *testing.T) {
	testSetup()
	
	req := NewListProjectWorkitemTypesRequest(WorkitemCategoryTask)
	data, err := req.Do()
	require.NoError(t, err)
	require.NotNil(t, data)

	b, _ := sonic.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}
