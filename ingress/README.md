

#It is built around the Kubernetes Ingress resource, using a ConfigMap to store the NGINX configuration.
#https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop


curl -o ingress-controller.yaml https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.0/deploy/static/provider/cloud/deploy.yaml

kubectl apply -f ingress-controller.yaml 


% kubectl get pods -n ingress-nginx
NAME                                       READY   STATUS              RESTARTS   AGE
ingress-nginx-admission-create-2x44f       0/1     ImagePullBackOff    0          93s
ingress-nginx-admission-patch-jhwl6        0/1     ImagePullBackOff    0          93s
ingress-nginx-controller-fd7bb8d66-qm82d   0/1     ContainerCreating   0          94s

docker pull k8s.gcr.io/ingress-nginx/controller:v1.0.0@sha256:0851b34f69f69352bf168e6ccf30e1e20714a264ab1ecd1933e4d8c0fc3215c6
docker pull k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.0@sha256:f3b6b39a6062328c095337b4cadcefd1612348fdd5190b1dcbcb9b9e90bd8068

#Error response from daemon: Get "https://k8s.gcr.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)

quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.1
docker pull jettech/kube-webhook-certgen:v1.2.0

https://blog.csdn.net/thermal_life/article/details/107005689
https://www.cnblogs.com/v-fan/p/13252372.html

kubectl delete -f ingress-controller.yaml 
kubectl apply -f ingress-controller.yaml 

kubectl get pods -n ingress-nginx 
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-lszdt        0/1     Completed   0          99s
ingress-nginx-admission-patch-f9cjg         0/1     Completed   1          99s
ingress-nginx-controller-5889998cb4-tbrkv   0/1     Error       4          99s

kubectl describe pod ingress-nginx-controller-5889998cb4-tbrkv -n ingress-nginx

kubectl logs pods ingress-nginx-controller-5889998cb4-cj998 -n ingress-nginx
Error from server (NotFound): pods "pods" not found

% kubectl top nodes -n ingress-nginx
W0825 13:50:11.974677   98784 top_node.go:119] Using json format to get metrics. Next release will switch to protocol-buffers, switch early by passing --use-protocol-buffers flag
error: Metrics API not available

#https://blog.csdn.net/TinyJian/article/details/109711164


jettech/kube-webhook-certgen:v1.5.1
quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.26.1

registry.cn-hangzhou.aliyuncs.com/acs/aliyun-ingress-controller:v0.44.0.1-5e842447b-aliyun