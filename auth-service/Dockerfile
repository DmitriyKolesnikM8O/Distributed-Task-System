FROM golang:latest

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

RUN go mod download

# Копируем остальные файлы микросервиса
COPY . .

RUN go build -o auth-service ./cmd/main/main.go

EXPOSE 8081

CMD ["./auth-service"]
