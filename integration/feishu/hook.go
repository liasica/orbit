// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-10, by liasica

package feishu

import (
	"context"

	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"
)

type HookCardActionTrigger func(ctx context.Context, event *callback.CardActionTriggerEvent) (*callback.CardActionTriggerResponse, error)
