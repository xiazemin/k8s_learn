docker pull rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64

https://hub.docker.com/r/rancher/rancher/tags



 % docker images |grep rancher
rancher/rancher                                                   v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64   27d5a1153c11   6 hours ago     1.17GB
registryo.shimo.im/rancher/force_team_share                       latest                                                      d3cad44bd2d7   13 days ago     978MB


docker run -d --restart=unless-stopped -p 80:80 -p 443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher-arm rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64   

 % docker container ls
CONTAINER ID   IMAGE                                                                       COMMAND                  CREATED              STATUS                          PORTS     NAMES
300f374aa735   rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64   "entrypoint.sh"          About a minute ago   Restarting (1) 37 seconds ago             rancher_arm

% docker rm rancher_arm
rancher_arm

docker run -d --restart=unless-stopped -p 8088:80 -p 8443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher_arm rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64  

 % docker exec -it c37ce8264cf3 /bin/bash
Error response from daemon: Container c37ce8264cf377829f3da73adfd48dd3a5013e8b34b8d8b84f81a2f63aa61f08 is restarting, wait until the container is running

https://www.jianshu.com/p/40f4fbe1ec22

 % docker logs c37ce8264cf3
 ERROR: Rancher must be ran with the --privileged flag when running outside of Kubernetes
ERROR: Rancher must be ran with the --privileged flag when running outside of Kubernetes
ERROR: Rancher must be ran with the --privileged flag when running outside of Kubernetes
ERROR: Rancher must be ran with the --privileged flag when running outside of Kubernetes
ERROR: Rancher must be ran with the --privileged flag when running outside of Kubernetes

docker run -d --restart=unless-stopped -p 8088:80 -p 8443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--privileged \
--name rancher_arm rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64  

 % docker logs dc944a185a32
 https://blog.csdn.net/pearl8899/article/details/116277123

 2021/11/18 07:15:57 [ERROR] error syncing 'rancher-partner-charts': handler helm-clusterrepo-ensure: git -C /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974 fetch origin 8d8156d70a6cd31e78c01f127fb2f24363942453 error: exit status 128, detail: error: Server does not allow request for unadvertised object 8d8156d70a6cd31e78c01f127fb2f24363942453
, requeuing
2021/11/18 07:16:00 [ERROR] error syncing 'rancher-partner-charts': handler helm-clusterrepo-ensure: git -C /var/lib/rancher-data/local-catalogs/v2/rancher-partner-charts/8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974 fetch origin 8d8156d70a6cd31e78c01f127fb2f24363942453 error: exit status 128, detail: error: Server does not allow request for unadvertised object 8d8156d70a6cd31e78c01f127fb2f24363942453
, requeuing

2021/11/18 07:17:09 [ERROR] error syncing 'mutating-webhook-configuration': handler need-a-cert: services "webhook-service" not found, requeuing

2021/11/18 07:20:39 [ERROR] error syncing 'mutating-webhook-configuration': handler need-a-cert: services "webhook-service" not found, requeuing
2021/11/18 07:20:39 [ERROR] error syncing 'validating-webhook-configuration': handler need-a-cert: services "webhook-service" not found, requeuing
2021/11/18 07:21:09 [ERROR] error syncing 'mutating-webhook-configuration': handler need-a-cert: services "webhook-service" not found, requeuing
2021/11/18 07:21:09 [ERROR] error syncing 'validating-webhook-configuration': handler need-a-cert: services "webhook-service" not found, requeuing


https://www.cnblogs.com/oolo/p/11778727.html

https://rancher.com/docs/rancher/v2.6/en/installation/resources/advanced/arm64-platform/

https://segmentfault.com/a/1190000021547208

https://github.com/jetstack/cert-manager/issues/2602

https://www.cnblogs.com/kevingrace/p/14617757.html

 % docker inspect --format '{{ .NetworkSettings.IPAddress }}' rancher_arm
172.17.0.3

https://www.cnblogs.com/duwamish/p/10944914.html



--cacert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt \
--key=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key \
--cert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt 


https://www.xtplayer.cn/rancher/import-k8s-cluster-update-ca/#%E9%87%8D%E5%BB%BA-Agent-Pod

 % kubectl get ns
^@^@^@^@Error from server (InternalError): an error on the server ("") has prevented the request from succeeding

export HTTPS_PROXY=""


docker run -d --restart=unless-stopped -p 7071:80 -p 7443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--privileged \
--name rancher_arm rancher/rancher:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64  

https://zhuanlan.zhihu.com/p/360824239

https://www.bookstack.cn/read/rancher-v2.x/faf9b1976947d11f.md
https://github.com/rancher/webhook

https://www.rancher.cn/products/rke/

RKE是一个用Golang编写的Kubernetes安装程序，极为简单易用，用户不再需要做大量的准备工作，即可拥有闪电般快速的Kubernetes安装部署体验。
https://segmentfault.com/a/1190000012288926

 While Rancher uses RKE under the hood, RKE can also be used to create Kubernetes clusters as a standalone piece of software.
 https://www.suse.com/c/rancher_blog/rancher-vs-rke-what-is-the-difference/


2021/11/19 02:19:00 [ERROR] failed to parse constraint for kubeversion <1.21: Could not parse Range "<1.21": Could not parse version "1.21" in "<1.21": No Major.Minor.Patch elements found
 https://github.com/rancher/rancher/issues/35573


 lsof发现端口冲突了，修改端口
 % curl  0.0.0.0:7071
<a href="https://0.0.0.0:7443/">Found</a>.

% curl  http://0.0.0.0:7071
<a href="https://0.0.0.0:7443/">Found</a>.

% curl  http://0.0.0.0:7443
Client sent an HTTP request to an HTTPS server.

% curl  https://0.0.0.0:7443
curl: (60) SSL certificate problem: unable to get local issuer certificate

https://cheapsslsecurity.com/blog/ssl-certificate-problem-unable-to-get-local-issuer-certificate/

浏览器信任就可以了
https://itnext.io/rancher-2-4-kubernetes-on-your-macos-laptop-with-docker-k3d-b578b1c7568b

https://127.0.0.1:7443/dashboard/auth/login

https://blog.csdn.net/chucangguan9546/article/details/101019831


https://www.bookstack.cn/read/rancher-v2.x/eb4dcbcdb337125c.md
 % docker exec -it f2149a1f379d reset-password
New password for default admin user (user-z64h7):
L9QkCDe6eJNrsoTl7o8P

admin
L9QkCDe6eJNrsoTl7o8P

kubectl apply -f https://172.17.0.3:8443/v3/import/htfjdmgnhjlzt69fzdknvjjwqrj62kvr72cz2b4fv4qwqp2r8nbcw9_c-m-dgdvgfs7.yaml


Unable to connect to the server: dial tcp 172.17.0.3:8443: i/o timeout




kubectl apply -f https://127.0.0.1:7443/v3/import/htfjdmgnhjlzt69fzdknvjjwqrj62kvr72cz2b4fv4qwqp2r8nbcw9_c-m-dgdvgfs7.yaml
Unable to connect to the server: x509: certificate signed by unknown authority

curl --insecure -sfL https://172.17.0.3:8443/v3/import/htfjdmgnhjlzt69fzdknvjjwqrj62kvr72cz2b4fv4qwqp2r8nbcw9_c-m-dgdvgfs7.yaml | kubectl apply -f -


curl --insecure -sfL https://127.0.0.1:7443/v3/import/htfjdmgnhjlzt69fzdknvjjwqrj62kvr72cz2b4fv4qwqp2r8nbcw9_c-m-dgdvgfs7.yaml | kubectl apply -f -

clusterrole.rbac.authorization.k8s.io/proxy-clusterrole-kubeapiserver created
clusterrolebinding.rbac.authorization.k8s.io/proxy-role-binding-kubernetes-master created
namespace/cattle-system created
serviceaccount/cattle created
clusterrolebinding.rbac.authorization.k8s.io/cattle-admin-binding created
secret/cattle-credentials-674e7b5 created
clusterrole.rbac.authorization.k8s.io/cattle-admin created
deployment.apps/cattle-cluster-agent created
service/cattle-cluster-agent created



 % kubectl get ns
NAME              STATUS   AGE
cattle-system     Active   18m
default           Active   88d
kube-node-lease   Active   88d
kube-public       Active   88d
kube-system       Active   88d

https://www.cnblogs.com/kevingrace/p/14617757.html

% kubectl get pods -n cattle-system
NAME                                   READY   STATUS             RESTARTS   AGE
cattle-cluster-agent-86d999777-vjwm2   0/1     ImagePullBackOff   0          18m



% kubectl get pods -n cattle-system
NAME                                   READY   STATUS             RESTARTS   AGE
cattle-cluster-agent-86d999777-vjwm2   0/1     ImagePullBackOff   0          18m

% kubectl describe pod cattle-cluster-agent-86d999777-vjwm2
Error from server (NotFound): pods "cattle-cluster-agent-86d999777-vjwm2" not found

% kubectl get svc -n cattle-system
NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
cattle-cluster-agent   ClusterIP   10.108.3.16   <none>        80/TCP,443/TCP   20m

% kubectl describe svc cattle-cluster-agent
Error from server (NotFound): services "cattle-cluster-agent" not found

% kubectl get deployment -n cattle-system
NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
cattle-cluster-agent   0/1     1            0           20m
% kubectl describe deployment cattle-cluster-agent
Error from server (NotFound): deployments.apps "cattle-cluster-agent" not found

% kubectl get daemonset -n cattle-system
No resources found in cattle-system namespace.


https://forums.rancher.com/t/register-an-imported-cluster-using-a-private-registry/17805



% kubectl  -n cattle-system describe deployment.apps/cattle-cluster-agent
Name:                   cattle-cluster-agent
Namespace:              cattle-system
CreationTimestamp:      Fri, 19 Nov 2021 11:19:37 +0800
Labels:                 <none>
Annotations:            deployment.kubernetes.io/revision: 1
                        management.cattle.io/scale-available: 2
Selector:               app=cattle-cluster-agent
Replicas:               1 desired | 1 updated | 1 total | 0 available | 1 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  0 max unavailable, 1 max surge
Pod Template:
  Labels:           app=cattle-cluster-agent
  Service Account:  cattle
  Containers:
   cluster-register:
    Image:      rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-head
    Port:       <none>
    Host Port:  <none>
    Environment:
      CATTLE_IS_RKE:             false
      CATTLE_SERVER:             https://172.17.0.3:8443
      CATTLE_CA_CHECKSUM:        a2f5ce583180431175ce0a08b283ef216eee183de4edee930ce414ae719b039e
      CATTLE_CLUSTER:            true
      CATTLE_K8S_MANAGED:        true
      CATTLE_CLUSTER_REGISTRY:
      CATTLE_SERVER_VERSION:     v2.6-67229511aeacd882ccb97e185826d664951c795c-head
      CATTLE_INSTALL_UUID:       2b4a4913-a091-48e9-9138-95501d8bf1dc
      CATTLE_INGRESS_IP_DOMAIN:  sslip.io
    Mounts:
      /cattle-credentials from cattle-credentials (ro)
  Volumes:
   cattle-credentials:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  cattle-credentials-674e7b5
    Optional:    false
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      False   MinimumReplicasUnavailable
  Progressing    False   ProgressDeadlineExceeded
OldReplicaSets:  <none>
NewReplicaSet:   cattle-cluster-agent-86d999777 (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  27m   deployment-controller  Scaled up replica set cattle-cluster-agent-86d999777 to 1




   % docker pull rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-head
^@^@Error response from daemon: manifest for rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-head not found: manifest unknown: manifest unknown


https://hub.docker.com/r/rancher/rancher-agent
docker pull rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64

"https://registry-1.docker.io/v2/": net/http: TLS handshake timeout


docker tag  rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-linux-arm64  rancher/rancher-agent:v2.6-67229511aeacd882ccb97e185826d664951c795c-head

 % kubectl get pods -n cattle-system
NAME                                   READY   STATUS   RESTARTS   AGE
cattle-cluster-agent-86d999777-vjwm2   0/1     Error    2          55m



 % kubectl logs cattle-cluster-agent-86d999777-vjwm2  -n cattle-system
INFO: Environment: CATTLE_ADDRESS=10.1.0.204 CATTLE_CA_CHECKSUM=a2f5ce583180431175ce0a08b283ef216eee183de4edee930ce414ae719b039e CATTLE_CLUSTER=true CATTLE_CLUSTER_AGENT_PORT=tcp://10.108.3.16:80 CATTLE_CLUSTER_AGENT_PORT_443_TCP=tcp://10.108.3.16:443 CATTLE_CLUSTER_AGENT_PORT_443_TCP_ADDR=10.108.3.16 CATTLE_CLUSTER_AGENT_PORT_443_TCP_PORT=443 CATTLE_CLUSTER_AGENT_PORT_443_TCP_PROTO=tcp CATTLE_CLUSTER_AGENT_PORT_80_TCP=tcp://10.108.3.16:80 CATTLE_CLUSTER_AGENT_PORT_80_TCP_ADDR=10.108.3.16 CATTLE_CLUSTER_AGENT_PORT_80_TCP_PORT=80 CATTLE_CLUSTER_AGENT_PORT_80_TCP_PROTO=tcp CATTLE_CLUSTER_AGENT_SERVICE_HOST=10.108.3.16 CATTLE_CLUSTER_AGENT_SERVICE_PORT=80 CATTLE_CLUSTER_AGENT_SERVICE_PORT_HTTP=80 CATTLE_CLUSTER_AGENT_SERVICE_PORT_HTTPS_INTERNAL=443 CATTLE_CLUSTER_REGISTRY= CATTLE_INGRESS_IP_DOMAIN=sslip.io CATTLE_INSTALL_UUID=2b4a4913-a091-48e9-9138-95501d8bf1dc CATTLE_INTERNAL_ADDRESS= CATTLE_IS_RKE=false CATTLE_K8S_MANAGED=true CATTLE_NODE_NAME=cattle-cluster-agent-86d999777-vjwm2 CATTLE_SERVER=https://172.17.0.3:8443 CATTLE_SERVER_VERSION=v2.6-67229511aeacd882ccb97e185826d664951c795c-head
INFO: Using resolv.conf: nameserver 10.96.0.10 search cattle-system.svc.cluster.local svc.cluster.local cluster.local options ndots:5
ERROR: https://172.17.0.3:8443/ping is not accessible (Failed to connect to 172.17.0.3 port 8443: Connection refused)


https://blog.csdn.net/fuck487/article/details/103161493



% export HTTPS_PROXY="" 
% kubectl get deployment -n cattle-system -o yaml >> cattle-system.yaml.bak 

% cp cattle-system.yaml.bak  cattle-system.yaml
修改
          - name: CATTLE_SERVER
            value: https://172.17.0.3:8443

          - name: CATTLE_SERVER
            value: https://127.0.0.1:7443

% kubectl apply -f cattle-system.yaml 

INFO: Using resolv.conf: nameserver 10.96.0.10 search cattle-system.svc.cluster.local svc.cluster.local cluster.local options ndots:5
ERROR: https://127.0.0.1:7443/ping is not accessible (Failed to connect to 127.0.0.1 port 7443: Connection refused)


全局设置不对
https://127.0.0.1:7443/dashboard/c/_/settings/management.cattle.io.setting

server-url Modified
Default Explorer install url. Must be HTTPS. All nodes in your cluster must be able to reach this.
https://172.17.0.3:8443


https://127.0.0.1:7443

删除，重新导入

kubectl apply -f https://127.0.0.1:7443/v3/import/gx4jwd5zgvlgv6mgpzzwnrf8gnq527blx2g98l55r6qzkmtd77mtgn_c-m-gflstz86.yaml

curl --insecure -sfL https://127.0.0.1:7443/v3/import/gx4jwd5zgvlgv6mgpzzwnrf8gnq527blx2g98l55r6qzkmtd77mtgn_c-m-gflstz86.yaml | kubectl apply -f -
clusterrole.rbac.authorization.k8s.io/proxy-clusterrole-kubeapiserver unchanged
clusterrolebinding.rbac.authorization.k8s.io/proxy-role-binding-kubernetes-master unchanged
namespace/cattle-system unchanged
serviceaccount/cattle unchanged
clusterrolebinding.rbac.authorization.k8s.io/cattle-admin-binding unchanged
secret/cattle-credentials-d353d74 created
clusterrole.rbac.authorization.k8s.io/cattle-admin unchanged
deployment.apps/cattle-cluster-agent configured
service/cattle-cluster-agent unchanged


清理干净，重启rancher
% curl --insecure -sfL https://127.0.0.1:7443/v3/import/fdcbqgtddf9gmdc5vhxwnwnlfzs2r5xz2dj7qhvb94j7mkffgcvbxk_c-m-m58c6wbr.yaml | kubectl apply -f -
clusterrole.rbac.authorization.k8s.io/proxy-clusterrole-kubeapiserver unchanged
clusterrolebinding.rbac.authorization.k8s.io/proxy-role-binding-kubernetes-master unchanged
namespace/cattle-system created
serviceaccount/cattle created
clusterrolebinding.rbac.authorization.k8s.io/cattle-admin-binding unchanged
secret/cattle-credentials-1db23e4 created
clusterrole.rbac.authorization.k8s.io/cattle-admin unchanged
deployment.apps/cattle-cluster-agent created
service/cattle-cluster-agent created

 % kubectl get pod -n cattle-system
NAME                                    READY   STATUS             RESTARTS   AGE
cattle-cluster-agent-7c57dd68cc-8vfv2   0/1     CrashLoopBackOff   2          46s

 % kubectl logs cattle-cluster-agent-7c57dd68cc-8vfv2  -n cattle-system

 https://blog.csdn.net/fuck487/article/details/103161493



 curl --insecure -sfL https://127.0.0.1:7443/v3/import/fdcbqgtddf9gmdc5vhxwnwnlfzs2r5xz2dj7qhvb94j7mkffgcvbxk_c-m-m58c6wbr.yaml >> rancher.yaml


  % kubectl get svc cattle-cluster-agent -n cattle-system
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
cattle-cluster-agent   ClusterIP   10.111.195.15   <none>        80/TCP,443/TCP   37m


 % kubectl -n cattle-system expose svc cattle-cluster-agent --port=7443 --target-port=443
Error from server (AlreadyExists): services "cattle-cluster-agent" already exists



把 type： ClusterIP 
改成 
 type: NodePort
 并且修改 nodePort

 % kubectl apply -f  cattle-cluster-agent.yaml
The Service "cattle-cluster-agent" is invalid: spec.ports[0].nodePort: Invalid value: 7071: provided port is not in the valid range. The range of valid ports is 30000-32767

nodePort: 30443
nodePort: 30071


修改全局设置
https://127.0.0.1:30443

 % kubectl apply -f  cattle-system.yaml
Error from server (Conflict): Operation cannot be fulfilled on deployments.apps "cattle-cluster-agent": the object has been modified; please apply your changes to the latest version and try again


