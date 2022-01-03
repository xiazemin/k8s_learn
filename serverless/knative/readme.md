https://github.com/knative/serving

k8s最底层，istio部署在k8s之上，knative最上面，依赖istio和k8s


Kubernetes 是基座，Istio和Knative都是通过CRD + controller模式 扩展出来的项目，分别用于不同的场景。
Istio 号称k8s 网络++，在网络上对k8s是一种增强。


Knative 则专注于构建这些工具和服务来提升现有的 Kubernetes 体验。这为不断增长的开发者（Kubernetes 用户）群体带来了即时利益以及轻松的无服务器开发体验。为此，Knative 使用与 Kubernetes 本身相同的模式（控制器）、API (kube-api) 和 Kubernetes 基础架构（Kubernetes 资源）构建而成。Knative 还提供“缩容至零”功能，支持真正零成本使用空闲应用程序，并支持采用蓝/绿部署来测试无服务器应用的新版本。

Knative是一个serverless的平台，简单来说，你只需要一个docker image，kantive会自动加上自动缩放，api，灰度



在 Knative 之前社区已经有很多 Serverless 解决方案，如下所示这些：
kubeless
Fission
OpenFaaS
Apache OpenWhisk
除了上面这些社区的开源解决方案以外各大云厂商也都有各自的 FAAS 产品的实现比如：

AWS Lambda
Google Cloud Functions
Microsoft Azure Functions
阿里云的函数计算

https://zhuanlan.zhihu.com/p/90430608



knative 核心概念和原理
为了实现 serverless 应用的管理，knative 把整个系统分成了三个部分：
Build：构建系统，把用户定义的函数和应用 build 成容器镜像
Serving：服务系统，用来配置应用的路由、升级策略、自动扩缩容等功能
Eventing：事件系统，用来自动完成事件的绑定和触发


https://blog.csdn.net/jiangbb8686/article/details/102295612

