// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-12, by liasica

package service

import "testing"

func TestStoreProjects(t *testing.T) {
	testSetup()
	
	NewGitlab().StoreProjects()
}
