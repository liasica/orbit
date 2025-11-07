// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package service

import (
	"regexp"
	"strings"
	"testing"

	"github.com/liasica/orbit/config"
)

const (
	testBranch1 = "dev/DAUR-317"
	testBranch2 = "feature/DAUR-317"
)

var testRegexpPattern = regexp.MustCompile(`(?m)dev/(.*)`)

func testSplit(str string) {
	if !strings.HasPrefix(str, config.GitlabBranchPrefix) {
		return
	}

	_ = strings.Split(str, "/")
}

func testSlice(str string) {
	if !strings.HasPrefix(str, config.GitlabBranchPrefix) {
		return
	}

	_ = str[len(config.GitlabBranchPrefix):]
}

func testRegexp(str string) {
	matches := testRegexpPattern.FindStringSubmatch(str)
	if len(matches) > 1 {
		_ = matches[1]
	}
}

func TestRegexp(t *testing.T) {
	testRegexp(testBranch1)
	testRegexp(testBranch2)
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testSplit(testBranch1)
		testSplit(testBranch2)
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testSlice(testBranch1)
		testSlice(testBranch2)
	}
}

func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testRegexp(testBranch1)
		testRegexp(testBranch2)
	}
}
