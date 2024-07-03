#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

cd "$1" || { echo "Failed to change directory to $1"; exit 1; }

DEPLOYMENT_NAME="$1-dep"
IMAGE_NAME="kontrakti/$1"
IMAGE_TAG="latest"

#echo "Deleting cluster"
#k3d cluster delete

#echo "Redeploying cluster"
#k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2

echo "Deleting Kubernetes deployment $DEPLOYMENT_NAME..."
kubectl delete deployments $DEPLOYMENT_NAME

echo "Building Docker image $IMAGE_NAME:$IMAGE_TAG..."
docker build . -t $IMAGE_NAME:$IMAGE_TAG || { echo "Docker build failed"; exit 1; }

echo "Pushing Docker image $IMAGE_NAME:$IMAGE_TAG to registry..."
docker push $IMAGE_NAME:$IMAGE_TAG || { echo "Docker push failed"; exit 1; }

echo "Creating service"
kubectl apply -f manifests/ || { echo "Failed to apply Kubernetes service"; exit 1; }

echo "Waiting for deployment $DEPLOYMENT_NAME to be available..."
kubectl wait --for=condition=available --timeout=600s deployment/$DEPLOYMENT_NAME || { echo "Deployment did not become available in time"; exit 1; }

echo "Checking pod status..."
kubectl get pods -l app=$DEPLOYMENT_NAME || { echo "Failed to get pod status"; exit 1; }

echo "Checking container logs..."
kubectl logs -l app=$DEPLOYMENT_NAME || { echo "Failed to get container logs"; exit 1; }

echo "Deployment updated successfully!"
