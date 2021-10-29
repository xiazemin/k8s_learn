1.使用当前系统的ca证书认证一个私有证书
# cd /etc/kubernetes/pki/
# openssl genrsa -out dev.key 2048
# openssl req -new -key dev.key -out dev.csr -subj "/CN=dev"
# openssl x509 -req -in dev.csr -CA ./ca.crt -CAkey ./ca.key -CAcreateserial -out dev.crt -days 3650
# openssl x509 -noout -text -in ./dev.crt

2.使用生成的证书创建一个用户
# kubectl config set-credentials dev --client-certificate=./dev.crt --client-key=./dev.key --embed-certs=true
User "dev" set.

3.定义一个context
# kubectl config set-context dev@kubernetes --cluster=kubernetes --user=dev --namespace=development
Context "dev@kubernetes" created.


4.切换context
# kubectl config use-context dev@kubernetes
Switched to context "dev@kubernetes".
# kubectl get pods
NAME                               READY   STATUS    RESTARTS   AGE
nginx-deployment-6dd86d77d-pqndm   1/1     Running   0          22m
nginx-deployment-6dd86d77d-q268r   1/1     Running   0          22m
nginx-deployment-6dd86d77d-zn4f4   1/1     Running   0          22m
# kubectl get pod -n default
Error from server (Forbidden): pods is forbidden: User "dev" cannot list resource "pods" in API group "" in the namespace "default"
# kubectl config use-context kubernetes-admin@kubernetes
Switched to context "kubernetes-admin@kubernetes".



https://zhuanlan.zhihu.com/p/43237959

K8S中有两种用户(User)——服务账号(ServiceAccount)和普通意义上的用户(User)
ServiceAccount是由K8S管理的，而User通常是在外部管理，K8S不存储用户列表——也就是说，添加/编辑/删除用户都是在外部进行，无需与K8S API交互，虽然K8S并不管理用户，但是在K8S接收API请求时，是可以认知到发出请求的用户的，实际上，所有对K8S的API请求都需要绑定身份信息(User或者ServiceAccount)


最主要的区别上面已经说过了，即ServiceAccount是K8S内部资源，而User是独立于K8S之外的


User通常是人来使用，而ServiceAccount是某个服务/资源/程序使用的
User独立在K8S之外，也就是说User是可以作用于全局的，在任何命名空间都可被认知，并且需要在全局唯一
而ServiceAccount作为K8S内部的某种资源，是存在于某个命名空间之中的，在不同命名空间中的同名ServiceAccount被认为是不同的资源
K8S不会管理User，所以User的创建/编辑/注销等，需要依赖外部的管理机制，K8S所能认知的只有一个用户名 ServiceAccount是由K8S管理的，创建等操作，都通过K8S完


角色绑定包含了一组相关主体（即 subject, 包括用户 ——User、用户组 ——Group、或者服务账户 ——Service Account）以及对被授予角色的引用。
https://jimmysong.io/kubernetes-handbook/concepts/rbac.html


尽管K8S认知用户靠的只是用户的名字，但是只需要一个名字就能请求K8S的API显然是不合理的，所以依然需要验证此用户的身份

在K8S中，有以下几种验证方式：

X509客户端证书
客户端证书验证通过为API Server指定--client-ca-file=xxx选项启用，API Server通过此ca文件来验证API请求携带的客户端证书的有效性，一旦验证成功，API Server就会将客户端证书Subject里的CN属性作为此次请求的用户名
静态token文件
通过指定--token-auth-file=SOMEFILE选项来启用bearer token验证方式，引用的文件是一个包含了 token,用户名,用户ID 的csv文件 请求时，带上Authorization: Bearer 31ada4fd-adec-460c-809a-9e56ceb75269头信息即可通过bearer token验证
静态密码文件
通过指定--basic-auth-file=SOMEFILE选项启用密码验证，类似的，引用的文件时一个包含 密码,用户名,用户ID 的csv文件 请求时需要将Authorization头设置为Basic BASE64ENCODED(USER:PASSWORD)




为用户生成证书
假设我们操作的用户名为tom

首先需要为此用户创建一个私钥
openssl genrsa -out tom.key 2048
接着用此私钥创建一个csr(证书签名请求)文件，其中我们需要在subject里带上用户信息(CN为用户名，O为用户组)
openssl req -new -key tom.key -out tom.csr -subj "/CN=tom/O=MGM"
其中/O参数可以出现多次，即可以有多个用户组
找到K8S集群(API Server)的CA证书文件，其位置取决于安装集群的方式，通常会在/etc/kubernetes/pki/路径下，会有两个文件，一个是CA证书(ca.crt)，一个是CA私钥(ca.key)
通过集群的CA证书和之前创建的csr文件，来为用户颁发证书
openssl x509 -req -in tom.csr -CA path/to/ca.crt -CAkey path/to/ca.key -CAcreateserial -out tom.crt -days 365
-CA和-CAkey参数需要指定集群CA证书所在位置，-days参数指定此证书的过期时间，这里为365天
最后将证书(tom.crt)和私钥(tom.key)保存起来，这两个文件将被用来验证API请求



openssl genrsa -out jane.key 2048
openssl req -new -key jane.key -out jane.csr -subj "/CN=jane/O=MGM"
openssl x509 -req -in jane.csr -CA ~/Library/Group\ Containers/group.com.docker/pki/apiserver.crt  -CAkey ~/Library/Group\ Containers/group.com.docker/pki/apiserver.key -out jane.crt  -days 365 


4376067392:error:02001002:system library:fopen:No such file or directory:crypto/bio/bss_file.c:69:fopen('/Users/xiazemin/Library/Group Containers/group.com.docker/pki/apiserver.srl','r')
4376067392:error:2006D080:BIO routines:BIO_new_file:no such file:crypto/bio/bss_file.c:76:
https://blog.csdn.net/lipviolet/article/details/109456104

-CAcreateserial——表示创建证书序列号文件(即上方提到的serial文件)，创建的序列号文件默认名称为-CA，指定的证书名称后加上.srl后缀

在x509指令中，有多种方式可以指定一个将要生成证书的序列号，可以使用set_serial选项来直接指定证书的序列号，也可以使用-CAserial选项来指定一个包含序列号的文件。所谓的序列号是一个包含一个十六进制正整数的文件，在默认情况下，该文件的名称为输入的证书名称加上.srl后缀，比如输入的证书文件为ca.cer，那么指令会试图从ca.srl文件中获取序列号，可以自己创建一个ca.srl文件，也可以通过-CAcreateserial选项来生成一个序列号文件。

openssl x509 -req -in jane.csr -CA ~/Library/Group\ Containers/group.com.docker/pki/apiserver.crt  -CAkey ~/Library/Group\ Containers/group.com.docker/pki/apiserver.key -CAcreateserial -out jane.crt  -days 365 

Signature ok
subject=CN = jane, O = MGM
Getting CA Private Key

openssl x509 -req -in jane.csr -CA ~/Library/Group\ Containers/group.com.docker/pki/ca.crt  -CAkey ~/Library/Group\ Containers/group.com.docker/pki/ca.key -CAcreateserial -out jane.crt  -days 365 

问题解决

Docker On MAC K8S证书过期
1.直接删除文件夹(先备份下)
ls ~/Library/Group\ Containers/group.com.docker/pki/
apiserver-etcd-client.crt       apiserver.crt                   etcd                            front-proxy-client.key
apiserver-etcd-client.key       apiserver.key                   front-proxy-ca.crt              sa.key
apiserver-kubelet-client.crt    ca.crt                          front-proxy-ca.key              sa.pub
apiserver-kubelet-client.key    ca.key                          front-proxy-client.crt

2.重启Docker on MAC
证书会自动延续1年
https://blog.csdn.net/yezi1993/article/details/106430390



kubectl describe pod kube-apiserver-docker-desktop  -n kube-system 
    Mounts:
      /etc/ca-certificates from etc-ca-certificates (ro)
      /etc/ssl/certs from ca-certs (ro)
      /run/config/pki from k8s-certs (ro)
      /usr/local/share/ca-certificates from usr-local-share-ca-certificates (ro)
      /usr/share/ca-certificates from usr-share-ca-certificates (ro)

% ls  /run/config/pki
ls: /run/config/pki: No such file or directory 
这是虚拟机的目录，在主机上是看不到的



https://docs.docker.com/desktop/mac/

Add custom CA certificates (server side)
All trusted CAs (root or intermediate) are supported. Docker Desktop creates a certificate bundle of all user-trusted CAs based on the Mac Keychain, and appends it to Moby trusted certificates. So if an enterprise SSL certificate is trusted by the user on the host, it is trusted by Docker Desktop.

To manually add a custom, self-signed certificate, start by adding the certificate to the macOS keychain, which is picked up by Docker Desktop. Here is an example:

 sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ca.crt
Or, if you prefer to add the certificate to your own local keychain only (rather than for all users), run this command instead:

 security add-trusted-cert -d -r trustRoot -k ~/Library/Keychains/login.keychain ca.crt
See also, Directory structures for certificates.

Note: You need to restart Docker Desktop after making any changes to the keychain or to the ~/.docker/certs.d directory in order for the changes to take effect.


Add client certificates
You can put your client certificates in ~/.docker/certs.d/<MyRegistry>:<Port>/client.cert and ~/.docker/certs.d/<MyRegistry>:<Port>/client.key.

When the Docker Desktop application starts, it copies the ~/.docker/certs.d folder on your Mac to the /etc/docker/certs.d directory on Moby (the Docker Desktop xhyve virtual machine).

You need to restart Docker Desktop after making any changes to the keychain or to the ~/.docker/certs.d directory in order for the changes to take effect.

The registry cannot be listed as an insecure registry (see Docker Engine. Docker Desktop ignores certificates listed under insecure registries, and does not send client certificates. Commands like docker run that attempt to pull from the registry produce error messages on the command line, as well as on the registry.

Directory structures for certificates
If you have this directory structure, you do not need to manually add the CA certificate to your Mac OS system login:

/Users/<user>/.docker/certs.d/
└── <MyRegistry>:<Port>
   ├── ca.crt
   ├── client.cert
   └── client.key


The following further illustrates and explains a configuration with custom certificates:

/etc/docker/certs.d/        <-- Certificate directory
└── localhost:5000          <-- Hostname:port
   ├── client.cert          <-- Client certificate
   ├── client.key           <-- Client key
   └── ca.crt               <-- Certificate authority that signed
                                the registry certificate
You can also have this directory structure, as long as the CA certificate is also in your keychain.

/Users/<user>/.docker/certs.d/
└── <MyRegistry>:<Port>
    ├── client.cert
    └── client.key




为kubectl配置用户
tom已经是管理员了，现在我们想要通过kubectl以tom的身份来操作集群，需要将tom的认证信息添加进kubectl的配置，即~/.kube/config中

这里假设config中已经配置好了k8s集群

通过命令
kubectl config set-credentials tom --client-certificate=path/to/tom.crt --client-key=path/to/tom.key
将用户tom的验证信息添加进kubectl的配置
此命令会在配置中添加一个名为tom的用户
kubectl config set-context tom@aliyun --cluster=aliyun --namespace=a-1 --user=tom
此命令添加了一个context配置——设定使用aliyun集群，默认使用a-1命名空间，使用用户tom进行验证
在命令中带上 kubectl --context=tom@aliyun ... 参数即可指定kubectl使用之前添加的名为tom@aliyun的context操作集群
也可以通过命令 kubectl config use-context tom@aliyun 来设置当前激活的context



kubectl config set-credentials jane --client-certificate=jane.crt --client-key=jane.key    
User "jane" set.
vi ~/.kube/config 
可以看到
- name: jane
  user:
    client-certificate: /Users/xiazemin/source/k8s_learn/rbac/user/jane/jane.crt
    client-key: /Users/xiazemin/source/k8s_learn/rbac/user/jane/jane.key
% kubectl cluster-info     
Kubernetes control plane is running at https://kubernetes.docker.internal:6443
CoreDNS is running at https://kubernetes.docker.internal:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

kubectl config set-context docker-desktop@jane --cluster=docker-desktop --user=jane
Context "docker-desktop@jane" created.

kubectl config use-context docker-desktop@jane
Switched to context "docker-desktop@jane".

% kubectl get pods                                                                   
error: You must be logged in to the server (Unauthorized)

认证不对，直接改~/.bube/config 里的认证，改成 docker-desktop 一样的 ，不行，因为被改后就是  docker-desktop这个用户了





将认证信息嵌入kubectl的配置中
通过kubectl config set-credentials命令添加的用户，其默认使用的是引用证书文件路径的方式，表现在~/.kube/config中，就是：

users:
- name: tom
  user:
    client-certificate: path/to/tom.crt
    client-key: path/to/tom.key
如果觉得这样总是带着两个证书文件不方便的话，可以将证书内容直接放到config文件里

将tom.crt/tom.key的内容用BASE64编码
cat tom.crt | base64 --wrap=0
cat tom.key | base64 --wrap=0
将获取的编码后的文本复制进config文件中
users:
- name: ich
  user:
    client-certificate-data: ...
    client-key-data: ...
这样就不再需要证书和私钥文件了，当然这两个文件还是保存起来比较好

https://www.jianshu.com/p/8aadab1c3185



https://github.com/lqshow/notes/issues/43

kubectl --server=https://192.168.99.100:8443 --token=xxx --insecure-skip-tls-verify=true get nodes
curl -k --header "Authorization: Bearer xxx" https://192.168.99.100:8443/api


https://github.com/lqshow/notes/issues/45

CURRENT_CONTEXT=$(kubectl config current-context)\n
CLUSTER_NAME=$(kubectl config get-contexts $CURRENT_CONTEXT | awk '{print $3}' | tail -n 1)
KUBE_APISERVER=$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"$CLUSTER_NAME\")].cluster.server}")
echo $KUBE_APISERVER



https://github.com/Wang-Kai/cherish-today/issues/83

第一性原则看问题
kubectl 本质上是一个与 kube-apiserver 做 7 层通信的客户端工具，因为 kube-apiserver 会做鉴权，所以 kubectl 使用过程中需要 kubeconfig 文件来保存一些配置信息，这些信息包括：

用来验证 kube-apiserver 的 CA 根证书
用来标识 kubectl 管理员的 证书 & 私钥，或者用来标识普通用户的 token
kubeconfig 是一个 YAML 格式的配置文件，其主要字段如下：

clusters 类型为数组，每个元素代表一个 k8s cluster
users 类型为数组，每个元素代表拥有访问权限的用户
contexts 类型为数组，每个元素表示要使用的 cluster & user 组合
current-context 当前正在使用的上下文
kubectl 做多集群的管理，本质上就是把集群的相关参数，以及用户的相关信息全都记录下来，然后通过 context 将其组合，通过 current-context 参数来标明当前正在使用的 context。

2. 向 kubeconfig 中填写参数
该步骤的任务是向 kubeconfig 文件中填写配置信息，主要包含 cluster 、user、context。不需要手动填写，kubectl config 会将我们传入的参数填入到配置文件中去。默认的，kubectl 会操作 $HOME/.kube/config 文件（没有 KUBECONFIG 环境变量做特殊声明的话），我们可以使用 --kubeconfig 参数来指明要操作的 kubeconfig 文件。

对 cluster 做配置

分别对 dev & test 集群配置，填写内容包括 kube-apiserver 的地址端口和集群的根证书路径。--embed-certs=true 参数可以把证书内容内嵌到配置文件中。

kubectl config --kubeconfig=kubeconfig set-cluster ${Cluster_Name} \
--server=${API_Server_Addr} 
--certificate-authority=${CA_PATH}
--embed-certs=true
对 user 做配置

分别对 test-admin & user-admin 做参数配置，填写的内容包括标识 admin 身份的证书和私钥。

kubectl config --kubeconfig=kubeconfig set-credentials ${User_Name} \
--client-certificate=${Client_Cert} \
--client-key=${Client_Private_Key} \
--embed-certs=true
对 context 做配置

context 本质上是 cluster、user、namespace 的组合，其中 namespace 选填，默认为 default。

# 配置 test 集群
kubectl config --kubeconfig=kubeconfig set-context test \
--cluster=test \
--user=test-admin

# 配置 dev 集群
kubectl config --kubeconfig=kubeconfig set-context dev \
--cluster=dev \
--user=dev-admin
选定要使用的集群

切换集群的本质就是改变 current-context 的值，通过 current-context 来标识当前 kubectl 访问的是哪个集群。

kubectl config use-context ${Context_Name}

https://github.com/Wang-Kai/cherish-today/issues/84
https://github.com/knightFly/knightFly-blogs/issues/1

https://github.com/kaybinwong/posts/issues/2

https://github.com/cisen/blog/issues/311

https://github.com/xizhibei/blog/issues/64



配置kubectl客户端通过token方式访问kube-apiserver
使用的变量
本文档用到的变量定义如下：

$ export MASTER_IP=XX.XX.XX.XX # 替换为 kubernetes master VIP
$ export KUBE_APISERVER="https://${MASTER_IP}:6443"
$
创建 kubectl config 文件
$ # 设置集群参数
$ kubectl config set-cluster kubernetes \
  --insecure-skip-tls-verify=true \
  --server=${KUBE_APISERVER} 
$ # 设置客户端认证参数
$ kubectl config set-credentials crd-admin \
 --token=7176d48e4e66ddb3557a82f2dd316a93 
$ # 设置上下文参数
$ kubectl config set-context kubernetes \
  --cluster=kubernetes \
  --user=crd-admin  \
  --namespace=crd 
$ # 设置默认上下文
$ kubectl config use-context kubernetes

https://www.cnblogs.com/tianshifu/p/7841007.html

kubectl get secret -n=kube-system
https://blog.csdn.net/u010063830/article/details/108572391


http://www.itttl.com/blog/docker_for_mac_add_certs.html


https://blog.csdn.net/m0_38112165/article/details/120117648


% kubectl get secret                                                   Error from server (Forbidden): secrets is forbidden: User "jane" cannot list resource "secrets" in API group "" in the namespace "default"
% kubectl get pods   
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     12         65d
ingress-nginx-admission-create-vjn92        0/1     Completed   0          62d
ingress-nginx-admission-patch-wlq6p         0/1     Completed   0          62d
ingress-nginx-controller-57648496fc-84wl8   1/1     Running     19         62d
minio-deployment-55bf5bff5d-cvq7v           1/1     Running     14         60d
redis-f9f74787-tq6tw                        1/1     Running     14         60d
% kubectl get svc 
Error from server (Forbidden): services is forbidden: User "jane" cannot list resource "services" in API group "" in the namespace "default"