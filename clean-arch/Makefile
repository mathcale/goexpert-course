.PHONY: build
include .env

latform := $(shell uname -s | tr '[:upper:]' '[:lower:]')

build:
	go build -o ./bin/ordersystem ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

run:
	air -c .air.toml

test:
	./scripts/test.sh

migrate-up:
	migrate -path ./migrations -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?query" -verbose up

migrate-down:
	migrate -path ./migrations -database "mysql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?query" -verbose down

install-deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/google/wire/cmd/wire@latest
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.$(platform)-amd64.tar.gz | tar xvz migrate && sudo mv migrate /usr/local/bin

update-di:
	wire ./cmd/api

setup: install-deps
