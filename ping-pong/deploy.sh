cd /home/miika/Projects/devopsK8s/ping-pong/ping-pong
docker build . -t kontrakti/ping-pong:latest
docker push kontrakti/ping-pong:latest

cd /home/miika/Projects/devopsK8s/ping-pong/

kubectl apply -f manifests/