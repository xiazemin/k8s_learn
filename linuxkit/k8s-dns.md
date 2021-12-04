% kubectl exec -it redis-f9f74787-tq6tw -- sh
# cat /etc/resolv.conf
nameserver 10.96.0.10
search default.svc.cluster.local svc.cluster.local cluster.local
options ndots:5

这个文件中，配置的 DNS Server，一般就是 K8S 中，kubedns 的 Service 的 ClusterIP，这个IP是虚拟IP，无法ping，但可以访问。


 % kubectl get svc -n kube-system
NAME                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
etcd-svc-docker-desktop-xzm   NodePort    10.111.136.178   <none>        2379:32389/TCP           25d
kube-dns                      ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   98d


可以看到nameserver 的ip 10.96.0.10 就是 kubedns的ip


 % docker run -it --privileged --pid=host alpine:latest nsenter sh
/ #  curl  10.96.0.10



在容器内发请求时，会根据 /etc/resolv.conf 进行解析流程。选择 nameserver 10.96.0.10 进行解析，然后用nginx-svc-old ，依次带入 /etc/resolve.conf 中的 search 域，进行DNS查找，分别是：

search 内容类似如下（不同的pod，第一个域会有所不同）

search default.svc.cluster.local svc.cluster.local cluster.local

search default.svc.cluster.local svc.cluster.local cluster.local
nginx-svc-old.default.svc.cluster.local -> nginx-svc-old.svc.cluster.local -> nginx-svc-old.cluster.local 

我们执行 ping nginx-svc-old，或者执行 ping nginx-svc-old.default，都可以完成DNS请求，这2个不同的操作，会分别进行不同的DNS查找步骤。



 % kubectl exec -i -t dnsutils -- nslookup kube-dns.kube-system
Server:		10.96.0.10
Address:	10.96.0.10#53

Name:	kube-dns.kube-system.svc.cluster.local
Address: 10.96.0.10




