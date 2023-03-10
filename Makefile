.PHONY: all clean

SHELL := /bin/bash
GO_DIR := $(shell git rev-parse --show-toplevel)

GO_SOURCE = \
	$(GO_DIR)/go.mod $(GO_DIR)/go.sum \
	$(find $(GO_DIR) -type f) \
	$(NULL)

VERSION := $(shell git describe --match 'v[0-9]*' --always --tags --abbrev=8)

all: ./build/bin/gpt-zmide-server

./build/bin/gpt-zmide-server: $(GO_SOURCE)
	cd $(GO_DIR); CGO_ENABLED=0 go build -o ./build/bin/gpt-zmide-server .

clean:
	rm -rf ./build/bin/gpt-zmide-server