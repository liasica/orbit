// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-07, by liasica

package yc

import (
	"os"
	"sync/atomic"

	"github.com/rs/zerolog/log"
	"sigs.k8s.io/yaml"
)

// Field 工作项字段
type Field string

const (
	FieldRepository Field = "repository"
)

const (
	// WorkitemStatusKey 工作项状态字段键
	WorkitemStatusKey = "status"
)

// WorkitemCategory 工作项类别
type WorkitemCategory string

const (
	WorkitemCategoryBug  WorkitemCategory = "Bug"  // 缺陷
	WorkitemCategoryTask WorkitemCategory = "Task" // 任务
)

func (wc WorkitemCategory) Text() string {
	switch wc {
	case WorkitemCategoryBug:
		return "缺陷"
	case WorkitemCategoryTask:
		return "任务"
	default:
		return "未知"
	}
}

func (wc WorkitemCategory) String() string {
	return string(wc)
}

// WorkflowStatus 工作流状态
type WorkflowStatus string

const (
	WorkflowStatusOpen        WorkflowStatus = "open"        // 打开
	WorkflowStatusReopen      WorkflowStatus = "reopen"      // 重新打开
	WorkflowStatusInProgress  WorkflowStatus = "inProgress"  // 处理中
	WorkflowStatusUnderReview WorkflowStatus = "underReview" // 待审查
	WorkflowStatusResolved    WorkflowStatus = "resolved"    // 已解决
	WorkflowStatusClosed      WorkflowStatus = "closed"      // 关闭
)

func (s WorkflowStatus) Text() string {
	switch s {
	case WorkflowStatusOpen:
		return "打开"
	case WorkflowStatusReopen:
		return "重新打开"
	case WorkflowStatusInProgress:
		return "处理中"
	case WorkflowStatusUnderReview:
		return "待审查"
	case WorkflowStatusResolved:
		return "已解决"
	case WorkflowStatusClosed:
		return "关闭"
	default:
		return "未知"
	}
}

type Workitem struct {
	TypeId   string                      `json:"typeId"`
	Fields   map[Field]WorkitemField     `json:"fields"`
	Workflow map[WorkflowStatus]Workflow `json:"workflow"`
}

type WorkitemField struct {
	Id   string `json:"id"`   // 字段ID
	Name string `json:"name"` // 字段名称
}

type Workflow struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Configure struct {
	p string
	v atomic.Value
}

var instance *Configure

func Setup(cfgPath string) {
	log.Info().Msgf("从文件 %s 加载云效配置...", cfgPath)

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal().Err(err)
	}

	instance = &Configure{
		p: cfgPath,
		v: atomic.Value{},
	}

	m := make(map[WorkitemCategory]Workitem)

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatal().Err(err)
	}

	Set(m)

	log.Info().Msg("云效配置文件加载完成")
}

// Get 获取配置
func Get() map[WorkitemCategory]Workitem {
	if instance == nil {
		instance = &Configure{}
	}
	if instance.v.Load() == nil {
		instance.v.Store(make(map[WorkitemCategory]Workitem))
	}
	return instance.v.Load().(map[WorkitemCategory]Workitem)
}

// Set 设置配置
func Set(v map[WorkitemCategory]Workitem) {
	instance.v.Store(v)
}

// Store 存储配置到文件
func Store(v map[WorkitemCategory]Workitem) error {
	Set(v)

	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	return os.WriteFile(instance.p, data, 0644)
}

// GetWorkitem 获取指定工作项类别的工作项配置
func GetWorkitem(cate WorkitemCategory) Workitem {
	return Get()[cate]
}

// GetWorkflowStatus 获取指定工作项类别和状态的工作流配置
func GetWorkflowStatus(cate WorkitemCategory, s WorkflowStatus) Workflow {
	return GetWorkitem(cate).Workflow[s]
}
