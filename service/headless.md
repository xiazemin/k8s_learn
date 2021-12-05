headless使用场景

第一种：自主选择权，有时候client想自己来决定使用哪个Real Server，可以通过查询DNS来获取Real Server的信息。

第二种：Headless Service的对应的每一个Endpoints，即每一个Pod，都会有对应的DNS域名；这样Pod之间就能互相访问，集群也能单独访问pod

https://www.cnblogs.com/wuchangblog/p/14032057.html


service的作用，主要是代理一组pod容器负载均衡服务，但是有时候我们不需要这种负载均衡场景，比如下面的两个例子。

比如kubernetes部署某个kafka集群，这种就不需要service来代理，客户端需要的是一组pod的所有的ip。
还有一种场景客户端自己处理负载均衡的逻辑，比如kubernates部署两个mysql，有客户端处理负载请求，或者根本不处理这种负载，就要两套mysql。
基于上面的两个例子，kubernates增加了headless serivces功能，字面意思无service其实就是改service对外无提供IP。

其实一般headless services一般结合StatefulSet来部署有状态的应用

clusterIP: None

https://zhuanlan.zhihu.com/p/114484789


% kubectl get endpoints
NAME                                 ENDPOINTS                                                     AGE
apple-service                        10.1.0.245:5678,10.1.0.249:5678,10.1.0.250:5678 + 9 more...   102d
ingress-nginx-controller             10.1.0.241:443,10.1.0.241:80                                  99d
ingress-nginx-controller-admission   10.1.0.241:8443                                               99d
kubernetes                           192.168.65.4:6443                                             104d
minio-service                        10.1.0.246:9001,10.1.0.246:9000                               97d
redis                                10.1.0.242:6379                                               98d


1.什么是headless service
headless service 是将service的发布文件中的clusterip=none ，不让其获取clusterip ， DNS解析的时候直接走pod



部署常规service
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  labels:
    app: nginx
spec:
  ports:
  - port: 88
    targetPort: 80
  selector:
    app: nginx
  type: NodePort


部署headlessservice

apiVersion: v1
kind: Service
metadata:
  name: headless-service
spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  clusterIP: None


service 的 FQDN： nginx-service.default.svc.cluster.local
headless service的FQDN： headless-service.default.svc.cluster.local
我们在容器里面ping FQDN ， service解析出的地址是clusterip

https://blog.csdn.net/textdemo123/article/details/102954489

https://www.jianshu.com/p/a6d8b28c88a2

 % kubectl apply -f service/headlessservice.yaml
service/apple-headless-service created

% kubectl get svc
NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
apple-headless-service               ClusterIP      None             <none>        5678/TCP                        16s

