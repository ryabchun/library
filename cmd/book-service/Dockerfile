FROM golang:1.19 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o book_service ./cmd/book-service/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/book_service /app/
EXPOSE 8082
CMD ["./book_service"]
