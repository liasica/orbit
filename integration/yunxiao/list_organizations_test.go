// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoListOrganizationsRequest(t *testing.T) {
	testSetup()

	req := NewListOrganizationsRequest()
	res, err := req.Do()
	require.NoError(t, err)
	require.NotNil(t, res)
}
