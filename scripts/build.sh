#!/bin/bash -e

echo "this script will build the container image"

docker build -t token-janitor:"$TAG" .
