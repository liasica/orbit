// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"sigs.k8s.io/yaml"
)

var (
	config  = &Config{}
	Version = "v0.0.0-dev"
)

func GetVersion() string {
	return Version
}

type Config struct {
	Database *Database
	PingCode *PingCode
	Yunxiao  *Yunxiao
}

// Database 数据库配置
type Database struct {
	Postgres struct {
		Dsn   string
		Debug bool
	}
}

type PingCode struct {
	BaseUrl  string
	ClientID string
	Secret   string
}

type Yunxiao struct {
	Debug           bool
	AccessKeyId     string
	AccessKeySecret string
	OrganizationId  string
	ProjectId       string
	Domain          string
	Token           string
	WorkitemTypes   map[string]string // key: workitem category, value: typeId
	WorkitemFields  map[string]string // key: workitem field, value: name
	WorkflowNames   map[string]string // key: workflow status, value: name
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

	log.Info().Msg("配置文件加载完成")
}

func Get() *Config {
	return config
}
