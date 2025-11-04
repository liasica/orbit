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

func TestDoGetWorkitemInfoRequest(t *testing.T) {
	testSetup()

	req := NewGetWorkitemInfoRequest("9d49be78377ff4a1e2573ce69b")
	data, err := req.Do()
	require.NoError(t, err)
	require.NotNil(t, data)
	b, _ := sonic.MarshalIndent(data.Workitem, "", "  ")
	fmt.Println(string(b))
}
