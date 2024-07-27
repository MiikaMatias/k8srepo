
source .env

export SOPS_AGE_KEY_FILE=$(pwd)/key.txt
sops --decrypt $(pwd)/ping-pong/manifests/secret.enc.yaml | kubectl apply -f -

kubectl config set-context --current --namespace=applications
docker rmi $(docker images -q)
kubectl apply -f manifests/application-manifests/

kubectl config set-context --current --namespace=applications
for dir in "${PROJECT_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  bash deploy.sh
done

kubectl config set-context --current --namespace=project
docker rmi $(docker images -q)
kubectl apply -f manifests/project-manifests/

kubectl config set-context --current --namespace=project
for dir in "${APPLICATION_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  bash deploy.sh
done

cd ${HOME_PATH}