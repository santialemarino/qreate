# syntax=docker/dockerfile:1

ARG GO_VERSION=1.24

FROM golang:${GO_VERSION}-alpine AS builder

# Install dependencies
RUN apk add --no-cache git alpine-sdk

WORKDIR /api

# Copy go mod files
COPY apps/api/go.mod apps/api/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY apps/api/ .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /api

# Copy binary and config
COPY --from=builder /api/app .

EXPOSE 8080
ENV PORT=8080

ENTRYPOINT ["./app"]
