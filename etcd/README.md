k8s的各个组件与apiServer交互操作各种资源对象，最终都会落入到etcd中。
http://www.dockone.io/article/2177

Storage创建
要了解etcd操作接口的实现，我们先需要了解下Master.GenericAPIServer.storage结构：
storage map[string]rest.Storage


该storage变量是个map，Key是REST API的path，Value是rest.Storage接口，该接口就是一个通用的符合Restful要求的资源存储接口。

https://zwforrest.github.io/post/kube-apiserver%E5%AD%98%E5%82%A8etcd%E6%9E%B6%E6%9E%84%E5%88%86%E6%9E%90/

https://www.jianshu.com/p/daa4ff387a78
https://juejin.cn/post/6870844477020307463
https://www.bookstack.cn/read/source-code-reading-notes/kubernetes-kube_apiserver.md

https://www.zhihu.com/column/c_1195294063723929600

K8s中所有元数据的增删改查都是由kube-apiserver来执行的，那么这些数据在ETCD中必然有一套存储规范，这样才能保证在集群中部署成千上万的应用时不会出差错。在此基础上可以认为，只要掌握了k8s在ETCD中存储数据的规范，便可以像k8s一样手动来操作ETCD了（虽然不建议这么做
https://jingwei.link/2018/11/25/kubernetes-etcd-data-save-specification.html

https://github.com/etcd-io/etcd

https://etcd.io/docs/v3.5/install/


git clone -b v3.5.0 https://github.com/etcd-io/etcd.git

 ./build.sh

export PATH="$PATH:`pwd`/bin"



 %  ETCDCTL_API=3 etcdctl get "" --prefix --keys-only |grep -Ev "^$"
{"level":"warn","ts":"2021-11-01T11:09:45.127+0800","logger":"etcd-client","caller":"v3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0x140006ba380/#initially=[127.0.0.1:2379]","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: last connection error: connection error: desc = \"transport: Error while dialing dial tcp 127.0.0.1:2379: connect: connection refused\""}
Error: context deadline exceeded




https://csunny.gitbook.io/etcd/revision