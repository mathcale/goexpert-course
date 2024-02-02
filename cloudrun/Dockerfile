FROM golang:1.21.6-alpine3.19 AS builder

WORKDIR /app
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/api/main.go"]
