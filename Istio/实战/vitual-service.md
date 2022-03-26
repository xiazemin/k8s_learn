所有进行通讯的服务必须都进行了istio注入

% kubectl apply -f ./vitural-service/client.yaml -n inject
deployment.apps/client created

% kubectl apply -f ./vitural-service/deploy.yaml -n inject
deployment.apps/httpd created
deployment.apps/tomcat created

% kubectl apply -f ./vitural-service/service.yaml -n inject
service/tomcat-svc created
service/httpd-svc created

% kubectl -n inject get endpoints
NAME         ENDPOINTS   AGE
httpd-svc    <none>      66s
tomcat-svc   <none>      67s

修改 label
 % kubectl -n inject get endpoints -o wide
NAME         ENDPOINTS   AGE
httpd-svc                3m17s
tomcat-svc               3m18s

 % docker run -it busybox sh      
/ # exit


% docker pull tomcat
Using default tag: latest
latest: Pulling from library/tomcat

docker.io/library/tomcat:latest







 % kubectl create ns vc   
namespace/vc created
%  kubectl apply -f ./vitural-service/client.yaml -n vc 
deployment.apps/client created
% kubectl apply -f ./vitural-service/deploy.yaml -n vc
deployment.apps/httpd created
deployment.apps/tomcat created
%  kubectl apply -f ./vitural-service/service.yaml -n  vc
service/tomcat-svc created
service/httpd-svc created


%  kubectl -n vc get pods
NAME                      READY   STATUS             RESTARTS      AGE
client-655576cdb9-plqgs   0/1     CrashLoopBackOff   3 (47s ago)   87s
httpd-6bff7d856f-kwh7m    0/1     Error              3 (40s ago)   76s
httpd-6bff7d856f-p5s6v    0/1     CrashLoopBackOff   3 (17s ago)   76s
tomcat-78b8b7c78c-dc2qx   1/1     Running            0             76s
tomcat-78b8b7c78c-dln5f   1/1     Running            0             76s


 Warning  BackOff    37s (x10 over 2m15s)  kubelet            Back-off restarting failed container
https://blog.csdn.net/qq_40017427/article/details/107150970
在deployment 镜像的后面加上命令
command: [ "/bin/bash", "-ce", "tail -f /dev/null" ]

 Warning  BackOff    85s (x26 over 6m30s)   kubelet            Back-off restarting failed container

  %  kubectl -n vc logs  httpd-6bff7d856f-p5s6v 
httpd: can't change directory to '/var/www/index.html': Not a directory

%  kubectl -n vc logs httpd-8b474ccfb-q6xr8   
/bin/sh: can't create /var/tmp/index.html: nonexistent directory

httpd: can't change directory to '/index.html': Not a directory


 % kubectl -n vc get pods -w                   
NAME                      READY   STATUS    RESTARTS   AGE
client-5469d56b7f-9kdz8   1/1     Running   0          8h
httpd-84cd78c884-fttxh    1/1     Running   0          27s
httpd-84cd78c884-wpbn7    1/1     Running   0          24s
tomcat-78b8b7c78c-9jjl7   1/1     Running   0          4m32s
tomcat-78b8b7c78c-lx69c   1/1     Running   0          4m32s
^C%                      

 % kubectl -n vc get svc    
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
httpd-svc    ClusterIP   10.111.99.251   <none>        8080/TCP   8h
tomcat-svc   ClusterIP   10.105.58.245   <none>        8080/TCP   8h

 % kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://10.111.99.251:8080 
hello httpd

% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://httpd-svc:8080    
hello httpd

 % kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://tomcat-svc:8080 
wget: server returned error: HTTP/1.1 404 
command terminated with exit code 1



% kubectl -n vc get endpoints
NAME         ENDPOINTS                                                     AGE
httpd-svc    10.1.4.164:8080,10.1.4.165:8080                               9h
tomcat-svc   10.1.4.162:8080,10.1.4.163:8080                               9h
web-svc      10.1.4.162:8080,10.1.4.163:8080,10.1.4.164:8080 + 1 more...   2m3s
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
hello httpd
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
hello httpd
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 
command terminated with exit code 1
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 
command terminated with exit code 1
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 
command terminated with exit code 1
% kubectl -n vc exec -it client-5469d56b7f-9kdz8 -- wget -q -O - http://web-svc:8080
hello httpd

实现了轮询


%  kubectl -n vc apply -f vitural-service/vs.yaml 
virtualservice.networking.istio.io/web-svc-vs created


% kubectl -n vc get service 
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
httpd-svc    ClusterIP   10.111.99.251    <none>        8080/TCP   9h
tomcat-svc   ClusterIP   10.105.58.245    <none>        8080/TCP   9h
web-svc      ClusterIP   10.108.116.125   <none>        8080/TCP   9m3s


 % kubectl -n vc get VirtualService 
NAME         GATEWAYS   HOSTS         AGE
web-svc-vs              ["web-svc"]   45s

 % istioctl kube-inject -f  vitural-service/client.yaml |kubectl -n vc apply -f -
deployment.apps/client configured

% istioctl kube-inject -f  vitural-service/deploy.yaml |kubectl -n vc apply -f -
deployment.apps/httpd configured
deployment.apps/tomcat configured

% kubectl get pods -n vc 
NAME                      READY   STATUS    RESTARTS   AGE
client-78b5c977-xb27z     2/2     Running   0          3m48s
httpd-6cbf88b69c-l4dwn    2/2     Running   0          2m43s
httpd-6cbf88b69c-sm5xn    2/2     Running   0          3m28s
tomcat-7d754cc9b4-49sd2   2/2     Running   0          3m28s
tomcat-7d754cc9b4-rqz5t   2/2     Running   0          2m28s

% kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 Not Found
command terminated with exit code 1
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 Not Found
command terminated with exit code 1
xiazemin@xiazemindeMBP 实战 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
hello httpd


host fields
routing rules


    前面优先级更高

 %  kubectl -n vc apply -f vitural-service/vs.headers.yaml
error: error validating "vitural-service/vs.headers.yaml": error validating data: ValidationError(VirtualService.spec.http[0].match[0]): unknown field "end-user" in io.istio.networking.v1alpha3.VirtualService.spec.http.match; if you choose to ignore these errors, turn validation off with --validate=false


https://zhuanlan.zhihu.com/p/262249783

 %  kubectl -n vc apply -f vitural-service/vs.headers.yaml
virtualservice.networking.istio.io/web-svc-vs-header created

https://istio.io/latest/docs/concepts/traffic-management/#virtual-services

 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080
wget: server returned error: HTTP/1.1 404 Not Found
command terminated with exit code 1

 % kubectl -n vc exec -it client-78b5c977-xb27z -- wget -q -O - http://web-svc:8080 --header 'end-user:xiazemin'
hello httpd

