// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/config/yc"
)

func TestGetConfigure(t *testing.T) {
	testSetup()

	c, err := GetConfigure()
	require.NoError(t, err)
	require.NotNil(t, c)

	err = yc.Store(c)
	require.NoError(t, err)
}
