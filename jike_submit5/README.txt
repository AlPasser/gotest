helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack

kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml
# config and reload prometheus
kubectl create -f rbac.yaml

kubectl apply -f gowebmetrics.yaml
kubectl apply -f service.yaml

kubectl port-forward svc/kube-prometheus-stack-prometheus 9090:9090 --address='0.0.0.0'
histogram_quantile(0.5, sum(rate(gowebmetics_execution_latency_seconds_bucket[5m])) by (le))

kubectl port-forward svc/kube-prometheus-stack-grafana 9000:80 --address='0.0.0.0'

