FROM golang:1.25-alpine

WORKDIR /app
COPY . .

# Сборка Go-модуля
RUN go mod tidy
RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]
