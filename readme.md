k3d cluster create \
    -p "9900:80@loadbalancer" \
    --k3s-arg '--kubelet-arg=eviction-hard=imagefs.available<1%,nodefs.available<1%@agent:*' \
    --k3s-arg '--kubelet-arg=eviction-minimum-reclaim=imagefs.available=1%,nodefs.available=1%@agent:*' \
    --servers 1 --agents 1 --registry-create mycluster-registry:0.0.0.0:5000

docker build -t go-web-server .

docker tag go-web-server:latest mycluster-registry:5000/go-web-server:local

docker push mycluster-registry:5000/go-web-server:local

<!--kubectl run --image mycluster-registry:5000/testimage:local testimage --command -- tail -f /dev/null -->

kubectl create deployment go-web-server --image=mycluster-registry:5000/go-web-server:local

kubectl create service clusterip go-web-server --tcp=9091:9091

kubectl apply -f ./ingress.yaml

visit http://localhost:9900/

k3d cluster delete k3s-default