// Copyright (C) moneta. 2025-present.
//
// Created at 2025-08-17, by liasica

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMyIP(t *testing.T) {
	ip, err := getMyIPQQ()
	require.NoError(t, err)
	t.Logf("QQ IP: %s", ip)

	ip, err = getMyIP189()
	require.NoError(t, err)
	t.Logf("189 IP: %s", ip)

	ip, err = getMyIP12306()
	require.NoError(t, err)
	t.Logf("12306 IP: %s", ip)

	ip, err = getMyIPiQiyi()
	require.NoError(t, err)
	t.Logf("iQiyi IP: %s", ip)
}
