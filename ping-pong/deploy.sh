cd /home/miika/Projects/devopsK8s/ping-pong/ping-pong
docker build . --no-cache -t kontrakti/ping-pong:latest
docker push kontrakti/ping-pong:latest

cd /home/miika/Projects/devopsK8s/ping-pong/ping-pong-db
docker build . --no-cache -t kontrakti/ping-pong-db:latest
docker push kontrakti/ping-pong-db:latest

cd /home/miika/Projects/devopsK8s/ping-pong/ping-pong-db
docker build . --no-cache -t kontrakti/ping-pong-db:latest
docker push kontrakti/ping-pong-db:latest

cd /home/miika/Projects/devopsK8s/ping-pong/

kubectl apply -f manifests/