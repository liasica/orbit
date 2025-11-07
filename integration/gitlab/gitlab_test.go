// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package gitlab

import (
	"os"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
)

func testSetup() {
	// 设置全局时区
	tz := "Asia/Shanghai"
	_ = os.Setenv("TZ", tz)
	loc, _ := time.LoadLocation(tz)
	time.Local = loc

	// 日志初始化
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05.000",
	}).With().CallerWithSkipFrameCount(2).Logger()

	config.Setup("../../configs/config.yaml")

	// 初始化数据库
	ent.Setup(config.Get().Database.Postgres.Dsn, config.Get().Database.Postgres.Debug)

	Setup()
}

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
