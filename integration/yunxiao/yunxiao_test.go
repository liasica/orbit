// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sigs.k8s.io/yaml"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/ent"
)

func testWorkitemConfigureYamlParser() (c *WorkitemConfigure, err error) {
	c = new(WorkitemConfigure)
	b, _ := os.ReadFile("../../configs/yunxiao_workitem_configure.yaml")
	_ = yaml.Unmarshal(b, c)
	return c, nil
}

func testWorkitemConfigureJsonParser() (c *WorkitemConfigure, err error) {
	c = new(WorkitemConfigure)
	b, _ := os.ReadFile("../../configs/yunxiao_workitem_configure.json")
	_ = sonic.Unmarshal(b, c)
	return c, nil
}

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

	cfg := config.Get().Yunxiao

	Setup(&Config{
		AccessKeyId:     cfg.AccessKeyId,
		AccessKeySecret: cfg.AccessKeySecret,
		Endpoint:        cfg.Endpoint,
		OrganizationId:  cfg.OrganizationId,
		ProjectId:       cfg.ProjectId,
	}, testWorkitemConfigureYamlParser)
}
