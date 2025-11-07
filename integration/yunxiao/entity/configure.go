// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package entity

// ConfigureWorkitemCustomField 工作项自定义属性配置
type ConfigureWorkitemCustomField = string

const (
	ConfigureWorkitemFieldRepository ConfigureWorkitemCustomField = "repository"
)

type ConfigureWorkflowStatus string

const (
	ConfigureWorkflowStatusOpen        ConfigureWorkflowStatus = "open"        // 打开
	ConfigureWorkflowStatusReopen      ConfigureWorkflowStatus = "reopen"      // 重新打开
	ConfigureWorkflowStatusInProgress  ConfigureWorkflowStatus = "inProgress"  // 处理中
	ConfigureWorkflowStatusUnderReview ConfigureWorkflowStatus = "underReview" // 待审查
	ConfigureWorkflowStatusResolved    ConfigureWorkflowStatus = "resolved"    // 已解决
	ConfigureWorkflowStatusClosed      ConfigureWorkflowStatus = "closed"      // 关闭
)

func (s ConfigureWorkflowStatus) Text() string {
	switch s {
	case ConfigureWorkflowStatusOpen:
		return "打开"
	case ConfigureWorkflowStatusReopen:
		return "重新打开"
	case ConfigureWorkflowStatusInProgress:
		return "处理中"
	case ConfigureWorkflowStatusUnderReview:
		return "待审查"
	case ConfigureWorkflowStatusResolved:
		return "已解决"
	case ConfigureWorkflowStatusClosed:
		return "关闭"
	default:
		return "未知"
	}
}

// ConfigureMap 配置集合
type ConfigureMap map[WorkitemCategory]*Configure

type Configure struct {
	Workitem *WorkitemConfigure `yaml:"workitem,omitempty" json:"workitem,omitempty"`
}

type WorkitemConfigure struct {
	Category         WorkitemCategory                                            `yaml:"category,omitempty" json:"category,omitempty"`
	TypeId           string                                                      `yaml:"typeId,omitempty" json:"typeId,omitempty"`
	Fields           map[ConfigureWorkitemCustomField]WorkitemFieldConfigure     `yaml:"fields,omitempty" json:"fields,omitempty"`
	WorkflowStatuses map[ConfigureWorkflowStatus]WorkitemWorkflowStatusConfigure `yaml:"workflowStatuses,omitempty" json:"workflowStatuses,omitempty"`
}

func (c *WorkitemConfigure) GetWorkflowStatus(s ConfigureWorkflowStatus) WorkitemWorkflowStatusConfigure {
	if c == nil {
		return WorkitemWorkflowStatusConfigure{}
	}

	result, ok := c.WorkflowStatuses[s]
	if !ok {
		return WorkitemWorkflowStatusConfigure{}
	}
	return result
}

type WorkitemWorkflowStatusConfigure struct {
	Id   string `yaml:"id,omitempty" json:"id,omitempty"`
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
}

type WorkitemFieldConfigure struct {
	Id          string `yaml:"id,omitempty" json:"id,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}
