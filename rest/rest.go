// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-30, by liasica

package rest

import (
	"fmt"
	"io"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Run(addr string) {
	e := echo.New()

	e.Any("/", func(c echo.Context) error {
		fmt.Println("==========================================")
		fmt.Printf("[URL]: %s\n\n", c.Request().RequestURI)

		fmt.Println("[Headers]:")
		for k, v := range c.Request().Header {
			fmt.Printf("\t%s: %v\n", k, strings.Join(v, ","))
		}
		fmt.Println()

		fmt.Println("[Body]:")
		body, _ := io.ReadAll(c.Request().Body)
		v := make(map[string]any)
		_ = sonic.Unmarshal(body, &v)
		b, _ := sonic.MarshalIndent(v, "", "  ")
		fmt.Printf("%s\n", string(b))
		fmt.Println()
		return nil
	})
	log.Fatal().Err(e.Start(addr))
}
