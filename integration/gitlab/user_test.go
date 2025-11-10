// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package gitlab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	testSetup()

	users, err := ListUsers(nil)
	require.NoError(t, err)
	require.NotNil(t, users)
	t.Logf("Users count: %d", len(users))
}
