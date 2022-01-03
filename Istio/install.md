https://www.bookstack.cn/read/serverless-handbook/knative-primer-installation.md

https://blog.csdn.net/weixin_30561177/article/details/98355350

软件下载地址：https://github.com/istio/istio/
安装 istioctl install --set profile=demo -y
https://www.cnblogs.com/hujinzhong/p/15012988.html

https://www.cnblogs.com/Mr-Axin/p/14752524.html

https://blog.csdn.net/qq_42747099/article/details/88941514

https://zhuanlan.zhihu.com/p/150643479

https://my.oschina.net/u/2306127/blog/1793740

https://www.howtoing.com/how-to-install-and-use-istio-with-kubernetes


 % istioctl manifest apply --set values.global.mtls.enabled=true 
Run the command with the --force flag if you want to ignore the validation error and proceed.
Error: generate config: unknown field "mtls" in v1alpha1.GlobalConfig

https://github.com/istio/istio/issues/20000




default: enables components according to the default settings of the IstioControlPlane API (recommend for production deployments). You can display the default setting by running the command istioctl profile dump.可用于部署生产环境的Istio

demo: configuration designed to showcase Istio functionality with modest resource requirements. It is suitable to run the Bookinfo application and associated tasks. This is the configuration that is installed with the quick start instructions, but you can later customize the configuration to enable additional features if you wish to explore more advanced tasks. 只是demo，资源request的很低，安装的功能很多

This profile enables high levels of tracing and access logging so it is not suitable for performance tests.

minimal: the minimal set of components necessary to use Istio’s traffic management features. 最小化安装

sds: similar to the default profile, but also enables Istio’s SDS (secret discovery service). This profile comes with additional authentication features enabled by default (Strict Mutual TLS). 对比default，只开启了SDS

remote: used for configuring remote clusters of a multicluster mesh with a shared control plane configuration. 多cluster共享一个控制面部署

https://knner.wang/2020/01/07/ServiceMesh-Istio-Series--install-istio-1-4-using-istioctl.html


https://github.com/istio/istio/issues/14365

https://istio.io/latest/zh/blog/2019/data-plane-setup/

https://istio.io/latest/zh/docs/reference/config/annotations/


给命名空间添加标签，指示 Istio 在部署应用的时候，自动注入 Envoy 边车代理：

 % kubectl label namespace default istio-injection=enabled
namespace/default labeled


 % docker images |grep proxyv2
istio/proxyv2                                                     1.12.1                                                      4f3c74acb37a   3 weeks ago     252MB


% kubectl -n istio-system get deployment istio-egressgateway -o yaml > istio-egressgateway.yaml



 % kubectl -n istio-system describe pod istio-egressgateway-687f4db598-tt72j 

  Warning  FailedScheduling  24m   default-scheduler  0/1 nodes are available: 1 node(s) didn't match Pod's node affinity/selector.



      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - preference:
              matchExpressions:
              - key: kubernetes.io/arch
                operator: In
                values:
                - amd64


没有arm64
在最后加上即可


 % kubectl delete -f istio-egressgateway.yaml
deployment.apps "istio-egressgateway" deleted

% kubectl apply -f istio-egressgateway.yaml 
deployment.apps/istio-egressgateway created

https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/


% kubectl -n istio-system get deploy istio-ingressgateway  -o yaml > istio-ingressgateway.yaml  

% kubectl delete -f  istio-ingressgateway.yaml  
deployment.apps "istio-ingressgateway" deleted


% kubectl apply -f  istio-ingressgateway.yaml  
deployment.apps/istio-ingressgateway created

发现服务已经起来了，我们修改下哪里呢


Istio/istio-1.12.1/manifests/charts/gateways/istio-ingress/values.yaml

Istio/istio-1.12.1/manifests/charts/gateways/istio-egress/values.yaml

  arch:
    amd64: 2
    s390x: 2
    ppc64le: 2
    arm64: 3

% istioctl install --set profile=demo  


% kubectl -n istio-system get pods
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-687f4db598-hkc5k    0/1     Pending   0          12m
istio-egressgateway-6ff67579c4-fk7l4    1/1     Running   0          14m
istio-ingressgateway-67b86984bb-sdfqb   1/1     Running   0          12m
istio-ingressgateway-78f69bd5db-4wp8x   0/1     Pending   0          12m
istiod-76d66d9876-2qbl8                 1/1     Running   0          52m


 % kubectl -n istio-system describe pod istio-egressgateway-687f4db598-hkc5k 
Events:
  Type     Reason            Age   From               Message
  ----     ------            ----  ----               -------
  Warning  FailedScheduling  12m   default-scheduler  0/1 nodes are available: 1 node(s) didn't match Pod's node affinity/selector.

 
  % kubectl -n istio-system delete  pod istio-egressgateway-687f4db598-hkc5k 
pod "istio-egressgateway-687f4db598-hkc5k" deleted


 % kubectl delete -f istio-egressgateway.yaml 
deployment.apps "istio-egressgateway" deleted

% kubectl delete -f istio-ingressgateway.yaml
deployment.apps "istio-ingressgateway" deleted


卸载Istio


# 使用 --purge 进行完全卸载，包含集群范围的资源，这些资源有可能是和其他Istio控制面共享的
istioctl x uninstall --purge

# 只卸载某个特定的Istio控制面
istioctl x uninstall <your original installation options>
# 或 
istioctl manifest generate <your original installation options> | kubectl delete -f -
https://www.jianshu.com/p/92c0d536d950



 % istioctl install --set profile=minimal -y                                     
✔ Istio core installed                                                                                                                  
✔ Istiod installed                                                                                                                      
- Pruning removed resources                                                                                                               Removed PodDisruptionBudget:istio-system:istio-ingressgateway.
  Removed PodDisruptionBudget:istio-system:istio-egressgateway.
  Removed Deployment:istio-system:istio-ingressgateway.
  Removed Deployment:istio-system:istio-egressgateway.
  Removed Service:istio-system:istio-ingressgateway.
  Removed Service:istio-system:istio-egressgateway.
  Removed ServiceAccount:istio-system:istio-ingressgateway-service-account.
  Removed ServiceAccount:istio-system:istio-egressgateway-service-account.
  Removed RoleBinding:istio-system:istio-ingressgateway-sds.
  Removed RoleBinding:istio-system:istio-egressgateway-sds.
  Removed Role:istio-system:istio-ingressgateway-sds.
  Removed Role:istio-system:istio-egressgateway-sds.
✔ Installation complete                                                                                                                 Making this installation the default for injection and validation.
2022-01-03T11:40:02.496411Z     error   klog    couldn't get resource list for custom.metrics.k8s.io/v1beta1: the server is currently unable to handle the request
2022-01-03T11:40:02.501888Z     error   klog    couldn't get resource list for custom.metrics.k8s.io/v1beta1: the server is currently unable to handle the request

Thank you for installing Istio 1.12.  Please take a few minutes to tell us about your install/upgrade experience!  https://forms.gle/FegQbc9UvePd4Z9z7


 % istioctl install --set profile=demo -y     
✔ Istio core installed                                                                                                                  
✔ Istiod installed                                                                                                                      
  Processing resources for Egress gateways, Ingress gateways. Waiting for Deployment/istio-system/istio-egressgateway, Deployment/ist..





% istioctl x uninstall --purge
All Istio resources will be pruned from the cluster
Proceed? (y/N) y
  Removed IstioOperator:istio-system:installed-state.
  Removed HorizontalPodAutoscaler:istio-system:istiod.
  Removed PodDisruptionBudget:istio-system:istio-egressgateway.
  Removed PodDisruptionBudget:istio-system:istio-ingressgateway.
  Removed PodDisruptionBudget:istio-system:istiod.
  Removed Deployment:istio-system:istio-egressgateway.
  Removed Deployment:istio-system:istio-ingressgateway.
  Removed Deployment:istio-system:istiod.
  Removed Service:istio-system:istio-egressgateway.
  Removed Service:istio-system:istio-ingressgateway.
  Removed Service:istio-system:istiod.
  Removed ConfigMap:istio-system:istio.
  Removed ConfigMap:istio-system:istio-sidecar-injector.
object: Pod:istio-system:istio-egressgateway-687f4db598-mpz42 is not being deleted because it no longer exists
  Removed Pod:istio-system:istio-egressgateway-687f4db598-mpz42.
object: Pod:istio-system:istio-ingressgateway-78f69bd5db-c57bb is not being deleted because it no longer exists
  Removed Pod:istio-system:istio-ingressgateway-78f69bd5db-c57bb.
  Removed Pod:istio-system:istiod-76d66d9876-hjmj4.
  Removed ServiceAccount:istio-system:istio-egressgateway-service-account.
  Removed ServiceAccount:istio-system:istio-ingressgateway-service-account.
  Removed ServiceAccount:istio-system:istio-reader-service-account.
  Removed ServiceAccount:istio-system:istiod.
  Removed ServiceAccount:istio-system:istiod-service-account.
  Removed RoleBinding:istio-system:istio-egressgateway-sds.
  Removed RoleBinding:istio-system:istio-ingressgateway-sds.
  Removed RoleBinding:istio-system:istiod.
  Removed RoleBinding:istio-system:istiod-istio-system.
  Removed Role:istio-system:istio-egressgateway-sds.
  Removed Role:istio-system:istio-ingressgateway-sds.
  Removed Role:istio-system:istiod.
  Removed Role:istio-system:istiod-istio-system.
  Removed EnvoyFilter:istio-system:stats-filter-1.10.
  Removed EnvoyFilter:istio-system:stats-filter-1.11.
  Removed EnvoyFilter:istio-system:stats-filter-1.12.
  Removed EnvoyFilter:istio-system:tcp-stats-filter-1.10.
  Removed EnvoyFilter:istio-system:tcp-stats-filter-1.11.
  Removed EnvoyFilter:istio-system:tcp-stats-filter-1.12.
  Removed MutatingWebhookConfiguration::istio-sidecar-injector.
  Removed ValidatingWebhookConfiguration::istio-validator-istio-system.
  Removed ValidatingWebhookConfiguration::istiod-default-validator.
  Removed ClusterRole::istio-reader-clusterrole-istio-system.
  Removed ClusterRole::istio-reader-istio-system.
  Removed ClusterRole::istiod-clusterrole-istio-system.
  Removed ClusterRole::istiod-gateway-controller-istio-system.
  Removed ClusterRole::istiod-istio-system.
  Removed ClusterRoleBinding::istio-reader-clusterrole-istio-system.
  Removed ClusterRoleBinding::istio-reader-istio-system.
  Removed ClusterRoleBinding::istiod-clusterrole-istio-system.
  Removed ClusterRoleBinding::istiod-gateway-controller-istio-system.
  Removed ClusterRoleBinding::istiod-istio-system.
  Removed CustomResourceDefinition::authorizationpolicies.security.istio.io.
  Removed CustomResourceDefinition::destinationrules.networking.istio.io.
  Removed CustomResourceDefinition::envoyfilters.networking.istio.io.
  Removed CustomResourceDefinition::gateways.networking.istio.io.
  Removed CustomResourceDefinition::istiooperators.install.istio.io.
  Removed CustomResourceDefinition::peerauthentications.security.istio.io.
  Removed CustomResourceDefinition::requestauthentications.security.istio.io.
  Removed CustomResourceDefinition::serviceentries.networking.istio.io.
  Removed CustomResourceDefinition::sidecars.networking.istio.io.
  Removed CustomResourceDefinition::telemetries.telemetry.istio.io.
  Removed CustomResourceDefinition::virtualservices.networking.istio.io.
  Removed CustomResourceDefinition::wasmplugins.extensions.istio.io.
  Removed CustomResourceDefinition::workloadentries.networking.istio.io.
  Removed CustomResourceDefinition::workloadgroups.networking.istio.io.
✔ Uninstall complete 


 % istioctl install --set profile=demo -y    
✔ Istio core installed                                                                                                                  
✔ Istiod installed         

https://istio.io/v1.4/docs/reference/config/installation-options/#global-options

istioctl manifest generate ... --set values.global.arch.amd64=3


 % istioctl install --set profile=demo --set values.global.arch.amd64=3 -y    
! values.global.arch is deprecated; use the affinity of k8s settings instead
✔ Istio core installed                                                                                                                  
✔ Istiod installed                                                                                                                      
- Processing resources for Egress gateways, Ingress gateways. Waiting for Deployment/istio-system/istio-egressgateway, Deployment/ist...^C

https://istio.io/latest/docs/setup/install/operator/#install


 % kubectl -n istio-system get pods
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-6ff67579c4-zzmrk    1/1     Running   0          39s
istio-ingressgateway-67b86984bb-nklhh   1/1     Running   0          20s
istiod-76d66d9876-8wdvt                 1/1     Running   0          22m


https://github.com/nowandme/k8s-istio-m1/blob/main/install-istio.yaml



% kubectl apply -f  ./istio-1.12.1/samples/bookinfo/platform/kube/bookinfo.yaml
service/details created
serviceaccount/bookinfo-details created
deployment.apps/details-v1 created
service/ratings created
serviceaccount/bookinfo-ratings created
deployment.apps/ratings-v1 created
service/reviews created
serviceaccount/bookinfo-reviews created
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage created
serviceaccount/bookinfo-productpage created
deployment.apps/productpage-v1 created



%  kubectl get services
NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
apple-headless-service               ClusterIP      None             <none>        5678/TCP                        28d
details                              ClusterIP      10.99.11.211     <none>        9080/TCP                        16s
ingress-nginx-controller             LoadBalancer   10.108.133.209   localhost     80:30701/TCP,443:32342/TCP      127d
ingress-nginx-controller-admission   ClusterIP      10.99.126.26     <none>        443/TCP                         127d
kubernetes                           ClusterIP      10.96.0.1        <none>        443/TCP                         133d
matrix-deployment                    ClusterIP      10.101.246.185   <none>        12345/TCP                       10d
minio-service                        NodePort       10.99.62.36      <none>        9000:30000/TCP,9001:30001/TCP   126d
productpage                          ClusterIP      10.106.83.156    <none>        9080/TCP                        15s
ratings                              ClusterIP      10.97.116.200    <none>        9080/TCP                        16s
redis                                NodePort       10.105.105.140   <none>        6379:30379/TCP                  126d
reviews                              ClusterIP      10.106.107.230   <none>        9080/TCP                        16s
sidecar-injector-webhook-svc         NodePort       10.105.102.109   <none>        7896:30896/TCP                  22d



% kubectl get pods
NAME                                                    READY   STATUS                  RESTARTS   AGE
admission-webhook-example-deployment-7b996d5c64-ddvgp   1/1     Running                 12         22d
alpine-job-xhlqk                                        0/1     Completed               0          28d
apple-configmap-volume                                  1/1     Running                 21         28d
apple-deployment-fb8cfb965-xsqbd                        1/1     Running                 13         29d
apple-rc-k88km                                          1/1     Running                 13         29d
apple-rc-mx825                                          1/1     Running                 1          12d
apple-rc-wgd4c                                          1/1     Running                 13         29d
apple-rctl-2hdpt                                        1/1     Running                 13         29d
apple-rctl-wrmwp                                        1/1     Running                 13         29d
apple-rs-7hzk9                                          1/1     Running                 13         29d
apple-rs-cjs27                                          1/1     Running                 13         29d
apple-rs-x44pd                                          1/1     Running                 13         29d
apple-ss-0                                              1/1     Running                 13         29d
apple-ss-1                                              1/1     Running                 13         29d
apple-ss-2                                              1/1     Running                 13         29d
details-v1-79f774bdb9-w55hv                             0/2     Init:CrashLoopBackOff   6          13m
dnsutils                                                1/1     Running                 287        36d
example-foo-9bbb75dc8-v7dzv                             1/1     Running                 16         60d
ingress-nginx-admission-create-vjn92                    0/1     Completed               0          127d
ingress-nginx-admission-patch-wlq6p                     0/1     Completed               0          127d
ingress-nginx-controller-57648496fc-84wl8               1/1     Running                 38         127d
matrix-deployment-5ccd76bb66-479v2                      1/1     Running                 0          10d
minio-deployment-55bf5bff5d-cvq7v                       1/1     Running                 33         126d
productpage-v1-6b746f74dc-564l5                         0/2     Init:CrashLoopBackOff   6          13m
ratings-v1-b6994bb9-vl8pn                               0/2     Init:CrashLoopBackOff   6          13m
redis-f9f74787-tq6tw                                    1/1     Running                 33         126d
reviews-v1-545db77b95-zstgv                             0/2     Init:CrashLoopBackOff   6          13m
reviews-v2-7bf8c9648f-wjmn8                             0/2     Init:CrashLoopBackOff   6          13m
reviews-v3-84779c7bbc-cpkh6                             0/2     Init:CrashLoopBackOff   6          13m


 % kubectl describe pod reviews-v2-7bf8c9648f-wjmn8 
  Warning  BackOff    12m (x10 over 14m)  kubelet            Back-off restarting failed container

 % kubectl logs reviews-v2-7bf8c9648f-wjmn8  
Error from server (BadRequest): container "reviews" in pod "reviews-v2-7bf8c9648f-wjmn8" is waiting to start: PodInitializing

%  kubectl get deploy productpage-v1 -o yaml > productpage-v1.yaml 
% docker pull docker.io/istio/examples-bookinfo-productpage-v1:1.16.2

%  kubectl get deploy details-v1 -o yaml > details-v1.yaml
% docker pull docker.io/istio/examples-bookinfo-details-v1:1.16.2

%  kubectl get deploy reviews-v1 -o yaml > reviews-v1.yaml
% docker pull docker.io/istio/examples-bookinfo-reviews-v1:1.16.2

%  kubectl get deploy reviews-v2 -o yaml > reviews-v2.yaml 
% docker pull docker.io/istio/examples-bookinfo-reviews-v2:1.16.2

%  kubectl get deploy reviews-v3 -o yaml > reviews-v3.yaml      
% docker pull docker.io/istio/examples-bookinfo-reviews-v3:1.16.2

 % kubectl get deploy ratings-v1 -o yaml > ratings-v1.yaml
% docker pull docker.io/istio/examples-bookinfo-ratings-v1:1.16.2


 kubectl -n istio-system logs istio-egressgateway-6ff67579c4-j4clf
 size:1.1kB resource:ROOTCA
2022-01-03T12:06:10.417036Z     info    ads     SDS: PUSH request for node:istio-egressgateway-6ff67579c4-j4clf.istio-system resources:1 size:4.0kB resource:default
[mutex.cc : 926] RAW: pthread_getschedparam failed: 1


% kubectl exec "$(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -s productpage:9080/productpage | grep -o "<title>.*</title>"


% kubectl apply -f ./istio-1.12.1/samples/bookinfo/networking/bookinfo-gateway.yaml
gateway.networking.istio.io/bookinfo-gateway created
virtualservice.networking.istio.io/bookinfo created

https://istio.io/latest/zh/docs/setup/getting-started/



% istioctl analyze
Warning [IST0103] (Pod default/admission-webhook-example-deployment-7b996d5c64-ddvgp) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.

 % istioctl   verify-install  -f ./istio-1.12.1/manifests/profiles/demo.yaml
✔ Deployment: istio-ingressgateway.istio-system checked successfully
✔ PodDisruptionBudget: istio-ingressgateway.istio-system checked successfully
✔ Role: istio-ingressgateway-sds.istio-system checked successfully
✔ RoleBinding: istio-ingressgateway-sds.istio-system checked successfully
✔ Service: istio-ingressgateway.istio-system checked successfully
✔ ServiceAccount: istio-ingressgateway-service-account.istio-system checked successfully
✔ Deployment: istio-egressgateway.istio-system checked successfully
✔ PodDisruptionBudget: istio-egressgateway.istio-system checked successfully
✔ Role: istio-egressgateway-sds.istio-system checked successfully
✔ RoleBinding: istio-egressgateway-sds.istio-system checked successfully
✔ Service: istio-egressgateway.istio-system checked successfully
✔ ServiceAccount: istio-egressgateway-service-account.istio-system checked successfully
✔ ClusterRole: istiod-istio-system.istio-system checked successfully
✔ ClusterRole: istio-reader-istio-system.istio-system checked successfully
✔ ClusterRoleBinding: istio-reader-istio-system.istio-system checked successfully
✔ ClusterRoleBinding: istiod-istio-system.istio-system checked successfully
✔ ServiceAccount: istio-reader-service-account.istio-system checked successfully
✔ Role: istiod-istio-system.istio-system checked successfully
✔ RoleBinding: istiod-istio-system.istio-system checked successfully
✔ ServiceAccount: istiod-service-account.istio-system checked successfully
✔ CustomResourceDefinition: wasmplugins.extensions.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: destinationrules.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: envoyfilters.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: gateways.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: serviceentries.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: sidecars.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: virtualservices.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: workloadentries.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: workloadgroups.networking.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: authorizationpolicies.security.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: peerauthentications.security.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: requestauthentications.security.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: telemetries.telemetry.istio.io.istio-system checked successfully
✔ CustomResourceDefinition: istiooperators.install.istio.io.istio-system checked successfully
✔ ClusterRole: istiod-clusterrole-istio-system.istio-system checked successfully
✔ ClusterRole: istiod-gateway-controller-istio-system.istio-system checked successfully
✔ ClusterRoleBinding: istiod-clusterrole-istio-system.istio-system checked successfully
✔ ClusterRoleBinding: istiod-gateway-controller-istio-system.istio-system checked successfully
✔ ConfigMap: istio.istio-system checked successfully
✔ Deployment: istiod.istio-system checked successfully
✔ ConfigMap: istio-sidecar-injector.istio-system checked successfully
✔ MutatingWebhookConfiguration: istio-sidecar-injector.istio-system checked successfully
✔ PodDisruptionBudget: istiod.istio-system checked successfully
✔ ClusterRole: istio-reader-clusterrole-istio-system.istio-system checked successfully
✔ ClusterRoleBinding: istio-reader-clusterrole-istio-system.istio-system checked successfully
✔ Role: istiod.istio-system checked successfully
✔ RoleBinding: istiod.istio-system checked successfully
✔ Service: istiod.istio-system checked successfully
✔ ServiceAccount: istiod.istio-system checked successfully
✔ EnvoyFilter: stats-filter-1.10.istio-system checked successfully
✔ EnvoyFilter: tcp-stats-filter-1.10.istio-system checked successfully
✔ EnvoyFilter: stats-filter-1.11.istio-system checked successfully
✔ EnvoyFilter: tcp-stats-filter-1.11.istio-system checked successfully
✔ EnvoyFilter: stats-filter-1.12.istio-system checked successfully
✔ EnvoyFilter: tcp-stats-filter-1.12.istio-system checked successfully
✔ ValidatingWebhookConfiguration: istio-validator-istio-system.istio-system checked successfully
✔ IstioOperator: .istio-system checked successfully
Checked 14 custom resource definitions
Checked 3 Istio Deployments
✔ Istio is installed and verified successfully



 % kubectl get pod details-v1-79f774bdb9-2lhzg -o yaml
  state:
      waiting:
        message: back-off 5m0s restarting failed container=istio-init pod=details-v1-79f774bdb9-2lhzg_default(c46c76ba-e500-4680-b9f2-54f0cee4c3ab)
        reason: CrashLoopBackOff



% istioctl version --remote                          
client version: 1.12.1
control plane version: 1.12.1
data plane version: 1.12.1 (2 proxies)


./bin/istioctl manifest generate --set profile=demo --set values.global.proxy.privileged=true | kubectl apply -f -

We need to use a filter to exclude istio-init container, in case you are using docker system prune command:
docker system prune -af --volumes --filter "label!=io.kubernetes.container.name=istio-init"

https://github.com/istio/istio/issues/19717

