// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"fmt"
	"os"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"

	"github.com/liasica/orbit/integration/yunxiao/entity"
)

func testConfigureYamlParser() (c entity.ConfigureMap, err error) {
	c = make(entity.ConfigureMap)
	b, _ := os.ReadFile("../../configs/yunxiao_workitem_configure.yaml")
	_ = yaml.Unmarshal(b, c)
	return c, nil
}

func testConfigureJsonParser() (c entity.ConfigureMap, err error) {
	c = make(entity.ConfigureMap)
	b, _ := os.ReadFile("../../configs/yunxiao_workitem_configure.json")
	_ = sonic.Unmarshal(b, c)
	return c, nil
}

func BenchmarkConfigureGetterYaml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = testConfigureYamlParser()
	}
}

func BenchmarkConfigureJsonParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = testConfigureJsonParser()
	}
}

func TestTestConfigureYamlParser(t *testing.T) {
	c, err := testConfigureYamlParser()
	require.NoError(t, err)
	fmt.Printf("%#v\n", c)
}

func TestTestConfigureJsonParser(t *testing.T) {
	c, err := testConfigureJsonParser()
	require.NoError(t, err)
	fmt.Printf("%#v\n", c)
}

func TestGetConfigure(t *testing.T) {
	testSetup()

	c, err := GetConfigure()
	require.NoError(t, err)
	require.NotNil(t, c)

	b, _ := sonic.MarshalIndent(c, "", "  ")
	fmt.Println(string(b))
}
