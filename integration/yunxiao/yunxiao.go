// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica
//
// API参考: https://help.aliyun.com/zh/yunxiao/developer-reference/service-access-point-domain?spm=a2c4g.11186623.help-menu-150040.d_5_0_0.41ef37e8YYh4dH
// OpenAPI: https://next.api.aliyun.com/api/devops/2021-06-25/ListWorkitems
// SDK: https://next.api.aliyun.com/api-tools/sdk/devops?version=2021-06-25&language=go-tea&tab=primer-doc
// RAM: https://help.aliyun.com/zh/yunxiao/user-guide/add-a-ram-user

package yunxiao

import (
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	OrganizationId  string
	ProjectId       string
}

type Yunxiao struct {
	config *Config
	client *devops.Client
	wcp    WorkitemConfigureParser
}

var (
	instance *Yunxiao
	once     sync.Once
)

func Setup(cfg *Config, workitemConfigureParser WorkitemConfigureParser) {
	once.Do(func() {
		instance = &Yunxiao{
			config: cfg,
			wcp:    workitemConfigureParser,
		}

		config := &openapi.Config{
			AccessKeyId:     tea.String(cfg.AccessKeyId),
			AccessKeySecret: tea.String(cfg.AccessKeySecret),
			Endpoint:        tea.String(cfg.Endpoint),
		}

		var err error
		instance.client, err = devops.NewClient(config)
		if err != nil {
			log.Panic().Err(err)
		}
	})
}
