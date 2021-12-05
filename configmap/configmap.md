给容器内应用程序传递参数的实现方式:
　　1. 将配置文件直接打包到镜像中，但这种方式不推荐使用，因为修改配置不够灵活。
　　2. 通过定义Pod清单时，指定自定义命令行参数，即设定 args:["命令参数"]，这种也
　　    可在启动Pod时，传参来修改Pod的应用程序的配置文件.
　　3. 使用环境变量来给Pod中应用传参修改配置
　　   但要使用此种方式，必须符合以下前提之一:
　　    1) Pod中的应用程序必须是Cloud Native的应用程序，即支持直接通过环境变量来加载配置信息。
　　    2) 通过定义Entrypoint脚本的预处理变量来修改Pod中应用程序的配置文件，这些Entrypoint脚本
　　　    可以使用set，sed，grep等工具来实现修改，但也要确保容器中有这些工具。
　　4.存储卷: 我们可将配置信息直接放到存储卷中，如PV中,Pod启动时，自动挂载存储卷到配置文件目录，
　　    来实现给Pod中应用提供不同的配置。
　　5. configMap 或 secret
　　　　

configMap的主要作用:
　　就是为了让镜像 和 配置文件解耦，以便实现镜像的可移植性和可复用性，因为一个configMap其实就是一系列配置信息的集合，将来可直接注入到Pod中的容器使用，而注入方式有两种，一种将configMap做为存储卷，一种是将configMap通过env中configMapKeyRef注入到容器中； configMap是KeyValve形式来保存数据的，如: name=zhangsan 或 nginx.conf="http{server{...}}" 对于configMap的Value的长度是没有限制的，所以它可以是一整个配置文件的信息。
configMap: 它是K8s中的标准组件,它通过两种方式实现给Pod传递配置参数:
　　A. 将环境变量直接定义在configMap中，当Pod启动时,通过env来引用configMap中定义的环境变量。
　　B. 将一个完整配置文件封装到configMap中,然后通过共享卷的方式挂载到Pod中,实现给应用传参。
secret: 它时一种相对安全的configMap，因为它将configMap通过base64做了编码, 让数据不是明文直接存储在configMap中，起到了一定的保护作用，但对Base64进行反编码，对专业人士来说，没有任何难度，因此它只是相对安全。

对于configMap中第一种，让Pod引用configMap中的环境变量的方式:
　　kubectl explain pods.spec.containers.env 　　 #env也可直接定义传递给Pod中容器的环境变量，这点需要记住。
　　　　env.valueFrom
　　　　   configMapKeyRef： 可用于定义Pod启动时引用的configMapKey是哪个。
　　　　   fieldRef: 也可引用一个字段，为Pod中容器内应用程序的每个环境变量值,如:
　　　　   metadata.name: 引用Pod的名称
　　　　   metadata.namespace: 引用Pod所在的名称空间名
　　　　   metadata.labels: 引用Pod的标签
　　　　   status.hostIP: 引用Pod所在节点的IP
　　　　   status.podIP: 引用Pod的IP
　　　　   resourceFieldRef: 引用一个资源需求 或 资源限制
　　　　   secretKeyRef: 引用一个secretKey来为Pod传参

在定义configMap时，通常仅需要定义它的data 或 binaryData(二进制数据)，它俩都是map[string]类型的，
所以它们的值都是以hash列表方式存储的，即key和value没有直接关系，key就是hash码。

https://www.cnblogs.com/wn1m/p/11288860.html

 (1)通过文件创建configmap
% echo hello > test1.txt
% echo world > test2.txt

% kubectl create configmap my-config --from-file=key1=test1.txt  --from-file=key2=test2.txt
configmap/my-config created

% kubectl describe configmap my-config
Name:         my-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
key1:
----
hello

key2:
----
world

Events:  <none>


通过文件夹创建configmap
$ mkdir config
$ echo hello > config/test1
$ echo world > config/test2

% kubectl create configmap dir-config --from-file=config/
configmap/dir-config created

kubectl describe configmap dir-config
Name:         dir-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
test1:
----
hello

test2:
----
world

Events:  <none>

通过键值对创建configmap
%  kubectl create configmap literal-config --from-literal=key1=hello --from-literal=key2=world 
configmap/literal-config created

%  kubectl describe configmap literal-config
Name:         literal-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
key1:
----
hello
key2:
----
world
Events:  <none>


通过yaml文件创建
kubectl create -f config.yaml
% kubectl create -f config.yaml
Error from server (AlreadyExists): error when creating "config.yaml": configmaps "my-config" already exists
% kubectl create -f config.yaml
configmap/my-yaml-config created

 % kubectl describe configmap my-yaml-config
Name:         my-yaml-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
key1:
----
hello
key2:
----
world
Events:  <none>

ConfigMap的使用
Pod的使用方式：

1. 将ConfigMap中的数据设置为容器的环境变量

2. 将ConfigMap中的数据设置为命令行参数

3. 使用Volume将ConfigMap作为文件或目录挂载

4. 编写代码在 Pod 中运行，使用 Kubernetes API 来读取 ConfigMap


env:
    - name: KEY1
      valueFrom:
       configMapKeyRef:
        name: my-config
        key: key1


将configmap挂载到容器中
% kubectl apply -f configmap-volume.yaml 
pod/apple-configmap-volume created

% kubectl exec -it apple-configmap-volume  sh
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
/ # 

/ # ls /projected-volume
key1  key2

通过volume挂载和环境变量的区别

通过Volume挂载到容器内部时，当该configmap的值发生变化时，容器内部具备自动更新的能力，但是通过环境变量设置到容器内部该值不具备自动更新的能力。

注意：

ConfigMap必须在Pod使用它之前创建

使用envFrom时，将会自动忽略无效的键

Pod只能使用同一个命名空间的ConfigMap

https://blog.csdn.net/skh2015java/article/details/109228836

https://zhuanlan.zhihu.com/p/109111916
https://www.jianshu.com/p/b1d516f02ecd



Configmap热更新原理
Kubernetes中提供configmap，用来管理应用的配置，configmap具备热更新的能力，但只有通过目录挂载的configmap才具备热更新能力，其余通过环境变量，通过subPath挂载的文件都不能动态更新。

https://blog.csdn.net/qingyafan/article/details/102848860


% kubectl get configmap -o wide
NAME                              DATA   AGE
dir-config                        2      13m
ingress-controller-leader-nginx   0      102d
ingress-nginx-controller          0      99d
kube-root-ca.crt                  1      104d
literal-config                    2      12m
my-config                         2      14m
my-yaml-config                    2      10m
redis-config                      1      98d



% kubectl apply -f configmap-volume.yaml 
The Pod "apple-configmap-volume" is invalid: spec: Forbidden: pod updates may not change fields other than `spec.containers[*].image`, `spec.initContainers[*].image`, `spec.activeDeadlineSeconds` or `spec.tolerations` (only additions to existing tolerations)
  core.PodSpec{


% kubectl delete -f configmap-volume.yaml
pod "apple-configmap-volume" deleted


是kubelet在做事
kubelet是每个节点都会安装的主要代理，负责维护节点上的所有容器，并监控容器的健康状况，同步容器需要的数据，数据可能来自配置文件，也可能来自etcd。kubelet有一个启动参数--sync-frequency，控制同步配置的时间间隔，它的默认值是1min，所以更新configmap的内容后，真正容器中的挂载内容变化可能在0~1min之后。修改一下这个值，修改为5s，然后更改configmap的数据，检查热更新延迟时间，都降低到了3s左右，但同时kubelet的资源消耗会上升，尤其运行比较多pod的node上，性能会显著下降。

怎么实现的呢
Kubelet是管理pod生命周期的主要组件，同时它也会维护pod所需的资源，其中之一就是configmap，实现定义在pkg/kubelet/configmap/中，kubelet主要是通过 configmap_manager 来管理每个pod所使用的configmap，configmap_manager有三种：

Simple Manager
TTL Based Manager
Watch Manager
默认使用 Watch Manager。其实Manager管理的主要是缓存中的configmap对象，而kubelet同步的是Pod和缓存中的configmap对象。

% kubectl delete -f configmap-volume.yaml
pod "apple-configmap-volume" deleted


% kubectl apply -f configmap-volume.yaml 
pod/apple-configmap-volume created

% echo "xiazemin" >> config/test1 

% cat config/test1 
hello
xiazemin


 kubectl exec -it apple-configmap-volume  sh

/ # cat /projected-volume/test1 
hello


% kubectl edit configmap dir-config 
configmap/dir-config edited

% kubectl exec -it apple-configmap-volume  sh
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
/ # cat /projected-volume/test1
hello xiazemin1

手动更新文件不行，需要用kubectl edit


