// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package feishu

import (
	"github.com/bytedance/sonic"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type Serialization struct{}

func NewSerialization() *Serialization {
	return &Serialization{}
}

func (s *Serialization) Serialize(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func (s *Serialization) Deserialize(data []byte, v interface{}) error {
	e, _ := sonic.Get(data, "error")
	if e.Exists() {
		var codeErr larkcore.CodeError
		if err := sonic.Unmarshal(data, &codeErr); err != nil {
			return err
		}
		return &codeErr
	}
	return sonic.Unmarshal(data, v)
}
