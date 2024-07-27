cd /home/miika/Projects/devopsK8s/to-do-backend
docker build . --no-cache -t kontrakti/to-do-backend:latest
docker push kontrakti/to-do-backend:latest
kubectl apply -f manifests/