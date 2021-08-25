https://kubernetes.github.io/ingress-nginx/examples/auth/basic/

cd ../../../apple

docker build -t apple:5678 .

htpasswd -c auth foo

kubectl create secret generic basic-auth --from-file=auth

kubectl get secret basic-auth -o yaml


kubectl apply -f apple.yaml 
```
pod/apple-app created
service/apple-service created
```

kubectl get svc -o wide
```
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE   SELECTOR
apple-service   ClusterIP   10.105.42.239   <none>        5678/TCP   65s   app=apple
kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP    47h   <none>
```

https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/
#kubectl expose deployment hello-world --type=LoadBalancer --name=my-service




External IP
外部访问Service的方式有两种：

1）通过设置nodePort映射到物理机，同时设置Service的类型为NodePort。

2）通过设置LoadBalancer映射到云服务上提供的LoadBalancer地址。这种用法仅用于公有云服务提供商的云平台设置Service的场景。对该Service的请求将会通过LoadBalancer转发到后端Pod上，负载分发的实现方式则依赖于云服务商提供的LoadBalancer的实现机制。

https://www.cnblogs.com/embedded-linux/p/12657128.html

https://blog.csdn.net/lixinkuan328/article/details/103993274

service 可以独立于deployment工作，但是需要一个个去创建pod，而不是像deployment那样一次性创建。

https://zhuanlan.zhihu.com/p/358916098


nodePort： 使用nodeIp：nodePort 从外部访问请求某个service

targetPort：是pod的端口，从port和nodePort来的流量经过kube-proxy流入到后端pod的targetPort上，最后进入容器

containerPort：是pod内部容器的端口，targetPort映射到containerPort


% kubectl get svc
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
apple-service   NodePort    10.105.42.239   <none>        5678:30080/TCP   29m
kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP          47h

 % curl http://127.0.0.1:30080/apple
/apple%     


kubectl apply -f basicAuth/ingress-with-auth.yaml 

% curl www.xzm.com
<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html>
<head><title>403 Forbidden</title></head>

#https://blog.csdn.net/woshizhangliang999/article/details/108762012


curl www.xzm.com/apple -u 'foo:bar' -H 'www.xzm.com'

curl http://foo:bar@www.xzm.com/apple


kubectl apply -f ../noAuth/ingress.yaml 
ingress.networking.k8s.io/ingress-with-auth configured

 
xiazemin@localhost auth % curl www.xzm.com
<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
<html>
<head><title>403 Forbidden</title></head>

原因：这个域名是真的域名，改下curl www.xzm-text789.com
curl: (6) Could not resolve host: www.xzm-text789.com

https://kubernetes.io/zh/docs/tasks/access-application-cluster/ingress-minikube/

https://www.bookstack.cn/read/feiskyer-kubernetes-handbook-202005/practice-minikube-ingress.md


https://stackoverflow.com/questions/48244233/starting-an-ingress-service-on-docker-for-mac

https://github.com/jnewland/local-dev-with-docker-for-mac-kubernetes


kubectl apply -f ingress-with-auth.yaml

Error from server (InternalError): error when creating "ingress-with-auth.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": an error on the server ("") has prevented the request from succeeding

实际上，我只是发现问题在于重新安装minikube时，没有删除Validation Webhook并因此创建问题，应使用以下命令将其删除.
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission 
https://www.kancloud.cn/panxin20/notes/1743165
https://www.it1352.com/2238346.html

https://github.com/kubernetes/ingress-nginx/issues/5401


kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted

curl 127.0.0.1
404 page not found

curl 127.0.0.1/apple
/apple%                                    
auth没有生效，所以不能干掉hook

https://www.cnblogs.com/tylerzhou/p/11080653.html
https://blog.csdn.net/weixin_26748959/article/details/109122462
https://www.cnblogs.com/oolo/p/11778727.html
https://www.bookstack.cn/read/istio-1.4-zh/1a476a06d9d1e03e.md

https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#prerequisites


https://kubernetes.github.io/ingress-nginx/troubleshooting/

% kubectl exec -it ingress-nginx-controller-57648496fc-bz5qb  -- cat /etc/nginx/nginx.conf

https://gitmemory.com/issue/kubernetes/ingress-nginx/5968/782092413


验证 $HOME/.kube/config 文件是否包含有效证书，并 在必要时重新生成证书。在 kubeconfig 文件中的证书是 base64 编码的。 该 base64 -d 命令可以用来解码证书，openssl x509 -text -noout 命令 可以用于查看证书信息。
https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/troubleshooting-kubeadm/



https://github.com/kubernetes/ingress-nginx/issues/5968

CA=$(kubectl get secret ingress-nginx-admission -ojsonpath='{.data.ca}')
kubectl patch validatingwebhookconfigurations ingress-nginx-admission --type='json' -p='[{"op": "add", "path": "/webhooks/0/clientConfig/caBundle", "value":"'$CA'"}]'

validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission patched

 % kubectl apply -f ingress-with-auth.yaml                                 
Error from server (InternalError): error when creating "ingress-with-auth.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": an error on the server ("") has prevented the request from succeeding


curl -O  https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.0/deploy/static/provider/aws/deploy-tls-termination.yaml


kubectl describe ValidatingWebhookConfiguration ingress-nginx-admission

https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#prerequisites

https://kubernetes.io/docs/concepts/services-networking/ingress/

https://stackoverflow.com/questions/61365202/nginx-ingress-service-ingress-nginx-controller-admission-not-found



kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted


https://stackoverflow.com/questions/59864497/basic-auth-doesnt-work-in-kubernetes-ingress
https://stackoverflow.com/questions/62720148/nginx-ingress-with-basic-authentication-breaks-preflight-requests

https://stackoverflow.com/questions/62689705/cant-configure-kubernetes-nginx-ingress-basic-auth

https://longhorn.io/docs/1.1.2/deploy/accessing-the-ui/longhorn-ingress/
https://linkerd.io/2.9/tasks/exposing-dashboard/

https://kubernetes.io/zh/docs/concepts/services-networking/ingress/



为啥没有密码也可以登陆？
https://cloud.tencent.com/developer/article/1630900

https://blog.csdn.net/hxpjava1/article/details/79580334


% kubectl get secret basic-auth -o yaml           
apiVersion: v1
data:
  auth: Zm9vOiRhcHIxJC5acmVoaU5DJDlsbUtzenFsS0Z0MnVFZ1dkR3F0UzEK
kind: Secret
metadata:
  creationTimestamp: "2021-08-25T15:26:20Z"
  name: basic-auth
  namespace: default
  resourceVersion: "128891"
  uid: 9b921b8f-31c1-40a6-be9e-954b35846cae
type: Opaque



https://kubernetes.io/zh/docs/concepts/configuration/secret/

需要重启 ingress controller 才能生效

% curl -iv foo:bar@127.0.0.1/apple
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 80 (#0)
* Server auth using Basic with user 'foo'
> GET /apple HTTP/1.1
> Host: 127.0.0.1
> Authorization: Basic Zm9vOmJhcg==
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 200 OK
HTTP/1.1 200 OK
< Date: Wed, 25 Aug 2021 16:06:13 GMT
Date: Wed, 25 Aug 2021 16:06:13 GMT
< Content-Type: text/plain; charset=utf-8
Content-Type: text/plain; charset=utf-8
< Content-Length: 6
Content-Length: 6
< Connection: keep-alive
Connection: keep-alive

< 
* Connection #0 to host 127.0.0.1 left intact
/apple* Closing connection 0



% curl -iv 127.0.0.1/apple        
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 80 (#0)
> GET /apple HTTP/1.1
> Host: 127.0.0.1
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 401 Unauthorized
HTTP/1.1 401 Unauthorized
< Date: Wed, 25 Aug 2021 16:06:19 GMT
Date: Wed, 25 Aug 2021 16:06:19 GMT
< Content-Type: text/html
Content-Type: text/html
< Content-Length: 172
Content-Length: 172
< Connection: keep-alive
Connection: keep-alive
< WWW-Authenticate: Basic realm="Authentication Required - foo"
WWW-Authenticate: Basic realm="Authentication Required - foo"

< 
<html>
<head><title>401 Authorization Required</title></head>
<body>
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx</center>
</body>
</html>
* Connection #0 to host 127.0.0.1 left intact
* Closing connection 0