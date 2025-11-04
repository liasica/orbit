// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-31, by liasica

package yunxiao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoListWorkitemWorkFlowStatusRequest(t *testing.T) {
	testSetup()

	req := NewListWorkitemWorkFlowStatusRequest(WorkitemCategoryBug)
	data, err := req.Do()
	require.NoError(t, err)
	require.NotNil(t, data)
	for _, s := range data.Statuses {
		fmt.Printf("[%s] %s: %s\n", *s.Source, *s.Name, *s.Identifier)
	}
}

func TestListAllWorkitemWorkFlowStatus(t *testing.T) {
	testSetup()

	items := ListAllWorkitemWorkFlowStatus()
	require.NotEmpty(t, items)
	for cate, statuses := range items {
		fmt.Println(cate)
		for _, s := range statuses {
			fmt.Printf("%s. [%s] %s: %s\n", *s.WorkflowStageIdentifier, *s.Source, *s.Name, *s.Identifier)
		}
		fmt.Println()
	}
}
