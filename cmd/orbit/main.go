// Copyright (C) orbit. 2025-present.
//
// Created at 2025-10-29, by liasica

package main

import (
	"github.com/liasica/orbit/boot"
	"github.com/liasica/orbit/cmd/orbit/internal/script"
)

func main() {
	boot.Bootstrap("configs/config.yaml")

	script.Run()
}
