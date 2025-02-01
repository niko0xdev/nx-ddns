# Stage 1: Build the application
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o nxddns ./cmd/nxddns

# Stage 2: Build the final image
FROM alpine:latest

WORKDIR /root/

# Install dependencies (if any)
RUN apk add --no-cache ca-certificates

# Copy the Pre-built binary from the builder stage
COPY --from=builder /app/nxddns .

# Expose port (optional, depending on your app's network configuration)
EXPOSE 8080

# Command to run the executable
CMD ["./nxddns"]
