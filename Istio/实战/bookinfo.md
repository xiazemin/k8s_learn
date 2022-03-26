 % kubectl create ns bookinfo
namespace/bookinfo created

 % kubectl -n bookinfo apply -f samples/bookinfo/networking/destination-rule-all.yaml
destinationrule.networking.istio.io/productpage created
destinationrule.networking.istio.io/reviews created
destinationrule.networking.istio.io/ratings created
destinationrule.networking.istio.io/details created


 % kubectl get all -n bookinfo
No resources found in bookinfo namespace.


 %  kubectl label namespace bookinfo istio-injection=enabled
error: 'istio-injection' already has a value (enabled), and --overwrite is false


%  kubectl -n bookinfo apply -f samples/bookinfo/platform/kube/bookinfo.yaml
service/details created
serviceaccount/bookinfo-details created
deployment.apps/details-v1 created
service/ratings created
serviceaccount/bookinfo-ratings created
deployment.apps/ratings-v1 created
service/reviews created
serviceaccount/bookinfo-reviews created
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage created
serviceaccount/bookinfo-productpage created
deployment.apps/productpage-v1 created



 % kubectl get all -n bookinfo
NAME                                  READY   STATUS    RESTARTS   AGE
pod/details-v1-79f774bdb9-dg7tp       2/2     Running   0          19s
pod/productpage-v1-6b746f74dc-k7lph   2/2     Running   0          18s
pod/ratings-v1-b6994bb9-85mtv         2/2     Running   0          19s
pod/reviews-v1-545db77b95-d79vb       2/2     Running   0          18s
pod/reviews-v2-7bf8c9648f-g8j2q       2/2     Running   0          18s
pod/reviews-v3-84779c7bbc-gjm7n       2/2     Running   0          18s

NAME                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/details       ClusterIP   10.96.149.58     <none>        9080/TCP   19s
service/productpage   ClusterIP   10.103.21.203    <none>        9080/TCP   19s
service/ratings       ClusterIP   10.108.210.196   <none>        9080/TCP   19s
service/reviews       ClusterIP   10.103.56.248    <none>        9080/TCP   19s

NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/details-v1       1/1     1            1           19s
deployment.apps/productpage-v1   1/1     1            1           18s
deployment.apps/ratings-v1       1/1     1            1           19s
deployment.apps/reviews-v1       1/1     1            1           19s
deployment.apps/reviews-v2       1/1     1            1           19s
deployment.apps/reviews-v3       1/1     1            1           19s

NAME                                        DESIRED   CURRENT   READY   AGE
replicaset.apps/details-v1-79f774bdb9       1         1         1       19s
replicaset.apps/productpage-v1-6b746f74dc   1         1         1       18s
replicaset.apps/ratings-v1-b6994bb9         1         1         1       19s
replicaset.apps/reviews-v1-545db77b95       1         1         1       19s
replicaset.apps/reviews-v2-7bf8c9648f       1         1         1       19s
replicaset.apps/reviews-v3-84779c7bbc       1         1         1       19s


 %  kubectl  -n bookinfo exec "$(kubectl  -n bookinfo get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"
<title>Simple Bookstore App</title>


 % istioctl dashboard kiali
Error: no Kiali pods found


%  kubectl apply -f samples/addons
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
clusterrole.rbac.authorization.k8s.io/kiali-viewer unchanged
clusterrole.rbac.authorization.k8s.io/kiali unchanged
clusterrolebinding.rbac.authorization.k8s.io/kiali unchanged
role.rbac.authorization.k8s.io/kiali-controlplane created
rolebinding.rbac.authorization.k8s.io/kiali-controlplane created
service/kiali created
deployment.apps/kiali created
serviceaccount/prometheus created
configmap/prometheus created
clusterrole.rbac.authorization.k8s.io/prometheus unchanged
clusterrolebinding.rbac.authorization.k8s.io/prometheus unchanged
service/prometheus created
deployment.apps/prometheus created



 % kubectl get all -n istio-system
NAME                                        READY   STATUS    RESTARTS   AGE
pod/grafana-6ccd56f4b6-rkzpx                1/1     Running   0          6m31s
pod/istio-egressgateway-687f4db598-8dmmb    1/1     Running   0          45m
pod/istio-ingressgateway-78f69bd5db-tdbtw   1/1     Running   0          45m
pod/istiod-76d66d9876-hgsmt                 1/1     Running   0          45m
pod/jaeger-5d44bc5c5d-bsppl                 1/1     Running   0          6m31s
pod/kiali-79b86ff5bc-xzwgl                  1/1     Running   0          6m30s
pod/prometheus-64fd8ccd65-bblgs             2/2     Running   0          6m29s

NAME                           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
service/grafana                ClusterIP   10.105.189.206   <none>        3000/TCP                                                                     6m32s
service/istio-egressgateway    ClusterIP   10.105.172.31    <none>        80/TCP,443/TCP                                                               45m
service/istio-ingressgateway   NodePort    10.104.14.138    <none>        15021:31594/TCP,80:31760/TCP,443:30272/TCP,31400:30137/TCP,15443:32583/TCP   45m
service/istiod                 ClusterIP   10.108.137.91    <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        45m
service/jaeger-collector       ClusterIP   10.102.255.227   <none>        14268/TCP,14250/TCP,9411/TCP                                                 6m31s
service/kiali                  ClusterIP   10.103.88.172    <none>        20001/TCP,9090/TCP                                                           6m30s
service/prometheus             ClusterIP   10.107.190.173   <none>        9090/TCP                                                                     6m29s
service/tracing                ClusterIP   10.100.68.250    <none>        80/TCP,16685/TCP                                                             6m31s
service/zipkin                 ClusterIP   10.100.196.198   <none>        9411/TCP                                                                     6m31s

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/grafana                1/1     1            1           6m31s
deployment.apps/istio-egressgateway    1/1     1            1           45m
deployment.apps/istio-ingressgateway   1/1     1            1           45m
deployment.apps/istiod                 1/1     1            1           45m
deployment.apps/jaeger                 1/1     1            1           6m31s
deployment.apps/kiali                  1/1     1            1           6m30s
deployment.apps/prometheus             1/1     1            1           6m29s

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/grafana-6ccd56f4b6                1         1         1       6m31s
replicaset.apps/istio-egressgateway-687f4db598    1         1         1       45m
replicaset.apps/istio-ingressgateway-78f69bd5db   1         1         1       45m
replicaset.apps/istiod-76d66d9876                 1         1         1       45m
replicaset.apps/jaeger-5d44bc5c5d                 1         1         1       6m31s
replicaset.apps/kiali-79b86ff5bc                  1         1         1       6m30s
replicaset.apps/prometheus-64fd8ccd65             1         1         1       6m29s
xiazemin@xiazemindeMBP istio-1.12.1 %





 % istioctl dashboard kiali
http://localhost:20001/kiali


% kubectl delete dr details
destinationrule.networking.istio.io "details" deleted
xiazemin@xiazemindeMBP MyBlogSrc % kubectl delete dr productpage ratings reviews
destinationrule.networking.istio.io "productpage" deleted
destinationrule.networking.istio.io "ratings" deleted
destinationrule.networking.istio.io "reviews" deleted

服务暴露给外网
 %  kubectl  -n bookinfo apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
gateway.networking.istio.io/bookinfo-gateway created
virtualservice.networking.istio.io/bookinfo created

 % kubectl get gateway  -n bookinfo
NAME               AGE
bookinfo-gateway   81s

https://istio.io/latest/docs/examples/bookinfo/


% kubectl get svc istio-ingressgateway -n istio-system
NAME                   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   NodePort   10.104.14.138   <none>        15021:31594/TCP,80:31760/TCP,443:30272/TCP,31400:30137/TCP,15443:32583/TCP   69m


http://localhost:31760/productpage

 % ab -n 100 http://localhost:31760/productpage

This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/


 % istioctl profile dump demo > demo.profile
 