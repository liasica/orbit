// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package pingcode

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"
)

var instance *PingCode

type Config struct {
	BaseUrl  string
	ClientID string
	Secret   string
}

type PingCode struct {
	config *Config

	client *resty.Client
}

type Logger struct {
}

func (l Logger) Errorf(format string, v ...any) {
	log.Error().Err(fmt.Errorf(format, v...))
}

func (l Logger) Warnf(format string, v ...any) {
	log.Warn().Err(fmt.Errorf(format, v...))
}

func (l Logger) Debugf(format string, v ...any) {
	log.Debug().Msgf(format, v...)
}

func Setup(cfg *Config) {
	instance = &PingCode{
		config: cfg,
		client: resty.New().
			SetBaseURL(cfg.BaseUrl).
			SetDebug(true).
			SetLogger(&Logger{}).
			SetResponseMiddlewares(
				func(c *resty.Client, res *resty.Response) error {
					// Response
					var resBody []byte
					if res.Body != nil { // Read
						resBody, _ = io.ReadAll(res.Body)
					}
					res.Body = io.NopCloser(bytes.NewBuffer(resBody)) // Reset

					// Request
					var reqBody []byte
					if res.Request.RawRequest != nil {
						reqBody, _ = sonic.Marshal(res.Request.Body)
					}

					if rawBody, _ := res.Request.RawRequest.GetBody(); rawBody != nil {
						b, _ := io.ReadAll(rawBody)
						if b != nil {
							reqBody = b
						}
					}

					log.Info().
						Str("url", res.Request.URL).
						Bytes("request", reqBody).
						Bytes("response", resBody).
						Msg("[S]")
					return nil
				},
				resty.AutoParseResponseMiddleware,
			),
	}
}
