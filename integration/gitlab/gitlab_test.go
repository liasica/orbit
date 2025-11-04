// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package gitlab

import (
	"os"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func TestWebhookPushEvent(t *testing.T) {
	b, _ := os.ReadFile("../../tests/push.json")
	var v gitlab.PushEvent
	err := sonic.Unmarshal(b, &v)
	require.NoError(t, err)
}

func TestWebhookMergeEvent(t *testing.T) {
	var v gitlab.MergeEvent
	b, _ := os.ReadFile("../../tests/pr.json")
	err := sonic.Unmarshal(b, &v)
	require.NoError(t, err)
}
