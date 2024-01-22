.PHONY: build
include .env

build:
	@go build -o ./bin/api ./cmd/api/main.go ./cmd/api/wire_gen.go

run:
	@air -c .air.toml

test:
	./scripts/test.sh

install-deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/google/wire/cmd/wire@latest

update-di:
	@wire ./cmd/api

setup: install-deps
