// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-05, by liasica

package entity

type Workflow struct {
	DefaultStatusId string           `json:"defaultStatusId,omitempty"`
	Id              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Statuses        []WorkflowStatus `json:"statuses,omitempty"`
}

type WorkflowStatus struct {
	DisplayName string `json:"displayName,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	NameEn      string `json:"nameEn,omitempty"`
}
