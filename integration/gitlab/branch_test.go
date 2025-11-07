// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package gitlab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBranch(t *testing.T) {
	testSetup()

	b, err := GetBranch("auroraride/aurservd", "development")
	require.NoError(t, err)
	require.NotNil(t, b)
}

func TestCreateBranch(t *testing.T) {
	testSetup()

	b, err := CreateBranch("auroraride/aurservd", "test-branch-from-api", "development")
	require.NoError(t, err)
	require.NotNil(t, b)
}
