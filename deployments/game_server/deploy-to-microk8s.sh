#!/usr/bin/env bash

# Build image
docker build . -t gameserver:local

# Transfer image to microk8s environment
docker save gameserver > gameserver.tar
microk8s.ctr -n k8s.io image import gameserver.tar

# Cleanup
rm gameserver.tar
docker image rm gameserver:local

# Apply deployment
microk8s.kubectl apply -f gameserver-deployment.yaml
