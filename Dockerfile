FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o nexora-api ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/nexora-api .
EXPOSE 8080

CMD ["./nexora-api"]