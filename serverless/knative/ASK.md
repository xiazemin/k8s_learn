阿里云Serverless Kubernetes（ASK）集群
https://www.jianshu.com/p/e43c1413cb51
Knative支持HPA的弹性能力。您可以在Knative Service中设置CPU指标阈值，满足在突发高负载的场景下，自动扩缩容资源的诉求
https://help.aliyun.com/document_detail/198683.html

在今年阿里云的Serverless Kubernetes做了一个演进，可以把Serverless 的编排模块变成一个纯单租的架构，每个Serverless K8s的master都是独立的，好处是不同的 serverless K8s 互相之间是并不会干扰，从而实现基本上等价于标准 K8s 能力的一个Serverless服务。可以更好的支撑场景，比如Knative、Istio等技术。此外，在单租的架构下，可以很好的跟用户VPC做点对点的打通，可以实现较好的安全防护。

https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/103397246


可参考：阿里云容器服务部署 Knative
https://zhuanlan.zhihu.com/p/92001734