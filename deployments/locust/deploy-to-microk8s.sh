#!/usr/bin/env bash

# Build image
docker build . -t locust:local

# Remove old container if it exists
microk8s.ctr -n k8s.io images remove docker.io/library/locust:local

# Transfer image to microk8s environment
docker save locust > locust.tar
microk8s.ctr -n k8s.io image import locust.tar

# Cleanup
rm locust.tar
docker image rm locust:local

# Apply deployment
microk8s.kubectl apply -f locust-deployment.yaml
