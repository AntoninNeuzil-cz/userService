# Build stage
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the binary
RUN go build -o userService cmd/server/server.go

# Runtime stage
FROM debian:bookworm-slim

# Install the required version of glibc
RUN apt-get update && apt-get install -y libc6

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/userService .

# Expose the HTTP port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./userService"]