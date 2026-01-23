# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN make build

# Final stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
# Chromium is required for Rod
RUN apk add --no-cache chromium

# Copy binary from builder
COPY --from=builder /app/bin/mime /usr/local/bin/mime

# Create user for security
RUN adduser -D mimeuser
USER mimeuser

# Expose any ports if necessary (MIME usually runs via stdio but can run as server)
# EXPOSE 8080

ENTRYPOINT ["mime"]
CMD ["serve"]
