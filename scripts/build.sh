#!/bin/bash -e

IMAGE_NAME="token-janitor"
TAG=$(date +"%Y%m%d-%H%M%S")

echo "######################################"
echo "Building $IMAGE_NAME:$TAG..."
echo "######################################"
echo

docker build -t "$IMAGE_NAME":"$TAG" .
