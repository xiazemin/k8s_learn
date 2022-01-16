 % kubectl create ns bookinfo
namespace/bookinfo created

kubectl apply -f istio-1.12.1/samples/bookinfo/platform/kube/bookinfo.yaml -n bookinfo 

 % kubectl -n bookinfo get deploy
NAME             READY   UP-TO-DATE   AVAILABLE   AGE
details-v1       1/1     1            1           22s
productpage-v1   1/1     1            1           21s
ratings-v1       1/1     1            1           22s
reviews-v1       1/1     1            1           22s
reviews-v2       1/1     1            1           21s
reviews-v3       1/1     1            1           21s

 % kubectl -n bookinfo get pods  
NAME                              READY   STATUS    RESTARTS   AGE
details-v1-79f774bdb9-7zl2h       1/1     Running   0          37s
productpage-v1-6b746f74dc-g7xcx   1/1     Running   0          36s
ratings-v1-b6994bb9-gkfcz         1/1     Running   0          37s
reviews-v1-545db77b95-jj7wb       1/1     Running   0          36s
reviews-v2-7bf8c9648f-pw9dt       1/1     Running   0          36s
reviews-v3-84779c7bbc-9jjtj       1/1     Running   0          36s

注意一定要有单独的namespace 否则相互干扰


 % kubectl -n bookinfo get svc 
NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
details       ClusterIP   10.96.9.127      <none>        9080/TCP   110s
productpage   ClusterIP   10.98.50.191     <none>        9080/TCP   108s
ratings       ClusterIP   10.100.204.206   <none>        9080/TCP   109s
reviews       ClusterIP   10.103.136.15    <none>        9080/TCP   109s

kubectl -n bookinfo  exec "$(kubectl -n bookinfo  get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"
<title>Simple Bookstore App</title>


kubectl -n bookinfo apply -f istio-1.12.1/samples/bookinfo/networking/bookinfo-gateway.yaml
gateway.networking.istio.io/bookinfo-gateway created
virtualservice.networking.istio.io/bookinfo created

% istioctl -n bookinfo analyze
Info [IST0102] (Namespace bookinfo) The namespace is not enabled for Istio injection. Run 'kubectl label namespace bookinfo istio-injection=enabled' to enable it, or 'kubectl label namespace bookinfo istio-injection=disabled' to explicitly mark it as not needing injection.

% kubectl label namespace bookinfo istio-injection=enabled
namespace/bookinfo labeled


% istioctl -n bookinfo analyze                            
Warning [IST0103] (Pod bookinfo/details-v1-79f774bdb9-7zl2h) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.
Warning [IST0103] (Pod bookinfo/productpage-v1-6b746f74dc-g7xcx) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.
Warning [IST0103] (Pod bookinfo/ratings-v1-b6994bb9-gkfcz) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.
Warning [IST0103] (Pod bookinfo/reviews-v1-545db77b95-jj7wb) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.
Warning [IST0103] (Pod bookinfo/reviews-v2-7bf8c9648f-pw9dt) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.
Warning [IST0103] (Pod bookinfo/reviews-v3-84779c7bbc-9jjtj) The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.


% kubectl delete  -f istio-1.12.1/samples/bookinfo/platform/kube/bookinfo.yaml -n bookinfo
service "details" deleted
serviceaccount "bookinfo-details" deleted
deployment.apps "details-v1" deleted
service "ratings" deleted
serviceaccount "bookinfo-ratings" deleted
deployment.apps "ratings-v1" deleted
service "reviews" deleted
serviceaccount "bookinfo-reviews" deleted
deployment.apps "reviews-v1" deleted
deployment.apps "reviews-v2" deleted
deployment.apps "reviews-v3" deleted
service "productpage" deleted
serviceaccount "bookinfo-productpage" deleted
deployment.apps "productpage-v1" deleted


 % kubectl apply -f istio-1.12.1/samples/bookinfo/platform/kube/bookinfo.yaml -n bookinfo 
service/details created
serviceaccount/bookinfo-details created
deployment.apps/details-v1 created
service/ratings created
serviceaccount/bookinfo-ratings created
deployment.apps/ratings-v1 created
service/reviews created
serviceaccount/bookinfo-reviews created
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage created
serviceaccount/bookinfo-productpage created
deployment.apps/productpage-v1 created

 % istioctl analyze -n bookinfo
✔ No validation issues found when analyzing namespace: bookinfo.


kubectl -n bookinfo delete -f istio-1.12.1/samples/bookinfo/networking/bookinfo-gateway.yaml
kubectl -n bookinfo apply -f istio-1.12.1/samples/bookinfo/networking/bookinfo-gateway.yaml

 % kubectl -n bookinfo get pods  
NAME                              READY   STATUS                  RESTARTS   AGE
details-v1-79f774bdb9-qpl74       0/2     Init:CrashLoopBackOff   4          3m7s
productpage-v1-6b746f74dc-p8n2r   0/2     Init:CrashLoopBackOff   4          3m7s
ratings-v1-b6994bb9-4m4bl         0/2     Init:CrashLoopBackOff   5          3m7s
reviews-v1-545db77b95-4hr64       0/2     Init:CrashLoopBackOff   4          3m7s
reviews-v2-7bf8c9648f-gkvfl       0/2     Init:CrashLoopBackOff   5          3m7s
reviews-v3-84779c7bbc-djblk       0/2     Init:Error              5          3m7s

% kubectl -n bookinfo logs details-v1-79f774bdb9-qpl74  istio-init
Error occurred at line: 1
Try `iptables-restore -h' or 'iptables-restore --help' for more information.
2022-01-09T05:00:51.438510Z     error   Failed to execute: iptables-restore --noflush /tmp/iptables-rules-1641704451390415592.txt1406266898, exit status 2

2022-01-09T05:00:51.437858Z     error   Command error output: xtables parameter problem: iptables-restore: unable to initialize table 'nat'

https://blog.csdn.net/weixin_50908696/article/details/122176154


istioctl manifest apply --set components.cni.enabled=true 


https://github.com/istio/istio/issues/23009

% istioctl manifest apply --set components.cni.enabled=true 
This will install the Istio 1.12.1 default profile with ["Istio core" "Istiod" "CNI" "Ingress gateways"] components into the cluster. Proceed? (y/N) y
✔ Istio core installed                                                                                               
✔ Istiod installed                                                                                                   
- Processing resources for CNI, Ingress gateways. Waiting for DaemonSet/istio-system/istio-cni-node, Deployment/is...


 modprobe br_netfilter ; modprobe nf_nat ; modprobe xt_REDIRECT ; modprobe xt_owner; modprobe iptable_nat; modprobe iptable_mangle; modprobe iptable_filter


 需要流量重定向的Pod

Istio CNI插件通过根据以下条件列表进行检查来查找需要重定向流量的Pod：

该pod不在配置的excludeNamespaces列表中的命名空间中
该pod 包含一个名为 istio-proxy 的容器
该pod 中有多个容器
该pod 无 sidecar.istio.io/inject 注解或该注解值为 true
Istio的CNI插件可作为CNI插件链运行。

https://blog.csdn.net/weixin_39858124/article/details/111643884


#当发现通过以上部署pod时，未创建pod且原因 为 sidecar容器自动失败，此时可以采用手动注入方式
#卸载bookinfo示例程序
kubectl delete -f /root/wf/istio/samples/bookinfo/platform/kube/ bookinfo.yaml
#手动注入sidecar容器
istioctl kube-inject -f /root/wf/istio/samples/bookinfo/platform/kube/ bookinfo.yaml |kubectl apply -f -
https://www.cnblogs.com/sandyflower/p/13843122.html

stio CNI 插件会处理 Kubernetes Pod 的创建和删除事件，并作出如下动作：
通过 Istio sidecars 识别 Istio 用户应用 Pods 是否需要流量重定向
对 pod 网络命名空间进行配置，将流量转向 Istio sidecar
https://blog.csdn.net/miss1181248983/article/details/116308254

Istio 会在网格中部署的 Pods 上注入一个 initContainer：istio-init。istio-init 容器会将 Pod 的网络流量劫持到 Istio sidecar 代理上。这需要用户或部署 Pods 的 Service Account 具有足够的部署 NET_ADMIN 容器的 Kubernetes RBAC 权限。Istio 用户权限的提升，对于某些组织的安全政策来说，可能是难以接受的。Istio CNI 插件就是一个能够替代 istio-init 容器来实现相同的网络功能但却不需要 Istio 用户申请额外的 Kubernetes RBAC 授权的方案。

Istio CNI 插件会在 Kubernetes Pod 生命周期的网络设置阶段完成 Istio 网格的 Pod 流量转发设置工作，因此用户在部署 Pods 到 Istio 网格中时，不再需要配置 NET_ADMIN 功能需求了。Istio CNI 插件代替了 istio-init 容器所实现的功能。

前提条件：
安装支持 CNI 的 Kubernetes 集群，并且 kubelet 使用 --network-plugin=cni 参数启用 CNI 插件。
ps aux |grep kubelet
1
Kubernetes 需要启用 ServiceAccount 准入控制器。

下列注入方式都是可以支持 Istio CNI 插件的：

自动 sidecar 注入

使用 istio-sidecar-injector configmap 进行手动注入

执行 istioctl kube-inject 直接使用 configmap：
istioctl kube-inject -f deployment.yaml -o deployment-injected.yaml --injectConfigMapName istio-sidecar-injector

kubectl apply -f deployment-injected.yaml
用 configmap 创建文件，用于执行 istioctl kube-inject：
kubectl get cm -n istio-system istio-sidecar-injector -o=jsonpath='{.data.config}' > inject-config.yaml

istioctl kube-inject -f deployment.yaml -o deployment-injected.yaml --injectConfigFile inject-config.yaml

kubectl apply -f deployment-injected.yaml


https://hub.docker.com/r/calico/cni

https://istio.io/latest/zh/docs/setup/additional-setup/cni/

https://github.com/containernetworking/cni/blob/master/SPEC.md#network-configuration-lists
