// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package feishu

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/config"
)

func TestFindUserByDepartment(t *testing.T) {
	testSetup()

	data, err := FindUserByDepartment(context.Background(), config.Get().Feishu.DepartmentId, FindUserByDepartmentWithPageSize(5))
	require.NoError(t, err)
	require.NotNil(t, data)
	t.Logf("Department Users: %+v", data)
}
