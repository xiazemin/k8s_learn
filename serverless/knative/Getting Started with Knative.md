https://juejin.cn/post/6844903794086510599
Knative 将重点放在三个关键组件上：build（构建）你的应用程序，为其提供流量serving（服务），以及确保应用程序能够轻松地生产和消费event（事件）。

Build（构建）通过灵活的插件化的构建系统将用户源代码构建成容器。目前已经支持多个构建系统，比如 Google 的 Kaniko，它无需运行 Docker daemon 就可以在 Kubernetes 集群上构建容器镜像。Serving（服务）基于负载自动伸缩，包括在没有负载时缩减到零。允许你为多个修订版本（revision）应用创建流量策略，从而能够通过 URL 轻松路由到目标应用程序。Event（事件）使得生产和消费事件变得容易。抽象出事件源，并允许操作人员使用自己选择的消息传递层。Knative 是以 Kubernetes 的一组自定义资源类型（CRD）的方式来安装的，因此只需使用几个 YAML 文件就可以轻松地开始使用 Knative 了。



为应用定义服务

为了定义一个从源代码构建和部署应用所需的最小配置，我们必须遵循以下步骤：

1. 选择适合你的用例的构建模板BuildTemplate，并将其应用于Kube群集。在我的例子中，我选择了Buildpack模板，因为它能够自动检测并安装我的应用程序的任何依赖项。

$ kubectl apply -f https://raw.githubusercontent.com/knative/build-templates/master/buildpack/buildpack.yaml

2. Knative需要将应用程序上传到Docker Registry，因此我们需要有一个注册器，为了简单起见，我决定使用Docker Hub。

3.要让Knative推送镜像，我们需要定义ServiceAccount用户名和密码才能使用Docker Hub，并在集群中配置它们

https://www.jdon.com/49816


 kubectl apply -f https://raw.githubusercontent.com/knative/build-templates/master/buildpack/buildpack.yaml

 https://blog.csdn.net/alisystemsoftware/article/details/94739576

https://github.com/knative/build
  Knative Build is deprecated in favor of Tekton Pipelines. There are no plans to produce future releases of this component.


https://github.com/tektoncd/pipeline

The Tekton Pipelines project provides k8s-style resources for declaring CI/CD-style pipelines.


https://github.com/tektoncd/pipeline/blob/main/docs/install.md
Run the following command to install Tekton Pipelines and its dependencies:

kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
