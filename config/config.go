// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package config

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"sigs.k8s.io/yaml"

	"github.com/liasica/orbit/config/yc"
)

const (
	GitlabBranchPrefix = "dev/"
)

var (
	config  = &Config{}
	Version = "v0.0.0-dev"
)

func GetVersion() string {
	return Version
}

type Config struct {
	Gitlab  *Gitlab  `json:"gitlab,omitempty"`
	Yunxiao *Yunxiao `json:"yunxiao,omitempty"`
}

type Gitlab struct {
	BaseUrl            string   `json:"baseUrl,omitempty"`
	Token              string   `json:"token,omitempty"`
	WebhookSecret      string   `json:"webhookSecret,omitempty"`
	MergeTargetBranchs []string `json:"mergeTargetBranchs,omitempty"`
}

type Yunxiao struct {
	ConfigPath string `json:"configPath,omitempty"` // 配置文件路径, 相对目录
	Debug      bool   `json:"debug,omitempty"`
	Webhook    struct {
		Secret string `json:"secret,omitempty"`
	} `json:"webhook,omitempty"`
	AccessKeyId     string `json:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty"`
	OrganizationId  string `json:"organizationId,omitempty"`
	ProjectId       string `json:"projectId,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Token           string `json:"token,omitempty"`
}

// Setup 读取并解析配置文件
func Setup(cfgPath string) {
	log.Info().Msg("加载配置文件: " + cfgPath)

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal().Err(err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatal().Err(err)
	}

	ycfg := filepath.Join(filepath.Dir(cfgPath), config.Yunxiao.ConfigPath)
	yc.Setup(ycfg)

	log.Info().Msg("配置文件加载完成")
}

func Get() *Config {
	return config
}
