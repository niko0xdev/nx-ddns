# Stage 1: Build the Go application
FROM golang:1.23.1-alpine3.20 as builder

LABEL maintainer="niko0xdev <niko0xdev@gmail.com>"

ARG OS

ARG ARCH

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app for multiple architectures
RUN go build -o nxddns ./cmd/nxddns/main.go

# Stage 2: Build the final image
FROM alpine:3.20

ARG VERSION
ARG user=nxddns
ARG group=nxddns
ARG uid=1000
ARG gid=1000

ARG API_PORT=8080

USER root

# Install necessary certificates (like CA certificates) for all platforms
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the pre-built binaries from the builder stage
COPY --from=builder /app/nxddns ./nxddns

RUN apk update && apk --no-cache add bash && addgroup -g ${gid} ${group} && adduser -h /app -u ${uid} -G ${group} -s /bin/bash -D ${user}
RUN chown ${user}:${group} /app/nxddns && chmod +x /app/nxddns

USER ${user}

# Expose ports if needed (change port if required)
EXPOSE ${API_PORT}

# Default command to run
CMD ["./nxddns"]
