函数服务主要开源项目
项目	简介
Serverless	Simple way to build serverless applications
OpenFaaS	Serverless Functions Made Simple
Kubeless	Kubernetes Native Serverless Framework
OpenWhisk	Serverless Functions Platform for Building Cloud Applications
Knative	Kubernetes-based, scale-to-zero, request-driven compute
Fission	Fast and Simple Serverless Functions for Kubernetes
Fn	Event-driven Functions-as-a-Service (FaaS) compute platform
Nuclio	Serverless for Real-Time Events and Data Processing


https://bbs.huaweicloud.com/blogs/247064



https://cloud.tencent.com/developer/article/1617013



Severless Framework
它不是一个平台，但它可以运行任何函数。它是无服务的一个 SDK。事实上，它本质上只是做了一层包装。但最爽的是，通过 Serverless framwork 来打包的函数，你可以将相同的代码部署到 Lambda，Google Functions，Azure Functions，OpenWhisk，OpenFaas，Kubeless 或 Fn 中。


OpenWhisk 利用了 CouchDB, Kafka, Nginx, Redis 和 Zookeeper，有很多底层的组件，所以增加了一定的复杂性。好处是开发者可以清晰地关注于可伸缩和弹性的服务，缺点是开发者和使用者都需要具备这些工具的知识和学习如何使用，另一个缺点是它重复实现了一些 Kubernetes 中已经存在的特性（比如自动扩缩容）。函数最终会和框架一起运行在 Docker 容器中。

Kubeless
我对 Kubeless 非常感兴趣，因为它是基于原生 Kubernetes 的。工作原理是在原生 Kubernetes 添加了 “函数” 这种自定义资源的 CRD。除了这个实现非常聪明，也意味着它将 Kubernetes 变成了一个函数运行器，而没有像其他框架那样添加了各种复杂的功能，比如消息机制。

https://blog.csdn.net/weixin_30432179/article/details/100092874


