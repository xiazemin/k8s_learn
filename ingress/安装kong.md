helm install kong-gateway stable/kong --version 0.36.7 \
    --set admin.useTLS=false \
    --set admin.nodePort=32444 \
    --set proxy.http.nodePort=32080 \
    --set proxy.tls.nodePort=32443 \
    --set replicaCount=2
WARNING: This chart is deprecated
Error: INSTALLATION FAILED: rendered manifests contain a resource that already exists. Unable to continue with install: could not get information about the resource: customresourcedefinitions.apiextensions.k8s.io "kongconsumers.configuration.konghq.com" is forbidden: User "u-tus2yqlyas" cannot get resource "customresourcedefinitions" in API group "apiextensions.k8s.io" at the cluster scope

在Kubernetes下部署Kong，Kong定义了多个自定义资源CRDs，Kong通过这些CRDs定义路由、插件等配置数据，这些配置数据通过Kubernetes API转发给Kong Ingress Controller，并转发给Kong数据平面，Kong继而控制客户端的请求流量。

CRDs数据应该存放在Kubernetes的etcd数据库中，这些数据通过Ingress Controller转换为Kong的配置数据，Kong数据平面缓存这些配置数据。从上面也可以看出Kong Pod中有两个容器：ingress-controller、kong。

https://blog.csdn.net/twingao/article/details/104073112


在Kubernetes下部署的有数据库的Kong，Kong定义了多个自定义资源CRDs，Kong通过这些CRDs定义路由、插件等配置，这些配置数据通过Kubernetes API转发给Kong Ingress Controller，并通过控制平面将数据保存在数据库中，数据平面从数据库中读取配置信息并缓存，Kong继而控制客户端的请求流量。

CRDs数据应该存放在Kubernetes的etcd数据库中，这些数据通过Ingress Controller转换为Kong的数据库配置方式，Kong数据平面从数据库中读取配置数据。现在这种方式有些复杂，感觉为了复用原来的Kong和数据库同步数据的方式，而且无数据库方式不支持传统的管理Rest API方式。Kubernetes并支持Ingress Controller的方式下支持数据库做持久存储有些多此一举。应该直接无数据库比较合理。当然Kong已经明确推荐在Kubernetes下采用无数据库方式。

无数据方式不支持传统的管理Rest API。为了向后兼容，Kubernetes保留了数据库的方式。


helm uninstall kong-gateway
Error: uninstall: Release not loaded: kong-gateway: release: not found

kubectl delete pvc data-gateway-postgresql-0

在NFS服务器端，执行：

cd /data/kong

rm -rf *

https://blog.csdn.net/twingao/article/details/104073159


 % helm install kong-gateway stable/kong --version 0.36.7 \
    --set admin.useTLS=false \
    --set admin.nodePort=32444 \
    --set proxy.http.nodePort=32080 \
    --set proxy.tls.nodePort=32443 \
    --set replicaCount=2
WARNING: This chart is deprecated
Error: INSTALLATION FAILED: unable to build kubernetes objects from release manifest: [unable to recognize "": no matches for kind "CustomResourceDefinition" in version "apiextensions.k8s.io/v1beta1", unable to recognize "": no matches for kind "ClusterRole" in version "rbac.authorization.k8s.io/v1beta1", unable to recognize "": no matches for kind "ClusterRoleBinding" in version "rbac.authorization.k8s.io/v1beta1", unable to recognize "": no matches for kind "Role" in version "rbac.authorization.k8s.io/v1beta1", unable to recognize "": no matches for kind "RoleBinding" in version "rbac.authorization.k8s.io/v1beta1"]

 % kubectl api-versions |grep apiextensions
apiextensions.k8s.io/v1


