// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package gitlab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProject(t *testing.T) {
	testSetup()

	p, err := GetProject("auroraride/aurservd")
	require.NoError(t, err)
	require.NotNil(t, p)
	t.Logf("default branch: %s", p.DefaultBranch)
}

func TestListProjects(t *testing.T) {
	testSetup()

	ps, err := ListProjects(nil)
	require.NoError(t, err)
	require.NotNil(t, ps)
	t.Logf("total projects: %d", len(ps))

	for _, p := range ps {
		t.Logf("project: %s | %s", p.Name, p.PathWithNamespace)
	}
}
