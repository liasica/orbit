// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-12, by liasica

package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/liasica/orbit/ent"
	"github.com/liasica/orbit/ent/message"
)

func TestSaveMessage(t *testing.T) {
	testSetup()

	err := ent.Database.Message.Create().
		SetMessageID("test").
		SetWorkitemID("workitemId").
		SetVaraibales([]byte(`""`)).
		SetCreatedAt(time.Now()).
		OnConflictColumns(message.FieldMessageID).
		UpdateNewValues().
		Exec(context.Background())
	require.NoError(t, err)
}
