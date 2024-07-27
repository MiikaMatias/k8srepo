PROJECT_PATH=/home/miika/Projects/devopsK8s
PROJECT_DIRS=(
  to-do-backend
)
cd ${PROJECT_PATH}

kubectl apply -f persistent-volumes/    

for dir in "${PROJECT_DIRS[@]}"; do
  cd "${PROJECT_PATH}/$dir" || exit
  bash deploy.sh
done

cd /home/miika/Projects/devopsK8s

kubectl apply -f manifests/
