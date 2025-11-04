// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package request

import "resty.dev/v3"

type Request struct {
	client *resty.Client
}

func New(baseUrl string) *Request {
	return &Request{
		client: resty.New().SetBaseURL(baseUrl),
	}
}
