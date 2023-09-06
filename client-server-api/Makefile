.PHONY: build

build:
	go build -o ./bin/client ./cmd/client/client.go
	go build -o ./bin/server ./cmd/server/server.go

init-db:
	cat scripts/db.sql | sqlite3 rates.db

start-client:
	go run ./cmd/client/client.go

start-server:
	go run ./cmd/server/server.go
