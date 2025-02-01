#!/bin/bash

# Exit on error
set -e

# Define the image name and tag
IMAGE_NAME="niko0xdev/nxddns"
TAG="$1"

# Platforms to build for
PLATFORMS="linux/amd64,linux/arm64,linux/arm/v7"

# Build the image using docker buildx
echo "Building Docker image for platforms: $PLATFORMS..."
docker buildx build --platform $PLATFORMS -t $IMAGE_NAME:$TAG  --push .

echo "Build completed successfully!"
