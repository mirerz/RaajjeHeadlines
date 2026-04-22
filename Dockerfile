# 🐳 Raajjé HEADLINES | Multi-Service Docker Infrastructure
# Build Stage
FROM golang:alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Prefetch modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build Binary Target 1: API
RUN go build -o /app/api_bin ./cmd/api/main.go

# Build Binary Target 2: Aggregator
RUN go build -o /app/aggregator_bin ./cmd/aggregator/main.go

# Production Stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# Copy binaries
COPY --from=builder /app/api_bin .
COPY --from=builder /app/aggregator_bin .

# Copy assets & configs
COPY ./web ./web
COPY ./configs ./configs

# Default entrypoint for API (Overridden in compose for aggregator)
EXPOSE 8080
ENTRYPOINT ["./api_bin"]
