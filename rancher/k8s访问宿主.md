https://www.cnblogs.com/yidiandhappy/p/13917680.html

https://www.cnblogs.com/ssss429170331/p/7686877.html


Rancher 纳管的集群部署有两种不同的 Agent：

cattle-cluster-agent
cattle-node-agent

http://rancher2.docs.rancher.cn/docs/rancher2/cluster-provisioning/rke-clusters/rancher-agents/_index/

cattle-cluster-agent用于连接集群的Rancher 部署的 Kubernetes 集群的 Kubernetes API。cattle-cluster-agent通过 Deployment 的方式部署。


http://rancher2.docs.rancher.cn/docs/rancher2/overview/architecture/_index/

 Rancher Server 管控 Rancher 部署的 Kubernetes 集群（RKE 集群）和托管的 Kubernetes 集群的（EKS）集群的流程。以用户下发指令为例，指令的流动路径如下：

首先，用户通过 Rancher UI（即 Rancher 控制台） Rancher 命令行工具（Rancher CLI）输入指令；直接调用 Rancher API 接口也可以达到相同的效果。
用户通过 Rancher 的代理认证后，指令会进一步下发到 Rancher Server 。
与此同时，Rancher Server 也会执行容灾备份，将数据备份到 etcd 节点。
然后 Rancher Server 把指令传递给集群控制器。集群控制器把指令传递到下游集群的 Agent，最终通过 Agent 把指令下发到指定的集群中。


在执行集群操作时，cattle-node-agent用于和Rancher 部署的 Kubernetes 集群中的节点进行交互。集群操作的示例包括升级 Kubernetes 版本、创建 etcd 快照和恢复 etcd 快照。cattle-node-agent通过 DaemonSet 的方式部署，以确保其在每个节点上运行。当cattle-cluster-agent不可用时，cattle-node-agent 将作为备选方案连接到Rancher 部署的 Kubernetes 集群中的 Kubernetes API。

https://blog.csdn.net/weixin_39841589/article/details/107151105
使用docker for mac搭建jenkins+gitlab，启动容器采用将端口映射到宿主机端口的方式，在指定git repository时，提示cannot assign requested address
原因：
docker官网说Docker for Mac是没有docker0网桥的，就算启动容器时为容器指定–network=host也没有用
解决方案：
以docker.for.mac.host.internal为ip访问

https://blog.csdn.net/weixin_33860528/article/details/91461648


设置全局地址：https://docker.for.mac.host.internal:7443


rm -rf ~/docker_volume/rancher_home/rancher

docker run -d --restart=unless-stopped -p 7071:80 -p 7443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--privileged \
--name rancher_arm rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64  


 % docker container ls |grep rancher
b5d69e2b9f70   rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64   "entrypoint.sh"          47 seconds ago   Up 38 seconds   0.0.0.0:7071->80/tcp, 0.0.0.0:7443->443/tcp   rancher_arm


docker exec -it b5d69e2b9f70 reset-password

% docker exec -it b5d69e2b9f70 reset-password
New password for default admin user (user-rwcln):
gTDBa40yF-ns4VSFRmq1


kubectl get secret --namespace cattle-system bootstrap-secret -o go-template='{{.data.bootstrapPassword|base64decode}}{{"\n"}}'


https://docker.for.mac.host.internal:7443


% curl --insecure -sfL https://docker.for.mac.host.internal:7443/v3/import/n9s6rlt79759vw6nf5djv7r9zpsdnvtn8sxpzfzx4dh55cqgkjclc8_c-m-r7xmtmss.yaml | kubectl apply -f -
error: no objects passed to apply


curl --insecure -sfL https://localhost:7443/v3/import/n9s6rlt79759vw6nf5djv7r9zpsdnvtn8sxpzfzx4dh55cqgkjclc8_c-m-r7xmtmss.yaml | kubectl apply -f -


 % curl --insecure -sfL https://localhost:7443/v3/import/n9s6rlt79759vw6nf5djv7r9zpsdnvtn8sxpzfzx4dh55cqgkjclc8_c-m-r7xmtmss.yaml | kubectl apply -f -
clusterrole.rbac.authorization.k8s.io/proxy-clusterrole-kubeapiserver unchanged
clusterrolebinding.rbac.authorization.k8s.io/proxy-role-binding-kubernetes-master unchanged
namespace/cattle-system created
serviceaccount/cattle created
clusterrolebinding.rbac.authorization.k8s.io/cattle-admin-binding unchanged
secret/cattle-credentials-b361dd4 created
clusterrole.rbac.authorization.k8s.io/cattle-admin unchanged
deployment.apps/cattle-cluster-agent created
service/cattle-cluster-agent created



 % kubectl -n cattle-system get pod
NAME                                    READY   STATUS        RESTARTS   AGE
cattle-cluster-agent-7f579477df-jrnr6   1/1     Running       0          27s
cattle-cluster-agent-8d667b659-nhh9f    1/1     Terminating   0          32s

 % kubectl -n cattle-system get pod
NAME                                    READY   STATUS    RESTARTS   AGE
cattle-cluster-agent-7f579477df-jrnr6   1/1     Running   0          50s


https://127.0.0.1:7443/dashboard/c/c-m-r7xmtmss/explorer

https://blog.csdn.net/one_chao/article/details/98673010?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link

https://www.jianshu.com/p/7f96e322cb79?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation

docker内部实际上实现了一个虚拟网桥docker0，需要通过网桥找到外部宿主机的在网桥的虚拟地址，也就是docker.for.mac.host.internal，就可以实现容器内访问外部宿主机。感兴趣的话可以了解下Docker的网络原理、计算机网络原理和docker compose等内容。

https://stackoverflow.com/questions/22944631/how-to-get-the-ip-address-of-the-docker-host-from-inside-a-docker-container

在使用Docker时，要注意平台之间实现的差异性，如Docker For Mac的实现和标准Docker规范有区别，Docker For Mac的Docker Daemon是运行于虚拟机(xhyve)中的, 而不是像Linux上那样作为进程运行于宿主机，因此Docker For Mac没有docker0网桥，不能实现host网络模式，host模式会使Container复用Daemon的网络栈(在xhyve虚拟机中)，而不是与Host主机网络栈，这样虽然其它容器仍然可通过xhyve网络栈进行交互，但却不是用的Host上的端口(在Host上无法访问)。bridge网络模式 -p 参数不受此影响，它能正常打开Host上的端口并映射到Container的对应Port。文档在这一点上并没有充分说明，容易踩坑。

https://www.cnblogs.com/yuyutianxia/p/8073411.html

https://www.jb51.net/article/149196.htm

2

From docker 18.03 onwards official recommendation is to connect to the special DNS name host.docker.internal, which resolves to the internal IP address used by the host, this is for development purpose

https://stackoverflow.com/questions/38504890/docker-for-mac-1-12-0-how-to-connect-to-host-from-container


I want to connect from a container to a service on the host
The host has a changing IP address (or none if you have no network access). We recommend that you connect to the special DNS name host.docker.internal which resolves to the internal IP address used by the host. This is for development purpose and will not work in a production environment outside of Docker Desktop for Mac.

You can also reach the gateway using gateway.docker.internal.

If you have installed Python on your machine, use the following instructions as an example to connect from a container to a service on the host:

https://docs.docker.com/desktop/mac/networking/

https://www.jianshu.com/p/6f20023e4586

The way docker does DNS by default is fundamentally broken. It should always use dnsmasq and the host's DNS configuration by default (and not use 8.8.8.8). Each container should (by default) resolve DNS by querying the host, which should forward the request to its own resolver, and provide resolution for all of .docker.internal, including host.docker.internal (which should be configurable to allow containers named host). All other accessible containers should also resolve in .docker.internal

https://github.com/moby/libnetwork/pull/2348

Docker 在 Mac 中的实现是通过 Hypervisor 创建一个轻量级的虚拟机，然后 将 docker 放入到虚拟机中实现。Mac OS 宿主机和 Docker 中的容器通过 /var/run/docker.sock 这种 socket 文件来通信，所以在 Mac OS 中 ping 容器的 IP，在容器中 ping 宿主机的 IP 就不通。


容器内访问宿主机，在 Docker 18.03 过后推荐使用 特殊的 DNS 记录 host.docker.internal 访问宿主机。但是注意，这个只是在 Docker Desktop for Mac 中作为开发时有效。 网关的 DNS 记录: gateway.docker.internal。

宿主机访问容器，使用本机 localhost 端口映射功能，使用 –publish（单个端口）, -p（单个端口）, -P（所有端口） 将本机的端口和容器的端口映射。

宿主机访问容器，使用 -p 参数映射端口。容器访问宿主机，可以在宿主机使用下面的命令获取 宿主机的 ip 地址：

ps -ef | grep -i docker | grep -i  "\-\-host\-ip" |awk -F "host-ip" '{print $2}' | awk -F '--lowest-ip' '{print $1}'


https://yuanmomo.net/2019/06/13/docker-network/