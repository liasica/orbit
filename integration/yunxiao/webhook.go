// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-06, by liasica

package yunxiao

const (
	HookHeaderAction     = "X-Action"
	HookHeaderSecret     = "X-Projex-Signature"
	HookHeaderStatusFrom = "X-Status-From"
	HookHeaderStatusTo   = "X-Status-To"
)

const (
	HookActionWorkitemCreated       = "workitem_created"
	HookActionWorkitemStatusChanged = "workitem_status_changed"
	HookActionWorkitemUnderReview   = "workitem_under_review"
	HookActionWorkitemReviewed      = "workitem_reviewed"
)

type WebhookWorkitemIdentifier struct {
	Identifier string `json:"identifier,omitempty"`
}
