Harbor由6个大的模块所组成：

Proxy: Harbor的registry、UI、token services等组件，都处在一个反向代理后边。该代理将来自浏览器、docker clients的请求转发到后端服务上。

Registry: 负责存储Docker镜像，以及处理Docker push/pull请求。因为Harbor强制要求对镜像的访问做权限控制， 在每一次push/pull请求时，Registry会强制要求客户端从token service那里获得一个有效的token。

Core services: Harbor的核心功能，主要包括如下3个服务:
1）UI: 作为Registry Webhook, 以图像用户界面的方式辅助用户管理镜像。
2) WebHook：WebHook是在registry中配置的一种机制， 当registry中镜像发生改变时，就可以通知到Harbor的webhook endpoint。Harbor使用webhook来更新日志、初始化同步job等。
3) Token 服务：负责根据用户权限给每个docker push/pull命令签发token. Docker 客户端向Regiøstry服务发起的请求,如果不包含token，会被重定向到这里，获得token后再重新向Registry进行请求。

 Database：为core services提供数据库服务，负责储存用户权限、审计日志、Docker image分组信息等数据。
Job services: 主要用于镜像复制，本地镜像可以被同步到远程Harbor实例上。

Log collector: 负责收集其他组件的日志到一个地方

这里我们与上面运行的7个容器对比，对harbor-adminserver感觉有些疑虑。其实这里harbor-adminserver主要是作为一个后端的配置数据管理，并没有太多的其他功能。harbor-ui所要操作的所有数据都通过harbor-adminserver这样一个数据配置管理中心来完成。

Harbor由7个容器组件所组成：

proxy: 通过nginx服务器来做反向代理

registry: docker官方发布的一个仓库镜像组件

ui: 整个架构的核心服务。该容器是Harbor工程的主要部分

adminserver: 作为Harbor工程的配置数据管理器使用

mysql: 通过官方Mysql镜像创建的数据库容器

job services: 通过状态机的形式将镜像复制到远程Harbor实例。镜像删除同样也可以被同步到远程Harbor实例中。

log: 运行rsyslogd的容器，主要用于收集其他容器的日志

https://www.cnblogs.com/wxwgk/p/13287336.html
https://github.com/goharbor/harbor

With support for both container images and Helm charts,

https://github.com/goharbor/harbor-helm

https://helm.sh/zh/docs/intro/install/

wget https://get.helm.sh/helm-v3.8.2-linux-amd64.tar.gz

https://github.com/helm/helm/releases

# tar -zxvf helm-v3.8.2-linux-amd64.tar.gz linux-amd64/
linux-amd64/
linux-amd64/helm
linux-amd64/LICENSE
linux-amd64/README.md

chmod +x linux-amd64/helm

 sudo cp linux-amd64/helm /usr/local/bin/
 https://www.cnblogs.com/wxx999/p/15982010.html
 https://helm.sh/zh/docs/intro/install/


 https://github.com/goharbor/harbor-helm

 # helm repo add harbor https://helm.goharbor.io

"harbor" has been added to your repositories

helm install my-harbor harbor/harbor
NAME: my-harbor
LAST DEPLOYED: Thu Apr 28 22:27:56 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Please wait for several minutes for Harbor deployment to complete.
Then you should be able to visit the Harbor portal at https://core.harbor.domain
For more details, please visit https://github.com/goharbor/harbor

# helm uninstall my-harbor
These resources were kept due to the resource policy:
[PersistentVolumeClaim] my-harbor-chartmuseum
[PersistentVolumeClaim] my-harbor-jobservice
[PersistentVolumeClaim] my-harbor-registry

release "my-harbor" uninstalled

# kubectl create ns harbor
namespace/harbor created

# helm install my-harbor harbor/harbor --set externalURL=http://124.220.185.35:30002 -n harbor
NAME: my-harbor
LAST DEPLOYED: Thu Apr 28 22:38:02 2022
NAMESPACE: harbor
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Please wait for several minutes for Harbor deployment to complete.
Then you should be able to visit the Harbor portal at http://124.220.185.35:30002
For more details, please visit https://github.com/goharbor/harbor

https://www.jianshu.com/p/69ac775f1fda

https://www.cnblogs.com/wangxu01/articles/11714308.html

harbor 核心服务外部访问 URL。主要用于：
# 1) 补全 portal 页面上面显示的 docker/helm 命令
# 2) 补全返回给 docker/notary 客户端的 token 服务 URL

# 格式：protocol://domain[:port]。
# 1) 如果 expose.type=ingress，"domain"的值就是 expose.ingress.hosts.core 的值 
# 2) 如果 expose.type=clusterIP，"domain"的值就是 expose.clusterIP.name 的值
# 3) 如果 expose.type=nodePort，"domain"的值就是 k8s 节点的 IP 地址

# 如果在代理后面部署 Harbor，请将其设置为代理的 URL
externalURL: https://harbor.wangxu.com


# helm install my-harbor harbor/harbor --set domain=124.220.185.35 --set expose.type=nodePort  -n harbor
Error: INSTALLATION FAILED: execution error at (harbor/templates/nginx/secret.yaml:3:12): The "expose.tls.auto.commonName" is required!

https://blog.51cto.com/u_15331726/5195706

https://www.bianchengquan.com/article/511685.html


 helm install my-harbor harbor/harbor --set domain=124.220.185.35 --set expose.type=nodePort  --set expose.tls.auto.commonName=harbor  -n harbor
NAME: my-harbor
LAST DEPLOYED: Thu Apr 28 22:50:14 2022
NAMESPACE: harbor
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Please wait for several minutes for Harbor deployment to complete.
Then you should be able to visit the Harbor portal at https://core.harbor.domain
For more details, please visit https://github.com/goharbor/harbor


# kubectl -n harbor get svc -o wide
NAME                      TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                                     AGE   SELECTOR
harbor                    NodePort    10.100.102.203   <none>        80:30002/TCP,443:30003/TCP,4443:30004/TCP   44s   app=harbor,component=nginx,release=my-harbor
my-harbor-chartmuseum     ClusterIP   10.96.133.15     <none>        80/TCP                                      44s   app=harbor,component=chartmuseum,release=my-harbor
my-harbor-core            ClusterIP   10.104.231.102   <none>        80/TCP                                      44s   app=harbor,component=core,release=my-harbor
my-harbor-database        ClusterIP   10.107.219.218   <none>        5432/TCP                                    44s   app=harbor,component=database,release=my-harbor
my-harbor-jobservice      ClusterIP   10.102.200.42    <none>        80/TCP                                      44s   app=harbor,component=jobservice,release=my-harbor
my-harbor-notary-server   ClusterIP   10.108.78.165    <none>        4443/TCP                                    44s   app=harbor,component=notary-server,release=my-harbor
my-harbor-notary-signer   ClusterIP   10.102.235.58    <none>        7899/TCP                                    44s   app=harbor,component=notary-signer,release=my-harbor
my-harbor-portal          ClusterIP   10.104.187.228   <none>        80/TCP                                      44s   app=harbor,component=portal,release=my-harbor
my-harbor-redis           ClusterIP   10.97.157.171    <none>        6379/TCP                                    44s   app=harbor,component=redis,release=my-harbor
my-harbor-registry        ClusterIP   10.110.85.123    <none>        5000/TCP,8080/TCP                           44s   app=harbor,component=registry,release=my-harbor
my-harbor-trivy           ClusterIP   10.98.10.231     <none>        8080/TCP                                    44s   app=harbor,component=trivy,release=my-harbor


# curl http://127.0.0.1:30002
<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.19.3</center>
</body>
</html>


https://www.bianchengquan.com/article/511685.html


http://124.220.185.35:30002/


# kubectl -n harbor edit svc harbor
service/harbor edited

spec:
  clusterIP: 10.100.102.203
  externalIPs:
  - 124.220.185.35
  externalTrafficPolicy: Cluster

# kubectl -n harbor get svc -o wide
NAME                      TYPE        CLUSTER-IP       EXTERNAL-IP      PORT(S)                                     AGE   SELECTOR
harbor                    NodePort    10.100.102.203   124.220.185.35   80:30002/TCP,443:30003/TCP,4443:30004/TCP   15m   app=harbor,component=nginx,release=my-harbor

https://www.modb.pro/db/215490

# curl 124.220.185.35:30002
^@^@^@


去腾讯云服务器开启30002端口
https://console.cloud.tencent.com/lighthouse/instance/detail?rid=4&id=lhins-dlsry45q&tab=firewall

# curl 124.220.185.35:30002
<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.19.3</center>
</body>
</html>

https://124.220.185.35:30003/account/sign-in?redirect_url=%2Fharbor%2Fprojects

harbor的默认用户名密码，默认为：admin/Harbor12345

https://blog.csdn.net/weixin_39620001/article/details/110254703

https://www.jianshu.com/p/44391aa234fc

% docker tag tools:0.0.3 124.220.185.35:30003/library/tools:v0.0.1

% docker push 124.220.185.35:30003/library/tools:v0.0.1
The push refers to repository [124.220.185.35:30003/library/tools]
Get "https://124.220.185.35:30003/v2/": x509: cannot validate certificate for 124.220.185.35 because it doesn't contain any IP SANs
这是因为在证书中，要包含一些信息，比如国家、机构等等，好像访问的私有仓库ip或者域名必须要有，否则不予通过，就会报上面的错误。
https://blog.csdn.net/zsd498537806/article/details/79290732


# cd /etc/pki/tls/

# ll
总用量 28
lrwxrwxrwx  1 root root    49 8月  12 2020 cert.pem -> /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem
drwxr-xr-x. 2 root root  4096 4月  28 20:06 certs
-rw-r--r--  1 root root   412 4月  24 2020 ct_log_list.cnf
drwxr-xr-x. 2 root root  4096 4月  24 2020 misc
-rw-r--r--  1 root root 11225 4月  24 2020 openssl.cnf
drwxr-xr-x. 2 root root  4096 4月  24 2020 private

# vi openssl.cnf

[ v3_ca ]
subjectAltName = IP:124.220.185.35

 % docker push 124.220.185.35:30003/library/tools:v0.0.1
The push refers to repository [124.220.185.35:30003/library/tools]
Get "https://124.220.185.35:30003/v2/": x509: cannot validate certificate for 124.220.185.35 because it doesn't contain any IP SANs

# cd certs/

# ls
b5213941.0  ca-bundle.crt  ca-bundle.trust.crt

# sudo openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout registry.key -out registry.crt
Generating a RSA private key
..............+++++
..........................................................+++++
writing new private key to 'registry.key'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [XX]:
State or Province Name (full name) []:
Locality Name (eg, city) [Default City]:
Organization Name (eg, company) [Default Company Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (eg, your name or your server's hostname) []:
Email Address []:


# ls
b5213941.0  ca-bundle.crt  ca-bundle.trust.crt  registry.crt  registry.key


# sudo openssl req -subj '/C=CN/ST=BeiJing/L=HaiDian/CN=<Ipaddress> ' -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout registry.key -out
req: Option -out needs a value
req: Use -help for summary.

# cat ./registry.crt >> /etc/pki/tls/certs/ca-bundle.crt

# systemctl restart docker



1、编辑docker的配置文件（/etc/docker/daemon.json）

  2、 添加"insecure-registries":["192.168.2.250:443"],指定Harbor镜像仓库的地址和端口。

  3、 重启docker服务，问题即可解决。


  % docker push 124.220.185.35:30003/library/tools:v0.0.1
The push refers to repository [124.220.185.35:30003/library/tools]
275387ba739a: Preparing 
275387ba739a: Retrying in 6 seconds 
Service Unavailable

https://www.icode9.com/content-4-1298999.html

https://github.com/goharbor

https://github.com/goharbor/harbor-helm


% docker login 124.220.185.35:30003/library
Username: xiazemin
Password: 
Error response from daemon: Get "https://124.220.185.35:30003/v2/": Get "https://core.harbor.domain/service/token?account=xiazemin&client_id=docker&offline_token=true&service=harbor-registry": Service Unavailable


# kubectl -n harbor edit svc my-harbor-registry
service/my-harbor-registry edited

spec:
  clusterIP: 10.110.85.123
  externalIPs:
  - 124.220.185.35


  type: ClusterIP


]# kubectl -n harbor get svc my-harbor-registry -o wide
NAME                 TYPE       CLUSTER-IP      EXTERNAL-IP      PORT(S)                         AGE   SELECTOR
my-harbor-registry   NodePort   10.110.85.123   124.220.185.35   5000:31236/TCP,8080:31587/TCP   66m   app=harbor,component=registry,release=my-harbor



注意：其他远程服务器docker login登陆harbor需要配置hosts解析和加载证书。
1、配置hosts解析

vim /etc/hosts
124.220.185.35 core.harbor.domain #增加一行域名解析

https://blog.csdn.net/qq_36227153/article/details/119632278


# kubectl edit cm coredns -n kube-system

configmap/coredns edited

 hosts {
    124.220.185.35 core.harbor.domain
    fallthrough
    }

https://www.cnblogs.com/keep-live/p/11395973.html

https://blog.csdn.net/weixin_42518838/article/details/112872075


% ping core.harbor.domain
PING core.harbor.domain (124.220.185.35): 56 data bytes
64 bytes from 124.220.185.35: icmp_seq=0 ttl=52 time=31.139 ms
64 bytes from 124.220.185.35: icmp_seq=1 ttl=52 time=30.973 ms
64 bytes from 124.220.185.35: icmp_seq=2 ttl=52 time=31.472 ms
64 bytes from 124.220.185.35: icmp_seq=3 ttl=52 time=33.681 ms
64 bytes from 124.220.185.35: icmp_seq=4 ttl=52 time=32.800 ms

harbor 主机上
vim /etc/hosts
124.220.185.35 core.harbor.domain #增加一行域名解析

# ping core.harbor.domain
PING core.harbor.domain (124.220.185.35) 56(84) bytes of data.
64 bytes from core.harbor.domain (124.220.185.35): icmp_seq=1 ttl=63 time=2.57 ms
64 bytes from core.harbor.domain (124.220.185.35): icmp_seq=2 ttl=63 time=2.20 ms

https://blog.csdn.net/a1017680279/article/details/108869004

# systemctl restart docker

# vi /etc/docker/daemon.json

 "124.220.185.35:30003"


# systemctl restart docker
Job for docker.service failed because the control process exited with error code.
See "systemctl status docker.service" and "journalctl -xe" for details.

# systemctl status docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; enabled; vendor preset: disabled)
   Active: failed (Result: exit-code) since Fri 2022-04-29 00:26:30 CST; 737ms ago
     Docs: https://docs.docker.com
  Process: 139593 ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock (code=ex>
 Main PID: 139593 (code=exited, status=1/FAILURE)

4月 29 00:26:30 VM-4-15-centos systemd[1]: docker.service: Service RestartSec=2s expired, scheduling restar>
4月 29 00:26:30 VM-4-15-centos systemd[1]: docker.service: Scheduled restart job, restart counter is at 3.
4月 29 00:26:30 VM-4-15-centos systemd[1]: Stopped Docker Application Container Engine.
4月 29 00:26:30 VM-4-15-centos systemd[1]: docker.service: Start request repeated too quickly.
4月 29 00:26:30 VM-4-15-centos systemd[1]: docker.service: Failed with result 'exit-code'.
4月 29 00:26:30 VM-4-15-centos systemd[1]: Failed to start Docker Application Container Engine.

https://blog.csdn.net/weixin_42152531/article/details/120416296

# systemctl stop docker

]# service docker start
Redirecting to /bin/systemctl start docker.service
Job for docker.service failed because the control process exited with error code.
See "systemctl status docker.service" and "journalctl -xe" for details.

https://blog.csdn.net/weixin_42400619/article/details/116580625

# docker info | grep "Docker Root Dir"
errors pretty printing info
[root@VM-4-15-centos certs]# docker info
Client:
 Context:    default
 Debug Mode: false
 Plugins:
  app: Docker App (Docker Inc., v0.9.1-beta3)
  buildx: Build with BuildKit (Docker Inc., v0.5.1-docker)

Server:
ERROR: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
errors pretty printing info

]# service docker start
Redirecting to /bin/systemctl start docker.service

% docker login 124.220.185.35:30003
Username: xiazemin
Password: 
Error response from daemon: login attempt to http://124.220.185.35:30003/v2/ failed with status: 503 Service Unavailable

# vi /etc/docker/daemon.json
https://124.220.185.35:30003
# systemctl restart docker


https://kubesphere.com.cn/forum/d/902-harbor-docker-login
# systemctl restart docker
{
    "registry-mirrors": [
"https://124.220.185.35:30003",
"https://mirror.ccs.tencentyun.com"
    ],
"insecure-registries": [
"124.220.185.35:30003"
]
}


 % docker login 124.220.185.35:30003
Username: admin
Password: 
Error response from daemon: Get "https://124.220.185.35:30003/v2/": Get "https://core.harbor.domain/service/token?account=admin&client_id=docker&offline_token=true&service=harbor-registry": EOF


{
    "registry-mirrors": [
      "https://mirror.ccs.tencentyun.com"
    ],
   "insecure-registries":[
      "124.220.185.35:30003"
    ]
}

# systemctl restart docker


# docker login 124.220.185.35:30003
Username: admin
Password:
Error response from daemon: Get https://124.220.185.35:30003/v2/: Get https://core.harbor.domain/service/token?account=admin&client_id=docker&offline_token=true&service=harbor-registry: dial tcp 124.220.185.35:443: connect: connection refused

# helm uninstall my-harbor -n harbor
These resources were kept due to the resource policy:
[PersistentVolumeClaim] my-harbor-chartmuseum
[PersistentVolumeClaim] my-harbor-jobservice
[PersistentVolumeClaim] my-harbor-registry

release "my-harbor" uninstalled

helm install my-harbor harbor/harbor --set domain=124.220.185.35 --set expose.type=nodePort  --set expose.tls.auto.commonName=harbor --set externalURL=https://124.220.185.35:30003 -n harbor

NAME: my-harbor
LAST DEPLOYED: Fri Apr 29 06:17:14 2022
NAMESPACE: harbor
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Please wait for several minutes for Harbor deployment to complete.
Then you should be able to visit the Harbor portal at https://124.220.185.35:30003
For more details, please visit https://github.com/goharbor/harbor


% docker login 124.220.185.35:30003
Username: admin
Password: 
Login Succeeded

% docker push 124.220.185.35:30003/library/tools:v0.0.1
The push refers to repository [124.220.185.35:30003/library/tools]
275387ba739a: Preparing 
275387ba739a: Pushing  61.51MB/411.5MB
c396586d4164: Pushed 
510f18150319: Pushed 


helm pull harbor/harbor

tar -zxvf harbor-1.9.0.tgz

 spec:
  clusterIP: 10.110.85.123
  externalIPs:
  - 124.220.185.35

https://blog.csdn.net/erhaiou2008/article/details/122034826


 vi ./harbor/values.yaml
 spec:
  externalIPs:
  - 124.220.185.35

spec.externalIPs

vi ./harbor/templates/registry/registry-svc.yaml
vi ./harbor/templates/core/core-svc.yaml
spec:
 externalIPs:
 {{ spec.externalIPs }} 


 helm的安装本地chart
指定本地chart目录：helm install .
指定本地chart压缩包：helm install nginx-1.2.3.tgz
https://blog.csdn.net/qq_34759180/article/details/120456050


# helm install ./harbor
Error: INSTALLATION FAILED: must either provide a name or specify --generate-name
# helm install my-harbor ./harbor
Error: INSTALLATION FAILED: parse error at (harbor/templates/registry/registry-svc.yaml:9): function "spec" not defined


spec:
 externalIPs:
 {{ template "spec.externalIPs" }}


# helm install my-harbor ./harbor
Error: INSTALLATION FAILED: template: harbor/templates/registry/registry-svc.yaml:9:13: executing "harbor/templates/registry/registry-svc.yaml" at <{{template "spec.externalIPs"}}>: template "spec.externalIPs" not defined


./harbor/templates/core/core-cm.yaml:  EXT_ENDPOINT: "{{ .Values.externalURL }}"

 grep -r externalURL ./


 spec:
  externalIPs:
  {{ .Values.spec.externalIPs }}


# helm install my-harbor ./harbor
Error: INSTALLATION FAILED: YAML parse error on harbor/templates/core/core-svc.yaml: error converting YAML to JSON: yaml: line 13: could not find expected ':'


6.6.调试
编写好chart包的模版之后，我们可以给helm命令加上–debug --dry-run 两个参数，让helm输出模版结果，但是不把模版输出结果交给k8s处理。

例子：

#helm install命令类似，加上–debug --dry-run两个参数即可

$ helm upgrade –debug --dry-run -i \

–set replicas=2 \

–set host=www.xxxx.com \

myapp ./myapp

https://blog.csdn.net/weixin_44121790/article/details/119868615


# helm install my-harbor ./harbor --dry-run
Error: INSTALLATION FAILED: YAML parse error on harbor/templates/core/core-svc.yaml: error converting YAML to JSON: yaml: line 13: could not find expected ':'


查看生成的yaml文件

#helm  template  helm_charts-0.1.1.tgz

https://blog.csdn.net/kali_yao/article/details/120900881

]# helm install my-harbor ./harbor --dry-run
Error: INSTALLATION FAILED: YAML parse error on harbor/templates/core/core-svc.yaml: error converting YAML to JSON: yaml: line 13: could not find expected ':'

# helm template ./harbor
Error: YAML parse error on harbor/templates/core/core-svc.yaml: error converting YAML to JSON: yaml: line 13: could not find expected ':'

Use --debug flag to render out invalid YAML

# helm template ./harbor --debug


spec:
  externalIPs:
  - {{ .Values.spec.externalIPs }}



# helm install my-harbor ./harbor
Error: INSTALLATION FAILED: unable to build kubernetes objects from release manifest: error validating "": error validating data: ValidationError(Service.spec.externalIPs[0]): invalid type for io.k8s.api.core.v1.ServiceSpec.externalIPs: got "array", expected "string"

 vi harbor/values.yaml

spec:
  externalIPs:
    124.220.185.35


# helm install my-harbor ./harbor -n harbor
^@^@^@NAME: my-harbor
LAST DEPLOYED: Fri Apr 29 15:47:54 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Please wait for several minutes for Harbor deployment to complete.
Then you should be able to visit the Harbor portal at https://124.220.185.35:30003
For more details, please visit https://github.com/goharbor/harbor

 vi ./harbor/templates/nginx/service.yaml
spec:
  externalIPs:
  - {{ .Values.spec.externalIPs }}

  vi ./harbor/templates/notary/notary-svc.yaml

  vi ./harbor/templates/portal/service.yaml

vi ./harbor/templates/exporter/exporter-svc.yaml


vi ./harbor/templates/nginx/service.yaml


# kubectl -n harbor edit svc harbor

