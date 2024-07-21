cd /home/miika/Projects/devopsK8s

bash uninstall.sh
kubectl delete all --all            
docker rmi $(docker images -q)

kubectl apply -f persistent-volumes/    

cd /home/miika/Projects/devopsK8s/log-output
bash deploy.sh

cd /home/miika/Projects/devopsK8s/ping-pong
bash deploy.sh

cd /home/miika/Projects/devopsK8s/project
bash deploy.sh

cd /home/miika/Projects/devopsK8s

kubectl apply -f manifests/
