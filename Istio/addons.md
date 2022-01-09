% kubectl apply -f samples/addons
serviceaccount/grafana created
configmap/grafana created
service/grafana created
deployment.apps/grafana created
configmap/istio-grafana-dashboards created
configmap/istio-services-grafana-dashboards created
deployment.apps/jaeger created
service/tracing created
service/zipkin created
service/jaeger-collector created
serviceaccount/kiali created
configmap/kiali created
clusterrole.rbac.authorization.k8s.io/kiali-viewer created
clusterrole.rbac.authorization.k8s.io/kiali created
clusterrolebinding.rbac.authorization.k8s.io/kiali created
role.rbac.authorization.k8s.io/kiali-controlplane created
rolebinding.rbac.authorization.k8s.io/kiali-controlplane created
service/kiali created
deployment.apps/kiali created
serviceaccount/prometheus created
configmap/prometheus created
clusterrole.rbac.authorization.k8s.io/prometheus configured
clusterrolebinding.rbac.authorization.k8s.io/prometheus configured
service/prometheus created
deployment.apps/prometheus created



 % kubectl rollout status deployment/kiali -n istio-system
Waiting for deployment "kiali" rollout to finish: 0 of 1 updated replicas are available...



 % istioctl dashboard kiali                               
Error: failure running port forward process: failure running port forward process: pod is not running. Status=Pending


% kubectl -n istio-system get deploy kiali  -o yaml |grep image
 image: quay.io/kiali/kiali:v1.42

% kubectl -n istio-system get deploy kiali  -o yaml |grep image     
 image: prom/prometheus:v2.26.0



https://hub.docker.com/r/rancher/kiali-kiali/tags
docker pull rancher/kiali-kiali:v1.29.0
f9125884adc5: Pull complete 
Digest: sha256:6adc6c99d279a09b138f5e1fdaccca2fe439412aca20949d4513b3c47d8d74fd
Status: Downloaded newer image for rancher/kiali-kiali:v1.29.0


 % istioctl dashboard kiali                          
http://localhost:20001/kiali


Kiali 为网格管理和可观察性提供了良好的用户体验的可视化工具;

Kiali 为我们提供了查看相关服务与配置提供了统一化的可视化界面，并且能在其中展示他们的关联；同时他还提供了界面让我们可以很方便的验证 istio 配置与错误提示;

https://zhuanlan.zhihu.com/p/41324294


kubectl config set-context default --namespace=${work_namespace}
https://blog.csdn.net/varyuan/article/details/112210085


https://istio.io/latest/docs/setup/getting-started/


https://www.jianshu.com/p/07455dbfd6bb
