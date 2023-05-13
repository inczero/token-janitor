#!/bin/bash -e

RELEASE_NAME="token-janitor"
NAMESPACE="pir-intrusion-detection"

echo "Deploying $RELEASE_NAME..."

helm upgrade --install "$RELEASE_NAME" --values ./real-values.yaml ./helm-chart --namespace "$NAMESPACE"
