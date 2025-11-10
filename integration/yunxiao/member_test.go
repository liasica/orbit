// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListProjectMembers(t *testing.T) {
	testSetup()

	data, err := ListProjectMembers()
	require.NoError(t, err)
	require.NotNil(t, data)
	t.Logf("Project Members: %+v", data)
}
