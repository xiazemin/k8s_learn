

 % docker ps |grep etcd
9d2f3cf9d7bc   05b738aa1bc6             "etcd --advertise-cl…"   2 days ago       Up 2 days                 k8s_etcd_etcd-docker-desktop_kube-system_5d9d97b8d8daed31d6fd5c6d386c29c5_13
5a77504ea18e   k8s.gcr.io/pause:3.4.1   "/pause"                 2 days ago       Up 2 days                 k8s_POD_etcd-docker-desktop_kube-system_5d9d97b8d8daed31d6fd5c6d386c29c5_13



https://www.codenong.com/cs106430390/



 % kubectl get pods -n kube-system
NAME                                     READY   STATUS    RESTARTS   AGE
coredns-558bd4d5db-qc6px                 1/1     Running   13         69d
coredns-558bd4d5db-wzcgb                 1/1     Running   13         69d
etcd-docker-desktop                      1/1     Running   13         69d
kube-apiserver-docker-desktop            1/1     Running   13         69d
kube-controller-manager-docker-desktop   1/1     Running   13         69d
kube-proxy-k6shm                         1/1     Running   13         69d
kube-scheduler-docker-desktop            1/1     Running   15         69d
storage-provisioner                      1/1     Running   26         69d
vpnkit-controller                        1/1     Running   1944       69d




% kubectl describe pod etcd-docker-desktop -n kube-system
Name:                 etcd-docker-desktop
Namespace:            kube-system
Priority:             2000001000
Priority Class Name:  system-node-critical
Node:                 docker-desktop/192.168.65.4
Start Time:           Fri, 29 Oct 2021 16:23:09 +0800
Labels:               component=etcd
                      tier=control-plane
Annotations:          kubeadm.kubernetes.io/etcd.advertise-client-urls: https://192.168.65.4:2379
                      kubernetes.io/config.hash: 5d9d97b8d8daed31d6fd5c6d386c29c5
                      kubernetes.io/config.mirror: 5d9d97b8d8daed31d6fd5c6d386c29c5
                      kubernetes.io/config.seen: 2021-08-23T03:18:46.987692420Z
                      kubernetes.io/config.source: file
Status:               Running
IP:                   192.168.65.4
IPs:
  IP:           192.168.65.4
Controlled By:  Node/docker-desktop
Containers:
  etcd:
    Container ID:  docker://9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e
    Image:         k8s.gcr.io/etcd:3.4.13-0
    Image ID:      docker://sha256:05b738aa1bc6355db8a2ee8639f3631b908286e43f584a3d2ee0c472de033c28
    Port:          <none>
    Host Port:     <none>
    Command:
      etcd
      --advertise-client-urls=https://192.168.65.4:2379
      --cert-file=/run/config/pki/etcd/server.crt
      --client-cert-auth=true
      --data-dir=/var/lib/etcd
      --initial-advertise-peer-urls=https://192.168.65.4:2380
      --initial-cluster=docker-desktop=https://192.168.65.4:2380
      --key-file=/run/config/pki/etcd/server.key
      --listen-client-urls=https://127.0.0.1:2379,https://192.168.65.4:2379
      --listen-metrics-urls=http://127.0.0.1:2381
      --listen-peer-urls=https://192.168.65.4:2380
      --name=docker-desktop
      --peer-cert-file=/run/config/pki/etcd/peer.crt
      --peer-client-cert-auth=true
      --peer-key-file=/run/config/pki/etcd/peer.key
      --peer-trusted-ca-file=/run/config/pki/etcd/ca.crt
      --snapshot-count=10000
      --trusted-ca-file=/run/config/pki/etcd/ca.crt
    State:          Running
      Started:      Fri, 29 Oct 2021 16:23:10 +0800
    Last State:     Terminated
      Reason:       Error
      Exit Code:    255
      Started:      Mon, 25 Oct 2021 09:55:39 +0800
      Finished:     Fri, 29 Oct 2021 16:21:13 +0800
    Ready:          True
    Restart Count:  13
    Requests:
      cpu:        100m
      memory:     100Mi
    Liveness:     http-get http://127.0.0.1:2381/health delay=10s timeout=15s period=10s #success=1 #failure=8
    Startup:      http-get http://127.0.0.1:2381/health delay=10s timeout=15s period=10s #success=1 #failure=24
    Environment:  <none>
    Mounts:
      /run/config/pki/etcd from etcd-certs (rw)
      /var/lib/etcd from etcd-data (rw)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  etcd-certs:
    Type:          HostPath (bare host directory volume)
    Path:          /run/config/pki/etcd
    HostPathType:  DirectoryOrCreate
  etcd-data:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/etcd
    HostPathType:  DirectoryOrCreate
QoS Class:         Burstable
Node-Selectors:    <none>
Tolerations:       :NoExecute op=Exists
Events:            <none>




进入第二步找到的etcd容器docker exec -it 04607f704430 sh，查看/run/config/pki/etcd目录下的证书



ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get /registry/namespaces --prefix -w=json



ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt --key=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key



 % ETCDCTL_API=3  /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 --cacert=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt --key=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.key get /registry/namespaces --prefix -w=json
Error: KeyFile and CertFile must both be present[key: ~/Library/Group Containers/group.com.docker/pki/etcd/peer.key, cert: ]



ETCDCTL_API=3  /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 --cacert=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt --key=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.key get /registry/namespaces --prefix -w=json
Error: KeyFile and CertFile must both be present[key: ~/Library/Group Containers/group.com.docker/pki/etcd/ca.key, cert: ]


访问docker-desktop-for-mac安装的kubernetes的etcd
https://www.codenong.com/cs106430390/

https://blog.csdn.net/yezi1993/article/details/106430390



 % ETCDCTL_API=3  /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 --cacert=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt --key=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key --cert=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt get /registry/namespaces --prefix -w=json



  ETCDCTL_API=3  /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 --cacert=~/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt --key=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key --cert=~/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt get /registry/namespaces --prefix -w=json
Error: open ~/Library/Group Containers/group.com.docker/pki/etcd/peer.crt: no such file or directory


换成容器内部的：

ETCDCTL_API=3  /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get /registry/namespaces --prefix -w=json
Error: open /run/config/pki/etcd/peer.crt: no such file or directory
注意空格，删除空格就好了

ETCDCTL_API=3 /Users/xiazemin/source/etcd/bin/etcdctl --endpoints=https://127.0.0.1:2379 \
--cacert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt \
--key=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key \
--cert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt \
get /registry/namespaces --prefix -w=json


{"level":"warn","ts":"2021-11-01T11:59:31.114+0800","logger":"etcd-client","caller":"v3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0x1400039c540/#initially=[https://127.0.0.1:2379]","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: last connection error: connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:2379: connect: connection refused\""}
Error: context deadline exceeded


 ETCDCTL_API=3 etcdctl --cacert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt \
--cert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt \
--key=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key \
--endpoints=https://127.0.0.1:2379 \
get /registry/namespaces/default -w=json | jq .


{"level":"warn","ts":"2021-11-01T11:56:30.216+0800","logger":"etcd-client","caller":"v3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0x1400043a380/#initially=[https://127.0.0.1:2379]","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: last connection error: connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:2379: connect: connection refused\""}
Error: context deadline exceeded


 % kubectl get svc -n kube-system
NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   70d

没有暴露给集群外部访问的服务，所以访问不了


 % kubectl describe pod etcd-docker-desktop -n kube-system |grep 'Container ID'
    Container ID:  docker://9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e

 % docker exec -it  9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e sh
sh-5.0#


ETCDCTL_API=3  etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get /registry/namespaces --prefix -w=json

{"header":{"cluster_id":4202879228857769416,"member_id":16554063148076462710,"revision":1911370,"raft_term":15},"kvs":[{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMvZGVmYXVsdA==","create_revision":208,"mod_revision":208,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEoYCCusBCgdkZWZhdWx0EgAaACIAKiQzYjg0ZWMzMy02MzQ2LTRiMjgtOTI2MS0yNzU3ZWYzNWM4NWIyADgAQggIo6eMiQYQAFomChtrdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUSB2RlZmF1bHR6AIoBewoOa3ViZS1hcGlzZXJ2ZXISBlVwZGF0ZRoCdjEiCAijp4yJBhAAMghGaWVsZHNWMTpJCkd7ImY6bWV0YWRhdGEiOnsiZjpsYWJlbHMiOnsiLiI6e30sImY6a3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lIjp7fX19fRIMCgprdWJlcm5ldGVzGggKBkFjdGl2ZRoAIgA="},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1ub2RlLWxlYXNl","create_revision":52,"mod_revision":52,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEpYCCvsBCg9rdWJlLW5vZGUtbGVhc2USABoAIgAqJDQ5OTMxMzYxLTkzMzYtNGIwZC04MDQxLWQ3YzRiZDc2ZmIyNjIAOABCCAihp4yJBhAAWi4KG2t1YmVybmV0ZXMuaW8vbWV0YWRhdGEubmFtZRIPa3ViZS1ub2RlLWxlYXNlegCKAXsKDmt1YmUtYXBpc2VydmVyEgZVcGRhdGUaAnYxIggIoaeMiQYQADIIRmllbGRzVjE6SQpHeyJmOm1ldGFkYXRhIjp7ImY6bGFiZWxzIjp7Ii4iOnt9LCJmOmt1YmVybmV0ZXMuaW8vbWV0YWRhdGEubmFtZSI6e319fX0SDAoKa3ViZXJuZXRlcxoICgZBY3RpdmUaACIA"},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1wdWJsaWM=","create_revision":45,"mod_revision":45,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEo4CCvMBCgtrdWJlLXB1YmxpYxIAGgAiACokNTIxM2Q4NzAtNWRmNC00MzhlLWI1ZDgtMzc4OTA0ZmI5ZGQ5MgA4AEIICKGnjIkGEABaKgoba3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lEgtrdWJlLXB1YmxpY3oAigF7Cg5rdWJlLWFwaXNlcnZlchIGVXBkYXRlGgJ2MSIICKGnjIkGEAAyCEZpZWxkc1YxOkkKR3siZjptZXRhZGF0YSI6eyJmOmxhYmVscyI6eyIuIjp7fSwiZjprdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUiOnt9fX19EgwKCmt1YmVybmV0ZXMaCAoGQWN0aXZlGgAiAA=="},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1zeXN0ZW0=","create_revision":13,"mod_revision":13,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEo4CCvMBCgtrdWJlLXN5c3RlbRIAGgAiACokZWJhZDc5YzMtNmI3OC00OGQwLWI1ZGItZmI4YzNiYjk1MWVhMgA4AEIICKGnjIkGEABaKgoba3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lEgtrdWJlLXN5c3RlbXoAigF7Cg5rdWJlLWFwaXNlcnZlchIGVXBkYXRlGgJ2MSIICKGnjIkGEAAyCEZpZWxkc1YxOkkKR3siZjptZXRhZGF0YSI6eyJmOmxhYmVscyI6eyIuIjp7fSwiZjprdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUiOnt9fX19EgwKCmt1YmVybmV0ZXMaCAoGQWN0aXZlGgAiAA=="}],"count":4}



https://jimmysong.io/kubernetes-handbook/guide/using-etcdctl-to-access-kubernetes-data.html


https://zhuanlan.zhihu.com/p/94685947

https://www.huweihuang.com/kubernetes-notes/etcd/k8s-etcd-data.html


https://jimmysong.io/kubernetes-handbook/guide/using-etcdctl-to-access-kubernetes-data.html

https://jimmysong.io/kubernetes-handbook/concepts/etcd.html


https://jingwei.link/2018/11/25/kubernetes-etcd-data-save-specification.html

k8s主要把自己的数据注册在/registry/前缀下面（在ETCD-v3版本后没有了目录的概念，只能一切皆前缀了）。
通过观察k8s中deployment、namespace、pod等在ETCD中的表示，可以知道这部分资源的key的格式为/registry/#{k8s对象}/#{命名空间}/#{具体实例名}。
存在一个与众不同的key值compact_rev_key，搜索可以知道这是apiserver/compact.go中用来记录无效数据版本使用的；运行etcdctl get compact_rev_key可以发现，输出的是一个整形数值。
在查看ETCD时，k8s中除了必要的网络插件canal，未部署其他的应用，此时ETCD中只有240条数据，个人觉得这个量级没有想象中的多。


https://jimmysong.io/kubernetes-handbook/guide/using-etcdctl-to-access-kubernetes-data.html



ETCDCTL_API=3  etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get "" --prefix --keys-only |grep -Ev "^$"


/registry/apiregistration.k8s.io/apiservices/v1.

/registry/apiregistration.k8s.io/apiservices/v1.admissionregistration.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.apiextensions.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.apps

/registry/apiregistration.k8s.io/apiservices/v1.authentication.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.authorization.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.autoscaling

/registry/apiregistration.k8s.io/apiservices/v1.batch

/registry/apiregistration.k8s.io/apiservices/v1.certificates.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.coordination.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.discovery.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.events.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.networking.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.node.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.policy

/registry/apiregistration.k8s.io/apiservices/v1.rbac.authorization.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.scheduling.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1.storage.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.admissionregistration.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.apiextensions.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.authentication.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.authorization.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.batch

/registry/apiregistration.k8s.io/apiservices/v1beta1.certificates.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.coordination.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.discovery.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.events.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.extensions

/registry/apiregistration.k8s.io/apiservices/v1beta1.flowcontrol.apiserver.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.networking.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.node.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.policy

/registry/apiregistration.k8s.io/apiservices/v1beta1.rbac.authorization.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.scheduling.k8s.io

/registry/apiregistration.k8s.io/apiservices/v1beta1.storage.k8s.io

/registry/apiregistration.k8s.io/apiservices/v2beta1.autoscaling

/registry/apiregistration.k8s.io/apiservices/v2beta2.autoscaling

/registry/clusterrolebindings/cluster-admin

/registry/clusterrolebindings/docker-for-desktop-binding

/registry/clusterrolebindings/ingress-nginx

/registry/clusterrolebindings/ingress-nginx-admission

/registry/clusterrolebindings/kubeadm:get-nodes

/registry/clusterrolebindings/kubeadm:kubelet-bootstrap

/registry/clusterrolebindings/kubeadm:node-autoapprove-bootstrap

/registry/clusterrolebindings/kubeadm:node-autoapprove-certificate-rotation

/registry/clusterrolebindings/kubeadm:node-proxier

/registry/clusterrolebindings/storage-provisioner

/registry/clusterrolebindings/system:basic-user

/registry/clusterrolebindings/system:controller:attachdetach-controller

/registry/clusterrolebindings/system:controller:certificate-controller

/registry/clusterrolebindings/system:controller:clusterrole-aggregation-controller

/registry/clusterrolebindings/system:controller:cronjob-controller

/registry/clusterrolebindings/system:controller:daemon-set-controller

/registry/clusterrolebindings/system:controller:deployment-controller

/registry/clusterrolebindings/system:controller:disruption-controller

/registry/clusterrolebindings/system:controller:endpoint-controller

/registry/clusterrolebindings/system:controller:endpointslice-controller

/registry/clusterrolebindings/system:controller:endpointslicemirroring-controller

/registry/clusterrolebindings/system:controller:ephemeral-volume-controller

/registry/clusterrolebindings/system:controller:expand-controller

/registry/clusterrolebindings/system:controller:generic-garbage-collector

/registry/clusterrolebindings/system:controller:horizontal-pod-autoscaler

/registry/clusterrolebindings/system:controller:job-controller

/registry/clusterrolebindings/system:controller:namespace-controller

/registry/clusterrolebindings/system:controller:node-controller

/registry/clusterrolebindings/system:controller:persistent-volume-binder

/registry/clusterrolebindings/system:controller:pod-garbage-collector

/registry/clusterrolebindings/system:controller:pv-protection-controller

/registry/clusterrolebindings/system:controller:pvc-protection-controller

/registry/clusterrolebindings/system:controller:replicaset-controller

/registry/clusterrolebindings/system:controller:replication-controller

/registry/clusterrolebindings/system:controller:resourcequota-controller

/registry/clusterrolebindings/system:controller:root-ca-cert-publisher

/registry/clusterrolebindings/system:controller:route-controller

/registry/clusterrolebindings/system:controller:service-account-controller

/registry/clusterrolebindings/system:controller:service-controller

/registry/clusterrolebindings/system:controller:statefulset-controller

/registry/clusterrolebindings/system:controller:ttl-after-finished-controller

/registry/clusterrolebindings/system:controller:ttl-controller

/registry/clusterrolebindings/system:coredns

/registry/clusterrolebindings/system:discovery

/registry/clusterrolebindings/system:kube-controller-manager

/registry/clusterrolebindings/system:kube-dns

/registry/clusterrolebindings/system:kube-scheduler

/registry/clusterrolebindings/system:monitoring

/registry/clusterrolebindings/system:node

/registry/clusterrolebindings/system:node-proxier

/registry/clusterrolebindings/system:public-info-viewer

/registry/clusterrolebindings/system:service-account-issuer-discovery

/registry/clusterrolebindings/system:volume-scheduler

/registry/clusterrolebindings/vpnkit-controller

/registry/clusterroles/admin

/registry/clusterroles/cluster-admin

/registry/clusterroles/edit

/registry/clusterroles/ingress-nginx

/registry/clusterroles/ingress-nginx-admission

/registry/clusterroles/kubeadm:get-nodes

/registry/clusterroles/secret-reader

/registry/clusterroles/system:aggregate-to-admin

/registry/clusterroles/system:aggregate-to-edit

/registry/clusterroles/system:aggregate-to-view

/registry/clusterroles/system:auth-delegator

/registry/clusterroles/system:basic-user

/registry/clusterroles/system:certificates.k8s.io:certificatesigningrequests:nodeclient

/registry/clusterroles/system:certificates.k8s.io:certificatesigningrequests:selfnodeclient

/registry/clusterroles/system:certificates.k8s.io:kube-apiserver-client-approver

/registry/clusterroles/system:certificates.k8s.io:kube-apiserver-client-kubelet-approver

/registry/clusterroles/system:certificates.k8s.io:kubelet-serving-approver

/registry/clusterroles/system:certificates.k8s.io:legacy-unknown-approver

/registry/clusterroles/system:controller:attachdetach-controller

/registry/clusterroles/system:controller:certificate-controller

/registry/clusterroles/system:controller:clusterrole-aggregation-controller

/registry/clusterroles/system:controller:cronjob-controller

/registry/clusterroles/system:controller:daemon-set-controller

/registry/clusterroles/system:controller:deployment-controller

/registry/clusterroles/system:controller:disruption-controller

/registry/clusterroles/system:controller:endpoint-controller

/registry/clusterroles/system:controller:endpointslice-controller

/registry/clusterroles/system:controller:endpointslicemirroring-controller

/registry/clusterroles/system:controller:ephemeral-volume-controller

/registry/clusterroles/system:controller:expand-controller

/registry/clusterroles/system:controller:generic-garbage-collector

/registry/clusterroles/system:controller:horizontal-pod-autoscaler

/registry/clusterroles/system:controller:job-controller

/registry/clusterroles/system:controller:namespace-controller

/registry/clusterroles/system:controller:node-controller

/registry/clusterroles/system:controller:persistent-volume-binder

/registry/clusterroles/system:controller:pod-garbage-collector

/registry/clusterroles/system:controller:pv-protection-controller

/registry/clusterroles/system:controller:pvc-protection-controller

/registry/clusterroles/system:controller:replicaset-controller

/registry/clusterroles/system:controller:replication-controller

/registry/clusterroles/system:controller:resourcequota-controller

/registry/clusterroles/system:controller:root-ca-cert-publisher

/registry/clusterroles/system:controller:route-controller

/registry/clusterroles/system:controller:service-account-controller

/registry/clusterroles/system:controller:service-controller

/registry/clusterroles/system:controller:statefulset-controller

/registry/clusterroles/system:controller:ttl-after-finished-controller

/registry/clusterroles/system:controller:ttl-controller

/registry/clusterroles/system:coredns

/registry/clusterroles/system:discovery

/registry/clusterroles/system:heapster

/registry/clusterroles/system:kube-aggregator

/registry/clusterroles/system:kube-controller-manager

/registry/clusterroles/system:kube-dns

/registry/clusterroles/system:kube-scheduler

/registry/clusterroles/system:kubelet-api-admin

/registry/clusterroles/system:monitoring

/registry/clusterroles/system:node

/registry/clusterroles/system:node-bootstrapper

/registry/clusterroles/system:node-problem-detector

/registry/clusterroles/system:node-proxier

/registry/clusterroles/system:persistent-volume-provisioner

/registry/clusterroles/system:public-info-viewer

/registry/clusterroles/system:service-account-issuer-discovery

/registry/clusterroles/system:volume-scheduler

/registry/clusterroles/view

/registry/clusterroles/vpnkit-controller

/registry/configmaps/default/ingress-controller-leader-nginx

/registry/configmaps/default/ingress-nginx-controller

/registry/configmaps/default/kube-root-ca.crt

/registry/configmaps/default/redis-config

/registry/configmaps/kube-node-lease/kube-root-ca.crt

/registry/configmaps/kube-public/cluster-info

/registry/configmaps/kube-public/kube-root-ca.crt

/registry/configmaps/kube-system/coredns

/registry/configmaps/kube-system/extension-apiserver-authentication

/registry/configmaps/kube-system/kube-proxy

/registry/configmaps/kube-system/kube-root-ca.crt

/registry/configmaps/kube-system/kubeadm-config

/registry/configmaps/kube-system/kubelet-config-1.21

/registry/controllerrevisions/kube-system/kube-proxy-7cdcb64568

/registry/csinodes/docker-desktop

/registry/daemonsets/kube-system/kube-proxy

/registry/deployments/default/ingress-nginx-controller

/registry/deployments/default/minio-deployment

/registry/deployments/default/redis

/registry/deployments/kube-system/coredns

/registry/endpointslices/default/apple-service-5f9zk

/registry/endpointslices/default/ingress-nginx-controller-admission-shdx6

/registry/endpointslices/default/ingress-nginx-controller-jjm9l

/registry/endpointslices/default/kubernetes

/registry/endpointslices/default/minio-service-fd4dl

/registry/endpointslices/default/redis-gztm7

/registry/endpointslices/kube-system/kube-dns-wqqx2

/registry/events/kube-system/vpnkit-controller.16b272c39baa616d

/registry/events/kube-system/vpnkit-controller.16b272c39cb40bae

/registry/events/kube-system/vpnkit-controller.16b272c3a1d78e07

/registry/events/kube-system/vpnkit-controller.16b2742b54d7628e

/registry/flowschemas/catch-all

/registry/flowschemas/exempt

/registry/flowschemas/global-default

/registry/flowschemas/kube-controller-manager

/registry/flowschemas/kube-scheduler

/registry/flowschemas/kube-system-service-accounts

/registry/flowschemas/probes

/registry/flowschemas/service-accounts

/registry/flowschemas/system-leader-election

/registry/flowschemas/system-nodes

/registry/flowschemas/workload-leader-election

/registry/ingress/default/ingress-with-auth

/registry/jobs/default/ingress-nginx-admission-create

/registry/jobs/default/ingress-nginx-admission-patch

/registry/leases/kube-node-lease/docker-desktop

/registry/leases/kube-system/kube-scheduler

/registry/masterleases/192.168.65.4

/registry/minions/docker-desktop

/registry/namespaces/default

/registry/namespaces/kube-node-lease

/registry/namespaces/kube-public

/registry/namespaces/kube-system

/registry/persistentvolumeclaims/default/minio-pv-claim

/registry/persistentvolumeclaims/default/redis

/registry/persistentvolumes/minio

/registry/persistentvolumes/redis

/registry/pods/default/apple-app

/registry/pods/default/ingress-nginx-admission-create-vjn92

/registry/pods/default/ingress-nginx-admission-patch-wlq6p

/registry/pods/default/ingress-nginx-controller-57648496fc-84wl8

/registry/pods/default/minio-deployment-55bf5bff5d-cvq7v

/registry/pods/default/redis-f9f74787-tq6tw

/registry/pods/kube-system/coredns-558bd4d5db-qc6px

/registry/pods/kube-system/coredns-558bd4d5db-wzcgb

/registry/pods/kube-system/etcd-docker-desktop

/registry/pods/kube-system/kube-apiserver-docker-desktop

/registry/pods/kube-system/kube-controller-manager-docker-desktop

/registry/pods/kube-system/kube-proxy-k6shm

/registry/pods/kube-system/kube-scheduler-docker-desktop

/registry/pods/kube-system/storage-provisioner

/registry/pods/kube-system/vpnkit-controller

/registry/priorityclasses/system-cluster-critical

/registry/priorityclasses/system-node-critical

/registry/prioritylevelconfigurations/catch-all

/registry/prioritylevelconfigurations/exempt

/registry/prioritylevelconfigurations/global-default

/registry/prioritylevelconfigurations/leader-election

/registry/prioritylevelconfigurations/system

/registry/prioritylevelconfigurations/workload-high

/registry/prioritylevelconfigurations/workload-low

/registry/ranges/serviceips

/registry/ranges/servicenodeports

/registry/replicasets/default/ingress-nginx-controller-57648496fc

/registry/replicasets/default/minio-deployment-55bf5bff5d

/registry/replicasets/default/minio-deployment-6cfc69548

/registry/replicasets/default/minio-deployment-857cc8bbdf

/registry/replicasets/default/redis-f9f74787

/registry/replicasets/kube-system/coredns-558bd4d5db

/registry/rolebindings/default/ingress-nginx

/registry/rolebindings/default/ingress-nginx-admission

/registry/rolebindings/default/read-pods

/registry/rolebindings/kube-public/kubeadm:bootstrap-signer-clusterinfo

/registry/rolebindings/kube-public/system:controller:bootstrap-signer

/registry/rolebindings/kube-system/kube-proxy

/registry/rolebindings/kube-system/kubeadm:kubelet-config-1.21

/registry/rolebindings/kube-system/kubeadm:nodes-kubeadm-config

/registry/rolebindings/kube-system/system::extension-apiserver-authentication-reader

/registry/rolebindings/kube-system/system::leader-locking-kube-controller-manager

/registry/rolebindings/kube-system/system::leader-locking-kube-scheduler

/registry/rolebindings/kube-system/system:controller:bootstrap-signer

/registry/rolebindings/kube-system/system:controller:cloud-provider

/registry/rolebindings/kube-system/system:controller:token-cleaner

/registry/roles/default/ingress-nginx

/registry/roles/default/ingress-nginx-admission

/registry/roles/default/pod-reader

/registry/roles/kube-public/kubeadm:bootstrap-signer-clusterinfo

/registry/roles/kube-public/system:controller:bootstrap-signer

/registry/roles/kube-system/extension-apiserver-authentication-reader

/registry/roles/kube-system/kube-proxy

/registry/roles/kube-system/kubeadm:kubelet-config-1.21

/registry/roles/kube-system/kubeadm:nodes-kubeadm-config

/registry/roles/kube-system/system::leader-locking-kube-controller-manager

/registry/roles/kube-system/system::leader-locking-kube-scheduler

/registry/roles/kube-system/system:controller:bootstrap-signer

/registry/roles/kube-system/system:controller:cloud-provider

/registry/roles/kube-system/system:controller:token-cleaner

/registry/secrets/default/basic-auth

/registry/secrets/default/default-token-cg2vq

/registry/secrets/default/ingress-nginx-admission

/registry/secrets/default/ingress-nginx-admission-token-xblnc

/registry/secrets/default/ingress-nginx-token-7clh8

/registry/secrets/default/tls-secret

/registry/secrets/kube-node-lease/default-token-q97sv

/registry/secrets/kube-public/default-token-jq4g9

/registry/secrets/kube-system/attachdetach-controller-token-6lg9w

/registry/secrets/kube-system/bootstrap-signer-token-nnrn5

/registry/secrets/kube-system/certificate-controller-token-mwsrf

/registry/secrets/kube-system/clusterrole-aggregation-controller-token-27mp8

/registry/secrets/kube-system/coredns-token-5r4fz

/registry/secrets/kube-system/cronjob-controller-token-bvgp6

/registry/secrets/kube-system/daemon-set-controller-token-xx2z8

/registry/secrets/kube-system/default-token-pbdlw

/registry/secrets/kube-system/deployment-controller-token-8j72x

/registry/secrets/kube-system/disruption-controller-token-w9mmh

/registry/secrets/kube-system/endpoint-controller-token-t8dwq

/registry/secrets/kube-system/endpointslice-controller-token-dnbfm

/registry/secrets/kube-system/endpointslicemirroring-controller-token-plknt

/registry/secrets/kube-system/ephemeral-volume-controller-token-xm679

/registry/secrets/kube-system/expand-controller-token-9wgwf

/registry/secrets/kube-system/generic-garbage-collector-token-smjfw

/registry/secrets/kube-system/horizontal-pod-autoscaler-token-6jsj2

/registry/secrets/kube-system/job-controller-token-cft9x

/registry/secrets/kube-system/kube-proxy-token-c9bth

/registry/secrets/kube-system/namespace-controller-token-6rtph

/registry/secrets/kube-system/node-controller-token-ztkld

/registry/secrets/kube-system/persistent-volume-binder-token-zntw5

/registry/secrets/kube-system/pod-garbage-collector-token-qt297

/registry/secrets/kube-system/pv-protection-controller-token-fq7jx

/registry/secrets/kube-system/pvc-protection-controller-token-4krd7

/registry/secrets/kube-system/replicaset-controller-token-jjtjx

/registry/secrets/kube-system/replication-controller-token-kxdqb

/registry/secrets/kube-system/resourcequota-controller-token-n9zqv

/registry/secrets/kube-system/root-ca-cert-publisher-token-wldpn

/registry/secrets/kube-system/service-account-controller-token-wzgng

/registry/secrets/kube-system/service-controller-token-djmtc

/registry/secrets/kube-system/statefulset-controller-token-96msw

/registry/secrets/kube-system/storage-provisioner-token-n575g

/registry/secrets/kube-system/token-cleaner-token-vqxd2

/registry/secrets/kube-system/ttl-after-finished-controller-token-lmxng

/registry/secrets/kube-system/ttl-controller-token-xg5kg

/registry/secrets/kube-system/vpnkit-controller-token-8ztzx

/registry/serviceaccounts/default/default

/registry/serviceaccounts/default/ingress-nginx

/registry/serviceaccounts/default/ingress-nginx-admission

/registry/serviceaccounts/kube-node-lease/default

/registry/serviceaccounts/kube-public/default

/registry/serviceaccounts/kube-system/attachdetach-controller

/registry/serviceaccounts/kube-system/bootstrap-signer

/registry/serviceaccounts/kube-system/certificate-controller

/registry/serviceaccounts/kube-system/clusterrole-aggregation-controller

/registry/serviceaccounts/kube-system/coredns

/registry/serviceaccounts/kube-system/cronjob-controller

/registry/serviceaccounts/kube-system/daemon-set-controller

/registry/serviceaccounts/kube-system/default

/registry/serviceaccounts/kube-system/deployment-controller

/registry/serviceaccounts/kube-system/disruption-controller

/registry/serviceaccounts/kube-system/endpoint-controller

/registry/serviceaccounts/kube-system/endpointslice-controller

/registry/serviceaccounts/kube-system/endpointslicemirroring-controller

/registry/serviceaccounts/kube-system/ephemeral-volume-controller

/registry/serviceaccounts/kube-system/expand-controller

/registry/serviceaccounts/kube-system/generic-garbage-collector

/registry/serviceaccounts/kube-system/horizontal-pod-autoscaler

/registry/serviceaccounts/kube-system/job-controller

/registry/serviceaccounts/kube-system/kube-proxy

/registry/serviceaccounts/kube-system/namespace-controller

/registry/serviceaccounts/kube-system/node-controller

/registry/serviceaccounts/kube-system/persistent-volume-binder

/registry/serviceaccounts/kube-system/pod-garbage-collector

/registry/serviceaccounts/kube-system/pv-protection-controller

/registry/serviceaccounts/kube-system/pvc-protection-controller

/registry/serviceaccounts/kube-system/replicaset-controller

/registry/serviceaccounts/kube-system/replication-controller

/registry/serviceaccounts/kube-system/resourcequota-controller

/registry/serviceaccounts/kube-system/root-ca-cert-publisher

/registry/serviceaccounts/kube-system/service-account-controller

/registry/serviceaccounts/kube-system/service-controller

/registry/serviceaccounts/kube-system/statefulset-controller

/registry/serviceaccounts/kube-system/storage-provisioner

/registry/serviceaccounts/kube-system/token-cleaner

/registry/serviceaccounts/kube-system/ttl-after-finished-controller

/registry/serviceaccounts/kube-system/ttl-controller

/registry/serviceaccounts/kube-system/vpnkit-controller

/registry/services/endpoints/default/apple-service

/registry/services/endpoints/default/ingress-nginx-controller

/registry/services/endpoints/default/ingress-nginx-controller-admission

/registry/services/endpoints/default/kubernetes

/registry/services/endpoints/default/minio-service

/registry/services/endpoints/default/redis

/registry/services/endpoints/kube-system/docker.io-hostpath

/registry/services/endpoints/kube-system/kube-dns

/registry/services/specs/default/apple-service

/registry/services/specs/default/ingress-nginx-controller

/registry/services/specs/default/ingress-nginx-controller-admission

/registry/services/specs/default/kubernetes

/registry/services/specs/default/minio-service

/registry/services/specs/default/redis

/registry/services/specs/kube-system/kube-dns

/registry/storageclasses/hostpath

/registry/validatingwebhookconfigurations/ingress-nginx-admission

compact_rev_key


ETCDCTL_API=3  etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get /registry/ranges/serviceips

/registry/ranges/serviceips
k8s

v1RangeAllocation�

"*28Bz
      10.96.0.0/12�"

可以很明显看出来，ETCD中保存的并不是输出友好的数据（比如json或xml等就是输出友好型数据）。当然，如果进一步研究可以知道，ETCD保存的是Protocol Buffers序列化后的值。

ETCDCTL_API=3  etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/run/config/pki/etcd/ca.crt --key=/run/config/pki/etcd/peer.key --cert=/run/config/pki/etcd/peer.crt get \
/registry/configmaps/kube-system/canal-config


https://jingwei.link/2018/11/25/kubernetes-etcd-data-save-specification.html