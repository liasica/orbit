.PHONY: build

export GO111MODULE=on
export CGO_ENABLED=0

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))

BIN_OUTPUT := build/orbit
MAIN_DIRECTORY := ./cmd/orbit

build:
	@echo Version: $(VERSION)
	go build -trimpath -tags=sonic,poll_opt -gcflags "all=-N -l" -ldflags '-X "github.com/liasica/orbit/config.Version=${VERSION}"' -o ${BIN_OUTPUT} ${MAIN_DIRECTORY}
	pwd
	ls -lh ${BIN_OUTPUT}
