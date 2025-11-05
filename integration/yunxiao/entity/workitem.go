// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package entity

const (
	// WorkitemStatusKey 工作项状态字段键
	WorkitemStatusKey = "status"
)

// WorkitemCategory 工作项类别
type WorkitemCategory = string

const (
	WorkitemCategoryBug  WorkitemCategory = "Bug"  // 缺陷
	WorkitemCategoryTask WorkitemCategory = "Task" // 任务
)

var WorkitemCategories = []WorkitemCategory{
	WorkitemCategoryBug,
	WorkitemCategoryTask,
}

// Workitem 工作项实体
// https://help.aliyun.com/zh/yunxiao/developer-reference/getworkitem
type Workitem struct {
	AssignedTo        *IDName            `json:"assignedTo,omitempty"`
	CategoryID        WorkitemCategory   `json:"categoryId,omitempty"`
	Creator           *IDName            `json:"creator,omitempty"`
	CustomFieldValues []CustomFieldValue `json:"customFieldValues,omitempty"`
	Description       string             `json:"description,omitempty"`
	FormatType        string             `json:"formatType,omitempty"`
	GmtCreate         int64              `json:"gmtCreate,omitempty"`
	GmtModified       int64              `json:"gmtModified,omitempty"`
	ID                string             `json:"id,omitempty"`
	IDPath            string             `json:"idPath,omitempty"`
	Labels            []Label            `json:"labels,omitempty"`
	LogicalStatus     string             `json:"logicalStatus,omitempty"`
	Modifier          *IDName            `json:"modifier,omitempty"`
	ParentID          string             `json:"parentId,omitempty"`
	Participants      []IDName           `json:"participants,omitempty"`
	SerialNumber      string             `json:"serialNumber,omitempty"`
	Space             *IDName            `json:"space,omitempty"`
	Sprint            *IDName            `json:"sprint,omitempty"`
	Status            *Status            `json:"status,omitempty"`
	StatusStageID     string             `json:"statusStageId,omitempty"`
	Subject           string             `json:"subject,omitempty"`
	Trackers          []IDName           `json:"trackers,omitempty"`
	UpdateStatusAt    int64              `json:"updateStatusAt,omitempty"`
	Verifier          *IDName            `json:"verifier,omitempty"`
	Versions          []IDName           `json:"versions,omitempty"`
	WorkitemType      *IDName            `json:"workitemType,omitempty"`
}

// WorkitemType 工作项类型实体
// https://help.aliyun.com/zh/yunxiao/developer-reference/listworkitemtypes
type WorkitemType struct {
	Name          string           `json:"name,omitempty"`
	NameEn        string           `json:"nameEn,omitempty"`
	SystemDefault bool             `json:"systemDefault,omitempty"`
	GmtCreate     int64            `json:"gmtCreate,omitempty"`
	Creator       *IDName          `json:"creator,omitempty"`
	Description   string           `json:"description,omitempty"`
	Enable        bool             `json:"enable,omitempty"`
	DefaultType   bool             `json:"defaultType,omitempty"`
	GmtAdd        int64            `json:"gmtAdd,omitempty"`
	AddUser       *IDName          `json:"addUser,omitempty"`
	Id            string           `json:"id,omitempty"`
	CategoryId    WorkitemCategory `json:"categoryId,omitempty"`
}
