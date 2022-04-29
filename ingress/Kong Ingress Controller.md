https://blog.csdn.net/twingao/article/details/104073289

Kong之前都是使用Admin API来进行管理的，Kong主要暴露两个端口管理端口8001和代理端口8000，管理Kong主要的是为上游服务配置Service、Routes、Plugins、Consumer等实体资源，Kong按照这些配置规则进行对上游服务的请求进行路由分发和控制。在Kubernetes集群环境下，Admin API方式不是很适应Kubernetes声明式管理方式。所以Kong在Kubernetes集群环境下推出Kong Ingress Controller。Kong Ingress Controller定义了四个CRDs（CustomResourceDefinitions），基本上涵盖了原Admin API的各个方面。

kongconsumers：Kong的用户，给不同的API用户提供不同的消费者身份。
kongcredentials：Kong用户的认证凭证。
kongingresses：定义代理行为规则，是对Ingress的补充配置。
kongplugins：插件的配置。


Kong创建的CRDs：
kubectl get crds
NAME                                       CREATED AT
kongconsumers.configuration.konghq.com     2019-12-15T08:02:29Z
kongcredentials.configuration.konghq.com   2019-12-15T08:02:29Z
kongingresses.configuration.konghq.com     2019-12-15T08:02:29Z
kongplugins.configuration.konghq.com       2019-12-15T08:02:29Z

Kong Pod其中有两个容器，一个为ingress-controller，一个为kong。Kong对外提供两个服务，gateway-kong-admin为管理服务，支持Admin API，gateway-kong-proxy为代理服务,这两个服务都由kong提供，而CRDs的API接口是ingress-controller容器提供的。

其实在Kubernetes集群中也可以直接部署Kong和PostgreSQL，那样是不支持Kong Ingress Controller，直接使用Admin API管理即可。


https://blog.csdn.net/twingao/article/details/104073289