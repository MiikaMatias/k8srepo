k3d cluster create \
  --k3s-arg '--kubelet-arg=eviction-hard=imagefs.available<1%,nodefs.available<1%@agent:*' \
  --k3s-arg '--kubelet-arg=eviction-minimum-reclaim=imagefs.available=1%,nodefs.available=1%@agent:*' \
  --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2

kubectl create namespace applications
kubectl create namespace project