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
	path string

	Database *Database `json:"database,omitempty"`
	Gitlab   *Gitlab   `json:"gitlab,omitempty"`
	Yunxiao  *Yunxiao  `json:"yunxiao,omitempty"`
	Feishu   *Feishu   `json:"feishu,omitempty"`
}

// Database 数据库配置
type Database struct {
	Postgres struct {
		Dsn   string
		Debug bool
	}
}

type Gitlab struct {
	BaseUrl            string   `json:"baseUrl,omitempty"`
	Token              string   `json:"token,omitempty"`
	WebhookSecret      string   `json:"webhookSecret,omitempty"`
	MergeTargetBranchs []string `json:"mergeTargetBranchs,omitempty"`
}

type Yunxiao struct {
	ConfigPath string `json:"configPath,omitempty"` // 配置文件路径, 相对 config.yaml 目录
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

type Feishu struct {
	Debug        bool   `json:"debug,omitempty"`
	AppId        string `json:"appId,omitempty"`
	AppSecret    string `json:"appSecret,omitempty"`
	DepartmentId string `json:"departmentId,omitempty"`
	Icons        struct {
		Bug  string `json:"bug,omitempty"`
		Task string `json:"task,omitempty"`
	} `json:"icons,omitempty"`
	Message struct {
		ApkRelease  FeishuMessage `json:"apkRelease,omitempty"`  // APK测试发包消息模板
		UnderReview FeishuMessage `json:"underReview,omitempty"` // 待审查消息模板
		Reviewed    FeishuMessage `json:"reviewed,omitempty"`    // 已审查消息模板
		Job         FeishuMessage `json:"job,omitempty"`         // 新工作消息模板
	} `json:"message,omitempty"`
}

type FeishuMessage struct {
	TemplateId    string `json:"templateId,omitempty"`
	ReceiveIdType string `json:"receiveIdType,omitempty"`
	ReceiveId     string `json:"receiveId,omitempty"`
	MsgType       string `json:"msgType,omitempty"`
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

	config.path = cfgPath

	// 初始化云效配置
	ycfg := filepath.Join(filepath.Dir(cfgPath), config.Yunxiao.ConfigPath)
	yc.Setup(ycfg)

	log.Info().Msg("配置文件加载完成")
}

func Get() *Config {
	return config
}

var _ = GetPath

// GetPath 返回配置文件路径
func GetPath() string {
	return config.path
}

var _ = GetAbsolutePath

// GetAbsolutePath 返回相对于配置文件目录的绝对路径
func GetAbsolutePath(relPath string) (absPath string) {
	absPath, _ = filepath.Abs(filepath.Join(filepath.Dir(config.path), relPath))
	return
}
