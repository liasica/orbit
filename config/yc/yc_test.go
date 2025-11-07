// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-07, by liasica

package yc

import "testing"

func TestGetWorkflowStatus(t *testing.T) {
	workflow := GetWorkflowStatus(WorkitemCategoryBug, WorkflowStatusOpen)
	t.Log(workflow)
}
