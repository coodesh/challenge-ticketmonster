# Multi-stage build for Go application
FROM golang:1.21-alpine AS builder
WORKDIR /ticketmonster

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ticketmonster cmd/main.go

FROM alpine:latest
WORKDIR /ticketmonster
COPY --from=builder /ticketmonster .
COPY .env.example .env
EXPOSE 8080
CMD ["./ticketmonster"]
