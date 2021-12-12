The Certificates API enables automation of X.509 credential provisioning by providing a programmatic interface for clients of the Kubernetes API to request and obtain X.509 certificates from a Certificate Authority (CA).

A CertificateSigningRequest (CSR) resource is used to request that a certificate be signed by a denoted signer, after which the request may be approved or denied before finally being signed.



https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/

https://www.cnblogs.com/yangyuliufeng/p/13548915.html

请求签名流程
CertificateSigningRequest 资源类型允许客户使用它申请发放 X.509 证书。 CertificateSigningRequest 对象 在 spec.request 中包含一个 PEM 编码的 PKCS#10 签名请求。 CertificateSigningRequest 使用 spec.signerName 字段标示 签名者（请求的接收方）。 注意，spec.signerName 在 certificates.k8s.io/v1 之后的 API 版本是必填项。 在 Kubernetes v1.22 和以后的版本，客户可以可选地设置 spec.expirationSeconds 字段来为颁发的证书设定一个特定的有效期。该字段的最小有效值是 600，也就是 10 分钟。

https://kubernetes.io/zh/docs/reference/access-authn-authz/certificate-signing-requests/


Warning: certificates.k8s.io/v1beta1 CertificateSigningRequest is deprecated in v1.19+, unavailable in v1.22+; use certificates.k8s.io/v1 CertificateSigningRequest



certificatesigningrequest.certificates.k8s.io "sidecar-injector-webhook-svc.default" deleted
error: error validating "STDIN": error validating data: ValidationError(CertificateSigningRequest.spec): missing required field "signerName" in io.k8s.api.certificates.v1.CertificateSigningRequestSpec; if you choose to ignore these errors, turn validation off with --validate=false

https://stackoverflow.com/questions/65587904/condition-failed-attempting-to-approve-csrs-with-certificates-k8s-io-v1



signerName

kubernetes.io/kube-apiserver-client: signs certificates that will be honored as client certificates by the API server. Never auto-approved by kube-controller-manager.

Trust distribution: signed certificates must be honored as client certificates by the API server. The CA bundle is not distributed by any other means.
Permitted subjects - no subject restrictions, but approvers and signers may choose not to approve or sign. Certain subjects like cluster-admin level users or groups vary between distributions and installations, but deserve additional scrutiny before approval and signing. The CertificateSubjectRestriction admission plugin is enabled by default to restrict system:masters, but it is often not the only cluster-admin subject in a cluster.
Permitted x509 extensions - honors subjectAltName and key usage extensions and discards other extensions.
Permitted key usages - must include ["client auth"]. Must not include key usages beyond ["digital signature", "key encipherment", "client auth"].
Expiration/certificate lifetime - for the kube-controller-manager implementation of this signer, set to the minimum of the --cluster-signing-duration option or, if specified, the spec.expirationSeconds field of the CSR object.
CA bit allowed/disallowed - not allowed.
kubernetes.io/kube-apiserver-client-kubelet: signs client certificates that will be honored as client certificates by the API server. May be auto-approved by kube-controller-manager.

Trust distribution: signed certificates must be honored as client certificates by the API server. The CA bundle is not distributed by any other means.
Permitted subjects - organizations are exactly ["system:nodes"], common name starts with "system:node:".
Permitted x509 extensions - honors key usage extensions, forbids subjectAltName extensions and drops other extensions.
Permitted key usages - exactly ["key encipherment", "digital signature", "client auth"].
Expiration/certificate lifetime - for the kube-controller-manager implementation of this signer, set to the minimum of the --cluster-signing-duration option or, if specified, the spec.expirationSeconds field of the CSR object.
CA bit allowed/disallowed - not allowed.
kubernetes.io/kubelet-serving: signs serving certificates that are honored as a valid kubelet serving certificate by the API server, but has no other guarantees. Never auto-approved by kube-controller-manager.

Trust distribution: signed certificates must be honored by the API server as valid to terminate connections to a kubelet. The CA bundle is not distributed by any other means.
Permitted subjects - organizations are exactly ["system:nodes"], common name starts with "system:node:".
Permitted x509 extensions - honors key usage and DNSName/IPAddress subjectAltName extensions, forbids EmailAddress and URI subjectAltName extensions, drops other extensions. At least one DNS or IP subjectAltName must be present.
Permitted key usages - exactly ["key encipherment", "digital signature", "server auth"].
Expiration/certificate lifetime - for the kube-controller-manager implementation of this signer, set to the minimum of the --cluster-signing-duration option or, if specified, the spec.expirationSeconds field of the CSR object.
CA bit allowed/disallowed - not allowed.
kubernetes.io/legacy-unknown: has no guarantees for trust at all. Some third-party distributions of Kubernetes may honor client certificates signed by it. The stable CertificateSigningRequest API (version certificates.k8s.io/v1 and later) does not allow to set the signerName as kubernetes.io/legacy-unknown. Never auto-approved by kube-controller-manager.

Trust distribution: None. There is no standard trust or distribution for this signer in a Kubernetes cluster.
Permitted subjects - any
Permitted x509 extensions - honors subjectAltName and key usage extensions and discards other extensions.
Permitted key usages - any
Expiration/certificate lifetime - for the kube-controller-manager implementation of this signer, set to the minimum of the --cluster-signing-duration option or, if specified, the spec.expirationSeconds field of the CSR object.
CA bit allowed/disallowed - not allowed.

https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/


 % sh deploy/webhook-create-signed-cert.sh
creating certs in tmpdir ./tmp 
Generating RSA private key, 2048 bit long modulus (2 primes)
.+++++
..................+++++
e is 65537 (0x010001)
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default created
NAME                                   AGE   SIGNERNAME                      REQUESTOR            CONDITION
sidecar-injector-webhook-svc.default   0s    kubernetes.io/kubelet-serving   docker-for-desktop   Pending
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default approved


ERROR: After approving csr sidecar-injector-webhook-svc.default, the signed certificate did not appear on the resource. Giving up after 10 attempts.

为什么需要TLS验证：

首先API服务器与其他程序通信需要保证安全性，所以他们之间要https加密通信。
webhook程序跟API服务器通信的时候可以理解为webhook程序是服务端，所以它要产生公私匙，将公匙交给客户端，让其加密后再发送数据。但是，客户端不知道服务端发公匙是不是伪造的，所以需要一个第三方机构CA进行证书签名。证书签名是使用CA的私匙加密，然后客户端用CA的公匙解密验证证书的真伪。
关于证书这块，可以两种方式实现，分为：

由k8s自带的CA进行签名，因为这样api服务器就可以用自带CA的公匙进行验证。
自己通过cfssl签发

https://segmentfault.com/a/1190000022501045?sort=newest


% kubectl get csr
NAME                                   AGE     SIGNERNAME                      REQUESTOR            CONDITION
sidecar-injector-webhook-svc.default   3m46s   kubernetes.io/kubelet-serving   docker-for-desktop   Approved,Failed


https://github.com/kubeflow/pipelines/issues/4505
kubernetes.io/kube-apiserver-client 替换 kubernetes.io/kubelet-serving

% sh deploy/webhook-create-signed-cert.sh
creating certs in tmpdir ./tmp 
Generating RSA private key, 2048 bit long modulus (2 primes)
......................................+++++
..............................+++++
e is 65537 (0x010001)
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default created
NAME                                   AGE   SIGNERNAME                            REQUESTOR            CONDITION
sidecar-injector-webhook-svc.default   0s    kubernetes.io/kube-apiserver-client   docker-for-desktop   Pending
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.default approved
ERROR: After approving csr sidecar-injector-webhook-svc.default, the signed certificate did not appear on the resource. Giving up after 10 attempts.


https://www.kubeflow.org/docs/components/pipelines/overview/caching-v2/

https://zhuanlan.zhihu.com/p/137070531

apiVersion: kubeadm.k8s.io/v1beta3
kind: ClusterConfiguration
controllerManager:
  extraArgs:
    cluster-signing-cert-file: /etc/kubernetes/pki/ca.crt
    cluster-signing-key-file: /etc/kubernetes/pki/ca.key

https://stackoverflow.com/questions/59795325/approved-kubernetes-csr-but-certificate-not-shown-in-status

https://dev.to/ineedale/writing-a-very-basic-kubernetes-mutating-admission-webhook-5b1

https://medium.com/ibm-cloud/diving-into-kubernetes-mutatingadmissionwebhook-6ef3c5695f74

改回certificates.k8s.io/v1beta1
 % ls ./tmp
csr.conf        server-cert.pem server-key.pem  server.csr

ADD ./tmp/server-cert.pem /etc/webhook/certs/cert.pem
ADD ./tmp/server-key.pem /etc/webhook/certs/cert.pem

 % docker build -t webhook-example:0.0.1  .
 => => exporting layers                                                                                                        0.0s
 => => writing image sha256:b0306a7171ed4d0d4762012003d04b5710087e7ca017413086ab3804a3e4deec                                   0.0s
 => => naming to docker.io/library/webhook-example:0.0.1  


  % ./webhook --tlsCertFile=./tmp/server-cert.pem --tlsKeyFile= ./tmp/server-key.pem 
E1212 09:19:32.493991   81847 main.go:26] Failed to load key pair: open : no such file or directory

 % go build -o webhook .
  ./webhook --tlsCertFile=./tmp/server-cert.pem --tlsKeyFile=./tmp/server-key.pem 

  % kubectl apply -f deploy/deployment-simple.yaml
deployment.apps/admission-webhook-example-deployment created


 % kubectl logs admission-webhook-example-deployment-c7bc58df4-g7429
Error from server (BadRequest): container "admission-webhook-example" in pod "admission-webhook-example-deployment-c7bc58df4-g7429" is waiting to start: trying and failing to pull image

webhook-example:0.0.1  改为
docker.io/library/webhook-example:0.0.1 

 % kubectl logs admission-webhook-example-deployment-8cc89c5ff-vjgzn
Error from server (BadRequest): container "admission-webhook-example" in pod "admission-webhook-example-deployment-8cc89c5ff-vjgzn" is waiting to start: trying and failing to pull image

% docker pull docker.io/library/webhook-example:0.0.1
Error response from daemon: pull access denied for webhook-example, repository does not exist or may require 'docker login': denied: requested access to the resource is denied

 % docker logout
Removing login credentials for https://index.docker.io/v1/

Always =>
 imagePullPolicy: Never  

  % kubectl get pod |grep hook
admission-webhook-example-deployment-5d66985494-2llww   0/1     RunContainerError   2          46s

