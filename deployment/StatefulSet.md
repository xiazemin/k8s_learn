那么在k8s中，在2种工作负载的调度有何异同？
【相同】

均以container以落地形式
均支持增加replica（副本）横向扩展
均支持挂载存储
【不同】

无状态：
副本无序随机命名，销毁后重新生成命名
并行扩容，并行缩容
有状态：
副本以0,1,2有序命名，销毁后名称不变，可通过访问副本名保持访问。
串行有序扩容，顺序为0,1,2。 串行有序缩容，顺序反之。

https://blog.csdn.net/tonadowang/article/details/117304372?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~aggregatepage~first_rank_ecpm_v1~rank_v31_ecpm-1-117304372.pc_agg_new_rank&utm_term=k8s%E4%B8%ADdeployment%E4%B8%8EStatefulSet%E5%8C%BA%E5%88%AB&spm=1000.2123.3001.4430


1、介绍
RC、Deployment、DaemonSet都是面向无状态的服务，它们所管理的Pod的IP、名字，启停顺序等都是随机的，而StatefulSet是什么？顾名思义，有状态的集合，管理所有有状态的服务，比如MySQL、MongoDB集群等。
StatefulSet本质上是Deployment的一种变体，在v1.9版本中已成为GA版本，它为了解决有状态服务的问题，它所管理的Pod拥有固定的Pod名称，启停顺序，在StatefulSet中，Pod名字称为网络标识(hostname)，还必须要用到共享存储。
在Deployment中，与之对应的服务是service，而在StatefulSet中与之对应的headless service，headless service，即无头服务，与service的区别就是它没有Cluster IP，解析它的名称时将返回该Headless Service对应的全部Pod的Endpoint列表。
除此之外，StatefulSet在Headless Service的基础上又为StatefulSet控制的每个Pod副本创建了一个DNS域名，这个域名的格式为：
$(podname).(headless server name)
FQDN：$(podname).(headless server name).namespace.svc.cluster.local

2、特点
Pod一致性：包含次序（启动、停止次序）、网络一致性。此一致性与Pod相关，与被调度到哪个node节点无关；
稳定的次序：对于N个副本的StatefulSet，每个Pod都在[0，N)的范围内分配一个数字序号，且是唯一的；
稳定的网络：Pod的hostname模式为( s t a t e f u l s e t 名 称 ) − (statefulset名称)-(statefulset名称)−(序号)；
稳定的存储：通过VolumeClaimTemplate为每个Pod创建一个PV。删除、减少副本，不会删除相关的卷。

3、组成部分
Headless Service：用来定义Pod网络标识( DNS domain)；
volumeClaimTemplates ：存储卷申请模板，创建PVC，指定pvc名称大小，将自动创建pvc，且pvc必须由存储类供应；
StatefulSet ：定义具体应用，名为Nginx，有三个Pod副本，并为每个Pod定义了一个域名部署statefulset。

为什么需要 headless service 无头服务？
在用Deployment时，每一个Pod名称是没有顺序的，是随机字符串，因此是Pod名称是无序的，但是在statefulset中要求必须是有序 ，每一个pod不能被随意取代，pod重建后pod名称还是一样的。而pod IP是变化的，所以是以Pod名称来识别。pod名称是pod唯一性的标识符，必须持久稳定有效。这时候要用到无头服务，它可以给每个Pod一个唯一的名称 。

为什么需要volumeClaimTemplate？
对于有状态的副本集都会用到持久存储，对于分布式系统来讲，它的最大特点是数据是不一样的，所以各个节点不能使用同一存储卷，每个节点有自已的专用存储，但是如果在Deployment中的Pod template里定义的存储卷，是所有副本集共用一个存储卷，数据是相同的，因为是基于模板来的 ，而statefulset中每个Pod都要自已的专有存储卷，所以statefulset的存储卷就不能再用Pod模板来创建了，于是statefulSet使用volumeClaimTemplate，称为卷申请模板，它会为每个Pod生成不同的pvc，并绑定pv，从而实现各pod有专用存储。这就是为什么要用volumeClaimTemplate的原因。

4、StatefulSet详解
kubectl explain sts.spec ：主要字段解释
replicas ：副本数
selector：那个pod是由自己管理的
serviceName：必须关联到一个无头服务商
template：定义pod模板（其中定义关联那个存储卷）
volumeClaimTemplates ：生成PVC

https://blog.csdn.net/weixin_44729138/article/details/106054025


重启pod会发现，pod中的ip已经发生变化，但是pod的名称并没有发生变化；这就是为什么不要在其他应用中使用 StatefulSet 中的 Pod 的 IP 地址进行连接，这点很重要



Pod 的序号、主机名、SRV 条目和记录名称没有改变，但和 Pod 相关联的 IP 地址可能发生了改变
如果你需要查找并连接一个 StatefulSet 的活动成员，你应该查询 Headless Service 的 CNAME。和 CNAME 相关联的 SRV 记录只会包含 StatefulSet 中处于 Running 和 Ready 状态的 Pod。

如果你的应用已经实现了用于测试 liveness 和 readiness 的连接逻辑，你可以使用 Pod 的 SRV 记录（web-0.nginx.nginx-ss.svc.cluster.local， web-1.nginx.nginx-ss.svc.cluster.local，web-2.nginx.nginx-ss.svc.cluster.local）。因为他们是稳定的，并且当你的 Pod 的状态变为 Running 和 Ready 时，你的应用就能够发现它们的地址。

扩容/缩容 StatefulSet
扩容/缩容StatefulSet 指增加或减少它的副本数。这通过更新replicas字段完成。你可以使用kubectl scale 或者kubectl patch来扩容/缩容一个 StatefulSet。

kubectl scale sts web --replicas=4 -n nginx-ss   #扩容
kubectl scale sts web --replicas=2 -n nginx-ss   #缩容



% kubectl apply -f deployment/StorageClass.yaml
storageclass.storage.k8s.io/apple-nfs-storage created



% kubectl apply -f deployment/statefulset.yaml
statefulset.apps/apple-ss created


% kubectl get statefulset
NAME       READY   AGE
apple-ss   0/3     22s


% kubectl apply -f deployment/statefulset.yaml 
The StatefulSet "apple-ss" is invalid: spec: Forbidden: updates to statefulset spec for fields other than 'replicas', 'template', and 'updateStrategy' are forbidden

 % kubectl delete -f deployment/statefulset.yaml
statefulset.apps "apple-ss" deleted

% kubectl apply -f deployment/statefulset.yaml 
statefulset.apps/apple-ss created


 % kubectl describe pod apple-ss-0        
Name:           apple-ss-0
Namespace:      default
Priority:       0
Node:           <none>
Labels:         app=apple
                controller-revision-hash=apple-ss-6bb7949cbd
                statefulset.kubernetes.io/pod-name=apple-ss-0
Annotations:    <none>
Status:         Pending
IP:             
IPs:            <none>
Controlled By:  StatefulSet/apple-ss
Containers:
  apple:
    Image:        apple:5678
    Port:         5678/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:
      /usr/share/apple/html from apple-pvc (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-qzvxw (ro)
Conditions:
  Type           Status
  PodScheduled   False 
Volumes:
  apple-pvc:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  apple-pvc-apple-ss-0
    ReadOnly:   false
  kube-api-access-qzvxw:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason            Age   From               Message
  ----     ------            ----  ----               -------
  Warning  FailedScheduling  39s   default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.
  Warning  FailedScheduling  38s   default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.


% kubectl get StorageClass
NAME                 PROVISIONER          RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
apple-nfs-storage    fuseim.pri/ifs       Delete          Immediate           false                  8m53s
hostpath (default)   docker.io/hostpath   Delete          Immediate           false                  104d

https://blog.csdn.net/yrx420909/article/details/106127489/

% kubectl get pvc         
NAME                   STATUS    VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS        AGE
apple-pvc-apple-ss-0   Pending                                      apple-nfs-storage   9m10s
minio-pv-claim         Bound     minio    2Gi        RWO            hostpath            97d
redis                  Bound     redis    1Gi        RWO            hostpath            97d


 % kubectl describe pvc apple-pvc-apple-ss-0 
Name:          apple-pvc-apple-ss-0
Namespace:     default
StorageClass:  apple-nfs-storage
Status:        Pending
Volume:        
Labels:        app=apple
Annotations:   volume.beta.kubernetes.io/storage-provisioner: fuseim.pri/ifs
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      
Access Modes:  
VolumeMode:    Filesystem
Used By:       apple-ss-0
Events:
  Type    Reason                Age                   From                         Message
  ----    ------                ----                  ----                         -------
  Normal  ExternalProvisioning  40s (x12 over 9m55s)  persistentvolume-controller  waiting for a volume to be created, either by external provisioner "fuseim.pri/ifs" or manually created by system administrator


后面External provisioner给他创建pv来提供。一会时间。但是之前那个是一直创建不了，一直pending，那就是External provisioner的pod(deployment)有问题。

如果你装nfs-provisioner，出现这种情况，那么就重启虚拟机，不行就重启电脑，再重新部署deployment（可能也不需要吧，就删除pod）

https://www.jianshu.com/p/5e565a8049fc


环境变量中有个PROVISIONER_NAME，值为fuseim.pri/ifs，后面的StorageClass就以“fuseim.pri/ifs”作为参数来定位到此Pod作为NFS服务的提供者；

https://blog.csdn.net/boling_cavalry/article/details/79598905

csi 是什么？csi 的全称是 container storage interface，它是 K8s 社区后面对存储插件实现 ( out of tree ) 的官方推荐方式。csi 的实现大体可以分为两部分：

第一部分是由 k8s 社区驱动实现的通用的部分，像我们这张图中的 csi-provisioner和 csi-attacher controller；
另外一种是由云存储厂商实践的，对接云存储厂商的 OpenApi，主要是实现真正的 create/delete/mount/unmount 存储的相关操作，对应到上图中的 csi-controller-server 和 csi-node-server。
接下来看一下，当用户提交 yaml 之后，k8s 内部的处理流程。用户在提交 PVCyaml 的时候，首先会在集群中生成一个 PVC 对象，然后 PVC 对象会被 csi-provisioner controller watch 到，csi-provisioner 会结合 PVC 对象以及 PVC 对象中声明的 storageClass，通过 GRPC 调用 csi-controller-server，然后，到云存储服务这边去创建真正的存储，并最终创建出来 PV 对象。最后，由集群中的 PV controller 将 PVC 和 PV 对象做 bound 之后，这个 PV 就可以被使用了。

https://baijiahao.baidu.com/s?id=1647439213428473676&wfr=spider&for=pc


apple-nfs-storage 改为docker.io/hostpath 

% kubectl delete storageclass apple-nfs-storage
storageclass.storage.k8s.io "apple-nfs-storage" deleted

% kubectl apply -f deployment/statefulset.yaml
statefulset.apps/apple-ss created

% kubectl get storageclass 
NAME                 PROVISIONER          RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
hostpath (default)   docker.io/hostpath   Delete          Immediate           false                  104d

% kubectl describe pod apple-ss-0 
Name:           apple-ss-0
Namespace:      default
Priority:       0
Node:           <none>
Labels:         app=apple
                controller-revision-hash=apple-ss-6bb7949cbd
                statefulset.kubernetes.io/pod-name=apple-ss-0
Annotations:    <none>
Status:         Pending
IP:             
IPs:            <none>
Controlled By:  StatefulSet/apple-ss
Containers:
  apple:
    Image:        apple:5678
    Port:         5678/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:
      /usr/share/apple/html from apple-pvc (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-fbtdn (ro)
Conditions:
  Type           Status
  PodScheduled   False 
Volumes:
  apple-pvc:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  apple-pvc-apple-ss-0
    ReadOnly:   false
  kube-api-access-fbtdn:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason            Age   From               Message
  ----     ------            ----  ----               -------
  Warning  FailedScheduling  60s   default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.
  Warning  FailedScheduling  58s   default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.


  docker.io/hostpath  改为 hostpath



  fuseim.pri/ifs  改为docker.io/hostpath 
parameters:

% kubectl apply -f deployment/StorageClass.yaml
storageclass.storage.k8s.io/apple-nfs-storage created


% kubectl get storageclass                     
NAME                 PROVISIONER          RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
apple-nfs-storage    docker.io/hostpath   Delete          Immediate           false                  4s
hostpath (default)   docker.io/hostpath   Delete          Immediate           false                  104d

% kubectl apply -f deployment/statefulset.yaml
statefulset.apps/apple-ss created

% kubectl get pv |grep apple
pvc-1fb2a0d8-7707-419c-a44a-12287e3a8457   1Ki        RWO            Delete           Bound    default/apple-pvc-apple-ss-1   apple-nfs-storage            40s
pvc-72014d43-7d72-4801-b0af-09727972720f   1Ki        RWO            Delete           Bound    default/apple-pvc-apple-ss-2   apple-nfs-storage            38s
pvc-a7758056-3dec-4e4a-a628-5a7fae5eb6e3   1Gi        RWO            Delete           Bound    default/apple-pvc-apple-ss-0   apple-nfs-storage            90s


 % kubectl get pod  |grep apple-ss
apple-ss-0                                  1/1     Running     0          64s
apple-ss-1                                  1/1     Running     0          64s
apple-ss-2                                  1/1     Running     0          62s

 % kubectl get statefulset
NAME       READY   AGE
apple-ss   3/3     83s

