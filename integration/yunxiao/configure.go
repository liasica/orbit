// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"github.com/liasica/orbit/config/yc"
	"github.com/liasica/orbit/integration/yunxiao/entity"
)

// GetConfigure 获取工作项配置
func GetConfigure() (m map[yc.WorkitemCategory]yc.Workitem, err error) {
	m = yc.Get()

	// 获取所有工作项类型
	for category, wc := range m {
		typeId := wc.TypeId

		// 获取字段配置
		var fields []entity.Field
		fields, err = GetWorkitemTypeFieldConfig(typeId)
		if err != nil {
			return
		}

		for _, field := range fields {
			for k, v := range wc.Fields {
				if v.Name == field.Name {
					m[category].Fields[k] = yc.WorkitemField{
						Id:   field.Id,
						Name: field.Name,
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
			for k, v := range wc.Workflow {
				if v.Name == status.Name {
					v.Id = status.Id
					m[category].Workflow[k] = v
				}
			}
		}
	}

	return
}
