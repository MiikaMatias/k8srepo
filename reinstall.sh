PROJECT_PATH=/home/miika/Projects/devopsK8s
PROJECT_DIRS=(
  log-output
  ping-pong
  to-do-backend
  project
)
cd ${PROJECT_PATH}

bash uninstall.sh
kubectl delete all --all            
docker rmi $(docker images -q)
kubectl apply -f persistent-volumes/    

for dir in "${PROJECT_DIRS[@]}"; do
  cd "${PROJECT_PATH}/$dir" || exit
  bash deploy.sh
done

cd /home/miika/Projects/devopsK8s

kubectl apply -f manifests/
