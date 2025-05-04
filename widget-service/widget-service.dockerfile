FROM golang:1.24.2-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o widget-service ./cmd/api

EXPOSE 80 81

CMD ["./widget-service"]
