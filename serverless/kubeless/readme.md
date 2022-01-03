https://github.com/vmware-archive/kubeless

https://kubeless.io/



three steps:

Download the kubeless CLI from the release page.
Create a kubeless namespace (used by default)
Then use one of the YAML manifests found in the release page to deploy kubeless. It will create a functions Custom Resource Definition and launch a controller.


https://kubeless.io/docs/quick-start/


由于 Kubeless 的功能特性是建立在Kubernetes上的，所以对于熟悉 Kubernetes的人来说非常容易部署 Kubeless， 其主要实现是将用户编写的函数在Kubernetes中转变为 CRD（ Custom Resource Definition，自定义资源）， 并以容器的方式运行在集群中。


能力包括：

敏捷构建 - 能够基于用户提交的源码迅速构建可执行的函数，简化部署流程；
灵活触发 - 能够方便地基于各类事件触发函数的执行，并能方便快捷地集成新的事件源；
自动伸缩 - 能够根据业务需求，自动完成扩容缩容，无须人工干预。


Kubeless 基于 K8s 提供了较为完整的 serverless 解决方案，但和一些商业 serverless 产品还存在一定差距：

Kubeless 并未在镜像拉取、代码下载、容器启动等方面做过多优化，导致函数冷启动时间过长；
Kubeless 并未过多考虑多租户的问题，如果希望多个用户的函数运行在同一个集群里，还需要进行二次开发。

https://www.cnblogs.com/yunqishequ/p/10057458.html


1 Kubeless基本组成
Kubeless主要由以下三部分组成：

Functions
Triggers
Runtime
下面针对这三个组成部分，进行详细介绍。

Functions
Functions 表示要执行的代码，即为函数，在Kubeless中函数包含有关其运行时的依赖、构建等元数据。函数具有独立生命周期，并支持以下方法：
（1） Deploy： Kubeless 将函数部署为 Pod的形式运行在Kubernetes集群中，此步骤会涉及构建函数镜像等操作。
（2）Execute：执行函数，不通过任何事件源调用。
（3）Update：修改函数元数据。
（4）Delete：在Kubernetes集群中删除函数的所有相关资源。
（5）List：显示函数列表。
（6）Logs：函数实例在Kubernetes中生成及运行的日志。

Triggers
Triggers表示函数的事件源，当事件发生时，Kubeless确保最多调用一次函数，Triggers可以与单个功能相关联，也可与多个功能相关联，具体取决于事件源类型。Triggers与函数的生命周期解耦，可以进行如下操作：
（1）Create：使用事件源和相关功能的详细信息创建 Triggers。
（2）Update: 更新 Triggers元数据。
（3）Delete：删除Triggers及为其配置的任何资源。
（4）List：显示Triggers列表。

Runtime
函数使用语言因不同用户的喜好通常多样化， Kubeless 为用户带来了几乎所有的主流函数运行时， 目前含有[3]：
（1） Python: 支持2.7、3.4、3.6版本。
（2） NodeJS: 支持6、8版本。

（3） Ruby: 支持2.4版本。
（4） PHP: 支持7.2版本。
（5） Golang: 支持1.10版本。
（6） .NET: 支持2.0版本。

（7） Ballerina: 支持0.975.0版本。
在Kubeless中，每个函数运行时都会以镜像的方式封装在容器镜像中，通过在Kubeless配置中引用这些镜像来使用，可以通过 Docker CLI 查看源代码。


Kubeless设计方式
与其它开发框架一样， Kubeless也有自己的设计方式，Kubeless利用Kubernetes中的许多概念来完成对函数实例的部署，主要使用了 Kubernetes以下特性【2】 ：

（1） CRD（ 自定义资源） 用于表示函数。
（2） 每个事件源都被当作为一个单独的 Trigger CRD 对象。
（3） CRD Controller 用于处理与 CRD 对象相应的 CRUD 操作。
（4） Deployment/Pod 运行相应的运行时。
（5） ConfigMap 将函数的代码注入运行时的 Pod。
（6） Init-container 加载函数的依赖项。
（7） 使用Service在集群中暴露函数（ ClusterIP）。
（8） 使用Ingress资源对象暴露函数到外部。
Kubernetes CRD 和 CRD Controller 构成了 Kubeless 的设计宗旨，对函数和 Triggers 使用不同的 CRD 可以明确区分关键点，使用单独的 CRD Controller 可以使代码解耦并模块化。

http://blog.nsfocus.net/kubeless/
https://www.cnblogs.com/zhaowei121/p/12029678.html
https://zhuanlan.zhihu.com/p/143690697