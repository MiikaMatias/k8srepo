#!/bin/bash

source .env

kubectl delete -f manifests/
kubectl delete -f persistent-volumes/

for dir in "${DIRS[@]}"; do
  cd "$dir" || exit
  kubectl delete -f manifests/
done
