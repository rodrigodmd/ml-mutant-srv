
kubectl create -f k8/elasticsearch-service.yaml
kubectl create -f k8/elasticsearch-deployment.yaml
kubectl create -f k8/consul-deployment.yaml
kubectl create -f k8/consul-service.yaml
kubectl create -f k8/web-service.yaml
kubectl create -f k8/web-deployment.yaml

kubectl create -f k8/api-service.yaml
kubectl create -f k8/api-deployment.yaml
kubectl create -f k8/mutant-deployment.yaml
kubectl create -f k8/stats-deployment.yaml
kubectl create -f k8/gateway-deployment.yaml









