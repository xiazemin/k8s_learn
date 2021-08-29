https://www.cnblogs.com/kevingrace/p/14412134.html

https://www.jianshu.com/p/a34790c730cf

https://segmentfault.com/a/1190000039196137

https://www.gotkx.com/?p=76



https://kubernetes.io/zh/docs/tutorials/configuration/configure-redis-using-configmap/
https://www.cnblogs.com/zisefeizhu/p/14282299.html
https://juejin.cn/post/6844903806719754254

http://www.mydlq.club/article/76/
https://www.jianshu.com/p/0fc0e1b869c6
https://www.codenong.com/cs109763447/

https://blog.csdn.net/xujiamin0022016/article/details/109763447
https://www.gl.sh.cn/2021/03/26/k8s_bu_shu_dan_jie_dian_redis.html

https://kuboard.cn/learning/k8s-practice/ocp/redis.html

https://blog.51cto.com/u_15100527/2616832

 % kubectl apply -f redis-config.yaml
configmap/redis-config created
% kubectl apply -f redis-deploy.yaml 
deployment.apps/redis created

kubectl apply -f redis-service.yaml 
service/redis created

 % docker run -it  --rm  redis:latest /bin/bash
root@948f69ef03a3:/data# 

 % kubectl get svc -o wide |grep redis 
redis                                NodePort       10.108.154.209   <none>        6379:30379/TCP               8s     app=redis


 %  redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> 

kubectl describe ing  ingress-with-auth 

kubectl logs ingress-nginx-controller-57648496fc-qn2lz  -n <namespace> 
https://www.cxyzjd.com/article/weixin_33788244/90227960

2021/08/28 07:08:44 [emerg] 75#75: io_setup() failed (38: Function not implemented)
192.168.65.3 - - [28/Aug/2021:07:08:56 +0000] "GET /apple HTTP/1.1" 500 39 "-" "curl/7.64.1" 91 0.009 [default-apple-service-5678] [] - - - - 1ebb5fd894d7bb64ff7c835c4012358a


挂载valume
    volumeMounts:
        - name: data
          mountPath: /data
      volumes:
      - name: data
        persistentVolumeClaim:
          claimName: redis


% kubectl delete -f exp1/redis-deploy.yaml 
deployment.apps "redis" deleted
% kubectl apply -f exp2/redis-deploy.yaml
deployment.apps/redis created         


% kubectl apply -f exp2/redis-storage.yaml 
persistentvolume/redis created
persistentvolumeclaim/redis created
% cp exp1/redis-service.yaml exp2 
% kubectl delete -f exp1/redis-service.yaml
service "redis" deleted
% kubectl apply -f exp2/redis-service.yaml 
service/redis created

redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> set 1 1


PersistentVolume（PV）和PersistentVolumeClaim（PVC）
这两个概念用于pod和volume之间解耦。Pod根据自己的需要提出数据卷的申请，k8s系统将符合条件的数据卷返回给pod。这样一来pod就无需直接和数据卷本身强绑定了。Pod无需知道数据卷的细节信息，比如具体用的是什么存储。

Pod中对数据卷的申请为PVC，用来和PVC绑定的数据卷称为PV。

PV可以有多种存储方式，比如NFS，iSCSI等。

PVC是使用PV资源的声明。

https://www.jianshu.com/p/b3665b72126e

LocalPersistentVolume
这种类型的PV把数据储存在集群中某个节点上。

PV类型
　　PV类型使用插件的形式来实现。Kubernetes现在支持以下插件：
　　GCEPersistentDisk
　　AWSElasticBlockStore
　　AzureFile
　　AzureDisk
　　FC (Fibre Channel)
　　Flocker
　　NFS
　　iSCSI
　　RBD (Ceph Block Device)
　　CephFS
　　Cinder (OpenStack block storage)
　　Glusterfs
　　VsphereVolume
　　Quobyte Volumes
　　HostPath (仅测试过单节点的情况——不支持任何形式的本地存储，多节点集群中不能工作)
　　VMware Photon
　　Portworx Volumes
　　ScaleIO Volumes

https://blog.csdn.net/styshoo/article/details/72235208
https://blog.51cto.com/jacksoner/2333318


Local Persistent Volumes代表了直接绑定在计算节点上的一块本地磁盘。
kubernetes提供了一套卷插件（volume plugin）标准，使得k8s集群的工作负载可以使用多种块存储和文件存储。大部分磁盘插件都使用了远程存储，这是为了让持久化的数据与计算节点彼此独立，但远程存储通常无法提供本地存储那么强的读写性能。有了Local PV 插件，kubernetes负载现在可以用同样的volume api，在容器中使用本地磁盘。（分布式文件系统和数据库一直是 Local PV 的主要用例）


这跟hostPath有什么区别
hostPath是一种volume，可以让pod挂载宿主机上的一个文件或目录（如果挂载路径不存在，则创建为目录或文件并挂载）。
最大的不同在于调度器是否能理解磁盘和node的对应关系，一个使用hostPath的pod，当他被重新调度时，很有可能被调度到与原先不同的node上，这就导致pod内数据丢失了。而使用Local PV的pod，总会被调度到同一个node上（否则就调度失败）。

https://blog.csdn.net/bh1231/article/details/100947983


配置 Pod 以使用 PersistentVolume 作为hostPath存储
https://blog.csdn.net/ljx1528/article/details/113618996

https://www.cnblogs.com/rongfengliang/p/9237832.html

https://blog.csdn.net/ljx1528/article/details/113618996

% kubectl delete -f exp2/redis-storage.yaml 
persistentvolume "redis" deleted
persistentvolumeclaim "redis" deleted

% kubectl apply -f exp2/hostPathPV.yaml 
persistentvolume/redis created

Warning: Detected changes to resource redis which is currently being deleted.
persistentvolumeclaim/redis unchanged


% kubectl delete -f exp2/redis-storage.yaml  --force --grace-period=0
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
persistentvolumeclaim "redis" force deleted
Error from server (NotFound): error when deleting "exp2/redis-storage.yaml": persistentvolumes "redis" not found

https://www.jianshu.com/p/828f85cdef5f



https://segmentfault.com/a/1190000016021217


 % redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> save
OK
127.0.0.1:30379> bgsave
Background saving started
127.0.0.1:30379> 

% redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> save
Error: Server closed the connection

3.rbd备份文件损坏检测.可以使用redis-check-rdb工具检测rdb文件,该工具默认在/usr/local/bin/目录下面.
```shell
[root@syncd redis-data]# /usr/local/bin/redis-check-rdb ./6379-rdb.rdb

https://www.cnblogs.com/qqblog/p/10507667.html

 % docker ps |grep redis
7d2c78f3046d   redis                    "sh -c 'redis-server…"   3 minutes ago   Up 3 minutes                                                                                                 k8s_redis_redis-f9f74787-r8nmc_default_76e879b5-a600-49eb-a866-64697561b4ab_0


docker exec -it 7d2c78f3046d /bin/bash
root@redis-f9f74787-r8nmc:/data# touch testVolume.txt

% ls -al test
total 0

root@redis-f9f74787-r8nmc:/data# which redis-cli
/usr/local/bin/redis-cli


# cat  /usr/local/redis/redis.conf
dir /srv
port 6379
bind 0.0.0.0
appendonly yes
daemonize no
#protected-mode no
requirepass 123456

# ls /srv/
appendonly.aof  dump.rdb  redis-6379.pid


路径不对

重启大法
 % docker exec -it bdbbb676f42e /bin/bash
root@redis-f9f74787-9fdn4:/data# ls
appendonly.aof  redis-6379.pid  testVolume.txt

% kubectl apply -f redis-pvc.yaml         
The PersistentVolumeClaim "redis" is invalid: spec.resources.requests.storage: Forbidden: field can not be less than previous value

https://blog.csdn.net/weixin_36156325/article/details/113582728


https://www.jianshu.com/p/b3665b72126e

http://team.jiunile.com/blog/2020/09/k8s-local-volume.html



% kubectl get pv redis 
NAME    CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
redis   1Gi        RWO            Retain           Available           manual                  136m
% kubectl get pvc  redis 
NAME    STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis   Bound    pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d   1Gi        RWO            hostpath       10m
% kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM           STORAGECLASS   REASON   AGE
pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d   1Gi        RWO            Delete           Bound       default/redis   hostpath                13m
redis                                      1Gi        RWO            Retain           Available                   manual                  140m
% kubectl get pvc
NAME    STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis   Bound    pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d   1Gi        RWO            hostpath       14m
% kubectl get pod
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     0          4d10h
ingress-nginx-admission-create-vjn92        0/1     Completed   0          24h
ingress-nginx-admission-patch-wlq6p         0/1     Completed   0          24h
ingress-nginx-controller-57648496fc-84wl8   1/1     Running     0          24h
redis-f9f74787-tn46t                        1/1     Running     0          14m
% kubectl describe pod redis-f9f74787-tn46t 
Name:         redis-f9f74787-tn46t
Namespace:    default
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Sun, 29 Aug 2021 20:56:10 +0800
Labels:       app=redis
              pod-template-hash=f9f74787
Annotations:  <none>
Status:       Running
IP:           10.1.0.76
IPs:
  IP:           10.1.0.76
Controlled By:  ReplicaSet/redis-f9f74787
Containers:
  redis:
    Container ID:  docker://be1745da6d43e8a9c0b53a0570e12ff720bd3dea03f0ec01a9f55ae233aefa8f
    Image:         redis:latest
    Image ID:      docker-pullable://redis@sha256:66ce9bc742609650afc3de7009658473ed601db4e926a5b16d239303383bacad
    Port:          6379/TCP
    Host Port:     0/TCP
    Command:
      sh
      -c
      redis-server /usr/local/redis/redis.conf
    State:          Running
      Started:      Sun, 29 Aug 2021 20:56:27 +0800
    Ready:          True
    Restart Count:  0
    Limits:
      cpu:     1
      memory:  1Gi
    Requests:
      cpu:        1
      memory:     1Gi
    Liveness:     tcp-socket :6379 delay=300s timeout=1s period=10s #success=1 #failure=3
    Readiness:    tcp-socket :6379 delay=5s timeout=1s period=10s #success=1 #failure=3
    Environment:  <none>
    Mounts:
      /data from data (rw)
      /usr/local/redis/redis.conf from config (rw,path="redis.conf")
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-f7njk (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  data:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  redis
    ReadOnly:   false
  config:
    Type:      ConfigMap (a volume populated by a ConfigMap)
    Name:      redis-config
    Optional:  false
  kube-api-access-f7njk:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Guaranteed
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  14m   default-scheduler  Successfully assigned default/redis-f9f74787-tn46t to docker-desktop
  Normal  Pulling    14m   kubelet            Pulling image "redis:latest"
  Normal  Pulled     14m   kubelet            Successfully pulled image "redis:latest" in 16.322640215s
  Normal  Created    14m   kubelet            Created container redis
  Normal  Started    14m   kubelet            Started container redis
% kubectl describe pvc redis
Name:          redis
Namespace:     default
StorageClass:  hostpath
Status:        Bound
Volume:        pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: docker.io/hostpath
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      1Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Used By:       redis-f9f74787-tn46t
Events:
  Type    Reason                 Age   From                                                                         Message
  ----    ------                 ----  ----                                                                         -------
  Normal  ExternalProvisioning   15m   persistentvolume-controller                                                  waiting for a volume to be created, either by external provisioner "docker.io/hostpath" or manually created by system administrator
  Normal  Provisioning           15m   docker.io/hostpath_storage-provisioner_d6db29be-248d-4504-b8a5-add2e93a761b  External provisioner is provisioning volume for claim "default/redis"
  Normal  ProvisioningSucceeded  15m   docker.io/hostpath_storage-provisioner_d6db29be-248d-4504-b8a5-add2e93a761b  Successfully provisioned volume pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
% kubectl describe pv redis 
Name:            redis
Labels:          type=local
Annotations:     <none>
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    manual
Status:          Available
Claim:           
Reclaim Policy:  Retain
Access Modes:    RWO
VolumeMode:      Filesystem
Capacity:        1Gi
Node Affinity:   <none>
Message:         
Source:
    Type:          HostPath (bare host directory volume)
    Path:          /Users/xiazemin/source/k8s_learn/redis/test
    HostPathType:  
Events:            <none>
% kubectl describe pv pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
Name:            pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
Labels:          <none>
Annotations:     docker.io/hostpath: /var/lib/k8s-pvs/redis/pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
                 pv.kubernetes.io/provisioned-by: docker.io/hostpath
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    hostpath
Status:          Bound
Claim:           default/redis
Reclaim Policy:  Delete
Access Modes:    RWO
VolumeMode:      Filesystem
Capacity:        1Gi
Node Affinity:   <none>
Message:         
Source:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/k8s-pvs/redis/pvc-67f08d53-eed7-41e2-a972-3cf6171f3c0d
    HostPathType:  
Events:            <none>


重新来一遍，发现% kubectl apply -f redis-pvc.yaml 

多了个pv

 % kubectl get pv 
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM           STORAGECLASS   REASON   AGE
pvc-0ae99a88-2988-4e96-91f2-1050d62a5c58   1Gi        RWO            Delete           Bound       default/redis   hostpath                6s
redis                                      1Gi        RWO            Retain           Available                   manual                  73s

 % kubectl get pvc
NAME    STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis   Bound    pvc-0ae99a88-2988-4e96-91f2-1050d62a5c58   1Gi        RWO            hostpath       5s



% kubectl describe pvc redis                                   
Name:          redis
Namespace:     default
StorageClass:  hostpath
Status:        Bound
Volume:        pvc-0ae99a88-2988-4e96-91f2-1050d62a5c58
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: docker.io/hostpath
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      1Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Used By:       <none>
Events:
  Type    Reason                 Age                  From                                                                         Message
  ----    ------                 ----                 ----                                                                         -------
  Normal  ExternalProvisioning   117s (x2 over 117s)  persistentvolume-controller                                                  waiting for a volume to be created, either by external provisioner "docker.io/hostpath" or manually created by system administrator
  Normal  Provisioning           117s                 docker.io/hostpath_storage-provisioner_d6db29be-248d-4504-b8a5-add2e93a761b  External provisioner is provisioning volume for claim "default/redis"
  Normal  ProvisioningSucceeded  117s                 docker.io/hostpath_storage-provisioner_d6db29be-248d-4504-b8a5-add2e93a761b  Successfully provisioned volume pvc-0ae99a88-2988-4e96-91f2-1050d62a5c58


如何控制kubernetes PersistentVolumeClaim绑定到特定的PersistentVolume？

https://cloud.google.com/kubernetes-engine/docs/how-to/persistent-volumes/preexisting-pd?hl=zh_cn

需要在pv 上指定pvc name
  claimRef:
    namespace: default
    name: PV_CLAIM_NAME

By specifying a PersistentVolume in a PersistentVolumeClaim, you declare a binding between that specific PV and PVC. If the PersistentVolume exists and has not reserved PersistentVolumeClaims through its claimRef field, then the PersistentVolume and PersistentVolumeClaim will be bound.

% kubectl get pvc
NAME    STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis   Bound    redis    1Gi        RWO            hostpath       14s
% kubectl get pvc
NAME    STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis   Bound    redis    1Gi        RWO            hostpath       14s

 % kubectl describe pvc redis
Name:          redis
Namespace:     default
StorageClass:  hostpath
Status:        Bound
Volume:        redis
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      1Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Used By:       <none>
Events:        <none>


% redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> set x 1
OK
127.0.0.1:30379> save
OK
127.0.0.1:30379> exit
% ls -al ../test                           
total 24
drwxr-xr-x  5 xiazemin  staff  160  8 29 21:37 .
drwxr-xr-x  6 xiazemin  staff  192  8 29 17:38 ..
-rw-r--r--  1 xiazemin  staff   50  8 29 21:37 appendonly.aof
-rw-r--r--  1 xiazemin  staff  102  8 29 21:37 dump.rdb
-rw-r--r--  1 xiazemin  staff    2  8 29 21:36 redis-6379.pid

https://kubernetes.io/zh/docs/tasks/configure-pod-container/configure-persistent-volume-storage/


创建 PersistentVolumeClaim 之后，Kubernetes 控制平面将查找满足申领要求的 PersistentVolume。 如果控制平面找到具有相同 StorageClass 的适当的 PersistentVolume， 则将 PersistentVolumeClaim 绑定到该 PersistentVolume 上。


% kubectl get pv 
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM           STORAGECLASS   REASON   AGE
pvc-0ae99a88-2988-4e96-91f2-1050d62a5c58   1Gi        RWO            Delete           Bound       default/redis   hostpath                11m
redis                                      1Gi        RWO            Retain           Available                   manual                  12m

Reclaiming
When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource. The reclaim policy for a PersistentVolume tells the cluster what to do with the volume after it has been released of its claim. Currently, volumes can either be Retained, Recycled, or Deleted.


https://kubernetes.io/docs/concepts/storage/persistent-volumes/

