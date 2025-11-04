// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-03, by liasica

package yunxiao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkWorkitemConfigureGetterYaml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = testWorkitemConfigureYamlParser()
	}
}

func BenchmarkWorkitemConfigureJsonParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = testWorkitemConfigureJsonParser()
	}
}

func TestTestWorkitemConfigureYamlParser(t *testing.T) {
	c, err := testWorkitemConfigureYamlParser()
	require.NoError(t, err)
	fmt.Printf("%#v\n", c)
}

func TestTestWorkitemConfigureJsonParser(t *testing.T) {
	c, err := testWorkitemConfigureJsonParser()
	require.NoError(t, err)
	fmt.Printf("%#v\n", c)
}

func TestGetWorkitemConfigure(t *testing.T) {
	testSetup()

	c := GetWorkitemConfigure()
	require.NotNil(t, c)
}
