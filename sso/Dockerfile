# Используйте официальный образ Go как образ сборки.
FROM golang:1.22.1 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go модули и их зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sso_service ./sso/main

# Используем облегченный образ для запуска
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем собранный исполняемый файл из предыдущего шага
COPY --from=builder /app/sso_service .

# Порт, который будет прослушивать приложение
EXPOSE 5145

# Запускаем приложение
CMD ["./sso_service"]

