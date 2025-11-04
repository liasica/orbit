// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	"regexp"

	devops "github.com/alibabacloud-go/devops-20210625/v5/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/rs/zerolog/log"
)

// WorkitemCategory 工作项类别
type WorkitemCategory string

const (
	WorkitemCategoryReq  WorkitemCategory = "Req"  // 需求
	WorkitemCategoryBug  WorkitemCategory = "Bug"  // 缺陷
	WorkitemCategoryTask WorkitemCategory = "Task" // 任务
	// WorkitemCategoryRisk WorkitemCategory = "Risk" // 风险
)

var WorkitemCategories = []WorkitemCategory{
	WorkitemCategoryReq,
	WorkitemCategoryBug,
	WorkitemCategoryTask,
	// WorkitemCategoryRisk,
}

type WorkflowStatus struct {
	Identifier string `yaml:"identifier,omitempty" json:"identifier,omitempty"`
	Name       string `yaml:"name,omitempty" json:"name,omitempty"`
}

type WorkflowStatusList struct {
	Reopen      *WorkflowStatus `yaml:"reopen,omitempty" json:"reopen,omitempty"`
	Open        *WorkflowStatus `yaml:"open,omitempty" json:"open,omitempty"`
	InProgress  *WorkflowStatus `yaml:"inProgress,omitempty" json:"inProgress,omitempty"`
	UnderReview *WorkflowStatus `yaml:"underReview,omitempty" json:"underReview,omitempty"`
	Resolved    *WorkflowStatus `yaml:"resolved,omitempty" json:"resolved,omitempty"`
	Closed      *WorkflowStatus `yaml:"closed,omitempty" json:"closed,omitempty"`
}

type Field struct {
	Identifier  string `yaml:"identifier,omitempty" json:"identifier,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type WorkitemFields struct {
	Status *Field `yaml:"status,omitempty"`
}

type WorkitemCategoryConfigure struct {
	WorkitemTypeIdentifier string              `yaml:"workitemTypeIdentifier,omitempty" json:"workitemTypeIdentifier,omitempty"`
	Fields                 *WorkitemFields     `yaml:"fields,omitempty" json:"fields,omitempty"`
	WorkflowStatus         *WorkflowStatusList `yaml:"workflowStatus,omitempty" json:"workflowStatus,omitempty"`
}

type WorkitemConfigure struct {
	Bug  *WorkitemCategoryConfigure `yaml:"bug,omitempty" json:"bug,omitempty"`
	Task *WorkitemCategoryConfigure `yaml:"task,omitempty" json:"task,omitempty"`
}

type WorkitemConfigureParser func() (*WorkitemConfigure, error)

// GetWorkitemConfigure 获取工作项配置
func GetWorkitemConfigure() (configure *WorkitemConfigure) {
	configure = &WorkitemConfigure{}

	data := make(map[WorkitemCategory]*WorkitemCategoryConfigure)

	// 获取 workitemTypeIdentifier
	for category, workitemTypes := range ListAllWorkitemTypes() {
		if len(workitemTypes) == 0 {
			log.Warn().Msg("未获取到工作项类型: " + string(category))
			continue
		}

		if workitemTypes[0].Identifier == nil {
			log.Warn().Msg("工作项类型 Identifier 为空: " + string(category))
			continue
		}

		identifier := *workitemTypes[0].Identifier
		data[category] = &WorkitemCategoryConfigure{
			WorkitemTypeIdentifier: identifier,
		}

		// 获取字段 identifier
		fieldBody, err := NewListWorkitemAllFieldsRequest(identifier).Do()
		if err != nil {
			log.Warn().Err(err).Msg("获取工作项字段失败: " + identifier)
			continue
		}

		fields := &WorkitemFields{}
		for _, field := range fieldBody.Fields {
			if field.Identifier == nil {
				log.Warn().Msg("工作项字段 Identifier 为空: " + identifier)
				continue
			}
			switch *field.Identifier {
			case "status":
				fields.Status = &Field{
					Identifier:  *field.Identifier,
					Description: tea.StringValue(field.Description),
				}
			}
		}

		// 获取工作流状态
		var statusBody *devops.ListWorkItemWorkFlowStatusResponseBody
		statusBody, err = NewListWorkitemWorkFlowStatusRequest(category).Do()
		if err != nil {
			log.Warn().Err(err).Msg("获取工作流状态失败: " + identifier)
			continue
		}

		workflowStatus := &WorkflowStatusList{}
		for _, status := range statusBody.Statuses {
			if status.Identifier == nil {
				log.Warn().Msg("工作流状态 Identifier 为空: " + identifier)
				continue
			}
			if status.Name == nil {
				log.Warn().Msg("工作流状态 Name 为空: " + identifier)
				continue
			}
			switch *status.Name {
			case "重新打开":
				workflowStatus.Reopen = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			case "打开":
				workflowStatus.Open = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			case "处理中":
				workflowStatus.InProgress = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			case "待审查":
				workflowStatus.UnderReview = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			case "已解决":
				workflowStatus.Resolved = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			case "已关闭":
				workflowStatus.Closed = &WorkflowStatus{
					Identifier: tea.StringValue(status.Identifier),
					Name:       tea.StringValue(status.Name),
				}
			}
		}

		data[category].Fields = fields
		data[category].WorkflowStatus = workflowStatus
	}

	if d, ok := data[WorkitemCategoryBug]; ok {
		configure.Bug = d
	}

	if d, ok := data[WorkitemCategoryTask]; ok {
		configure.Task = d
	}

	return
}

// IsWorkitemIdentifier 判断字符串是否是 workitemIdentifier
// 如果字符串是 hexadecimal string, 则认为是 workitemIdentifier, 否则调用 GetWorkitemInfo 接口获取 workitemIdentifier
func IsWorkitemIdentifier(str string) bool {
	// 允许大写和小写，且长度为偶数
	matched, _ := regexp.MatchString("^[0-9a-fA-F]+$", str)
	return matched && len(str)%2 == 0
}
