https://zhuanlan.zhihu.com/p/404764407

检查是否开启了动态准入控制

# 获取apiserver pod名字
apiserver_pod_name=`kubectl get --no-headers=true po -n kube-system | grep kube-apiserver | awk '{ print $1 }'`
# 查看api server的启动参数plugin
kubectl get po $apiserver_pod_name -n kube-system -o yaml | grep plugin
如果输出如下，说明已经开启

- --enable-admission-plugins=NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook


% apiserver_pod_name=`kubectl get --no-headers=true po -n kube-system | grep kube-apiserver | awk '{ print $1 }'`
% kubectl get po $apiserver_pod_name -n kube-system -o yaml | grep plugin
    - --enable-admission-plugins=NodeRestriction

否则，需要修改启动参数，请不然直接修改Pod的参数，这样修改不会成功，请修改配置文件
/etc/kubernetes/manifests/kube-apiserver.yaml，加上相应的插件参数后保存，APIServer的Pod会监控该文件的变化，然后重新启动。


https://github.com/morvencao/kube-mutating-webhook-tutorial

https://zhuanlan.zhihu.com/p/404764407


kubeadm 安装的kubernetes修改apiserver参数如下

find / -name kube-apiserver.yaml

vi /etc/kubernetes/manifests/kube-apiserver.yaml
# kubeadm安装的apiserver是Static Pod，它的配置文件被修改后，立即生效。

# Kubelet 会监听该文件的变化，当您修改了 /etc/kubenetes/manifest/kube-apiserver.yaml 文件之后，kubelet 将自动终止原有的 kube-apiserver-{nodename} 的 Pod，并自动创建一个使用了新配置参数的 Pod 作为替代。
# 如果您有多个 Kubernetes Master 节点，您需要在每一个 Master 节点上都修改该文件，并使各节点上的参数保持一致。


https://blog.csdn.net/xujiamin0022016/article/details/108249660


https://www.songma.com/news/txtlist_i65976v.html


 % kubectl -n kube-system describe pod kube-apiserver-docker-desktop


 % kubectl -n kube-system edit pod kube-apiserver-docker-desktop
error: pods "kube-apiserver-docker-desktop" is invalid
A copy of your changes has been stored to "/var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-4jmn0.yaml"
error: Edit cancelled, no valid changes were saved.

  - --enable-admission-plugins=NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook



docker run -it --privileged --pid=host alpine:latest nsenter -t 1 -m -u -n -i sh

vi /etc/kubernetes/manifests/kube-apiserver.yaml


 % kubectl get po $apiserver_pod_name -n kube-system -o yaml | grep plugin
    - --enable-admission-plugins=NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook


 % kubectl api-versions | grep admissionregistration.k8s.io
admissionregistration.k8s.io/v1
admissionregistration.k8s.io/v1beta1



K8S集群默认是HTTPS通信的，所以APiserver调用webhook的过程也是HTTPS的，所以需要进行证书认证，证书认证相当于是给Service的域名进行认证（Service后面会创建），将Service域名放到认证请求server.csr文件中，然后创建一个K8S证书签署请求资源CertificateSigningRequest，APIServer签署该证书后生成server-cert.pem，再将最初创建的私钥server-key.pem和签署好的证书server-cert.pem放到Secret中供Deployment调用，详细过程看脚本
webhook-create-signed-cert.sh

% sh ./deploy/webhook-create-signed-cert.sh
creating certs in tmpdir /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/tmp.HdNUigEt 
Generating RSA private key, 2048 bit long modulus (2 primes)
.................................+++++
..................+++++
e is 65537 (0x010001)
Warning: certificates.k8s.io/v1beta1 CertificateSigningRequest is deprecated in v1.19+, unavailable in v1.22+; use certificates.k8s.io/v1 CertificateSigningRequest
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default created
NAME                                   AGE   SIGNERNAME                     REQUESTOR            CONDITION
sidecar-injector-webhook-svc.default   0s    kubernetes.io/legacy-unknown   docker-for-desktop   Pending
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default approved
W1212 00:16:54.025092   58797 helpers.go:557] --dry-run is deprecated and can be replaced with --dry-run=client.
secret/sidecar-injector-webhook-certs created



 % kubectl get CertificateSigningRequest
NAME                                   AGE     SIGNERNAME                     REQUESTOR            CONDITION
sidecar-injector-webhook-svc.default   4m33s   kubernetes.io/legacy-unknown   docker-for-desktop   Approved,Issued


 % kubectl apply -f deploy/deployment.yaml
Error from server (NotFound): error when creating "deploy/deployment.yaml": namespaces "sidecar-injector" not found

% kubectl apply -f deploy/deployment.yaml
deployment.apps/sidecar-injector-webhook-deployment created


% kubectl logs sidecar-injector-webhook-deployment-7689959d4b-hkq9z
Error from server (BadRequest): container "sidecar-injector" in pod "sidecar-injector-webhook-deployment-7689959d4b-hkq9z" is waiting to start: ContainerCreating




% kubectl get secret
NAME                                  TYPE                                  DATA   AGE
basic-auth                            Opaque                                1      108d
default-token-cg2vq                   kubernetes.io/service-account-token   3      110d
ingress-nginx-admission               Opaque                                3      108d
ingress-nginx-admission-token-xblnc   kubernetes.io/service-account-token   3      105d
ingress-nginx-token-7clh8             kubernetes.io/service-account-token   3      105d
sidecar-injector-webhook-certs        Opaque                                2      11m
tls-secret                            kubernetes.io/tls                     2      108d


% kubectl get serviceAccount
NAME                      SECRETS   AGE
default                   1         110d
ingress-nginx             1         105d
ingress-nginx-admission   1         105d


go build -o webhook .

% docker build -t webhook-example .

% kubectl apply -f deploy/deployment-simple.yaml
Error from server (InternalError): an error on the server ("") has prevented the request from succeeding

https://blog.csdn.net/qq_35396734/article/details/120247811



 % kubectl describe CertificateSigningRequest  sidecar-injector-webhook-svc.default
Name:               sidecar-injector-webhook-svc.default
Labels:             <none>
Annotations:        <none>
CreationTimestamp:  Sun, 12 Dec 2021 00:16:53 +0800
Requesting User:    docker-for-desktop
Signer:             kubernetes.io/legacy-unknown
Status:             Approved,Issued
Subject:
  Common Name:    sidecar-injector-webhook-svc.default.svc
  Serial Number:
Subject Alternative Names:
         DNS Names:  sidecar-injector-webhook-svc
                     sidecar-injector-webhook-svc.default
                     sidecar-injector-webhook-svc.default.svc
Events:  <none>



/ # ls /etc/kubernetes/pki/
apiserver-etcd-client.crt     apiserver-kubelet-client.key  ca.crt                        front-proxy-ca.crt            front-proxy-client.key
apiserver-etcd-client.key     apiserver.crt                 ca.key                        front-proxy-ca.key            sa.key
apiserver-kubelet-client.crt  apiserver.key                 etcd                          front-proxy-client.crt        sa.pub

% docker pull morvencao/sidecar-injector:latest
latest: Pulling from morvencao/sidecar-injector
aad63a933944: Pull complete 
ae33b7fcc866: Pull complete 
ca35ba41b554: Pull complete 
9d99eb3e63f4: Pull complete 
Digest: sha256:1eaf41352646627a78c6b4c7fcf948ded6f41bed5d34650fe6c678d36291d174
Status: Downloaded newer image for morvencao/sidecar-injector:latest
docker.io/morvencao/sidecar-injector:latest


% kubectl apply -f deploy/deployment.yaml 
deployment.apps/sidecar-injector-webhook-deployment created

% kubectl logs sidecar-injector-webhook-deployment-5857898d85-s266h
Error from server (BadRequest): container "sidecar-injector" in pod "sidecar-injector-webhook-deployment-5857898d85-s266h" is waiting to start: ContainerCreating