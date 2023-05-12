#!/bin/bash -e

IMAGE_NAME="token-janitor"
TAG=$(date +"%Y%m%d-%H%M%S")

echo "Building $IMAGE_NAME:$TAG..."

docker build -t "$IMAGE_NAME":"$TAG" .
