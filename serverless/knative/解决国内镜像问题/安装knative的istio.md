https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#install-the-knative-serving-component


1,
kubectl apply -l knative.dev/crd-install=true -f https://github.com/knative/net-istio/releases/download/knative-v1.3.0/istio.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.3.0/istio.yaml


2,kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.3.0/net-istio.yaml


3,kubectl --namespace istio-system get service istio-ingressgateway



 % kubectl --namespace istio-system get service istio-ingressgateway

NAME                   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   NodePort   10.104.14.138   <none>        15021:31594/TCP,80:31760/TCP,443:30272/TCP,31400:30137/TCP,15443:32583/TCP   43h



https://knative.dev/docs/install/installing-istio/



由于可能和现有的istio冲突，所以选择安装krouter

kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.3.0/kourier.yaml


#https://blog.51cto.com/u_14625168/2526945



The following commands install Kourier and enable its Knative integration.

Install the Knative Kourier controller by running the command:


kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.3.0/kourier.yaml
Configure Knative Serving to use Kourier by default by running the command:


kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
Fetch the External IP address or CNAME by running the command:


kubectl --namespace kourier-system get service kourier




 % kubectl apply -f serverless/knative/解决国内镜像问题/kourier.yaml
namespace/kourier-system created
configmap/kourier-bootstrap created
configmap/config-kourier created
serviceaccount/net-kourier created
clusterrole.rbac.authorization.k8s.io/net-kourier configured
clusterrolebinding.rbac.authorization.k8s.io/net-kourier configured
deployment.apps/net-kourier-controller created
service/net-kourier-controller created
deployment.apps/3scale-kourier-gateway created
service/kourier created
service/kourier-internal created



依赖两个镜像
- image: gcr.io/knative-releases/knative.dev/net-kourier/cmd/kourier@sha256:84af1fba93bcc1d504ee6fc110a49be80440f08d461ccb0702621b7b62d0f7b6

image: docker.io/envoyproxy/envoy:v1.18-latest



% kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'

configmap/config-network patched



% kubectl --namespace kourier-system get service kourier
NAME      TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier   LoadBalancer   10.108.41.244   localhost     80:31859/TCP,443:30753/TCP   34s


% kubectl get pods -n knative-serving
NAME                                      READY   STATUS    RESTARTS   AGE
activator-7cf4bd8548-gg5cm                1/1     Running   0          29m
autoscaler-577d766bdd-5xmkp               1/1     Running   0          29m
controller-5b74bfcc9f-z6kpf               1/1     Running   0          29m
domain-mapping-5b4f5f66b5-g8cmt           1/1     Running   0          29m
domainmapping-webhook-5d7fb6566d-59blp    1/1     Running   0          29m
net-kourier-controller-766c565d78-5gqpz   1/1     Running   0          50s
webhook-699fc555bf-4t9nk                  1/1     Running   0          29m


 % kubectl -n kourier-system  get pods
NAME                                      READY   STATUS    RESTARTS   AGE
3scale-kourier-gateway-5f96966d45-n5tgg   1/1     Running   0          115s


https://github.com/knative/docs/tree/main/code-samples/serving/hello-world/helloworld-go


go mod init helloworld
go mod tidy
docker build -t xiazemin/helloworld-go .
 => => naming to docker.io/xiazemin/helloworld-go   

% docker images |grep xiazemin
xiazemin/helloworld-go                     



% kubectl apply -f helloworld.yaml
error: error validating "helloworld.yaml": error validating data: [ValidationError(Service): unknown field "name" in dev.knative.serving.v1.Service, ValidationError(Service): unknown field "namespace" in dev.knative.serving.v1.Service, ValidationError(Service): unknown field "template" in dev.knative.serving.v1.Service]; if you choose to ignore these errors, turn validation off with --validate=false



 % kubectl apply -f helloworld.yaml
namespace/helloworld created
error: unable to recognize "helloworld.yaml": no matches for kind "Service" in version "serving.knative.dev/v1alpha1"


% kubectl apply -f helloworld.yaml
namespace/helloworld unchanged
error: error validating "helloworld.yaml": error validating data: ValidationError(Service.spec.template): unknown field "annotations" in dev.knative.serving.v1.Service.spec.template; if you choose to ignore these errors, turn validation off with --validate=false


% kubectl apply -f helloworld.yaml
namespace/helloworld unchanged
Error from server (BadRequest): error when creating "helloworld.yaml": admission webhook "webhook.serving.knative.dev" denied the request: mutation failed: cannot decode incoming new object: json: unknown field "env"


% kubectl apply -f helloworld.yaml
namespace/helloworld unchanged
service.serving.knative.dev/hello created


 % kubectl get ksvc -n helloworld
NAME    URL                                   LATESTCREATED   LATESTREADY   READY   REASON
hello   http://hello.helloworld.example.com   hello-00001                   False   RevisionMissing


% kubectl -n helloworld  get ksvc hello --output=custom-columns=NAME:.metadata.name,URL:.status.url
NAME    URL
hello   http://hello.helloworld.example.com

curl -H "Host: hello.helloworld.example.com"  xxxx

#https://blog.51cto.com/u_14625168/2526945

#https://github.com/knative/docs/tree/main/code-samples/serving/hello-world/helloworld-go

kn -n helloworld service create hello --image=docker.io/xiazemin/helloworld-go --env TARGET="Go Sample v1"

% kn -n helloworld service create hello --image=docker.io/xiazemin/helloworld-go --env TARGET="Go Sample v1"
Error: cannot create service 'hello' in namespace 'helloworld' because the service already exists and no --force option was given
Run 'kn --help' for usage


kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.3.0/serving-default-domain.yaml


需要替换镜像
gcr.io/knative-releases/knative.dev/serving/cmd/default-domain@sha256:a5ea2fa55e5cbf34f5cbf0c3080f082425c838c8a260863f25093747132574fd

% kubectl apply -f  serving-default-domain.yaml 
job.batch/default-domain created
service/default-domain-service created

Knative provides a Kubernetes Job called default-domain that configures Knative Serving to use sslip.io as the default DNS suffix.

% kubectl get pods -n knative-serving
NAME                                      READY   STATUS    RESTARTS   AGE
activator-7cf4bd8548-gg5cm                1/1     Running   0          60m
autoscaler-577d766bdd-5xmkp               1/1     Running   0          60m
controller-5b74bfcc9f-z6kpf               1/1     Running   0          60m
default-domain--1-47zcl                   0/1     Error     0          117s
default-domain--1-hzwpb                   0/1     Error     0          2m7s
default-domain--1-klc86                   0/1     Error     0          97s
default-domain--1-mt6mg                   0/1     Error     0          2m12s
default-domain--1-sxghr                   0/1     Error     0          57s
domain-mapping-5b4f5f66b5-g8cmt           1/1     Running   0          60m
domainmapping-webhook-5d7fb6566d-59blp    1/1     Running   0          60m
net-kourier-controller-766c565d78-5gqpz   1/1     Running   0          31m
webhook-699fc555bf-4t9nk                  1/1     Running   0          60m



% kubectl -n helloworld  get ksvc hello --output=custom-columns=NAME:.metadata.name,URL:.status.url
NAME    URL
hello   http://hello.helloworld.example.com


 % kubectl -n knative-serving describe pod default-domain--1-hzwpb
Name:         default-domain--1-hzwpb
Namespace:    knative-serving
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Sat, 26 Mar 2022 17:56:48 +0800
Labels:       app=default-domain
              app.kubernetes.io/component=default-domain-job
              app.kubernetes.io/name=knative-serving
              app.kubernetes.io/version=1.3.0
              controller-uid=46d17c63-0e80-4dfe-bad3-9189338cfccc
              job-name=default-domain
Annotations:  sidecar.istio.io/inject: false
Status:       Failed
IP:           10.1.4.225
IPs:
  IP:           10.1.4.225
Controlled By:  Job/default-domain
Containers:
  default-domain:
    Container ID:  docker://144e5e53e8c72f616b833d8487ebc4eb99789ce8d6b25e9adde472951ce7d8aa
    Image:         gcr.io/knative-releases/knative.dev/serving/cmd/default-domain:latest
    Image ID:      docker://sha256:29b9cec714ff4e6ebbfa5317eed558edfe7702936bafaae7ba1845076e9bbc74
    Port:          8080/TCP
    Host Port:     0/TCP
    Args:
      -magic-dns=sslip.io
    State:          Terminated
      Reason:       Error
      Exit Code:    1
      Started:      Sat, 26 Mar 2022 17:56:49 +0800
      Finished:     Sat, 26 Mar 2022 17:56:49 +0800
    Ready:          False
    Restart Count:  0
    Limits:
      cpu:     1
      memory:  1000Mi
    Requests:
      cpu:      100m
      memory:   100Mi
    Liveness:   http-get http://:8080/ delay=0s timeout=1s period=10s #success=1 #failure=6
    Readiness:  http-get http://:8080/ delay=0s timeout=1s period=10s #success=1 #failure=3
    Environment:
      POD_NAME:          default-domain--1-hzwpb (v1:metadata.name)
      SYSTEM_NAMESPACE:  knative-serving (v1:metadata.namespace)
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-v6jm7 (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  kube-api-access-v6jm7:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  3m4s  default-scheduler  Successfully assigned knative-serving/default-domain--1-hzwpb to docker-desktop
  Normal  Pulled     3m4s  kubelet            Container image "gcr.io/knative-releases/knative.dev/serving/cmd/default-domain:latest" already present on machine
  Normal  Created    3m4s  kubelet            Created container default-domain
  Normal  Started    3m4s  kubelet            Started container default-domain



% kubectl delete -f  serving-default-domain.yaml
job.batch "default-domain" deleted
service "default-domain-service" deleted



% kubectl get svc istio-ingressgateway --namespace istio-system 
NAME                   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   NodePort   10.104.14.138   <none>        15021:31594/TCP,80:31760/TCP,443:30272/TCP,31400:30137/TCP,15443:32583/TCP   44h


 % curl -iv http://127.0.0.1:31594
*   Trying 127.0.0.1:31594...
* Connected to 127.0.0.1 (127.0.0.1) port 31594 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:31594
> User-Agent: curl/7.77.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 404 Not Found
HTTP/1.1 404 Not Found
< date: Sat, 26 Mar 2022 10:05:21 GMT
date: Sat, 26 Mar 2022 10:05:21 GMT
< server: envoy
server: envoy
< content-length: 0
content-length: 0

< 
* Connection #0 to host 127.0.0.1 left intact


 % kubectl get route hello -n helloworld
NAME    URL                                   READY   REASON
hello   http://hello.helloworld.example.com   False   RevisionMissing

 % curl -H 'Host: hello.helloworld.example.com' http://127.0.0.1:31594    

#https://zhuanlan.zhihu.com/p/68664517


% kubectl --namespace kourier-system get service kourier

NAME      TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier   LoadBalancer   10.108.41.244   localhost     80:31859/TCP,443:30753/TCP   41m



 % kubectl -n helloworld get all 
NAME                              URL                                   READY   REASON
route.serving.knative.dev/hello   http://hello.helloworld.example.com   False   RevisionMissing

NAME                                       CONFIG NAME   K8S SERVICE NAME   GENERATION   READY   REASON             ACTUAL REPLICAS   DESIRED REPLICAS
revision.serving.knative.dev/hello-00001   hello                            1            False   ContainerMissing                     

NAME                                      LATESTCREATED   LATESTREADY   READY   REASON
configuration.serving.knative.dev/hello   hello-00001                   False   RevisionFailed

NAME                                URL                                   LATESTCREATED   LATESTREADY   READY   REASON
service.serving.knative.dev/hello   http://hello.helloworld.example.com   hello-00001                   False   RevisionMissing

#https://www.5axxw.com/questions/content/yczgi5



% kubectl -n helloworld get all
NAME                                      LATESTCREATED   LATESTREADY   READY   REASON
configuration.serving.knative.dev/hello   hello-00002                   False   RevisionFailed

NAME                                URL                                   LATESTCREATED   LATESTREADY   READY   REASON
service.serving.knative.dev/hello   http://hello.helloworld.example.com   hello-00002                   False   RevisionMissing

NAME                              URL                                   READY   REASON
route.serving.knative.dev/hello   http://hello.helloworld.example.com   False   RevisionMissing

NAME                                       CONFIG NAME   K8S SERVICE NAME   GENERATION   READY   REASON             ACTUAL REPLICAS   DESIRED REPLICAS
revision.serving.knative.dev/hello-00001   hello                            1            False   ContainerMissing                     
revision.serving.knative.dev/hello-00002   hello                            2            False   ContainerMissing                     




 % kubectl describe ksvc -n helloworld
Name:         hello
Namespace:    helloworld
Labels:       <none>
Annotations:  serving.knative.dev/creator: docker-for-desktop
              serving.knative.dev/lastModifier: docker-for-desktop
API Version:  serving.knative.dev/v1
Kind:         Service
Metadata:
  Creation Timestamp:  2022-03-26T09:58:04Z
  Generation:          3
  Managed Fields:
    API Version:  serving.knative.dev/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:status:
        .:
        f:conditions:
        f:latestCreatedRevisionName:
        f:observedGeneration:
        f:url:
    Manager:      controller
    Operation:    Update
    Subresource:  status
    Time:         2022-03-26T09:58:04Z
    API Version:  serving.knative.dev/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:metadata:
        f:annotations:
          .:
          f:kubectl.kubernetes.io/last-applied-configuration:
      f:spec:
        .:
        f:template:
          .:
          f:metadata:
            .:
            f:labels:
              .:
              f:app:
          f:spec:
            .:
            f:containers:
    Manager:         kubectl-client-side-apply
    Operation:       Update
    Time:            2022-03-26T09:58:04Z
  Resource Version:  2261685
  UID:               b96afd85-667a-42f3-95d6-f279ba418235
Spec:
  Template:
    Metadata:
      Creation Timestamp:  <nil>
      Labels:
        App:  hello
    Spec:
      Container Concurrency:  0
      Containers:
        Env:
          Name:             TARGET
          Value:            World!
        Image:              xiazemin/helloworld-go
        Image Pull Policy:  IfNotPresent
        Name:               user-container
        Readiness Probe:
          Success Threshold:  1
          Tcp Socket:
            Port:  0
        Resources:
      Enable Service Links:  false
      Timeout Seconds:       300
  Traffic:
    Latest Revision:  true
    Percent:          100
Status:
  Conditions:
    Last Transition Time:        2022-03-26T10:22:14Z
    Message:                     Revision "hello-00003" failed with message: Unable to fetch image "xiazemin/helloworld-go": failed to resolve image to digest: HEAD https://index.docker.io/v2/xiazemin/helloworld-go/manifests/latest: unexpected status code 401 Unauthorized (HEAD responses have no body, use GET for details).
    Reason:                      RevisionFailed
    Status:                      False
    Type:                        ConfigurationsReady
    Last Transition Time:        2022-03-26T10:22:14Z
    Message:                     Configuration "hello" does not have any ready Revision.
    Reason:                      RevisionMissing
    Status:                      False
    Type:                        Ready
    Last Transition Time:        2022-03-26T10:22:14Z
    Message:                     Configuration "hello" does not have any ready Revision.
    Reason:                      RevisionMissing
    Status:                      False
    Type:                        RoutesReady
  Latest Created Revision Name:  hello-00003
  Observed Generation:           3
  URL:                           http://hello.helloworld.example.com
Events:
  Type    Reason   Age   From                Message
  ----    ------   ----  ----                -------
  Normal  Created  26m   service-controller  Created Configuration "hello"
  Normal  Created  26m   service-controller  Created Route "hello"



https://github.com/knative/serving/issues/6705

https://www.cnblogs.com/shaohef/p/10688770.html
https://haralduebele.github.io/2020/06/03/serverless-and-knative-part-2-knative-serving/
https://www.cxybb.com/article/weixin_30814319/98040299


% docker push xiazemin/helloworld-go
Using default tag: latest
The push refers to repository [docker.io/xiazemin/helloworld-go]
bf1551a2f01a: Preparing 
b9eadfdd4be0: Preparing 
ff768a1413ba: Preparing 
denied: requested access to the resource is denied


 % docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: xiazemin
Password: 
Login Succeeded

Logging in with your password grants your terminal complete access to your account. 
For better security, log in with a limited-privilege personal access token. Learn more at https://docs.docker.com/go/access-tokens/

 % docker push xiazemin/helloworld-go
Using default tag: latest
The push refers to repository [docker.io/xiazemin/helloworld-go]
bf1551a2f01a: Pushed 
b9eadfdd4be0: Pushed 
ff768a1413ba: Mounted from library/alpine 
latest: digest: sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405 size: 949

% docker pull xiazemin/helloworld-go:latest
latest: Pulling from xiazemin/helloworld-go
Digest: sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405
Status: Image is up to date for xiazemin/helloworld-go:latest
docker.io/xiazemin/helloworld-go:latest


 % kubectl -n helloworld delete revision.serving.knative.dev/hello-00001
revision.serving.knative.dev "hello-00001" deleted


 % kubectl -n helloworld logs revision.serving.knative.dev/hello-00005 
error: no kind "Revision" is registered for version "serving.knative.dev/v1" in scheme "k8s.io/kubectl/pkg/scheme/scheme.go:28"


#https://github.com/knative/serving/issues/6085


%  kubectl -n helloworld get pods                                            
NAME                                      READY   STATUS             RESTARTS   AGE
hello-00005-deployment-77c7d84b98-gjwlr   1/2     ImagePullBackOff   0          11m
hello-00006-deployment-74f889dbb7-72kcm   1/2     ImagePullBackOff   0          2m51s



%  kubectl -n helloworld describe pod hello-00005-deployment-77c7d84b98-gjwlr
Name:         hello-00005-deployment-77c7d84b98-gjwlr
Namespace:    helloworld
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Sat, 26 Mar 2022 18:46:22 +0800
Labels:       app=hello
              pod-template-hash=77c7d84b98
              serving.knative.dev/configuration=hello
              serving.knative.dev/configurationGeneration=5
              serving.knative.dev/configurationUID=ee31bffa-a1ac-4001-83b6-4ead5457d4e5
              serving.knative.dev/revision=hello-00005
              serving.knative.dev/revisionUID=fd0d447e-94cd-4da6-b02a-50b4d1fd8032
              serving.knative.dev/service=hello
              serving.knative.dev/serviceUID=b96afd85-667a-42f3-95d6-f279ba418235
Annotations:  serving.knative.dev/creator: docker-for-desktop
Status:       Pending
IP:           10.1.4.234
IPs:
  IP:           10.1.4.234
Controlled By:  ReplicaSet/hello-00005-deployment-77c7d84b98
Containers:
  user-container:
    Container ID:   docker://bb422eb2d6bf0d6bdec082597814fc86391855f678e684d2b473b05bd5f68888
    Image:          index.docker.io/xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405
    Image ID:       docker-pullable://xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sat, 26 Mar 2022 18:46:24 +0800
    Ready:          True
    Restart Count:  0
    Environment:
      TARGET:           World!
      PORT:             8080
      K_REVISION:       hello-00005
      K_CONFIGURATION:  hello
      K_SERVICE:        hello
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-h77nr (ro)
  queue-proxy:
    Container ID:   
    Image:          gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest
    Image ID:       
    Ports:          8022/TCP, 9090/TCP, 9091/TCP, 8012/TCP
    Host Ports:     0/TCP, 0/TCP, 0/TCP, 0/TCP
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:      25m
    Readiness:  http-get http://:8012/ delay=0s timeout=1s period=10s #success=1 #failure=3
    Environment:
      SERVING_NAMESPACE:                 helloworld
      SERVING_SERVICE:                   hello
      SERVING_CONFIGURATION:             hello
      SERVING_REVISION:                  hello-00005
      QUEUE_SERVING_PORT:                8012
      CONTAINER_CONCURRENCY:             0
      REVISION_TIMEOUT_SECONDS:          300
      SERVING_POD:                       hello-00005-deployment-77c7d84b98-gjwlr (v1:metadata.name)
      SERVING_POD_IP:                     (v1:status.podIP)
      SERVING_LOGGING_CONFIG:            
      SERVING_LOGGING_LEVEL:             
      SERVING_REQUEST_LOG_TEMPLATE:      {"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl": "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}", "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent": "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}", "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}", "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"}, "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}
      SERVING_ENABLE_REQUEST_LOG:        false
      SERVING_REQUEST_METRICS_BACKEND:   prometheus
      TRACING_CONFIG_BACKEND:            none
      TRACING_CONFIG_ZIPKIN_ENDPOINT:    
      TRACING_CONFIG_DEBUG:              false
      TRACING_CONFIG_SAMPLE_RATE:        0.1
      USER_PORT:                         8080
      SYSTEM_NAMESPACE:                  knative-serving
      METRICS_DOMAIN:                    knative.dev/internal/serving
      SERVING_READINESS_PROBE:           {"tcpSocket":{"port":8080,"host":"127.0.0.1"},"successThreshold":1}
      ENABLE_PROFILING:                  false
      SERVING_ENABLE_PROBE_REQUEST_LOG:  false
      METRICS_COLLECTOR_ADDRESS:         
      CONCURRENCY_STATE_ENDPOINT:        
      CONCURRENCY_STATE_TOKEN_PATH:      /var/run/secrets/tokens/state-token
      HOST_IP:                            (v1:status.hostIP)
      ENABLE_HTTP2_AUTO_DETECTION:       false
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-h77nr (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  kube-api-access-h77nr:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age                  From               Message
  ----     ------     ----                 ----               -------
  Normal   Scheduled  11m                  default-scheduler  Successfully assigned helloworld/hello-00005-deployment-77c7d84b98-gjwlr to docker-desktop
  Normal   Pulled     11m                  kubelet            Container image "index.docker.io/xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405" already present on machine
  Normal   Created    11m                  kubelet            Created container user-container
  Normal   Started    11m                  kubelet            Started container user-container
  Warning  Failed     9m59s (x2 over 11m)  kubelet            Failed to pull image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest": rpc error: code = Unknown desc = Error response from daemon: Get "https://gcr.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
  Warning  Failed     9m22s (x5 over 11m)  kubelet            Error: ImagePullBackOff
  Normal   Pulling    9m7s (x4 over 11m)   kubelet            Pulling image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest"
  Warning  Failed     8m52s (x4 over 11m)  kubelet            Error: ErrImagePull
  Warning  Failed     8m52s (x2 over 10m)  kubelet            Failed to pull image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest": rpc error: code = Unknown desc = Error response from daemon: Get "https://gcr.io/v2/": context deadline exceeded
  Normal   BackOff    85s (x35 over 11m)   kubelet            Back-off pulling image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest"


serverless/knative/setup/serving-core.yaml

queueSidecarImage: gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest
imagePullPolicy: IfNotPresent  


 % kubectl apply -f ../setup/serving-core.yaml

  %  kubectl -n helloworld get pods
NAME                                      READY   STATUS             RESTARTS   AGE
hello-00001-deployment-566bb9bfc9-k7stf   1/2     ImagePullBackOff   0          2m48s



Error from server (NotFound): error when deleting "../setup/serving-core.yaml": the server could not find the requested resource (delete images.caching.internal.knative.dev queue-proxy)


curl -k -H "Content-Type: application/json" -X PUT --data-binary @a.json http://127.0.0.1:8001/api/v1/namespaces/knative-serving/finalize


% kubectl -n helloworld get deploy                                          
NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
hello-00001-deployment   0/1     1            0           71s



% kubectl -n helloworld edit deploy hello-00001-deployment 
deployment.apps/hello-00001-deployment edited

image: gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest
imagePullPolicy: Always

IfNotPresent


 % kubectl -n helloworld get deploy
NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
hello-00001-deployment   1/1     1            1           3m19s


 % kubectl -n helloworld get pods  
NAME                                      READY   STATUS             RESTARTS   AGE
hello-00001-deployment-6697b46f85-dqnnw   1/2     ImagePullBackOff   0          3m28s
hello-00001-deployment-797ddcfd59-8t46p   2/2     Running            0          42s


% kubectl -n helloworld  delete pod hello-00001-deployment-6697b46f85-dqnnw 
pod "hello-00001-deployment-6697b46f85-dqnnw" deleted


 % kubectl -n helloworld get pods   
No resources found in helloworld namespace.


 % kubectl -n helloworld get pods                                           
NAME                                     READY   STATUS         RESTARTS   AGE
hello-00001-deployment-b859854bc-mt5n6   1/2     ErrImagePull   0          64s
hello-00001-deployment-c4ffc5489-c9884   2/2     Running        0          27s




% kubectl -n helloworld describe pod hello-00001-deployment-b859854bc-mt5n6
Name:         hello-00001-deployment-b859854bc-mt5n6
Namespace:    helloworld
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Sat, 26 Mar 2022 19:23:14 +0800
Labels:       app=hello
              pod-template-hash=b859854bc
              serving.knative.dev/configuration=hello
              serving.knative.dev/configurationGeneration=1
              serving.knative.dev/configurationUID=8d8b30f7-21f1-4435-8610-45276cf31c4c
              serving.knative.dev/revision=hello-00001
              serving.knative.dev/revisionUID=48290df7-fa7c-4903-a809-218ef82a19d0
              serving.knative.dev/service=hello
              serving.knative.dev/serviceUID=7b35c9f5-eb1a-46a8-9391-c40aa76970d0
Annotations:  autoscaling.knative.dev/target: 10
              serving.knative.dev/creator: docker-for-desktop
Status:       Pending
IP:           10.1.4.248
IPs:
  IP:           10.1.4.248
Controlled By:  ReplicaSet/hello-00001-deployment-b859854bc
Containers:
  user-container:
    Container ID:   docker://0d6053e3b01ac6013341f21810caad9180d419ad26d7522b50e9a2580a6c48ba
    Image:          index.docker.io/xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405
    Image ID:       docker-pullable://xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sat, 26 Mar 2022 19:23:18 +0800
    Ready:          True
    Restart Count:  0
    Environment:
      TARGET:           World!
      PORT:             8080
      K_REVISION:       hello-00001
      K_CONFIGURATION:  hello
      K_SERVICE:        hello
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-ftsf2 (ro)
  queue-proxy:
    Container ID:   
    Image:          gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest
    Image ID:       
    Ports:          8022/TCP, 9090/TCP, 9091/TCP, 8012/TCP
    Host Ports:     0/TCP, 0/TCP, 0/TCP, 0/TCP
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:      25m
    Readiness:  http-get http://:8012/ delay=0s timeout=1s period=10s #success=1 #failure=3
    Environment:
      SERVING_NAMESPACE:                 helloworld
      SERVING_SERVICE:                   hello
      SERVING_CONFIGURATION:             hello
      SERVING_REVISION:                  hello-00001
      QUEUE_SERVING_PORT:                8012
      CONTAINER_CONCURRENCY:             0
      REVISION_TIMEOUT_SECONDS:          300
      SERVING_POD:                       hello-00001-deployment-b859854bc-mt5n6 (v1:metadata.name)
      SERVING_POD_IP:                     (v1:status.podIP)
      SERVING_LOGGING_CONFIG:            
      SERVING_LOGGING_LEVEL:             
      SERVING_REQUEST_LOG_TEMPLATE:      {"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl": "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}", "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent": "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}", "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}", "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"}, "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}
      SERVING_ENABLE_REQUEST_LOG:        false
      SERVING_REQUEST_METRICS_BACKEND:   prometheus
      TRACING_CONFIG_BACKEND:            none
      TRACING_CONFIG_ZIPKIN_ENDPOINT:    
      TRACING_CONFIG_DEBUG:              false
      TRACING_CONFIG_SAMPLE_RATE:        0.1
      USER_PORT:                         8080
      SYSTEM_NAMESPACE:                  knative-serving
      METRICS_DOMAIN:                    knative.dev/internal/serving
      SERVING_READINESS_PROBE:           {"tcpSocket":{"port":8080,"host":"127.0.0.1"},"successThreshold":1}
      ENABLE_PROFILING:                  false
      SERVING_ENABLE_PROBE_REQUEST_LOG:  false
      METRICS_COLLECTOR_ADDRESS:         
      CONCURRENCY_STATE_ENDPOINT:        
      CONCURRENCY_STATE_TOKEN_PATH:      /var/run/secrets/tokens/state-token
      HOST_IP:                            (v1:status.hostIP)
      ENABLE_HTTP2_AUTO_DETECTION:       false
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-ftsf2 (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  kube-api-access-ftsf2:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age                From               Message
  ----     ------     ----               ----               -------
  Normal   Scheduled  77s                default-scheduler  Successfully assigned helloworld/hello-00001-deployment-b859854bc-mt5n6 to docker-desktop
  Normal   Pulled     75s                kubelet            Container image "index.docker.io/xiazemin/helloworld-go@sha256:23b742c725fa51786bb71603a28c34e2723f9b65384a4886349145df532c1405" already present on machine
  Normal   Created    75s                kubelet            Created container user-container
  Normal   Started    74s                kubelet            Started container user-container
  Warning  Failed     29s (x2 over 59s)  kubelet            Failed to pull image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest": rpc error: code = Unknown desc = Error response from daemon: Get "https://gcr.io/v2/": context deadline exceeded
  Warning  Failed     29s (x2 over 59s)  kubelet            Error: ErrImagePull
  Normal   BackOff    18s (x2 over 58s)  kubelet            Back-off pulling image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest"
  Warning  Failed     18s (x2 over 58s)  kubelet            Error: ImagePullBackOff
  Normal   Pulling    6s (x3 over 74s)   kubelet            Pulling image "gcr.io/knative-releases/knative.dev/serving/cmd/queue:latest"


% kubectl -n helloworld get pods
No resources found in helloworld namespace.


 % kubectl -n helloworld get all
NAME                          TYPE           CLUSTER-IP       EXTERNAL-IP                    PORT(S)                                      AGE
service/hello                 ExternalName   <none>           hello.helloworld.example.com   80/TCP                                       2m36s
service/hello-00001           ClusterIP      10.105.213.205   <none>                         80/TCP                                       3m18s
service/hello-00001-private   ClusterIP      10.104.37.28     <none>                         80/TCP,9090/TCP,9091/TCP,8022/TCP,8012/TCP   3m19s

NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/hello-00001-deployment   0/0     0            0           3m19s

NAME                                               DESIRED   CURRENT   READY   AGE
replicaset.apps/hello-00001-deployment-b859854bc   0         0         0       3m19s
replicaset.apps/hello-00001-deployment-c4ffc5489   0         0         0       2m42s

NAME                                      LATESTCREATED   LATESTREADY   READY   REASON
configuration.serving.knative.dev/hello   hello-00001     hello-00001   True    

NAME                                URL                                   LATESTCREATED   LATESTREADY   READY     REASON
service.serving.knative.dev/hello   http://hello.helloworld.example.com   hello-00001     hello-00001   Unknown   IngressNotConfigured

NAME                              URL                                   READY     REASON
route.serving.knative.dev/hello   http://hello.helloworld.example.com   Unknown   IngressNotConfigured

NAME                                       CONFIG NAME   K8S SERVICE NAME   GENERATION   READY   REASON   ACTUAL REPLICAS   DESIRED REPLICAS
revision.serving.knative.dev/hello-00001   hello                            1            True             0                 0



 % kubectl get ksvc -n helloworld
NAME    URL                                   LATESTCREATED   LATESTREADY   READY     REASON
hello   http://hello.helloworld.example.com   hello-00001     hello-00001   Unknown   IngressNotConfigured

https://github.com/knative/serving/issues/6346
https://mlog.club/article/3321300

 % kubectl -n helloworld describe ksvc helloworld-go
Error from server (NotFound): services.serving.knative.dev "helloworld-go" not found





 % kubectl -n helloworld describe ksvc hello        
Name:         hello
Namespace:    helloworld
Labels:       <none>
Annotations:  serving.knative.dev/creator: docker-for-desktop
              serving.knative.dev/lastModifier: docker-for-desktop
API Version:  serving.knative.dev/v1
Kind:         Service
Metadata:
  Creation Timestamp:  2022-03-26T11:23:11Z
  Generation:          1
  Managed Fields:
    API Version:  serving.knative.dev/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:metadata:
        f:annotations:
          .:
          f:kubectl.kubernetes.io/last-applied-configuration:
      f:spec:
        .:
        f:template:
          .:
          f:metadata:
            .:
            f:annotations:
              .:
              f:autoscaling.knative.dev/target:
            f:labels:
              .:
              f:app:
          f:spec:
            .:
            f:containers:
    Manager:      kubectl-client-side-apply
    Operation:    Update
    Time:         2022-03-26T11:23:11Z
    API Version:  serving.knative.dev/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:status:
        .:
        f:address:
          .:
          f:url:
        f:conditions:
        f:latestCreatedRevisionName:
        f:latestReadyRevisionName:
        f:observedGeneration:
        f:traffic:
        f:url:
    Manager:         controller
    Operation:       Update
    Subresource:     status
    Time:            2022-03-26T11:23:57Z
  Resource Version:  2276482
  UID:               7b35c9f5-eb1a-46a8-9391-c40aa76970d0
Spec:
  Template:
    Metadata:
      Annotations:
        autoscaling.knative.dev/target:  10
      Creation Timestamp:                <nil>
      Labels:
        App:  hello
    Spec:
      Container Concurrency:  0
      Containers:
        Env:
          Name:   TARGET
          Value:  World!
        Image:    docker.io/xiazemin/helloworld-go
        Name:     user-container
        Readiness Probe:
          Success Threshold:  1
          Tcp Socket:
            Port:  0
        Resources:
      Enable Service Links:  false
      Timeout Seconds:       300
  Traffic:
    Latest Revision:  true
    Percent:          100
Status:
  Address:
    URL:  http://hello.helloworld.svc.cluster.local
  Conditions:
    Last Transition Time:        2022-03-26T11:23:57Z
    Status:                      True
    Type:                        ConfigurationsReady
    Last Transition Time:        2022-03-26T11:23:57Z
    Message:                     Ingress has not yet been reconciled.
    Reason:                      IngressNotConfigured
    Status:                      Unknown
    Type:                        Ready
    Last Transition Time:        2022-03-26T11:23:57Z
    Message:                     Ingress has not yet been reconciled.
    Reason:                      IngressNotConfigured
    Status:                      Unknown
    Type:                        RoutesReady
  Latest Created Revision Name:  hello-00001
  Latest Ready Revision Name:    hello-00001
  Observed Generation:           1
  Traffic:
    Latest Revision:  true
    Percent:          100
    Revision Name:    hello-00001
  URL:                http://hello.helloworld.example.com
Events:
  Type    Reason   Age   From                Message
  ----    ------   ----  ----                -------
  Normal  Created  10m   service-controller  Created Configuration "hello"
  Normal  Created  10m   service-controller  Created Route "hello"


Adding the label serving.knative.dev/visibility=cluster-local makes the problem go away, but then it is only accessible internally without SSL.


% kubectl -n helloworld label kservice hello serving.knative.dev/visibility=cluster-local
service.serving.knative.dev/hello labeled

% kubectl -n helloworld label service hello serving.knative.dev/visibility=cluster-local                          
service/hello labeled

 % kubectl -n helloworld label route hello serving.knative.dev/visibility=cluster-local
error: 'serving.knative.dev/visibility' already has a value (cluster-local), and --overwrite is false

https://knative-v1.netlify.app/docs/serving/cluster-local-route/


https://knative.dev/docs/serving/services/private-services/#example



 % kubectl get svc -n kourier-system 
NAME               TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier            LoadBalancer   10.108.41.244   localhost     80:31859/TCP,443:30753/TCP   143m
kourier-internal   ClusterIP      10.96.62.112    <none>        80/TCP                       143m


 % curl -H "Host: hello.helloworld.example.com" 127.0.0.1:31859
curl: (7) Failed to connect to 127.0.0.1 port 31859: Connection refused

% curl -H "Host: hello.helloworld.example.com" 10.96.62.112:80


% kubectl --namespace kourier-system get service kourier

NAME      TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier   LoadBalancer   10.108.41.244   localhost     80:31859/TCP,443:30753/TCP   3h10m


 % kubectl --namespace kourier-system edit service kourier
service/kourier edited

  type: LoadBalancer

    type: NodePort


% kubectl --namespace kourier-system get service         
NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier            NodePort    10.108.41.244   <none>        80:31859/TCP,443:30753/TCP   3h13m
kourier-internal   ClusterIP   10.96.62.112    <none>        80/TCP                       3h13m


 % kubectl get configurations.serving.knative.dev -n helloworld
NAME    LATESTCREATED   LATESTREADY   READY   REASON
hello   hello-00001     hello-00001   True    


 % kubectl get routes.serving.knative.dev  -n helloworld
Error from server (NotFound): routes.serving.knative.dev "\u00a0-n" not found
Error from server (NotFound): routes.serving.knative.dev "helloworld" not found

 % kubectl get revisions.serving.knative.dev -n helloworld
NAME          CONFIG NAME   K8S SERVICE NAME   GENERATION   READY   REASON   ACTUAL REPLICAS   DESIRED REPLICAS
hello-00001   hello                            1            True             0                 0

#https://blog.51cto.com/u_14625168/2527615



 % kubectl apply -f hello-world.yaml 
service.serving.knative.dev/helloworld-go created



 %   kn service create helloworld-go --image=docker.io/xiazemin/helloworld-go --env TARGET="Go Sample v1"
Creating service 'helloworld-go' in namespace 'default':

  0.134s The Route is still working to reflect the latest desired specification.
  0.250s ...
  0.321s Configuration "helloworld-go" is waiting for a Revision to become ready.


https://github.com/knative/docs/tree/main/code-samples/serving/hello-world/helloworld-go


bash kn service delete helloworld-go


 % kubectl -n knative-serving logs -f pod/activator-7cf4bd8548-scdtk 
{"severity":"WARNING","timestamp":"2022-03-26T13:01:30.553694004Z","logger":"activator","caller":"net/revision_backends.go:340","message":"Failed probing pods","commit":"328000e","knative.dev/controller":"activator","knative.dev/pod":"activator-7cf4bd8548-scdtk","knative.dev/key":"default/helloworld-go-00001","curDests":{"ready":"","notReady":"10.1.4.254:8012"},"error":"error roundtripping http://10.1.4.254:8012/healthz: dial tcp 10.1.4.254:8012: connect: connection refused"}
{"severity":"ERROR","timestamp":"2022-03-26T13:01:30.855320592Z","logger":"activator","caller":"net/revision_backends.go:396","message":"Failed to probe clusterIP 10.99.229.201:80","commit":"328000e","knative.dev/controller":"activator","knative.dev/pod":"activator-7cf4bd8548-scdtk","knative.dev/key":"default/helloworld-go-00001","error":"error roundtripping http://10.99.229.201:80/healthz: context deadline exceeded","stacktrace":"knative.dev/serving/pkg/activator/net.(*revisionWatcher).checkDests\n\tknative.dev/serving/pkg/activator/net/revision_backends.go:396\nknative.dev/serving/pkg/activator/net.(*revisionWatcher).run\n\tknative.dev/serving/pkg/activator/net/revision_backends.go:441"}


https://access.redhat.com/solutions/6492681