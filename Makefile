.PHONY: all clean

SHELL := /bin/bash
GO_DIR := $(shell git rev-parse --show-toplevel)
UI_DIR := $(shell git rev-parse --show-toplevel)

GO_SOURCE = \
	$(GO_DIR)/go.mod $(GO_DIR)/go.sum \
	$(shell find $(GO_DIR) -type f ! -path "*node_modules/*" ! -path "*/.git/*" ! -path "*/docs/*" ) \
	$(NULL)

UI_SOURCE = \
	$(UI_DIR)/package.json \
	$(shell find $(UI_DIR)/static/ -type f) \
	$(shell find $(UI_DIR)/src/ -type f) \
	$(shell find $(UI_DIR)/views/ -type f) \
	$(NULL)

VERSION := $(shell git describe --match 'v[0-9]*' --always --tags --abbrev=8)

all: ./build/bin/gpt-zmide-server 

# npm install
$(UI_DIR)/node_modules/.package-lock.json: $(UI_DIR)/package.json
	cd $(UI_DIR); npm install

./dist/: $(UI_DIR)/node_modules/.package-lock.json $(UI_SOURCE)
	npm run build

./build/bin/gpt-zmide-server: ./dist/ $(GO_SOURCE)
	cd $(GO_DIR); CGO_ENABLED=0 go build -ldflags="-s -w" -o ./build/bin/gpt-zmide-server .

clean:
	rm -rf ./build/bin/gpt-zmide-server ./dist/
	go clean -i .