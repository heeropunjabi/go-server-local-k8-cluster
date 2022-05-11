# How Things are working locally

![image](https://user-images.githubusercontent.com/3514019/167405868-ec901348-5bc4-4301-bcd9-2cf9f398d4fb.png)


## run go server locally
    go run main.go
## Build docker image from docker file
    docker build -t go-web-server .

## map following entry in your /etc/hosts file
    127.0.0.1	mycluster-registry

## Start the cluster, expose load balancer to local port 9900,create your own registery called mycluster-registry  

    k3d cluster create \
        -p "9900:80@loadbalancer" \
        --k3s-arg '--kubelet-arg=eviction-hard=imagefs.available<1%,nodefs.available<1%@agent:*' \
        --k3s-arg '--kubelet-arg=eviction-minimum-reclaim=imagefs.available=1%,nodefs.available=1%@agent:*' \
        --servers 1 --agents 1 --registry-create mycluster-registry:0.0.0.0:5000

## Tag docker image with my-cluster registry
    docker tag go-web-server:latest mycluster-registry:5000/go-web-server:local

## push docker image to local my-cluster registry
    docker push mycluster-registry:5000/go-web-server:local

## Create deployment from the image created from previous step, pull from the local registry 
    kubectl create deployment go-web-server --image=mycluster-registry:5000/go-web-server:local
## create servce which is pointing to deployment
    kubectl create service clusterip go-web-server --tcp=9091:9091

## deploy ingress reverse proxy controller 
    kubectl apply -f ./ingress.yaml

## verify your Go server is up & running
http://localhost:9900/

## Delete local cluster using k3d
    k3d cluster delete k3s-default
