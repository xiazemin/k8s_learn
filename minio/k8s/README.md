https://www.cnblogs.com/dai-zhe/p/14741905.html

https://artifacthub.io/packages/helm/minio/minio

https://blog.51cto.com/u_4988084/2461119

https://blog.csdn.net/weixin_51954021/article/details/112790204

https://blog.csdn.net/networken/article/details/111469223


基于官方helm chat进行部署，安装 MinIO chart

helm repo add minio https://helm.min.io/
1
独立模式部署minio，使用deployment方式部署单个pod：

helm install minio \
  --namespace minio --create-namespace \
  --set accessKey=minio,secretKey=minio123 \
  --set mode=standalone \
  --set service.type=NodePort \
  --set persistence.enabled=true \
  --set persistence.size=10Gi \
  --set persistence.storageClass=longhorn \
  minio/minio


Operator方式部署
MinIO Operator为Kubernetes带来了对MinIO，图形控制台和加密的原生支持。

参考：https://github.com/minio/operator

kubectl安装 krew插件

(
  set -x; cd "$(mktemp -d)" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" &&
  tar zxvf krew.tar.gz &&
  KREW=./krew-"$(uname | tr '[:upper:]' '[:lower:]')_$(uname -m | sed -e 's/x86_64/amd64/' -e 's/arm.*$/arm/')" &&
  "$KREW" install krew
)


http://docs.minio.org.cn/docs/master/deploy-minio-on-kubernetes


https://www.jianshu.com/p/2d45990dd652

https://www.cnblogs.com/rongfengliang/p/14967833.html



如果一个参数的选项格式是[],比如
-H=[]host
-p=[]portdirection
这都意味着这个flag可以多次出现，所以此处可以多次指定端口映射规则。

例：docker run -d -p 80:80 -p 22:22

https://blog.csdn.net/asdfgh0077/article/details/106438150

https://blog.csdn.net/qq_30038111/article/details/111823184



https://wiki.shileizcc.com/confluence/pages/viewpage.action?pageId=2523139
name：设置端口名称，必须 Pod 内唯一，当只配置一个端口的时候，这是一个可选项，当配置多个端口的时候，这是一个必须项。
containerPort：必须项，设置在容器内的端口，有效范围 1 ~ 65535。
protocol：可选项，设置端口的协议，TCP 或者 UDP，默认是 TCP。
hostIP：可选项，是指在宿主机上的IP，默认绑定到所有可用的 IP 上，即 0.0.0.0 。
hostPort：可选项，设置在宿主机上的端口，如果设置则进行端口映射，有效范围 1 ~ 65535。


% curl -iv http://127.0.0.1:30000
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 30000 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:30000
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 403 Forbidden
HTTP/1.1 403 Forbidden
< Accept-Ranges: bytes
Accept-Ranges: bytes
< Content-Length: 226
Content-Length: 226
< Content-Security-Policy: block-all-mixed-content
Content-Security-Policy: block-all-mixed-content
< Content-Type: application/xml
Content-Type: application/xml
< Server: MinIO
Server: MinIO
< Strict-Transport-Security: max-age=31536000; includeSubDomains
Strict-Transport-Security: max-age=31536000; includeSubDomains
< Vary: Origin
Vary: Origin
< Vary: Accept-Encoding
Vary: Accept-Encoding
< X-Amz-Request-Id: 16A012A7AEDC24F7
X-Amz-Request-Id: 16A012A7AEDC24F7
< X-Content-Type-Options: nosniff
X-Content-Type-Options: nosniff
< X-Xss-Protection: 1; mode=block
X-Xss-Protection: 1; mode=block
< Date: Mon, 30 Aug 2021 11:39:36 GMT
Date: Mon, 30 Aug 2021 11:39:36 GMT

< 
<?xml version="1.0" encoding="UTF-8"?>
* Connection #0 to host 127.0.0.1 left intact
<Error><Code>AccessDenied</Code><Message>Access Denied.</Message><Resource>/</Resource><RequestId>16A012A7AEDC24F7</RequestId><HostId>bd78d817-8f16-4589-a8fd-e7db8569033e</HostId></Error>* Closing connection 0


说明api接口是ok的

 % curl -iv http://127.0.0.1:30001
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 30001 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:30001
> User-Agent: curl/7.64.1
> Accept: */*
> 
* Empty reply from server
* Connection #0 to host 127.0.0.1 left intact
curl: (52) Empty reply from server
* Closing connection 0
console接口不ok


http://www.linuxea.com/1942.html
https://selinux.cn/k8s-yaml-port/


        args:
        - server
        - /storage
        - --console-address
        - ":9001"

http://127.0.0.1:30001/dashboard