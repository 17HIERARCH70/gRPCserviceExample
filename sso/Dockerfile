# Используйте официальный образ Go как образ сборки.
FROM golang:1.22.1 as builder

# Установите рабочую директорию в контейнере.
WORKDIR /app

# Скопируйте модульные файлы Go и загрузите зависимости.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Скопируйте исходный код проекта.
COPY . .

# Соберите бинарный файл.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o sso ./cmd/sso

# Используйте образ Distroless для минимального рабочего окружения.
FROM gcr.io/distroless/base-debian10

WORKDIR /

# Скопируйте бинарный файл и конфигурационный файл в контейнер.
COPY --from=builder /app/sso .
COPY --from=builder /app/config/config_local.yaml /config/config.yaml

# Укажите переменную среды для пути к конфигурационному файлу.
ENV CONFIG_PATH=/config/config.yaml

# Укажите порт, на котором будет работать сервис.
EXPOSE 5145

# Запустите сервис SSO.
CMD ["/sso"]
