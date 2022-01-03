% istioctl install --set profile=minimal -y                                               ✔ Istio core installed                                                                                                                  
✔ Istiod installed                                                                                                                      
- Pruning removed resources                                                                                                               Removed HorizontalPodAutoscaler:istio-system:istio-ingressgateway.
  Removed PodDisruptionBudget:istio-system:istio-ingressgateway.
  Removed Deployment:istio-system:istio-ingressgateway.
  Removed Service:istio-system:istio-ingressgateway.
  Removed ServiceAccount:istio-system:istio-ingressgateway-service-account.
  Removed RoleBinding:istio-system:istio-ingressgateway-sds.
  Removed Role:istio-system:istio-ingressgateway-sds.
✔ Installation complete                                                                                                                 Making this installation the default for injection and validation.
2022-01-03T10:36:53.889205Z     error   klog    couldn't get resource list for custom.metrics.k8s.io/v1beta1: the server is currently unable to handle the request
2022-01-03T10:36:53.891600Z     error   klog    couldn't get resource list for custom.metrics.k8s.io/v1beta1: the server is currently unable to handle the request

Thank you for installing Istio 1.12.  Please take a few minutes to tell us about your install/upgrade experience!  https://forms.gle/FegQbc9UvePd4Z9z7


 % kubectl get pods -n istio-system
NAME                      READY   STATUS    RESTARTS   AGE
istiod-58d79b7bff-222hd   1/1     Running   0          146m


 % kubectl get deploy -n istio-system
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
istiod   1/1     1            1           7h42m


 % kubectl get svc istio-ingressgateway -n istio-system
NAME                   TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   LoadBalancer   10.100.141.28   localhost     15021:30184/TCP,80:31618/TCP,443:31343/TCP,31400:30542/TCP,15443:30291/TCP   65s


% kubectl -n istio-system describe deploy istio-ingressgateway |grep fail
    Readiness:  http-get http://:15021/healthz/ready delay=1s timeout=1s period=2s #success=1 #failure=30


 % istioctl version
client version: 1.12.1
control plane version: 1.12.1
data plane version: none

