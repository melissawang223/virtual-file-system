# Makefile
DC :=
GO       ?= go

MAIN_CMD = "main.go"

BINARY_NAME = "vfs"
TARGET_FILE = "./local.txt"

.PHONY: build clean tool lint help test test_coverage

all: build

build:
	@go build -o ${BINARY_NAME}

init:
	@echo "[]" > $(TARGET_FILE)

run: build
	@./${BINARY_NAME}

clean:
	@go clean
	@rm ${BINARY_NAME}

test:
	go test ./...


test_coverage:
	go test ./... -coverprofile=coverage.out


