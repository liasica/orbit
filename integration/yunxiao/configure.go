// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"github.com/liasica/orbit/config"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

// ConfigureParser 配置解析器
type ConfigureParser func() (entity.ConfigureMap, error)

// GetConfigure 获取工作项配置
func GetConfigure() (configure entity.ConfigureMap, err error) {
	configure = make(entity.ConfigureMap)

	// 获取所有工作项类型
	for category, typeId := range config.Get().Yunxiao.WorkitemTypes {
		configure[category] = &entity.Configure{
			Workitem: entity.WorkitemConfigure{
				Category:         category,
				TypeId:           typeId,
				Fields:           make(map[entity.ConfigureWorkitemCustomField]entity.WorkitemFieldConfigure),
				WorkflowStatuses: make(map[entity.ConfigureWorkflowStatus]entity.WorkitemWorkflowStatusConfigure),
			},
		}

		// 获取字段配置
		var fields []entity.Field
		fields, err = GetWorkitemTypeFieldConfig(typeId)
		if err != nil {
			return
		}

		for _, field := range fields {
			for k, v := range config.Get().Yunxiao.WorkitemFields {
				if v == field.Name {
					configure[category].Workitem.Fields[k] = entity.WorkitemFieldConfigure{
						Id:          field.Id,
						Description: field.Description,
					}
				}
			}
		}

		// 获取工作流配置
		var workflow entity.Workflow
		workflow, err = GetWorkitemWorkflow(typeId)
		if err != nil {
			return
		}

		for _, status := range workflow.Statuses {
			for k, v := range config.Get().Yunxiao.WorkflowNames {
				if v == status.Name {
					configure[category].Workitem.WorkflowStatuses[k] = entity.WorkitemWorkflowStatusConfigure{
						Id:   status.Id,
						Name: status.Name,
					}
				}
			}
		}
	}

	return
}
