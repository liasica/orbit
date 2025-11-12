// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func (l Logger) Errorf(format string, v ...any) {
	log.Error().CallerSkipFrame(1).Msg(fmt.Sprintf(format, v...))
}

func (l Logger) Warnf(format string, v ...any) {
	log.Warn().CallerSkipFrame(1).Msg(fmt.Sprintf(format, v...))
}

func (l Logger) Debugf(format string, v ...any) {
	log.Debug().CallerSkipFrame(1).Msg(fmt.Sprintf(format, v...))
}
