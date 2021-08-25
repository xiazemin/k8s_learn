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