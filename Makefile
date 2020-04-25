VERSION = 0.1.0
TARGET_NAME = stg.exe
TARGET_OS=windows
TARGET_ARCH=amd64

PROJECT_DIR := $(CURDIR)
TMP_DIR := $(PROJECT_DIR)/tmp
BIN_DIR := $(TMP_DIR)/bin

GOIMPORTS := $(BIN_DIR)/goimports
STATICCHECK := $(BIN_DIR)/staticcheck

.PHONY: all
all: build

## Basic

.PHONY: setup
setup: ## Setup necessary tools.
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: clean
clean: ## Clean files.
	rm -rf $(TMP_DIR)/*
	-rmdir $(TMP_DIR)

## Build

.PHONY: build
build:
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go build -trimpath -ldflags "-s -w -X main.version=$(VERSION)" -o $(BIN_DIR)/$(TARGET_NAME) ./cmd/game

## Test

.PHONY: check-generate
check-generate: ## Generate code, and check if diff exists.
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go mod tidy
	git diff --exit-code --name-only

.PHONY: lint
lint:
	test -z "$$($(GOIMPORTS) -l $$(find . -name '*.pb.go' -prune -o -name '*.go' -print) | tee /dev/stderr)"
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) $(STATICCHECK) ./...
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go vet ./...

.PHONY: test
test:
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go test -v ./...

.PHONY: run
run:
	$(BIN_DIR)/$(TARGET_NAME)