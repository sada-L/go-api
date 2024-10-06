# Используем образ Golang
FROM golang:1.23-alpine

# Устанавливаем рабочий каталог внутри контейнера
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Копируем файлы проекта
COPY . .

# Собираем Go приложение
RUN go build -o /app/main

# Устанавливаем переменные среды для подключения к PostgreSQL
ENV DB_HOST=postgres \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_PASSWORD=postgres \
    DB_NAME=postgres

# Открываем порт 8080
EXPOSE 8080

# Запуск приложения
CMD ["/app/main"]
