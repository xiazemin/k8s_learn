repilca set 是更高级的replication contorller 功能基本一样 只是在seletor label 增加集合的概念，支持基于等式的seletor。
replica set 主要功能：
确保Pod数量：它会确保Kubernetes中有指定数量的Pod在运行，如果少于指定数量的Pod，RC就会创建新的，反之这会删除多余的，保证Pod的副本数量不变。

确保Pod健康：当Pod不健康，比如运行出错了，总之无法提供正常服务时，RC也会杀死不健康的Pod，重新创建新的。

弹性伸缩：在业务高峰或者低峰的时候，可以用过RC来动态的调整Pod数量来提供资源的利用率，当然我们也提到过如果使用HPA这种资源对象的话可以做到自动伸缩。

滚动升级：滚动升级是一种平滑的升级方式，通过逐步替换的策略，保证整体系统的稳定性。

https://blog.csdn.net/qq_35837864/article/details/90778026


https://www.cnblogs.com/fanqisoft/p/11573598.html

% kubectl apply -f deployment/ReplicaSet.yaml
error: unable to recognize "deployment/ReplicaSet.yaml": no matches for kind "ReplicaSet" in version "extensions/v1beta1"

解决：查看版本
%  kubectl api-versions
admissionregistration.k8s.io/v1
admissionregistration.k8s.io/v1beta1
apiextensions.k8s.io/v1
apiextensions.k8s.io/v1beta1
apiregistration.k8s.io/v1
apiregistration.k8s.io/v1beta1
apps/v1
authentication.k8s.io/v1
authentication.k8s.io/v1beta1
authorization.k8s.io/v1
authorization.k8s.io/v1beta1
autoscaling/v1
autoscaling/v2beta1
autoscaling/v2beta2
batch/v1
batch/v1beta1
catalog.cattle.io/v1
certificates.k8s.io/v1
certificates.k8s.io/v1beta1
coordination.k8s.io/v1
coordination.k8s.io/v1beta1
discovery.k8s.io/v1
discovery.k8s.io/v1beta1
events.k8s.io/v1
events.k8s.io/v1beta1
extensions/v1beta1
flowcontrol.apiserver.k8s.io/v1beta1
management.cattle.io/v3
networking.k8s.io/v1
networking.k8s.io/v1beta1
node.k8s.io/v1
node.k8s.io/v1beta1
policy/v1
policy/v1beta1
rbac.authorization.k8s.io/v1
rbac.authorization.k8s.io/v1beta1
samplecontroller.k8s.io/v1alpha1
scheduling.k8s.io/v1
scheduling.k8s.io/v1beta1
storage.k8s.io/v1
storage.k8s.io/v1beta1
ui.cattle.io/v1
v1

修改成

apiVersion: apps/v1


% kubectl apply -f deployment/ReplicaSet.yaml
error: error validating "deployment/ReplicaSet.yaml": error validating data: ValidationError(ReplicaSet.spec.template.metadata): unknown field "1abels" in io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta; if you choose to ignore these errors, turn validation off with --validate=false

拼写错误将1ables改成lables问题解决

 % kubectl apply -f deployment/ReplicaSet.yaml
replicaset.apps/apple-rc created


% kubectl get rs 
NAME                                  DESIRED   CURRENT   READY   AGE
apple-rc                              3         3         3       56s

% kubectl get pod
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     18         102d
apple-rc-k88km                              1/1     Running     0          44s
apple-rc-wgd4c                              1/1     Running     0          44s


