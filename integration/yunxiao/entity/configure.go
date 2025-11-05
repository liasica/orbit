// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package entity

// ConfigureWorkitemCustomField 工作项自定义属性配置
type ConfigureWorkitemCustomField = string

const (
	ConfigureWorkitemFieldRepository ConfigureWorkitemCustomField = "repository"
)

type ConfigureWorkflowStatus = string

const (
	ConfigureWorkflowStatusOpen        ConfigureWorkflowStatus = "open"
	ConfigureWorkflowStatusReopen      ConfigureWorkflowStatus = "reopen"
	ConfigureWorkflowStatusInProgress  ConfigureWorkflowStatus = "inProgress"
	ConfigureWorkflowStatusUnderReview ConfigureWorkflowStatus = "underReview"
	ConfigureWorkflowStatusResolved    ConfigureWorkflowStatus = "resolved"
	ConfigureWorkflowStatusClosed      ConfigureWorkflowStatus = "closed"
)

// ConfigureMap 配置集合
type ConfigureMap map[WorkitemCategory]*Configure

type Configure struct {
	Workitem WorkitemConfigure `yaml:"workitem,omitempty" json:"workitem,omitempty"`
}

type WorkitemConfigure struct {
	Category         WorkitemCategory                                            `yaml:"category,omitempty" json:"category,omitempty"`
	TypeId           string                                                      `yaml:"typeId,omitempty" json:"typeId,omitempty"`
	Fields           map[ConfigureWorkitemCustomField]WorkitemFieldConfigure     `yaml:"fields,omitempty" json:"fields,omitempty"`
	WorkflowStatuses map[ConfigureWorkflowStatus]WorkitemWorkflowStatusConfigure `yaml:"workflowStatuses,omitempty" json:"workflowStatuses,omitempty"`
}

type WorkitemWorkflowStatusConfigure struct {
	Id   string `yaml:"id,omitempty" json:"id,omitempty"`
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

type WorkitemFieldConfigure struct {
	Id          string `yaml:"id,omitempty" json:"id,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}
