% kubectl -n kube-system  get pod |grep dns
coredns-558bd4d5db-qc6px                 1/1     Running   17         97d
coredns-558bd4d5db-wzcgb                 1/1     Running   17         97d



 % kubectl -n kube-system  get pod  coredns-558bd4d5db-qc6px -o yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: "2021-08-23T03:19:05Z"
  generateName: coredns-558bd4d5db-
  labels:
    k8s-app: kube-dns
    pod-template-hash: 558bd4d5db
  name: coredns-558bd4d5db-qc6px
  namespace: kube-system
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: coredns-558bd4d5db
    uid: 554cb42b-c19b-44cd-a9a4-21b3d1b72de8
  resourceVersion: "2732012"
  uid: b6f89815-be23-4fca-93be-389202fba784
spec:
  containers:
  - args:
    - -conf
    - /etc/coredns/Corefile
    image: k8s.gcr.io/coredns/coredns:v1.8.0
    imagePullPolicy: IfNotPresent
    livenessProbe:
      failureThreshold: 5
      httpGet:
        path: /health
        port: 8080
        scheme: HTTP
      initialDelaySeconds: 60
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 5
    name: coredns
    ports:
    - containerPort: 53
      name: dns
      protocol: UDP
    - containerPort: 53
      name: dns-tcp
      protocol: TCP
    - containerPort: 9153
      name: metrics
      protocol: TCP
    readinessProbe:
      failureThreshold: 3
      httpGet:
        path: /ready
        port: 8181
        scheme: HTTP
      periodSeconds: 10
      successThreshold: 1
      timeoutSeconds: 1
    resources:
      limits:
        memory: 170Mi
      requests:
        cpu: 100m
        memory: 70Mi
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        add:
        - NET_BIND_SERVICE
        drop:
        - all
      readOnlyRootFilesystem: true
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /etc/coredns
      name: config-volume
      readOnly: true
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: kube-api-access-575ps
      readOnly: true
  dnsPolicy: Default
  enableServiceLinks: true
  nodeName: docker-desktop
  nodeSelector:
    kubernetes.io/os: linux
  preemptionPolicy: PreemptLowerPriority
  priority: 2000000000
  priorityClassName: system-cluster-critical
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: coredns
  serviceAccountName: coredns
  terminationGracePeriodSeconds: 30
  tolerations:
  - key: CriticalAddonsOnly
    operator: Exists
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
  - effect: NoSchedule
    key: node-role.kubernetes.io/control-plane
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - configMap:
      defaultMode: 420
      items:
      - key: Corefile
        path: Corefile
      name: coredns
    name: config-volume
  - name: kube-api-access-575ps
    projected:
      defaultMode: 420
      sources:
      - serviceAccountToken:
          expirationSeconds: 3607
          path: token
      - configMap:
          items:
          - key: ca.crt
            path: ca.crt
          name: kube-root-ca.crt
      - downwardAPI:
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
            path: namespace
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2021-08-23T03:19:33Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2021-11-22T01:57:16Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2021-11-22T01:57:16Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2021-08-23T03:19:33Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://d36d22a9c52caab9b4cd7e12433bb9cc3c0a843cd604ed7ccc916c1b83744f0a
    image: k8s.gcr.io/coredns/coredns:v1.8.0
    imageID: docker://sha256:1a1f05a2cd7c2fbfa7b45b21128c8a4880c003ca482460081dc12d76bfa863e8
    lastState:
      terminated:
        containerID: docker://a40315f1b7f43a8806eadc5a19cacede7720eb3a1b967eeb8e1268cd00eea2ec
        exitCode: 255
        finishedAt: "2021-11-22T01:56:09Z"
        reason: Error
        startedAt: "2021-11-04T03:48:39Z"
    name: coredns
    ready: true
    restartCount: 17
    started: true
    state:
      running:
        startedAt: "2021-11-22T01:56:36Z"
  hostIP: 192.168.65.4
  phase: Running
  podIP: 10.1.0.218
  podIPs:
  - ip: 10.1.0.218
  qosClass: Burstable
  startTime: "2021-08-23T03:19:33Z"




kube-proxy: 负责为Service提供cluster内部的服务发现和负载均衡；
controller manager: 负责维护集群的状态，比如故障检测、自动扩展、滚动更新等；
apiserver: 提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制；
scheduler: 负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上；
coredns: 基于DNS的服务发现
etcd: 保存了整个集群的状态；
pause: 在pod中担任Linux命名空间共享的基础


https://hub.docker.com/r/leipengkai/k8s-dns-kube-dns-amd64

https://github.com/docker/for-mac/issues/2646




Kubernetes master is running at https://localhost:6443
KubeDNS is running at https://localhost:6443/api/v1/namespaces/kube-system/services/kube-dns/proxy
Copy

https://logz.io/blog/kubernetes-docker-mac/

DNS 查询可以使用 Pod 中的 /etc/resolv.conf 展开。kubelet 会为每个 Pod 生成此文件。例如，对 data 的查询可能被展开为 data.test.cluster.local。 search 选项的取值会被用来展开查询。要进一步了解 DNS 查询

哪些对象会获得 DNS 记录呢？

Services
Pods


服务
A/AAAA 记录
“普通” 服务（除了无头服务）会以 my-svc.my-namespace.svc.cluster-domain.example 这种名字的形式被分配一个 DNS A 或 AAAA 记录，取决于服务的 IP 协议族。 该名称会解析成对应服务的集群 IP。

“无头（Headless）” 服务（没有集群 IP）也会以 my-svc.my-namespace.svc.cluster-domain.example 这种名字的形式被指派一个 DNS A 或 AAAA 记录， 具体取决于服务的 IP 协议族。 与普通服务不同，这一记录会被解析成对应服务所选择的 Pod 集合的 IP。 客户端要能够使用这组 IP，或者使用标准的轮转策略从这组 IP 中进行选择。

SRV 记录 
Kubernetes 会为命名端口创建 SRV 记录，这些端口是普通服务或 无头服务的一部分。 对每个命名端口，SRV 记录具有 _my-port-name._my-port-protocol.my-svc.my-namespace.svc.cluster-domain.example 这种形式。 对普通服务，该记录会被解析成端口号和域名：my-svc.my-namespace.svc.cluster-domain.example。 对无头服务，该记录会被解析成多个结果，服务对应的每个后端 Pod 各一个； 其中包含 Pod 端口号和形为 auto-generated-name.my-svc.my-namespace.svc.cluster-domain.example 的域名。


Pods
一般而言，Pod 会对应如下 DNS 名字解析：

pod-ip-address.my-namespace.pod.cluster-domain.example

例如，对于一个位于 default 名字空间，IP 地址为 172.17.0.3 的 Pod， 如果集群的域名为 cluster.local，则 Pod 会对应 DNS 名称：

172-17-0-3.default.pod.cluster.local.

Deployment 或通过 Service 暴露出来的 DaemonSet 所创建的 Pod 会有如下 DNS 解析名称可用：

pod-ip-address.deployment-name.my-namespace.svc.cluster-domain.example.


扩展 DNS 配置 
FEATURE STATE: Kubernetes 1.22 [alpha]
对于 Pod DNS 配置，Kubernetes 默认允许最多 6 个 搜索域（ Search Domain） 以及一个最多 256 个字符的搜索域列表。

如果启用 kube-apiserver 和 kubelet 的特性门控 ExpandedDNSConfig，Kubernetes 将可以有最多 32 个 搜索域以及一个最多 2048 个字符的搜索域列表。


https://kubernetes.io/zh/docs/concepts/services-networking/dns-pod-service/



从 Kubernetes v1.12 开始，CoreDNS 是推荐的 DNS 服务器，取代了 kube-dns。 如果 你的集群原来使用 kube-dns，你可能部署的仍然是 kube-dns 而不是 CoreDNS。

https://kubernetes.io/zh/docs/tasks/administer-cluster/dns-custom-nameservers/

如果 Pod 的 dnsPolicy 设置为 "default"，则它将从 Pod 运行所在节点继承名称解析配置。 Pod 的 DNS 解析行为应该与节点相同。 但请参阅已知问题。

如果你不想这样做，或者想要为 Pod 使用其他 DNS 配置，则可以 使用 kubelet 的 --resolv-conf 标志。 将此标志设置为 "" 可以避免 Pod 继承 DNS。 将其设置为有别于 /etc/resolv.conf 的有效文件路径可以设定 DNS 继承不同的配置。

CoreDNS 是模块化且可插拔的 DNS 服务器，每个插件都为 CoreDNS 添加了新功能。 可以通过维护 Corefile，即 CoreDNS 配置文件， 来定制其行为。 集群管理员可以修改 CoreDNS Corefile 的 ConfigMap，以更改服务发现的工作方式。

Corefile 配置包括以下 CoreDNS 插件：

errors：错误记录到标准输出。

health：在 http://localhost:8080/health 处提供 CoreDNS 的健康报告。

ready：在端口 8181 上提供的一个 HTTP 末端，当所有能够 表达自身就绪的插件都已就绪时，在此末端返回 200 OK。

kubernetes：CoreDNS 将基于 Kubernetes 的服务和 Pod 的 IP 答复 DNS 查询。你可以在 CoreDNS 网站阅读更多细节。 你可以使用 ttl 来定制响应的 TTL。默认值是 5 秒钟。TTL 的最小值可以是 0 秒钟， 最大值为 3600 秒。将 TTL 设置为 0 可以禁止对 DNS 记录进行缓存。

pods insecure 选项是为了与 kube-dns 向后兼容。你可以使用 pods verified 选项，该选项使得 仅在相同名称空间中存在具有匹配 IP 的 Pod 时才返回 A 记录。如果你不使用 Pod 记录，则可以使用 pods disabled 选项。

prometheus：CoreDNS 的度量指标值以 Prometheus 格式在 http://localhost:9153/metrics 上提供。
forward: 不在 Kubernetes 集群域内的任何查询都将转发到 预定义的解析器 (/etc/resolv.conf).
cache：启用前端缓存。
loop：检测到简单的转发环，如果发现死循环，则中止 CoreDNS 进程。
reload：允许自动重新加载已更改的 Corefile。 编辑 ConfigMap 配置后，请等待两分钟，以使更改生效。
loadbalance：这是一个轮转式 DNS 负载均衡器， 它在应答中随机分配 A、AAAA 和 MX 记录的顺序。





 % kubectl -n kube-system  get ConfigMap coredns -o yaml
apiVersion: v1
data:
  Corefile: |
    .:53 {
        errors
        health {
           lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
           max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
kind: ConfigMap
metadata:
  creationTimestamp: "2021-08-23T03:19:00Z"
  name: coredns
  namespace: kube-system
  resourceVersion: "278"
  uid: cdbb1028-afa9-4d1f-9eac-6375c3bbf1a9


  使用kube-dns发现服务
kube-dns可以解决Service的发现问题，k8s将Service的名称当做域名注册到kube-dns中，通过Service的名称就可以访问其提供的服务。

可能有人会问如果集群中没有部署kube-dns怎么办？没关系，实际上kube-dns插件只是运行在kube-system命名空间下的Pod，完全可以手动创建它。可以在k8s源码（v1.2）的cluster/addons/dns目录下找到两个模板（skydns-rc.yaml.in和skydns-svc.yaml.in）来创建

SkyDNS是用于服务发现的开源框架，构建于etcd之上。作用是为kubernetes集群中的Pod提供DNS查询接口。项目托管于https://github.com/skynetservices/skydns
etcd是一种开源的分布式key-value存储，其功能与ZooKeeper类似。在kube-dns中的作用为存储SkyDNS需要的各种数据，写入方为kube2sky，读取方为SkyDNS。项目托管于https://github.com/coreos/etcd。
kube2sky是k8s实现的一个适配程序，它通过名为kubernetes的Service（通过kubectl get svc可以查看到该Service，由集群自动创建）调用k8s的list和watch API来监听k8s Service资源的变更，从而修改etcd中的SkyDNS记录。代码可以在k8s源码（v1.2）的cluster/addons/dns/kube2sky/目录中找到。
exec-healthz是k8s提供的一种辅助容器，多用于side car模式中。它的原理是定期执行指定的Linux指令，从而判断当前Pod中关键容器的健康状态。在kube-dns中的作用就是通过nslookup指令检查DNS查询服务的健康状态，k8s livenessProbe通过访问exec-healthz提供的Http API了解健康状态，并在出现故障时重启容器。其源码位于https://github.com/kubernetes/contrib/tree/master/exec-healthz。
从图中可以发现，Pod查询DNS是通过ServiceName.Namespace子域名来查询的，但在之前的示例中只用了Service名称，什么原理呢？其实当我们只使用Service名称时会默认Namespace为default，而上面示例中的my-nginx Service就是在default Namespace中，因此是可以正常运行的。关于这一点，后续再深入介绍。
skydns-rc.yaml中可以发现livenessProbe是设置在kube2sky容器中的，其意图应该是希望通过重启kube2sky来重新写入DNS规则。

https://www.kubernetes.org.cn/273.html



kube-dns：DNS 服务的核心组件，主要由 KubeDNS 和 SkyDNS 组成KubeDNS 负责监听 Service 和 Endpoint 的变化情况，并将相关的信息更新到 SkyDNS 中SkyDNS 负责 DNS 解析，监听在 10053 端口 (tcp/udp)，同时也监听在 10055 端口提供 metricskube-dns 还监听了 8081 端口，以供健康检查使用
dnsmasq-nanny：负责启动 dnsmasq，并在配置发生变化时重启 dnsmasqdnsmasq 的 upstream 为 SkyDNS，即集群内部的 DNS 解析由 SkyDNS 负责
sidecar：负责健康检查和提供 DNS metrics（监听在 10054 端口）

https://kubernetes.feisky.xyz/concepts/components/kube-dns
https://cizixs.com/2017/04/11/kubernetes-intro-kube-dns/

Pod dnsPolicy
Kubernetes 集群中支持通过 dnsPolicy 字段为每个 Pod 配置不同的 DNS 策略。目前支持四种策略：

ClusterFirst：通过集群 DNS 服务来做域名解析，Pod 内 /etc/resolv.conf 配置的 DNS 服务地址是集群 DNS 服务的 kube-dns 地址。该策略是集群工作负载的默认策略。
None：忽略集群 DNS 策略，需要您提供 dnsConfig 字段来指定 DNS 配置信息。
Default：Pod 直接继承集群节点的域名解析配置。即在集群直接使用节点的 /etc/resolv.conf 文件。
ClusterFirstWithHostNetwork：强制在 hostNetWork 网络模式下使用 ClusterFirst 策略（默认使用 Default 策略）。


CoreDNS
CoreDNS 目前是 Kubernetes 标准的服务发现组件，dnsPolicy: ClusterFirst 模式的 Pod 会使用 CoreDNS 来解析集群内外部域名。

在命名空间 kube-system 下，集群有一个名为 coredns 的 configmap。其 Corefile 字段的文件配置内容如下（CoreDNS 功能都是通过 Corefile 内的插件提供）

https://developer.aliyun.com/article/779121

在 Kubernetes 中，服务发现有几种方式：
①：基于环境变量的方式
②：基于内部域名的方式

Kubernetes 中的域名是如何解析的
在 Kubernetes 中，比如服务 a 访问服务 b，对于同一个 Namespace下，可以直接在 pod 中，通过 curl b 来访问。对于跨 Namespace 的情况，服务名后边对应 Namespace即可。比如 curl b.default。那么，使用者这里边会有几个问题：

①：服务名是什么？
②：为什么同一个 Namespace 下，直接访问服务名即可？不同 Namespace 下，需要带上 Namespace 才行？
③：为什么内部的域名可以做解析，原理是什么？

DNS 如何解析，依赖容器内 resolv 文件的配置

cat /etc/resolv.conf

nameserver 10.233.0.3
search default.svc.cluster.local svc.cluster.local cluster.local
这个文件中，配置的 DNS Server，一般就是 K8S 中，kubedns 的 Service 的 ClusterIP，这个IP是虚拟IP，无法ping，但可以访问。

[root@node4 user1]# kubectl get svc -n kube-system
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
kube-dns               ClusterIP   10.233.0.3      <none>        53/UDP,53/TCP   270d
kubernetes-dashboard   ClusterIP   10.233.22.223   <none>        443/TCP         124d
所以，所有域名的解析，其实都要经过 kubedns 的虚拟IP 10.233.0.3 进行解析，不论是 Kubernetes 内部域名还是外部的域名。

Kubernetes 中，域名的全称，必须是 service-name.namespace.svc.cluster.local 这种模式，服务名，就是Kubernetes中 Service 的名称，所以，当我们执行下面的命令时：

curl b
必须得有一个 Service 名称为 b，这是前提。
在容器内，会根据 /etc/resolve.conf 进行解析流程。选择 nameserver 10.233.0.3 进行解析，然后，用字符串 “b”，依次带入 /etc/resolve.conf 中的 search 域，进行DNS查找，分别是：

// search 内容类似如下（不同的pod，第一个域会有所不同）
search default.svc.cluster.local svc.cluster.local cluster.local
b.default.svc.cluster.local -> b.svc.cluster.local -> b.cluster.local ，直到找到为止。
所以，我们执行 curl b，或者执行 curl b.default，都可以完成DNS请求，这2个不同的操作，会分别进行不同的DNS查找步骤：

// curl b，可以一次性找到（b +default.svc.cluster.local）
b.default.svc.cluster.local

// curl b.default，第一次找不到（ b.default + default.svc.cluster.local）
b.default.default.svc.cluster.local
// 第二次查找（ b.default + svc.cluster.local），可以找到
b.default.svc.cluster.local

https://hansedong.github.io/2018/11/20/9/

我们可以看到，在真正解析 youku.com 之前，经历了 youku.com.default.svc.cluster.local. -> youku.com.svc.cluster.local. -> youku.com.cluster.local. -> youku.com.

这也就意味着有3次DNS请求，是浪费的无意义的请求。

为何会出现DNS请求浪费的情况
这是因为，在 Kubernetes 中，其实 /etc/resolv.conf 这个文件，并不止包含 nameserver 和 search 域，还包含了非常重要的一项：ndots。我们之前没有提及这个项，也是希望再次能引起读者重视。

[root@xxxx-67f54c6dff-h4zxq /]# cat /etc/resolv.conf 
nameserver 10.233.0.3
search cicd.svc.cluster.local svc.cluster.local cluster.local
options ndots:5
ndots:5，表示：如果查询的域名包含的点“.”，不到5个，那么进行DNS查找，将使用非完全限定名称（或者叫绝对域名），如果你查询的域名包含点数大于等于5，那么DNS查询，默认会使用绝对域名进行查询。

DNS 服务器 – nameserver
我们先从nameserver 10.96.0.10来看, 为什么请求这个地址可以进行 DNS 解析. 这个答案就是 iptables, 我仅截取 UDP 的 53 端口, 以下内容可以通过iptables-save获得.

-A KUBE-SERVICES -d 10.96..10/32 -p udp -m comment --comment "kube-system/kube-dns:dns cluster IP" -m udp --dport 53 -j KUBE-SVC-TCOU7JCQXEZGVUNU
# 简单解释下, 这条规则表示, 如果目标地址是 10.96.0.10的udp53端口, 那么就会跳转到这条链上`KUBE-SVC-TCOU7JCQXEZGVUNU`

Kubernetes 的 Deployment
再看下我们的 Kubernetes 中 Pod 的 IP 地址, 也就是说, DNS 请求实际上会到我们的 Coredns 容器中被处理.'

可能有人会有疑问, 现在是 2 个 Pod 可以均分流量, 如果是 3 个, 4 个 Pod, Iptables 是如何做转发的呢, 正好我有这个疑问, 因此我就再加了 2 个 Pod, 看看iptables是怎么实现对于 4 个 Pod 均分流量的.

这是最后的实现方式:

-A KUBE-SVC-TCOU7JCQXEZGVUNU -m statistic --mode random --probability 0.25000000000 -j KUBE-SEP-HTZHQHQPOHVVNWZS
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m statistic --mode random --probability 0.33333333349 -j KUBE-SEP-3VNFB2SPYQJRRPK6
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-Q3HNNZPXUAYYDXW2
-A KUBE-SVC-TCOU7JCQXEZGVUNU -j KUBE-SEP-BBR3Z5NWFGXGVHEZ
这些语句的意思应该是:

前 1/4 的流量到一条链中, 剩 3/4
剩下 3/4 的流量, 1/3到一条链, 剩 2/4
剩下 2/4 的浏览, 1/2到一条链, 剩 1/4
最后 1/4 到一条链
通过这样的方式对流量进行了均分, 还是挺巧妙的, 这样, 5个,10个也是可以依次去分的.

假如没有这个search参数, 我们查找时:

> ping kube-dns
ping: kube-dns: Name or service not known
如果增加了search参数后, 再去查找:

> ping kube-dns
PING kube-dns.kube-system.svc.psigor-dev.nease.net (10.96.0.10) 56(84) bytes of data.

https://z.itpub.net/article/detail/4A6E5050C092D3D3FFB12B69A2A39547

https://www.bookstack.cn/read/feiskyer-kubernetes-handbook/components-kube-dns-internal.md

https://segmentfault.com/a/1190000007342180

CoreDNS 的实现原理
https://draveness.me/dns-coredns/


 dig -t A draveness.me +trace
Bash
我们可以使用 dig 命令追踪 draveness.me 域名对应 IP 地址是如何被解析出来的，首先会向预置的 13 组根域名服务器发出请求获取顶级域名的地址：

根域名服务器是 DNS 中最高级别的域名服务器，这些服务器负责返回顶级域的权威域名服务器地址，这些域名服务器的数量总共有 13 组，域名的格式从上面返回的结果可以看到是 .root-servers.net，每个根域名服务器中只存储了顶级域服务器的 IP 地址，大小其实也只有 2MB 左右，虽然域名服务器总共只有 13 组，但是每一组服务器都通过提供了镜像服务，全球大概也有几百台的根域名服务器在运行。

我们引入了胶水记录（Glue Record）这一概念，也就是在出现循环依赖时，直接在上一级作用域返回 DNS 服务器的 IP 地址：

也就是同时返回 NS 记录和 A（或 AAAA） 记录，这样就能够解决域名解析出现的循环依赖问题。


在微服务架构中，服务注册的方式其实大体上也只有两种，一种是使用 Zookeeper 和 etcd 等配置管理中心，另一种是使用 DNS 服务，比如说 Kubernetes 中的 CoreDNS 服务。


使用 DNS 在集群中做服务发现其实是一件比较容易的事情，这主要是因为绝大多数的计算机上都会安装 DNS 服务，所以这其实就是一种内置的、默认的服务发现方式，不过使用 DNS 做服务发现也会有一些问题，因为在默认情况下 DNS 记录的失效时间是 600s，这对于集群来讲其实并不是一个可以接受的时间，在实践中我们往往会启动单独的 DNS 服务满足服务发现的需求。


 CoreDNS 服务都建立在一个使用 Go 编写的 HTTP/2 Web 服务器 Caddy

 作为基于 Caddy 的 Web 服务器，CoreDNS 实现了一个插件链的架构，将很多 DNS 相关的逻辑都抽象成了一层一层的插件，包括 Kubernetes 等功能，每一个插件都是一个遵循如下协议的结构体：

 Corefile
另一个 CoreDNS 的特点就是它能够通过简单易懂的 DSL 定义 DNS 服务，在 Corefile 中就可以组合多个插件对外提供服务：

原理
CoreDNS 可以通过四种方式对外直接提供 DNS 服务，分别是 UDP、gRPC、HTTPS 和 TLS：

但是无论哪种类型的 DNS 服务，最终都会调用以下的 ServeDNS 方法，为服务的调用者提供 DNS 服务：


example.org {
    file /usr/local/etc/coredns/example.org
    prometheus     # enable metrics
    errors         # show errors
    log            # enable query logs
}
Text
那么在 CoreDNS 服务启动时，对于当前的 example.org 这个组，它会依次加载 file、log、errors 和 prometheus 几个插件


