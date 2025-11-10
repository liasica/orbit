// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-08, by liasica

package feishu

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(_ context.Context, i ...interface{}) {
	log.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(i...))
}

func (l *Logger) Info(_ context.Context, i ...interface{}) {
	log.Info().CallerSkipFrame(1).Msg(fmt.Sprint(i...))
}

func (l *Logger) Warn(_ context.Context, i ...interface{}) {
	log.Warn().CallerSkipFrame(1).Msg(fmt.Sprint(i...))
}

func (l *Logger) Error(_ context.Context, i ...interface{}) {
	log.Error().CallerSkipFrame(1).Msg(fmt.Sprint(i...))
}
