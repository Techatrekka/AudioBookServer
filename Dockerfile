# Use official Golang image as a builder
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy application source
COPY . .

# Ensure the binary is built for Alpine Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go

# Use a lightweight Alpine Linux image
FROM alpine:latest
WORKDIR /root/

# Install dependencies
RUN apk --no-cache add ca-certificates

# Copy the compiled binary and grant execution permissions
COPY --from=builder /app/server .
RUN chmod +x /root/server

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/root/server"]