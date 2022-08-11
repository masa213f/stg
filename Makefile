VERSION := devel
TARGET_EXE_NAME := stg.exe
TARGET_ZIP_NAME := stg-$(VERSION).zip
TARGET_OS := windows
TARGET_ARCH := amd64

BIN_DIR := $(CURDIR)/bin
OUT_DIR := $(CURDIR)/out

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
	rm -rf $(BIN_DIR)/* $(OUT_DIR)/*
	-rmdir $(BIN_DIR) $(OUT_DIR)

## Build

.PHONY: build
build:
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go build -trimpath -ldflags "-s -w -X main.version=$(VERSION)" -o $(OUT_DIR)/$(TARGET_EXE_NAME) ./cmd/game

.PHONY: zip
zip:
	zip -j $(OUT_DIR)/$(TARGET_ZIP_NAME) $(OUT_DIR)/$(TARGET_EXE_NAME)

## Test

.PHONY: check-generate
check-generate: ## Generate code, and check if diff exists.
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go mod tidy
	git diff --exit-code --name-only

.PHONY: lint
lint:
	test -z "$$($(GOIMPORTS) -l $$(find . -name '*.pb.go' -prune -o -name '*.go' -print) | tee /dev/stderr)"
	$(STATICCHECK) ./...
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	$(OUT_DIR)/$(TARGET_EXE_NAME) -debug
