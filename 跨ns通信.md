k8s集群跨命名空间通信

直接svc:namespace:端口
这样可以通信


ping svc-test.demo-com 

比如我想 访问 demo--com 这个namespace下的 svc-test 服务

kind: Service
apiVersion: v1
metadata:
  name: test-xxx
  namespace: default
spec:
  ports:
    - name: http
      protocol: TCP
      port: 9090
      targetPort: 9090
  type: ExternalName
  sessionAffinity: None
  externalName: prometheus.yournamespace.svc.cluster.local

  在ingress的那个namespace创建一个svc 引入externalName externalName里面指定你真实的namespace下svc



% kubectl get svc -n my-ns
NAME            TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
apple-service   ClusterIP   10.98.15.230   <none>        5678/TCP   30d

kubectl apply -f default-svc.yaml 
service/test-cross-ns created


% kubectl get svc -o wide
NAME             TYPE           CLUSTER-IP      EXTERNAL-IP                             PORT(S)    AGE   SELECTOR
banana-service   ClusterIP      10.111.66.167   <none>                                  5678/TCP   33d   app=banana
kubernetes       ClusterIP      10.96.0.1       <none>                                  443/TCP    36d   <none>
test-cross-ns    ExternalName   <none>          apple-service.my-ns.svc.cluster.local   9090/TCP   48s   <none>

 % curl -i http://127.0.0.1:9090
curl: (7) Failed to connect to 127.0.0.1 port 9090: Connection refused

% kubectl apply -f default-ingress.yaml 
ingress.networking.k8s.io/example-ingress created

% kubectl apply -f default-banana.yaml 
pod/banana-app created
service/banana-service unchanged

% kubectl delete -f default-banana.yaml 
pod "banana-app" deleted
service "banana-service" deleted


% curl -i http://localhost/apple  
HTTP/1.1 200 OK
Date: Mon, 26 Apr 2021 15:44:31 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 6
Connection: keep-alive

/apple


 % kubectl get svc -A -o wide
NAMESPACE       NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE   SELECTOR
default         kubernetes                           ClusterIP      10.96.0.1        <none>        443/TCP                      36d   <none>
ingress-nginx   ingress-nginx-controller             LoadBalancer   10.104.184.232   localhost     80:30992/TCP,443:32364/TCP   31d   app.kubernetes.io/component=controller,app.kubernetes.io/instance=ingress-nginx,app.kubernetes.io/name=ingress-nginx
ingress-nginx   ingress-nginx-controller-admission   ClusterIP      10.110.217.145   <none>        443/TCP                      31d   app.kubernetes.io/component=controller,app.kubernetes.io/instance=ingress-nginx,app.kubernetes.io/name=ingress-nginx
kube-system     kube-dns                             ClusterIP      10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP       36d   k8s-app=kube-dns
my-ns           apple-service                        ClusterIP      10.98.15.230     <none>        5678/TCP                     30d   app=apple


% kubectl get Ingress -A -o wide
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAMESPACE   NAME                  CLASS    HOSTS       ADDRESS     PORTS   AGE
my-ns       rancher-k8s-ingress   <none>   localhost   localhost   80      30d


% 
xiazemin@bogon k8s_learn % kubectl describe  Ingress rancher-k8s-ingress -n my-ns
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Name:             rancher-k8s-ingress
Namespace:        my-ns
Address:          localhost
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host        Path  Backends
  ----        ----  --------
  localhost   
              /apple   apple-service:5678 (10.1.1.100:5678)
Annotations:  ingress.kubernetes.io/server-snippet:
                if ($uri ~* "/apple1/.*") {
                    rewrite ^/apple1/(.*) /$1 break;
                }
              kubernetes.io/ingress.class: nginx
              nginx.ingress.kubernetes.io/use-regex: true
Events:       <none>

