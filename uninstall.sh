#!/bin/bash
source .env
cd ${HOME_PATH}

kubectl delete -f manifests/
kubectl delete -f persistent-volumes/

kubectl config set-context --current --namespace=applications
for dir in "${PROJECT_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  kubectl delete -f manifests/
done

kubectl config set-context --current --namespace=project
for dir in "${APPLICATION_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  kubectl delete -f manifests/
done