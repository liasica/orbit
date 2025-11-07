// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package yunxiao

import "github.com/liasica/orbit/integration/yunxiao/entity"

// UpdateCustomField 更新自定义字段
func UpdateCustomField(fieldId string, body *entity.CustomField) (err error) {
	_, err = instance.client.R().
		SetPathParam("organizationId", instance.organizationId).
		SetPathParam("id", fieldId).
		SetBody(body).
		Put("/oapi/v1/projex/organizations/{organizationId}/customField/{id}")
	return
}
