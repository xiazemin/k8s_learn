% kubectl get svc -o wide
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE     SELECTOR
apple-service    ClusterIP   10.103.33.174   <none>        5678/TCP   5h6m    app=apple
banana-service   ClusterIP   10.111.66.167   <none>        5678/TCP   5h6m    app=banana
kubernetes       ClusterIP   10.96.0.1       <none>        443/TCP    2d20h   <none>


 %  kubectl get svc -o wide -n cattle-system
NAME              TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE   SELECTOR
rancher-k8s       ClusterIP   10.106.15.8      <none>        80/TCP,443/TCP   40m   app=rancher-k8s
rancher-webhook   ClusterIP   10.105.124.206   <none>        443/TCP          34m   app=rancher-webhook


 % kubectl apply -f sample/ingress-rancher-k8s.yaml
ingress.networking.k8s.io/example-ingress configured

% kubectl apply -f sample/ingress-rancher-k8s.yaml
Error from server (BadRequest): error when creating "sample/ingress-rancher-k8s.yaml": admission webhook "validate.nginx.ingress.kubernetes.io" denied the request: host "_" and path "/rancher" is already defined in ingress default/example-ingress

% kubectl apply -f sample/ingress-rancher-k8s.yaml
ingress.networking.k8s.io/rancher-k8s-ingress created


https://docs.nginx.com/nginx-ingress-controller/configuration/ingress-resources/cross-namespace-configuration/

https://github.com/nginxinc/kubernetes-ingress/tree/v1.10.1/examples-of-custom-resources/cross-namespace-configuration

% kubectl apply -f sample/rancher-k8s-visual-server-route.yaml
error: unable to recognize "sample/rancher-k8s-visual-server-route.yaml": no matches for kind "VirtualServerRoute" in version "k8s.nginx.org/v1"

https://github.com/nginxinc/kubernetes-ingress/tree/v1.10.1/examples/mergeable-ingress-types


https://stackoverflow.com/questions/59844622/ingress-configuration-for-k8s-in-different-namespaces

externalName

apiVersion: v1
kind: Service
metadata:
  namespace: *NAMESPACEB*
  name: checks-service-integration-service
  annotations:
    "alb.ingress.kubernetes.io/healthcheck-path": "/v1/health_check"
    "alb.ingress.kubernetes.io/backend-protocol": "HTTP"
   "alb.ingress.kubernetes.io/target-type": "ip"
spec:
  externalName: checks-service-integration-service.NAMESPACEA.svc.cluster.local
  ports:
    - port: 80

rancher-k8s.cattle-system.svc.cluster.local

% kubectl apply -f sample/ingress-rancher-k8s.yaml
The Ingress "rancher-k8s-ingress" is invalid: spec.rules[0].http.paths[0].backend.service.name: Invalid value: "cattle-system/rancher-k8s": a DNS-1035 label must consist of lower case alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character (e.g. 'my-name',  or 'abc-123', regex used for validation is '[a-z]([-a-z0-9]*[a-z0-9])?')



 % kubectl apply -f sample/ingress-rancher-k8s.yaml
ingress.networking.k8s.io/rancher-k8s-ingress created

https://github.com/kubernetes/ingress-nginx/blob/master/docs/examples/rewrite/README.md


% kubectl get endpoints
NAME             ENDPOINTS           AGE
apple-service    10.1.0.105:5678     6h42m
banana-service   10.1.0.106:5678     6h42m
kubernetes       192.168.65.4:6443   2d21h

 % kubectl get endpoints  -n cattle-system
NAME              ENDPOINTS                                               AGE
rancher-k8s       10.1.0.113:80,10.1.0.114:80,10.1.0.115:80 + 3 more...   155m
rancher-webhook   10.1.0.119:9443                                         149m

% kubectl get Ingress
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME                  CLASS    HOSTS         ADDRESS     PORTS   AGE
ingress-with-auth     <none>   foo.bar.com   localhost   80      29h
rancher-k8s-ingress   <none>   *             localhost   80      83m

% kubectl get Ingress -n cattle-system
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME                  CLASS    HOSTS            ADDRESS     PORTS     AGE
rancher-k8s           <none>   rancher.my.org   localhost   80, 443   154m
rancher-k8s-ingress   <none>   *                localhost   80        29m


不同ns下的 rancher-k8s-ingress  是不一样的，声明周期不一样



 % kubectl get Ingress  -A
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAMESPACE       NAME                  CLASS    HOSTS       ADDRESS     PORTS   AGE
cattle-system   rancher-k8s-ingress   <none>   localhost   localhost   80      2m25s
default         example-ingress       <none>   *           localhost   80      39s

 % kubectl apply -f sample/ingress.v1.yaml
ingress.networking.k8s.io/example-ingress created

 % kubectl apply -f sample/ingress-rancher-k8s.yaml
ingress.networking.k8s.io/rancher-k8s-ingress configured



 % kubectl describe Ingress  rancher-k8s-ingress -n cattle-system
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Name:             rancher-k8s-ingress
Namespace:        cattle-system
Address:          localhost
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)


 % kubectl describe Ingress   example-ingress
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Name:             example-ingress
Namespace:        default
Address:          localhost
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)

 % kubectl delete deploy cattle-cluster-agent  -n cattle-system
deployment.apps "cattle-cluster-agent" deleted



https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md

https://blog.csdn.net/qingyafan/article/details/82692509
https://www.jianshu.com/p/665ef97bf977?utm_source=oschina-app

https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/



https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/

https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/

https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/custom-template/

