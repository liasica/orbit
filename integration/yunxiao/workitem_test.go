// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/config/yc"
)

func TestListWorkitemTypes(t *testing.T) {
	testSetup()

	data, err := ListWorkitemTypes(yc.WorkitemCategoryBug)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestListAllWorkitemTypes(t *testing.T) {
	testSetup()

	data, err := ListAllWorkitemTypes()
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetWorkitemTypeFieldConfig(t *testing.T) {
	testSetup()

	data, err := GetWorkitemTypeFieldConfig("37da3a07df4d08aef2e3b393")
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetWorkitemWorkflow(t *testing.T) {
	testSetup()

	data, err := GetWorkitemWorkflow("37da3a07df4d08aef2e3b393")
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetWorkitem(t *testing.T) {
	testSetup()

	data, err := GetWorkitem("DAUR-316")
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestCreateWorkitemComment(t *testing.T) {
	testSetup()

	err := CreateWorkitemComment("DAUR-317", "")
	require.NoError(t, err)
}

func TestUpdateWorkitem(t *testing.T) {
	testSetup()

	err := UpdateWorkitem("DAUR-317", map[string]string{
		yc.FieldStatus: "100010",
	})
	require.NoError(t, err)
}
