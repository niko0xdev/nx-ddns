# Stage 1: Build the Go application
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app for multiple architectures
RUN GOOS=linux GOARCH=amd64 go build -o nxddns-amd64 ./cmd/nxddns
RUN GOOS=linux GOARCH=arm64 go build -o nxddns-arm64 ./cmd/nxddns
RUN GOOS=linux GOARCH=armv7 go build -o nxddns-armv7 ./cmd/nxddns

# Stage 2: Build the final image
FROM alpine:latest

WORKDIR /root/

# Install necessary certificates (like CA certificates) for all platforms
RUN apk add --no-cache ca-certificates

# Copy the pre-built binaries from the builder stage
COPY --from=builder /app/nxddns-amd64 ./nxddns-amd64
COPY --from=builder /app/nxddns-arm64 ./nxddns-arm64
COPY --from=builder /app/nxddns-armv7 ./nxddns-armv7

# Expose ports if needed (change port if required)
EXPOSE 8080

# Default command to run based on architecture
CMD ["sh", "-c", "if [ $(uname -m) = 'x86_64' ]; then ./nxddns-amd64; elif [ $(uname -m) = 'aarch64' ]; then ./nxddns-arm64; else ./nxddns-armv7; fi"]
