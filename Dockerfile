# Используем официальный образ Go как базовый
FROM golang:1.24 AS builder

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для скачивания зависимостей
COPY poll_bot/go.mod poll_bot/go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Переходим в папку с исходным кодом для сборки
WORKDIR /app/poll_bot

# Устанавливаем необходимые библиотеки для работы с OpenSSL в Debian
RUN apt-get update && apt-get install -y \
    build-essential \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Собираем бинарный файл (включаем CGO)
RUN GOOS=linux GOARCH=amd64 go build -o /app/app .

# Создаем конечный образ на базе Alpine Linux для минимизации размера
FROM alpine:latest

# Установка необходимых пакетов (например, tzdata для корректной работы времени)
RUN apk add --no-cache tzdata libc6-compat

# Создаем рабочую директорию
WORKDIR /app

# Копируем собранный бинарный файл из стадии сборки
COPY --from=builder /app/app .

# Указываем команду запуска
CMD ["./app"]

# Определяем порт, который будет открыт
EXPOSE 8080
