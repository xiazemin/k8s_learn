 top
 PID    COMMAND      %CPU  TIME     #TH   #WQ  #PORT MEM    PURG   CMPRS  PGRP
37210  com.docker.h 155.4 31:23.69 13/4  0    36    11G+   0B     7922M- 36866
发现docker占用内存11G

然后具体看下哪个pod占用比较多
  % kubectl top pod
error: Metrics API not available
https://blog.csdn.net/TinyJian/article/details/109711164
https://www.cnblogs.com/lfl17718347843/p/14283796.html
默认没有安装metric-server安装下
https://github.com/kubernetes-sigs/metrics-server/releases
https://www.jianshu.com/p/bfe780ce14ce

 
 % kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.6.0/components.yaml
serviceaccount/metrics-server created
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader created
clusterrole.rbac.authorization.k8s.io/system:metrics-server created
rolebinding.rbac.authorization.k8s.io/metrics-server-auth-reader created
clusterrolebinding.rbac.authorization.k8s.io/metrics-server:system:auth-delegator created
clusterrolebinding.rbac.authorization.k8s.io/system:metrics-server created
service/metrics-server created
deployment.apps/metrics-server created
apiservice.apiregistration.k8s.io/v1beta1.metrics.k8s.io created


https://github.com/kubernetes-sigs/metrics-server/releases

% kubectl -n kube-system get pod metrics-server-65b979d49c-kdtg2 
NAME                              READY   STATUS             RESTARTS   AGE
metrics-server-65b979d49c-kdtg2   0/1     ImagePullBackOff   0          7m34s
镜像拉不下来

 % kubectl top pod
Error from server (ServiceUnavailable): the server is currently unable to handle the request (get pods.metrics.k8s.io)


image: k8s.gcr.io/metrics-server/metrics-server:v0.6.0


 % docker pull rancher/metrics-server:v0.4.1
v0.4.1: Pulling from rancher/metrics-server
e59bd8947ac7: Already exists 
cdbcff7dade2: Pull complete 
Digest: sha256:b99989f8b6a18a838737a155e0b7fd0fa237e239034a6bc9b6330879ad001aa1
Status: Downloaded newer image for rancher/metrics-server:v0.4.1
docker.io/rancher/metrics-server:v0.4.1

 % docker tag docker.io/rancher/metrics-server:v0.4.1 k8s.gcr.io/metrics-server/metrics-server:v0.6.0



 % docker rmi  docker.io/rancher/metrics-server:v0.4.1
Untagged: rancher/metrics-server:v0.4.1
Untagged: rancher/metrics-server@sha256:b99989f8b6a18a838737a155e0b7fd0fa237e239034a6bc9b6330879ad001aa1


% kubectl -n kube-system describe pod metrics-server-65b979d49c-d9wzg 
 Warning  Unhealthy  8s (x2 over 18s)  kubelet            Liveness probe failed: HTTP probe failed with statuscode: 500




% docker pull ricardbejarano/metrics-server:0.6.0
0.6.0: Pulling from ricardbejarano/metrics-server
9cdca96b0b76: Pull complete 
Digest: sha256:b43acfcb4babd657607dd7170bc8d1f608b068c97cab37a299b8e20553209fad
Status: Downloaded newer image for ricardbejarano/metrics-server:0.6.0
docker.io/ricardbejarano/metrics-server:0.6.0

 % docker tag docker.io/ricardbejarano/metrics-server:0.6.0  k8s.gcr.io/metrics-server/metrics-server:v0.6.0


% kubectl -n kube-system logs metrics-server-65b979d49c-5hrnz 
 panic: unable to get openapi models: [3,11] $root.info is missing required property: version

goroutine 1 [running]:
main.main()
        /go/src/sigs.k8s.io/metrics-server/cmd/metrics-server/metrics-server.go:39 +0xb9


https://blog.csdn.net/shenhonglei1234/article/details/111171525


解决方案： 
修改 YAML 文件，Pod 下的 args 增加参数 --kubelet-insecure-tls 以禁用证书校验 

% kubectl apply -f 资源占用控制/metrics-server.yaml 
serviceaccount/metrics-server unchanged
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader unchanged
clusterrole.rbac.authorization.k8s.io/system:metrics-server unchanged
rolebinding.rbac.authorization.k8s.io/metrics-server-auth-reader unchanged
clusterrolebinding.rbac.authorization.k8s.io/metrics-server:system:auth-delegator unchanged
clusterrolebinding.rbac.authorization.k8s.io/system:metrics-server unchanged
service/metrics-server unchanged
deployment.apps/metrics-server configured
apiservice.apiregistration.k8s.io/v1beta1.metrics.k8s.io unchanged


 % docker pull ricardbejarano/metrics-server:latest
latest: Pulling from ricardbejarano/metrics-server
a1b4a8c58490: Pull complete 
Digest: sha256:4b843846ce8f9a57805183ab5fefe5dab3e29671ee27a2576cbd38f0fae53ebb
Status: Downloaded newer image for ricardbejarano/metrics-server:latest
docker.io/ricardbejarano/metrics-server:latest

 % docker tag docker.io/ricardbejarano/metrics-server:latest  k8s.gcr.io/metrics-server/metrics-server:v0.6.0

https://github.com/kubernetes-sigs/metrics-server


% docker pull bitnami/metrics-server:latest
latest: Pulling from bitnami/metrics-server
edb9df82989e: Pull complete 
fd4ce8181ca9: Pull complete 
dd7359543396: Pull complete 
b515c7504db3: Pull complete 
a4af5d5f09ef: Pull complete 
48f8bc37bc18: Pull complete 
37af7a161a3c: Pull complete 
Digest: sha256:6f8bdb06de083988b13b6c08dee3e40e061194e07e6b780fa3d4de6c487fbd25
Status: Downloaded newer image for bitnami/metrics-server:latest
docker.io/bitnami/metrics-server:latest


 % kubectl -n kube-system logs metrics-server-57f695cd54-zf8fz 
 .go:188] "Failed probe" probe="metric-storage-ready" err="not metrics to serve"
E0123 03:38:08.027869       1 scraper.go:139] "Failed to scrape node" err="GET \"https://192.168.65.4:10250/stats/summary?only_cpu_and_memory=true\": bad status code \"403 Forbidden\"" node="docker-desktop"
I0123 03:38:13.498399       1 server.go:188] "Failed probe" probe="metric-storage-ready" err="not metrics to serve"


        - --kubelet-insecure-tls=true
        - --kubelet-preferred-address-types=InternalIP


https://stackoverflow.com/questions/68648198/metrics-service-in-kubernetes-not-working
https://blog.csdn.net/liuyanwuyu/article/details/119793631


% kubectl apply -f system:metrics-server.yaml
clusterrole.rbac.authorization.k8s.io/system:metrics-server configured

https://github.com/kubernetes-sigs/metrics-server/issues/95

% kubectl -n kube-system get pods metrics-server-664ffd5b45-mn4r8 
NAME                              READY   STATUS    RESTARTS   AGE
metrics-server-664ffd5b45-mn4r8   1/1     Running   0          6m22s


% kubectl top pod -n istio-system
NAME                                   CPU(cores)   MEMORY(bytes)
grafana-6ccd56f4b6-wp798               6m           42Mi
istio-egressgateway-687f4db598-rj2mg   5m           45Mi
istio-ingressgateway-8d5985f54-cvlcw   6m           42Mi
istiod-5698b78654-lm7r6                4m           49Mi
istiod-5698b78654-lwg4p                6m           52Mi
istiod-5698b78654-r2vp8                7m           46Mi
jaeger-5d44bc5c5d-mt9cq                3m           25Mi
kiali-79b86ff5bc-srf9k                 1m           12Mi
prometheus-64fd8ccd65-rn9lc            24m          406Mi


% kubectl top pod -n bookinfo
NAME                              CPU(cores)   MEMORY(bytes)
details-v1-79f774bdb9-dvl89       5m           60Mi
productpage-v1-6b746f74dc-9qfzk   20m          103Mi
ratings-v1-b6994bb9-z6zmm         9m           51Mi
reviews-v1-545db77b95-4vlbq       12m          132Mi
reviews-v2-7bf8c9648f-q7dfw       11m          166Mi
reviews-v3-84779c7bbc-tpfbs       8m           142Mi


% for ns in `kubectl get ns |grep -v NAME |awk '{print $1}'`
do
kubectl top pod -n $ns
done

 % kubectl top node docker-desktop
NAME             CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
docker-desktop   902m         22%    4147Mi          52%













https://blog.csdn.net/weixin_38754564/article/details/103193594
https://www.servicemesher.com/blog/201911-envoy-memory-optimize/
https://www.cnblogs.com/haoyunlaile/p/12874199.html

