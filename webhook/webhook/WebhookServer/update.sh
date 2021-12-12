CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o webhook .
docker build -t webhook-example:0.0.5 .
kubectl delete -f deploy/ValidatingWebhookConfiguration.yaml
kubectl delete -f deploy/deployment-simple.yaml
kubectl apply -f deploy/deployment-simple.yaml
kubectl delete -f deploy/service.yaml 
kubectl apply -f deploy/service.yaml
sleep 1
kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml