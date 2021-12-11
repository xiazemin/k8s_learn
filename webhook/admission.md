准入 Webhook 是一种用于接收准入请求并对其进行处理的 HTTP 回调机制。 可以定义两种类型的准入 webhook，即 验证性质的准入 Webhook 和 修改性质的准入 Webhook。 修改性质的准入 Webhook 会先被调用。它们可以更改发送到 API 服务器的对象以执行自定义的设置默认值操作。


对 apiservers 进行身份认证
如果你的 webhook 需要身份验证，则可以将 apiserver 配置为使用基本身份验证、持有者令牌或证书来向 webhook 提供身份证明。完成此配置需要三个步骤。

启动 apiserver 时，通过 --admission-control-config-file 参数指定准入控制配置文件的位置。

在准入控制配置文件中，指定 MutatingAdmissionWebhook 控制器和 ValidatingAdmissionWebhook 控制器应该读取凭据的位置。 凭证存储在 kubeConfig 文件中（是​​的，与 kubectl 使用的模式相同），因此字段名称为 kubeConfigFile。 

https://kubernetes.io/zh/docs/reference/access-authn-authz/extensible-admission-controllers/

https://istio.io/docs/setup/kubernetes/sidecar-injection/

https://github.com/kubernetes/kubernetes/blob/v1.13.0/test/images/webhook/main.go

先要看下api-server的启动参数里有没有开启准入配置
–enable-admission-plugins=NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook

https://blog.csdn.net/zyxpaomian/article/details/117443864


AdmissionWebhook
我们知道k8s在各个方面都具备可扩展性，比如通过cni实现多种网络模型，通过csi实现多种存储引擎，通过cri实现多种容器运行时等等。而AdmissionWebhook就是另外一种可扩展的手段。 除了已编译的Admission插件外，可以开发自己的Admission插件作为扩展，并在运行时配置为webhook。

Admission webhooks是HTTP回调，它接收Admission请求并对它们做一些事情。可以定义两种类型的Admission webhook，ValidatingAdmissionWebhook和MutatingAdmissionWebhook。

如果启用了MutatingAdmission，当开始创建一种k8s资源对象的时候，创建请求会发到你所编写的controller中，然后我们就可以做一系列的操作。比如我们的场景中，我们会统一做一些功能性增强，当业务开发创建了新的deployment，我们会执行一些注入的操作，比如敏感信息aksk，或是一些优化的init脚本。

而与此类似，只不过ValidatingAdmissionWebhook 是按照你自定义的逻辑是否允许资源的创建。比如，我们在实际生产k8s集群中，处于稳定性考虑，我们要求创建的deployment 必须设置request和limit。

https://zhuanlan.zhihu.com/p/136173524

http://docs.kubernetes.org.cn/706.html

https://cloud.tencent.com/developer/article/1445760


webhook如何工作的

注册webhook server

资源操作请求通过API Server Auth验证

根据注册信息回调对应的webhook server

webhook名称

② 描述api-server操作什么资源什么动作时调用webhook插件

③ webhook service所在的namespace

④ webhook service name

⑤ 调用webhook api的地址

⑥ 提供和webhook通信的TLS链接信息, 生成的证书必须支持<svc_name>.<svc_namespace>.svc，这个证书可以直接使用k8s集群的ca.crt（ kubectl config view --raw -o json | jq -r '.clusters[0].cluster."certificate-authority-data"' | tr -d '"'）

https://www.cnblogs.com/zhangmingcheng/p/14066884.html

https://blog.csdn.net/u012803274/article/details/115035744

