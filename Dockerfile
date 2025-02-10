FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем весь проект и скрипт ожидания
COPY . .
COPY wait-for-it.sh /usr/local/bin/wait-for-it

# Переходим в директорию, где находится main.go
WORKDIR /app

# Собираем приложение
RUN go build -o /app/main .

# Указываем порт
EXPOSE 8080

# Запускаем приложение с ожиданием базы данных
CMD ["wait-for-it", "db:5432", "--", "/app/main"]
