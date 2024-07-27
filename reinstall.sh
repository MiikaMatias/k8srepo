<<<<<<< HEAD
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
=======
source .env
cd ${HOME_PATH}

kubectl config set-context --current --namespace=applications
bash uninstall.sh
kubectl delete all --all            
docker rmi $(docker images -q)
kubectl apply -f manifests/application-manifests/

kubectl config set-context --current --namespace=project
bash uninstall.sh
kubectl delete all --all            
docker rmi $(docker images -q)
kubectl apply -f manifests/project-manifests/

kubectl config set-context --current --namespace=applications
for dir in "${PROJECT_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  bash deploy.sh
done

kubectl config set-context --current --namespace=project
for dir in "${APPLICATION_DIRS[@]}"; do
  cd "${HOME_PATH}/$dir" || exit
  bash deploy.sh
done

cd /home/miika/Projects/devopsK8s
>>>>>>> 18d42f1 (2.03 and 2.04: namespace sep)
