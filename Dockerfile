# ==== Сборка ====
FROM golang:1.24.4 AS builder

WORKDIR /app

ENV NAME=server \
    CGO_ENABLED=0 \
    GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o ${NAME} cmd/main.go

# ==== Запуск ====
FROM ubuntu:24.04

WORKDIR /app

ENV NAME=server \
    TZ="Asia/Tashkent" \
    DEBIAN_FRONTEND=noninteractive

COPY --from=builder /app/${NAME} .

RUN  apt-get update -qq \
    && apt-get install -yq \
                        tzdata \
                        apparmor-utils \
                        ca-certificates \
                        curl && \
                         ln -fs /usr/share/zoneinfo/Asia/Tashkent /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*;

CMD ["sh", "-c", "./${NAME}"]