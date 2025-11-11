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
	"fmt"
	"sync"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"

	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/integration/yunxiao/entity"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
)

type Yunxiao struct {
	accessKeyId     string
	accessKeySecret string

	organizationId string
	projectId      string

	webhookSecret string

	client *resty.Client
	sdk    *devops.Client
}

var (
	instance *Yunxiao
	once     sync.Once
)

func Setup() {
	once.Do(func() {
		cfg := config.Get().Yunxiao

		instance = &Yunxiao{
			accessKeyId:     cfg.AccessKeyId,
			accessKeySecret: cfg.AccessKeySecret,
			organizationId:  cfg.OrganizationId,
			projectId:       cfg.ProjectId,
			webhookSecret:   cfg.Webhook.Secret,
			client: resty.New().
				SetDebug(cfg.Debug).
				SetLogger(&Logger{}).
				SetError(&entity.ErrorResponse{}).
				SetBaseURL(fmt.Sprintf("https://%s", cfg.Domain)).
				SetHeader("x-yunxiao-token", cfg.Token).
				SetHeader("Content-Type", "application/json").
				AddResponseMiddleware(func(_ *resty.Client, res *resty.Response) error {
					if res.IsError() {
						switch v := res.Error().(type) {
						case error:
							return v
						default:
							return fmt.Errorf("请求失败: %v", v)
						}
					}
					return nil
				}),
		}

		var err error
		instance.sdk, err = devops.NewClient(&openapi.Config{
			AccessKeyId:     tea.String(cfg.AccessKeyId),
			AccessKeySecret: tea.String(cfg.AccessKeySecret),
			Endpoint:        tea.String("devops.cn-hangzhou.aliyuncs.com"),
		})
		if err != nil {
			log.Panic().Err(err)
		}
	})
}

// VerifyWebhookSecret 校验 GitLab Webhook Secret
func VerifyWebhookSecret(secret string) bool {
	return instance.webhookSecret == secret
}
