cd /home/miika/Projects/devopsK8s/log-output/log-reader
docker build . -t kontrakti/log-reader:latest
docker push kontrakti/log-reader:latest
cd /home/miika/Projects/devopsK8s/log-output/log-writer
docker build . -t kontrakti/log-writer:latest
docker push kontrakti/log-writer:latest

cd /home/miika/Projects/devopsK8s/log-output/
kubectl apply -f manifests/