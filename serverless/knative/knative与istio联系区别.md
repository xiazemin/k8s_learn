Knative 则专注于构建这些工具和服务来提升现有的 Kubernetes 体验。这为不断增长的开发者（Kubernetes 用户）群体带来了即时利益以及轻松的无服务器开发体验。为此，Knative 使用与 Kubernetes 本身相同的模式（控制器）、API (kube-api) 和 Kubernetes 基础架构（Kubernetes 资源）构建而成。Knative 还提供“缩容至零”功能，支持真正零成本使用空闲应用程序，并支持采用蓝/绿部署来测试无服务器应用的新版本。Knative 和Istio 并没有什么强关联，只是istio 的gateway 组件可以被knative用来做流量接入。事实上，knative 还支持solo，ambassador，contour等其他gateway来做流量接入。

Knative是一个serverless的平台，简单来说，你只需要一个docker image，kantive会自动加上自动缩放，api，灰度
Knative需要一个组件来实现流量分发相关的功能，主要是7层，这时候可以使用Istio，也可以使用别的，类似nginx ingress之类的，不是必须使用Istio。

https://www.zhihu.com/question/363951810


为什么使用 Knative
使用 Knative 作为 Faas 或者说 Serverless 的基础设施，这是国内的技术团队较为主流的做法。大家希望从knative中拿到的价值主要有这样了两个：
免除云服务供应商锁定。相对于AWS lambda 必须使用AWS上多个其他组件，造成巨大迁移壁垒，knative至少是云服务商中立的，控制权在自己手上。
自动扩容和缩容至零实例。极致的弹性带来了两个直观结果: 运维成本大大降低、硬件成本降低。
https://blog.csdn.net/weixin_38754564/article/details/103867231

https://www.servicemesher.com/blog/gloo-by-solo-io-is-the-first-alternative-to-istio-on-knative/



Knative的伸缩是依赖修改deployment的replica数实现的。

如何采集请求数？
启动revision的pod时，也会启动一个autoscaler（一个knative revision只启动一个autoscaler），autoscaler自己本身也会scale到0，用于接收请求数统计和处理伸缩容。

业务pod中，会注入queue-proxy sidecar，用于接收请求，在这里会统计并发数，每秒向autoscaler汇报，接收到的请求会转发给业务container。

注：单租户模式下一个revision启动一个autoscaler，多租户共用一个autoscaler

计算需要pod的个数？
autoscaler接收到并发统计的时候，会根据算法计算需要的pod个数。

算法中有两种模式，分别是panic和stable模式，一个是短时间，一个是长时间，为了解决短时间内请求突增的场景，需要快速扩容。

文档中描述的算法是，默认的target concurrency是1，如果一个revision 35QPS，每个请求花费0.25秒，Knative Serving 觉得需要 9 个 pod。

ceil(35 * .25) = ceil(8.75) = 9
Stable Mode（稳定模式）
在稳定模式下，Autoscaler 根据每个pod期望的并发来调整Deployment的副本个数。根据每个pod在60秒窗口内的平均并发来计算，而不是根据现有副本个数计算，因为pod的数量增加和pod变为可服务和提供指标数据有一定时间间隔。

Panic Mode （恐慌模式）
Panic时间窗口默认是6秒，如果在6秒内达到2倍期望的并发，则转换到恐慌模式下。在恐慌模式下，Autoscaler根据这6秒的时间窗口计算，这样更能及时的响应突发的流量请求。每2秒调整Deployment的副本数达到想要的pod个数（或者最大10倍当前pod的数量），为了避免pod数量频繁变动，在恐慌模式下只能增加，不会减少。60秒后会恢复回稳定模式

https://cloud.tencent.com/developer/article/1549965




