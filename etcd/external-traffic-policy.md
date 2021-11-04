在k8s的Service对象（申明一条访问通道）中，有一个“externalTrafficPolicy”字段可以设置。有2个值可以设置：Cluster或者Local。

1）Cluster表示：流量可以转发到其他节点上的Pod。

2）Local表示：流量只发给本机的Pod。

选择（1）Cluster
注：这个是默认模式，Kube-proxy不管容器实例在哪，公平转发。

选择（2）Local
这种情况下，只转发给本机的容器，绝不跨节点转发。


https://www.cnblogs.com/zisefeizhu/p/13262239.html