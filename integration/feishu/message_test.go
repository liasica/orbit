// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-09, by liasica

package feishu

import (
	"context"
	"testing"

	"github.com/google/uuid"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	v1 "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/config"
)

func TestSendMessage(t *testing.T) {
	testSetup()

	cfg := config.Get().Feishu.Message.ApkRelease

	req := v1.NewCreateMessageReqBuilder().
		ReceiveIdType(cfg.ReceiveIdType).
		Body(&v1.CreateMessageReqBody{
			ReceiveId: &cfg.ReceiveId,
			MsgType:   &cfg.MsgType,
			Content: NewInteractiveTemplateMessage[ApkMessage](cfg.TemplateId, &ApkMessage{
				ID:       "id",
				AppName:  "test",
				Message:  "test-message",
				Intranet: "RCIntranet",
				Extranet: "RCExtranet",
				Version:  "1.0.0",
			}).StringPtr(),
			Uuid: larkcore.StringPtr(uuid.New().String()),
		})

	resp, err := SendMessage(context.Background(), req.Build())
	require.NoError(t, err)
	t.Logf("%+v", resp)
}
