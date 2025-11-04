// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package pingcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestAuthToken(t *testing.T) {
	testSetup(t)

	res, err := RequestAuthToken()
	require.NoError(t, err)
	t.Logf("%+v", res)

	StoreAuthToken(res)
}

func TestGetAuthToken(t *testing.T) {
	testSetup(t)

	token := GetAuthToken()
	require.NotEmpty(t, token)
	t.Logf("token: %s", token)
}
