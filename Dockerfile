# ==== Сборка ====
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go install github.com/a-h/templ/cmd/templ@latest
# RUN templ generate

RUN go build -o server cmd/server/main.go


# ==== Запуск ====
FROM ubuntu:24.04

RUN apt-get update && apt-get install -y \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server .
# COPY --from=builder /app/templates ./templates
# COPY --from=builder /app/static ./static

EXPOSE ${SERVER_PORT}