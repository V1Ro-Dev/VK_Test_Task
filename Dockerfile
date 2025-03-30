# Используем официальный образ Go как базовый
FROM golang:1.24 AS builder

WORKDIR /app

COPY poll_bot/go.mod poll_bot/go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/poll_bot

RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

RUN GOOS=linux GOARCH=amd64 go build -o /app/app .

FROM alpine:latest

RUN apk add --no-cache tzdata libc6-compat

# Создаем рабочую директорию
WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
