FROM golang:1.24.2-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o frontend ./cmd/web

EXPOSE 80

CMD ["./frontend"]
