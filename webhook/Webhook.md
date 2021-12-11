WebHook 是一种 HTTP 回调：某些条件下触发的 HTTP POST 请求；通过 HTTP POST 发送的简单事件通知。一个基于 web 应用实现的 WebHook 会在特定事件发生时把消息发送给特定的 URL。

https://kubernetes.io/zh/docs/reference/access-authn-authz/webhook/

Webhook 模式需要一个 HTTP 配置文件，通过 --authorization-webhook-config-file=SOME_FILENAME 的参数声明。

配置文件的格式使用 kubeconfig。在文件中，"users" 代表着 API 服务器的 webhook，而 "cluster" 代表着远程服务。



# Kubernetes API 版本   
apiVersion: v1
# API 对象种类
kind: Config
# clusters 代表远程服务。
clusters:
  - name: name-of-remote-authz-service
    cluster:
      # 对远程服务进行身份认证的 CA。
      certificate-authority: /path/to/ca.pem
      # 远程服务的查询 URL。必须使用 'https'。
      server: https://authz.example.com/authorize

# users 代表 API 服务器的 webhook 配置
users:
  - name: name-of-api-server
    user:
      client-certificate: /path/to/cert.pem # webhook plugin 使用 cert
      client-key: /path/to/key.pem          # cert 所对应的 key

# kubeconfig 文件必须有 context。需要提供一个给 API 服务器。
current-context: webhook
contexts:
- context:
    cluster: name-of-remote-authz-service
    user: name-of-api-server
  name: webhook


  在做认证决策时，API 服务器会 POST 一个 JSON 序列化的 authorization.k8s.io/v1beta1 SubjectAccessReview 对象来描述这个动作。这个对象包含了描述用户请求的字段，同时也包含了需要被访问资源或请求特征的具体信息。

需要注意的是 webhook API 对象与其他 Kubernetes API 对象一样都同样都服从版本兼容规则。实施人员应该了解 beta 对象的更宽松的兼容性承诺，同时确认请求的 "apiVersion" 字段能被正确地反序列化。此外，API 服务器还必须启用 authorization.k8s.io/v1beta1 API 扩展组 (--runtime-config=authorization.k8s.io/v1beta1=true)。



请求内容的例子：

{
  "apiVersion": "authorization.k8s.io/v1beta1",
  "kind": "SubjectAccessReview",
  "spec": {
    "resourceAttributes": {
      "namespace": "kittensandponies",
      "verb": "get",
      "group": "unicorn.example.org",
      "resource": "pods"
    },
    "user": "jane",
    "group": [
      "group1",
      "group2"
    ]
  }
}

https://github.com/kubernetes/kubernetes/blob/main/staging/src/k8s.io/apiserver/plugin/pkg/authorizer/webhook/webhook.go

https://blog.csdn.net/u012803274/article/details/115035744

Webhook就是一种HTTP回调，用于在某种情况下执行某些动作，Webhook不是K8S独有的，很多场景下都可以进行Webhook，比如在提交完代码后调用一个Webhook自动构建docker镜像

K8S中提供了自定义资源类型和自定义控制器来扩展功能，还提供了动态准入控制，其实就是通过Webhook来实现准入控制，分为两种：验证性质的准入 Webhook （Validating Admission Webhook） 和 修改性质的准入 Webhook （Mutating Admission Webhook）

Admission Webhook有哪些使用场景？如下

在资源持久化到ETCD之前进行修改（Mutating Webhook），比如增加init Container或者sidecar Container
在资源持久化到ETCD之前进行校验（Validating Webhook），不满足条件的资源直接拒绝并给出相应信息
现在非常火热的的 Service Mesh 应用istio就是通过 mutating webhooks 来自动将Envoy这个 sidecar 容器注入到 Pod 中去的：


Webhook可以理解成Java Web开发中的Filter，每个请求都会经过Filter处理，从图中可以看到，先执行的是Mutating Webhook，它可以对资源进行修改，然后执行的是Validating Webhook，它可以拒绝或者接受请求，但是它不能修改请求

K8S中有已经实现了的Admission Webhook列表，详情参考每个准入控制器的作用是什么？

从0到1开发K8S Webhook最佳实践
我们以一个简单的Webhook作为例子，该Webhook会在创建Deployment资源的时候检查它是否有相应的标签，如果没有的话，则加上（Mutating Webhook），然后在检验它是否有相应的标签（Validating Webhook），有则创建该Deployment，否则拒绝并给出相应错误提示

https://zhuanlan.zhihu.com/p/404764407



https://github.com/cnych/admission-webhook-example

https://github.com/morvencao/kube-mutating-webhook-tutorial

https://github.com/kubernetes/kubernetes/blob/release-1.21/test/images/agnhost/webhook/main.go
https://github.com/morvencao/kube-mutating-webhook-tutorial

https://blog.csdn.net/weixin_38320674/article/details/107148403

https://blog.csdn.net/weixin_43114954/article/details/120627868

前提条件
k8s版本需至少v1.9
确保启用了 MutatingAdmissionWebhook and ValidatingAdmissionWebhook admission controllers
确定 启用了http://admissionregistration.k8s.io/v1beta1


webhook处理apiservers发送的AdmissionReview请求，并将其决定作为AdmissionReview对象发送回去。


动态配置admission webhooks
您可以通过ValidatingWebhookConfiguration或MutatingWebhookConfiguration动态配置哪些资源受入口webhooks的限制。


最后我们来总结下 webhook Admission 的优势：

webhook 可动态扩展 Admission 能力，满足自定义客户的需求
不需要重启 API Server，可通过创建 webhook configuration 热加载 webhook admission

https://zhuanlan.zhihu.com/p/407074852

https://kubernetes.io/zh/docs/reference/access-authn-authz/extensible-admission-controllers/

示例准入 Webhook 服务器置 ClientAuth 字段为 空， 默认为 NoClientCert 。这意味着 webhook 服务器不会验证客户端的身份，认为其是 apiservers。 如果你需要双向 TLS 或其他方式来验证客户端，请参阅 如何对 apiservers 进行身份认证


对 apiservers 进行身份认证
如果你的 webhook 需要身份验证，则可以将 apiserver 配置为使用基本身份验证、持有者令牌或证书来向 webhook 提供身份证明。完成此配置需要三个步骤。

启动 apiserver 时，通过 --admission-control-config-file 参数指定准入控制配置文件的位置。

在准入控制配置文件中，指定 MutatingAdmissionWebhook 控制器和 ValidatingAdmissionWebhook 控制器应该读取凭据的位置。 凭证存储在 kubeConfig 文件中（是​​的，与 kubectl 使用的模式相同），因此字段名称为 kubeConfigFile。


URL 
url 以标准 URL 形式给出 webhook 的位置（scheme://host:port/path）。

host 不应引用集群中运行的服务；通过指定 service 字段来使用服务引用。 主机可以通过某些 apiserver 中的外部 DNS 进行解析。 （例如，kube-apiserver 无法解析集群内 DNS，因为这将违反分层规则）。host 也可以是 IP 地址。

请注意，将 localhost 或 127.0.0.1 用作 host 是有风险的， 除非你非常小心地在所有运行 apiserver 的、可能需要对此 webhook 进行调用的主机上运行。这样的安装方式可能不具有可移植性，即很难在新集群中启用。

scheme 必须为 "https"；URL 必须以 "https://" 开头。

使用用户或基本身份验证（例如："user:password@"）是不允许的。 使用片段（"#..."）和查询参数（"?..."）也是不允许的。

这是配置为调用 URL 的修改性质的 Webhook 的示例 （并且期望使用系统信任根证书来验证 TLS 证书，因此不指定 caBundle）：

