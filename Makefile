SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH)

BINARY_BASENAME=qotm

.PHONY: all build build.image clean fmt run test.fast

all: clean fmt test.fast build

build: fmt
	$(GOBUILD) ./...
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

run: build
	bin/qotm

build.image:
	docker build \
	-t plombardi89/quote \
	-t plombardi89/quote:$(GIT_COMMIT) \
	-f Dockerfile \
	.

push.image:
	docker push plombardi89/quote
	docker push plombardi89/quote:$(GIT_COMMIT)

clean:
	rm -rf bin

fmt:
	go fmt ./...

test.fast:
	go test -v ./...
