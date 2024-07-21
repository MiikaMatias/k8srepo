cd /home/miika/Projects/devopsK8s/project
docker build . --no-cache -t kontrakti/project:latest
docker push kontrakti/project:latest
kubectl apply -f manifests/