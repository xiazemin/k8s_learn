https://zhuanlan.zhihu.com/p/51882683

Kubernetes没有为无服务器平台提供更精细的部署模型。通过Knative Serving，Kubernetes具有了运行无服务器工作负载所需的功能。

Knative Serving基于Istio，为应用服务提供了部署模型，以便运行无服务器工作负载。

让我们开始如何在Kubernetes上部署第一个无服务器应用程序。这里将使用OKD，这是Kubernetes的社区发行版，Red Hat OpenShift提供支持。要在本地运行OKD，可以使用minishift，一个可用于开发目的的单节点OKD集群。

https://cloud.redhat.com/blog/knative-serving-your-serverless-services


Knative 安装
因为 Knative 依赖 Istio，首先我们需要先安装 Istio。

安装 Knative
https://www.bookstack.cn/read/serverless-handbook/knative-primer-installation.md


如何部署一个Knative Service
https://blog.csdn.net/hzbooks/article/details/115987716


knative 核心概念和原理


为了实现 serverless 应用的管理，knative 把整个系统分成了三个部分：
Build：构建系统，把用户定义的函数和应用 build 成容器镜像
Serving：服务系统，用来配置应用的路由、升级策略、自动扩缩容等功能
Eventing：事件系统，用来自动完成事件的绑定和触发
https://zhuanlan.zhihu.com/p/47916416


Knative 将重点放在两个关键组件上：为其提供流量serving（服务），以及确保应用程序能够轻松地生产和消费event（事件）。

Serving（服务）

基于负载自动伸缩，包括在没有负载时缩减到零。允许你为多个修订版本（revision）应用创建流量策略，从而能够通过 URL 轻松路由到目标应用程序。
Event（事件）

使得生产和消费事件变得容易。抽象出事件源，并允许操作人员使用自己选择的消息传递层。
关于Serving和Event详细介绍，我们会在后续的文章中一一解读。

knative的优势：

便利性：Knative 以 Kubernetes 作为其底层框架，因此无论是线上还是线下，任何 Kubernetes 集群，无论是云上 Kubernetes 服务还是自建 Kubernetes 集群，都能通过安装 knative 插件快速的搭建 serverless 平台。
标准化：Knative 联合 CNCF，把所有事件标准化，统一为 CloudEvent，提供事件的跨平台，同时让函数和具体的调用方能够解耦。
服务间解耦：使用 Knative 使得应用不在与底层依赖服务强绑定，可以跨云实现业务互通
成熟的生态：Knative 基于 Kubernetes 体系构建，与 kubernetes 生态结合更紧密；
自动伸缩：监控应用的请求，并自动扩缩容, 借助于istio(ambassador,gloo等)天生支持蓝绿发布、回滚功能，方便应用发布流程。
应用监控：支持日志的收集、查找和分析，并支持 VAmetrics 数据展示、调用关系 tracing

https://zhuanlan.zhihu.com/p/139671487

https://github.com/knative/serving

Knative 事件结构：

1. 事件源  - 代表事件的生产者（例如GitHub）

2. 事件类型  - 描述不同事件源支持的事件类型（例如，上面提到的GitHub源的Webhook）

3. 事件消费者  - 代表你Action目标（即Knative定义的任何路线）

4. 事件feed  - 是将事件类型连接到操作的绑定或配置
https://www.jdon.com/49700