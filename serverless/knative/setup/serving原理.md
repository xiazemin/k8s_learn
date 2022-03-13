 https://zhuanlan.zhihu.com/p/139592932

 Knative Serving 用户应该熟悉以下四个主要资源： Service（服务）、 Route（路由）、Configuration（配置）和 Revision（修订）。

 Service：service.serving.knative.dev资源会自动管理您的工作负载的整个生命周期。它控制其他对象的创建，以确保您的应用为服务的每次更新都具有路由，配置和新修订版。可以将服务定义为始终将流量路由到最新修订版或固定修订版。
Route：route.serving.knative.dev资源将网络端点映射到一个或多个修订版。您可以通过几种方式管理流量，包括部分流量和命名路由。
Configuration：configuration.serving.knative.dev资源维护部署的所需状态。它在代码和配置之间提供了清晰的分隔，并遵循了十二要素应用程序方法。修改配置会创建一个新修订。
Revision：reversion.serving.knative.dev资源是对工作负载进行的每次修改的代码和配置的时间点快照。修订是不可变的对象，可以保留很长时间。可以根据传入流量自动缩放实例数。

https://zhuanlan.zhihu.com/p/139592932

Route（路由）
Knative 中的 Route 提供了一种将流量路由到正在运行的代码的机制。它将一个命名的，HTTP 可寻址端点映射到一个或者多个 Revision。Configuration 本身并不定义 Route。

Autoscaler（自动伸缩器）和 Activator（激活器）
Serverless 的一个关键原则是可以按需扩容以满足需要和缩容以节省资源。Serverless 负载应当可以一直缩容至零。这意味着如果没有请求进入，则不会运行容器实例。Knative 使用两个关键组件以实现该功能。它将 Autoscaler 和 Activator 实现为集群中的 Pod。您可以看到它们伴随其他 Serving 组件一起运行在 knative-serving 命名空间中

https://blog.csdn.net/weixin_33955681/article/details/91371087

我们可以分为五层来考虑：观测层(Queue和Activator)、决策层(Autoscaler)、控制层(Controller)、准入层(Webhook)、路由层(Istio INgress)

https://www.cnblogs.com/buyicoding/p/12767356.html


%  kubectl get all -n knative-serving |grep '0/1'
pod/default-domain--1-54fw6                  0/1     Error     0              25h
pod/default-domain--1-5lx8v                  0/1     Error     0              23h
pod/default-domain--1-5vn9t                  0/1     Error     0              22h
pod/default-domain--1-6d5xf                  0/1     Error     0              24h
pod/default-domain--1-9jhb2                  0/1     Error     0              22h
pod/default-domain--1-bkbt8                  0/1     Error     0              6d14h
pod/default-domain--1-d6smw                  0/1     Error     0              22h
pod/default-domain--1-frsvg                  0/1     Error     0              24h
pod/default-domain--1-jxgth                  0/1     Error     0              23h
pod/default-domain--1-lqwdz                  0/1     Error     0              25h
pod/default-domain--1-v8khv                  0/1     Error     0              24h
job.batch/default-domain   0/1           6d14h      6d14h

 % kubectl  -n knative-serving get pod |grep '0/1' |awk '{print$1}'|xargs -I{} kubectl  -n knative-serving delete pod {}


 % kubectl -n knative-serving delete job default-domain  
job.batch "default-domain" deleted


启动默认的例子
serverless/knative/docs/code-samples/serving/hello-world/helloworld-go/service.yaml


% kubectl -n knative-go get all 
NAME                                        URL   LATESTCREATED   LATESTREADY   READY   REASON
service.serving.knative.dev/helloworld-go                                               

NAME                                               CONFIG NAME     K8S SERVICE NAME   GENERATION   READY   REASON   ACTUAL REPLICAS   DESIRED REPLICAS
revision.serving.knative.dev/helloworld-go-00001   helloworld-go                      1                                               

NAME                                      URL   READY   REASON
route.serving.knative.dev/helloworld-go                 

NAME                                              LATESTCREATED   LATESTREADY   READY   REASON
configuration.serving.knative.dev/helloworld-go                                         



https://developer.aliyun.com/article/739581?spm=a2c6h.12883283.1362933.32.9ca5201cxTTveu
https://www.shangmayuan.com/a/b13b46a44f2a43d0b655a322.html

%   kubectl get ksvc helloworld-go -n knative-go
NAME            URL   LATESTCREATED   LATESTREADY   READY   REASON
helloworld-go                        

发现没有路由这是因为我们的ns knative-go 没有开启自动路由注入，istio不会注入

% kubectl label namespace knative-go istio-injection=enabled
namespace/knative-go labeled


% kubectl -n  knative-go apply  -f serverless/knative/docs/code-samples/serving/hello-world/helloworld-go/service.yaml
service.serving.knative.dev/helloworld-go created


%  kubectl -n knative-go describe route.serving.knative.dev/helloworld-go
  Warning  UpdateFailed     63s (x15 over 2m25s)  route-controller  Failed to update status for "helloworld-go": admission webhook "webhook.serving.knative.dev" denied the request: mutation failed: cannot decode incoming new object: json: unknown field "subresource"

  https://github.com/ubuntu/microk8s/issues/2720
https://github.com/knative/serving/issues/11448
https://github.com/knative/serving/pull/11805

This is caused by apiextensions.k8s.io/v1beta1 CustomResiurceDefinition being unavailable in Kubernetes v1.22+ found in net-istio.yaml. Migrating to apiextensions.k8s.io/v1 CustomResiurceDefinition` solves the issue.


% kubectl apply -f serverless/knative/knative-serving/net-istio.yaml


 % kn version
Version:      v1.1.0
Build Date:   2021-12-14 13:59:14
Git Revision: 530841f1
Supported APIs:
* Serving
  - serving.knative.dev/v1 (knative-serving v0.28.0)
* Eventing
  - sources.knative.dev/v1 (knative-eventing v0.28.0)
  - eventing.knative.dev/v1 (knative-eventing v0.28.0)



Knative Serving defines a set of objects as Kubernetes Custom Resource Definitions (CRDs). These objects are used to define and control how your serverless workload behaves on the cluster:

Service: The service.serving.knative.dev resource automatically manages the whole lifecycle of your workload. It controls the creation of other objects to ensure that your app has a route, a configuration, and a new revision for each update of the service. Service can be defined to always route traffic to the latest revision or to a pinned revision.
Route: The route.serving.knative.dev resource maps a network endpoint to one or more revisions. You can manage the traffic in several ways, including fractional traffic and named routes.
Configuration: The configuration.serving.knative.dev resource maintains the desired state for your deployment. It provides a clean separation between code and configuration and follows the Twelve-Factor App methodology. Modifying a configuration creates a new revision.
Revision: The revision.serving.knative.dev resource is a point-in-time snapshot of the code and configuration for each modification made to the workload. Revisions are immutable objects and can be retained for as long as useful. Knative Serving Revisions can be automatically scaled up and down according to incoming traffic. See Configuring the Autoscaler for more information.

https://knative.dev/docs/serving/



https://knative.dev/docs/serving/services/private-services/
% kubectl  -n  knative-go label kservice helloworld-go networking.knative.dev/visibility=cluster-local
service.serving.knative.dev/helloworld-go labeled

 % kubectl -n  knative-go get kservice helloworld-go 
NAME            URL   LATESTCREATED   LATESTREADY   READY   REASON
helloworld-go                   

https://github.com/knative-sandbox/sample-controller

https://developer.aliyun.com/article/739581?spm=a2c6h.12883283.1362933.32.9ca5201cxTTveu

https://github.com/google/go-containerregistry

https://github.com/knative/serving/issues/11448

https://github.com/knative/pkg/pull/2249

https://app.codecov.io/login/gh

https://about.codecov.io/
https://www.jianshu.com/p/146c4769d4b1
https://github.com/baukh789/GridManager


 % kubectl -n knative-go describe pod service.serving.knative.dev/helloworld-go
error: there is no need to specify a resource type as a separate argument when passing arguments in resource/name form (e.g. 'kubectl get resource/<resource_name>' instead of 'kubectl get resource resource/<resource_name>'


 % kubectl -n knative-go describe  service.serving.knative.dev/helloworld-go
  Warning  UpdateFailed  45s (x29 over 75m)  service-controller  Failed to update status for "helloworld-go": admission webhook "webhook.serving.knative.dev" denied the request: mutation failed: cannot decode incoming new object: json: unknown field "subresource"


docker.io/fandy00001/knative-serving-activator:latest
docker.io/fandy00001/knative-serving-autoscaler:latest


% docker pull shixiongqi/activator-ecd51ca5034883acbe737fde417a3d86:latest
latest: Pulling from shixiongqi/activator-ecd51ca5034883acbe737fde417a3d86
Digest: sha256:a1da9f52a48de84217b3c43f9854cdc13981a6a0d4d3bf16a363e7d7c7544f59
Status: Image is up to date for shixiongqi/activator-ecd51ca5034883acbe737fde417a3d86:latest
docker.io/shixiongqi/activator-ecd51ca5034883acbe737fde417a3d86:latest

docker pull shixiongqi/autoscaler-12c0fa24db31956a7cfa673210e4fa13:latest
latest: Pulling from shixiongqi/autoscaler-12c0fa24db31956a7cfa673210e4fa13
2df365faf0e3: Already exists 
aa91ba1b569d: Already exists 
160b5dc0a20d: Pull complete 
Digest: sha256:7c180dbd86f6d92b2f08bb6ceb8f6ed539fa7420abe2d08529410bd3c66be9a0
Status: Downloaded newer image for shixiongqi/autoscaler-12c0fa24db31956a7cfa673210e4fa13:latest
docker.io/shixiongqi/autoscaler-12c0fa24db31956a7cfa673210e4fa13:latest



% docker pull jzy99/net-istio-controller:v1.0.0
v1.0.0: Pulling from jzy99/net-istio-controller
e8614d09b7be: Pull complete 
f91ca0366a0d: Pull complete 
761e931ab2a4: Pull complete 
Digest: sha256:f6787518cac0b5bdf3fb4dcc829263140341f0fb9aacbe988ba81c07cd3b0780
Status: Downloaded newer image for jzy99/net-istio-controller:v1.0.0
docker.io/jzy99/net-istio-controller:v1.0.0

% docker pull jzy99/net-istio-webhook:v1.0.0
v1.0.0: Pulling from jzy99/net-istio-webhook
e8614d09b7be: Already exists 
f91ca0366a0d: Already exists 
264da01863fd: Pull complete 
Digest: sha256:25c3ccb5c0ee008d35394bf3881cfdf8f1e12e021e61e4daff9fcfc5b77078ff
Status: Downloaded newer image for jzy99/net-istio-webhook:v1.0.0
docker.io/jzy99/net-istio-webhook:v1.0.0


替换掉
docker.io/xunfeng/net-istio-controller:v0.17.3
docker.io/xunfeng/knative-net-istio-webhook:v0.25.2

 kubectl  -n knative-go apply -f docs/code-samples/serving/hello-world/helloworld-go/service.yaml

 https://zhuanlan.zhihu.com/p/141551023

  % kn -n knative-go service list
NAME            URL   LATEST   AGE   CONDITIONS   READY       REASON
helloworld-go                  27m   0 OK / 0     <unknown>   <unknown>


% kubectl -n knative-go logs route.serving.knative.dev/helloworld-go
error: no kind "Route" is registered for version "serving.knative.dev/v1" in scheme "k8s.io/kubectl/pkg/scheme/scheme.go:28"

% kubectl -n knative-go describe route.serving.knative.dev/helloworld-go

Warning  UpdateFailed     32m                   route-controller  Failed to update status for "helloworld-go": Internal error occurred: failed calling webhook "webhook.serving.knative.dev": failed to call webhook: Post "https://webhook.knative-serving.svc:443/defaulting?timeout=10s": dial tcp 10.102.237.123:443: connect: connection refused
  Warning  UpdateFailed     2m15s (x27 over 38m)  route-controller  Failed to update status for "helloworld-go": admission webhook "webhook.serving.knative.dev" denied the request: mutation failed: cannot decode incoming new object: json: unknown field "subresource"


 % kubectl -n knative-serving describe pod webhook-5468c9bd74-mvb7p
  Warning  Unhealthy  4m26s (x2 over 37m)  kubelet            Liveness probe failed: Get "https://10.1.1.134:8443/": context deadline exceeded


% kubectl -n knative-serving  get ValidatingWebhookConfiguration
NAME                                                   WEBHOOKS   AGE
config.webhook.eventing.knative.dev                    1          7d
config.webhook.istio.networking.internal.knative.dev   1          4h40m
config.webhook.pipeline.tekton.dev                     1          7d2h
config.webhook.serving.knative.dev                     1          7d2h
istio-validator-istio-system                           1          6d20h
istiod-default-validator                               1          7d3h
validation.webhook.domainmapping.serving.knative.dev   1          7d2h
validation.webhook.eventing.knative.dev                1          7d
validation.webhook.pipeline.tekton.dev                 1          7d2h
validation.webhook.serving.knative.dev                 1          7d2h


 % kubectl -n knative-serving get ValidatingWebhookConfiguration config.webhook.serving.knative.dev  -o yaml > config.webhook.serving.knative.dev.yaml
 % kubectl -n knative-serving get ValidatingWebhookConfiguration validation.webhook.serving.knative.dev  -o yaml > validation.webhook.serving.knative.dev.yaml  


 % kubectl get all -o yaml |grep 'apiextensions.k8s.io'

% kubectl api-versions |grep extensions
apiextensions.k8s.io/v1
extensions.istio.io/v1alpha1

% kubectl api-versions |grep serving 
serving.knative.dev/v1
serving.knative.dev/v1alpha1
serving.knative.dev/v1beta1

 % kubectl api-versions |grep eventing
eventing.knative.dev/v1
eventing.knative.dev/v1beta1


serving.knative.dev/v1 =》
serving.knative.dev/v1beta1

 % kubectl -n knative-go apply -f ../docs/code-samples/serving/hello-world/helloworld-go/service.yaml

error: unable to recognize "../docs/code-samples/serving/hello-world/helloworld-go/service.yaml": no matches for kind "Service" in version "serving.knative.dev/v1beta1"

https://www.cnblogs.com/zhaowei121/p/11995522.html
https://www.infoq.cn/article/ja2hpZCLxMiswx2Nm2I1/

serving.knative.dev/v1 =》
serving.knative.dev/v1alpha1

% kubectl -n knative-go apply -f ../docs/code-samples/serving/hello-world/helloworld-go/service.yaml
error: unable to recognize "../docs/code-samples/serving/hello-world/helloworld-go/service.yaml": no matches for kind "Service" in version "serving.knative.dev/v1alpha1"



serverless/knative/knative-serving/validation.webhook.serving.knative.dev.yaml
  - serving.knative.dev
    apiVersions:
    - v1

  - serving.knative.dev
    apiVersions:
    - v1
    - v1alpha1
    - v1beta1

% kubectl -n knative-go apply -f ../docs/code-samples/serving/hello-world/helloworld-go/service.yaml
error: unable to recognize "../docs/code-samples/serving/hello-world/helloworld-go/service.yaml": no matches for kind "Service" in version "serving.knative.dev/v1alpha1"

https://github.com/knative/serving/issues/8812
https://github.com/knative/serving/issues/2878


https://github.com/knative/serving/releases

https://github.com/knative-sandbox/net-kourier

kubectl apply -f https://github.com/knative/net-kourier/releases/latest/download/kourier.yaml

https://github.com/knative/serving/releases/tag/v0.25.2
https://github.com/knative-sandbox/net-kourier


% kubectl apply -f  kourier.yaml  
namespace/kourier-system created
configmap/kourier-bootstrap created
configmap/config-kourier created
serviceaccount/net-kourier created
clusterrole.rbac.authorization.k8s.io/net-kourier created
clusterrolebinding.rbac.authorization.k8s.io/net-kourier created
deployment.apps/net-kourier-controller created
service/net-kourier-controller created
deployment.apps/3scale-kourier-gateway created
service/kourier created
service/kourier-internal created



% kubectl get all -n kourier-system
NAME                                          READY   STATUS              RESTARTS   AGE
pod/3scale-kourier-gateway-58856c6cc7-r6vcl   0/1     ContainerCreating   0          91s

NAME                       TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
service/kourier            LoadBalancer   10.97.1.29     <pending>     80:30801/TCP,443:32191/TCP   87s
service/kourier-internal   ClusterIP      10.96.28.155   <none>        80/TCP                       80s

NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/3scale-kourier-gateway   0/1     1            0           93s

NAME                                                DESIRED   CURRENT   READY   AGE
replicaset.apps/3scale-kourier-gateway-58856c6cc7   1         1         0       92s



 % kubectl -n kourier-system describe pod 3scale-kourier-gateway-58856c6cc7-r6vcl 
  Warning  Unhealthy  4s (x17 over 79s)  kubelet            Readiness probe failed: Get "http://10.1.1.143:8081/ready": dial tcp 10.1.1.143:8081: connect: connection refused


 % kubectl patch configmap/config-network \
  -n knative-serving \
  --type merge \
  -p '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
configmap/config-network patched


 % kubectl patch configmap/config-domain \
  -n knative-serving \
  --type merge \
  -p '{"data":{"127.0.0.1.nip.io":""}}'
configmap/config-domain patched


kubectl port-forward --namespace kourier-system $(kubectl get pod -n kourier-system -l "app=3scale-kourier-gateway" --output=jsonpath="{.items[0].metadata.name}") 8080:8080 19000:19000 8443:8443

curl -v -H "Host: helloworld-go.default.127.0.0.1.nip.io" http://localhost:8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Forwarding from 127.0.0.1:19000 -> 19000
Forwarding from [::1]:19000 -> 19000
Forwarding from 127.0.0.1:8443 -> 8443
Forwarding from [::1]:8443 -> 8443

https://github.com/knative-sandbox/net-kourier/issues/111


% kubectl label namespace kourier-system istio-injection=enabled
namespace/kourier-system labeled

 %  kubectl  -n kourier-system delete pod 3scale-kourier-gateway-58856c6cc7-r6vcl
pod "3scale-kourier-gateway-58856c6cc7-r6vcl" deleted


 %  kubectl  get pod -n kourier-system
NAME                                      READY   STATUS    RESTARTS   AGE
3scale-kourier-gateway-58856c6cc7-p9ljt   1/2     Running   0          34s

 % kubectl  -n kourier-system logs  3scale-kourier-gateway-58856c6cc7-p9ljt 
[2022-01-24 15:50:50.935][1][warning][config] [bazel-out/k8-opt/bin/source/common/config/_virtual_includes/grpc_stream_lib/common/config/grpc_stream.h:101] StreamAggregatedResources gRPC config stream closed: 14, upstream connect error or disconnect/reset before headers. reset reason: connection failure, transport failure reason: delayed connect error: 111

https://github.com/knative-sandbox/net-kourier/issues/111

kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.1.0/kourier.yaml



% kubectl -n knative-serving get pod net-kourier-controller-7df5866d8d-q2j4b  -o yaml |grep image
    image: gcr.io/knative-releases/knative.dev/net-kourier/cmd/kourier@sha256:d10a5b94d1d9f59156493cc13ec710b4354c43b0aba71d7d57a43bc33024a715
    imagePullPolicy: IfNotPresent
  - image: gcr.io/knative-releases/knative.dev/net-kourier/cmd/kourier@sha256:d10a5b94d1d9f59156493cc13ec710b4354c43b0aba71d7d57a43bc33024a715
    imageID: ""
        message: Back-off pulling image "gcr.io/knative-releases/knative.dev/net-kourier/cmd/kourier@sha256:d10a5b94d1d9f59156493cc13ec710b4354c43b0aba71d7d57a43bc33024a715"



% docker pull ramydocker/kourier-b74c3918b7eee585f87df62ccd297dc8:latest
latest: Pulling from ramydocker/kourier-b74c3918b7eee585f87df62ccd297dc8
2df365faf0e3: Already exists 
250c06f7c38e: Pull complete 
ddf4df90d680: Pull complete 
Digest: sha256:56001126fe0c39615567229da66e0e30e53f69289656bca5bd42603b05c7de27
Status: Downloaded newer image for ramydocker/kourier-b74c3918b7eee585f87df62ccd297dc8:latest
docker.io/ramydocker/kourier-b74c3918b7eee585f87df62ccd297dc8:latest


gcr.io/knative-releases/knative.dev/net-kourier/cmd/kourier@sha256:d10a5b94d1d9f59156493cc13ec710b4354c43b0aba71d7d57a43bc33024a715

serverless/knative/setup/kourier.yaml

 % kubectl --namespace kourier-system get pod            
NAME                                      READY   STATUS    RESTARTS   AGE
3scale-kourier-gateway-58856c6cc7-p9ljt   2/2     Running   0          13m


% kubectl -n knative-go apply -f docs/code-samples/serving/hello-world/helloworld-go/service.yamlservice.serving.knative.dev/helloworld-go created


kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.1.0/serving-default-domain.yaml


% kubectl apply -f serving-default-domain.yaml
job.batch/default-domain created
service/default-domain-service unchanged


https://knative.dev/docs/install/serving/install-serving-with-yaml/

% kubectl  -n knative-serving describe pod net-kourier-controller-c995f7b9c-tjtf7  
  Warning  Unhealthy  19m                   kubelet            Readiness probe failed: 2022/01/24 16:01:21 failed to connect to service at ":18000": context deadline exceeded
  Warning  Unhealthy  4m13s (x51 over 19m)  kubelet            Readiness probe failed:

  https://github.com/knative/serving/issues/11448
  