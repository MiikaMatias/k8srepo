cd /home/miika/Projects/devopsK8s/log-output
bash deploy.sh

cd /home/miika/Projects/devopsK8s/ping-pong
bash deploy.sh

cd /home/miika/Projects/devopsK8s

kubectl apply -f manifests/
