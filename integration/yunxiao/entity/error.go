// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package entity

import (
	"strconv"
	"strings"
)

type ErrorResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorMsg     string `json:"errorMsg"`
	Success      bool   `json:"success"`
	TraceId      string `json:"traceId"`
}

func (r ErrorResponse) Error() string {
	builder := strings.Builder{}
	builder.WriteString("code: ")
	builder.WriteString(r.ErrorCode)
	builder.WriteString(", message: ")
	builder.WriteString(r.ErrorMessage)
	builder.WriteString(", errorMsg: ")
	builder.WriteString(r.ErrorMsg)
	builder.WriteString(", success: ")
	builder.WriteString(strconv.FormatBool(r.Success))
	builder.WriteString(", traceId: ")
	builder.WriteString(r.TraceId)
	return builder.String()
}
