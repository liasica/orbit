// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"errors"

	"github.com/bytedance/sonic"
	"resty.dev/v3"

	"github.com/liasica/orbit/integration/yunxiao/entity"
)

// Unmarshal 检查 json 是否包含 errorCode 字段, 若存在则返回错误, 否则正常反序列化到 v
func Unmarshal[T any](data []byte, v *T) error {
	if val, err := sonic.Get(data, "errorCode"); err == nil && val.Exists() {
		var errResp entity.ErrorResponse
		err = sonic.Unmarshal(data, &errResp)
		if err != nil {
			return err
		}
		return errResp
	}

	return sonic.Unmarshal(data, v)
}

type Request[T any] struct {
	*resty.Request
}

func NewRequest[T any]() *Request[T] {
	return &Request[T]{
		Request: instance.client.R(),
	}
}

// Do 发送请求并处理响应
func (r *Request[T]) Do() (data T, err error) {
	var res *resty.Response
	res, err = r.Send()

	if err != nil {
		return
	}

	if res == nil {
		err = errors.New("请求失败")
		return
	}

	b := res.Bytes()

	if res.IsError() {
		// 是否已经反序列化为 ErrorResponse
		switch v := res.Error().(type) {
		case *entity.ErrorResponse:
			err = v
		}
		if err != nil {
			return
		}

		// 如果响应包含 errorCode 字段, 则反序列化为 ErrorResponse
		if res.Err != nil {
			err = res.Err
			return
		}

		err = errors.New("请求失败，未知错误")
		return
	}

	err = sonic.Unmarshal(b, &data)
	return
}
