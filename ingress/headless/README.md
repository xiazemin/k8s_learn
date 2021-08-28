有时不需要或不想要负载均衡，以及单独的 Service IP。遇到这种情况，可以通过指定 Cluster IP（spec.clusterIP）的值为 "None" 来创建 Headless Service。

https://cloud.tencent.com/developer/article/1638722

https://kubernetes.io/zh/docs/concepts/services-networking/service/


当我们把redis以pod的形式部署在k8s中时，每个pod里缓存的数据都是不一样的，而且pod的IP是会随时变化，这时候如果使用普通的deployment和service来部署redis-cluster就会出现很多问题，因此需要改用StatefulSet + Headless Service来解决

redis虽然是基于内存的缓存，但还是需要依赖于磁盘进行数据的持久化，以便服务出现问题重启时可以恢复已经缓存的数据。在集群中，我们需要使用共享文件系统 + PV（持久卷）的方式来让整个集群中的所有pod都可以共享同一份持久化储存

StatefulSet是k8s中专门用于解决有状态应用部署的一种资源，总的来说可以认为它是Deployment/RC的一个变种，它有以下几个特性：


StatefulSet管理的每个Pod都有唯一的文档/网络标识，并且按照数字规律生成，而不是像Deployment中那样名称和IP都是随机的（比如StatefulSet名字为redis，那么pod名就是redis-0, redis-1 ...）


StatefulSet中ReplicaSet的启停顺序是严格受控的，操作第N个pod一定要等前N-1个执行完才可以


StatefulSet中的Pod采用稳定的持久化储存，并且对应的PV不会随着Pod的删除而被销毁


另外需要说明的是，StatefulSet必须要配合Headless Service使用，它会在Headless Service提供的DNS映射上再加一层，最终形成精确到每个pod的域名映射，格式如下：
$(podname).$(headless service name)