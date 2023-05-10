#!/bin/bash -e

echo "this script will deploy the helm chart"
echo ""

helm install token-janitor ./helm-chart --namespace pir-intrusion-detection
helm upgrade token-janitor ./helm-chart --namespace pir-intrusion-detection
#helm uninstall token-janitor --namespace pir-intrusion-detection
