// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package yunxiao

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/integration/yunxiao/entity"
)

func TestUpdateCustomField(t *testing.T) {
	testSetup()

	err := UpdateCustomField("1e216b5770aac61d8778651c86", &entity.CustomField{
		Name: "代码仓库",
		Options: []string{
			"auroraride/aurservd",
		},
	})
	require.NoError(t, err)
}
