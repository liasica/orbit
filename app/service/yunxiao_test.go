// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-11, by liasica

package service

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/integration/yunxiao"
)

func TestHookActionWorkitemCreated(t *testing.T) {
	testSetup()

	workitem, err := yunxiao.GetWorkitem("35a6ff43bea38b13e53e929b02")
	require.NoError(t, err)
	NewYunxiao().hookActionWorkitemCreated(nil, workitem)
}

func TestHookActionWorkitemUnderReview(t *testing.T) {
	testSetup()

	workitem, err := yunxiao.GetWorkitem("35a6ff43bea38b13e53e929b02")
	require.NoError(t, err)
	NewYunxiao().hookActionWorkitemUnderReview(nil, workitem)
}
