#!/bin/bash -e

RELEASE_NAME="token-janitor"
NAMESPACE="pir-intrusion-detection"

echo "Uninstalling previous release..."

helm uninstall "$RELEASE_NAME" --namespace "$NAMESPACE"
