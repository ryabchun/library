FROM golang:1.19 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o user_service ./cmd/user-service/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/user_service /app/
EXPOSE 8081
CMD ["./user_service"]
