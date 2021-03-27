K8S如何更新Terminating状态的命名空间？
更新时间：2021/01/04 GMT+08:00
查看PDF分享
k8s中namespace有两种常见的状态，即Active和Terminating状态，其中Terminating状态一般会比较少见，当对应的命名空间下还存在运行的资源，但该命名空间被删除时才会出现所谓的Terminating状态，这种情况下只要等待k8s本身将命名空间下的资源回收后，该命名空间将会被系统自动删除。

但是在某些情况下，即使命名空间下没有运行的资源，但依然无法删除Terminating状态的命名空间的情况，它会一直卡在Terminating状态下。

https://www.cnblogs.com/zisefeizhu/p/13786053.html


https://segmentfault.com/a/1190000016924414

% kubectl edit ns fleet-clusters-system

  finalizers:
  - controller.cattle.io/namespace-auth

namespace/fleet-clusters-system edited



https://support.huaweicloud.com/cce_faq/cce_faq_00277.html

# 查看k8s集群中可以使用命名空间隔离的资源
$ kubectl api-resources -o name --verbs=list --namespaced | xargs -n 1 kubectl get --show-kind --ignore-not-found -n local

Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME                                                           AGE
clusteralertgroup.management.cattle.io/cluster-scan-alert      10h
clusteralertgroup.management.cattle.io/etcd-alert              10h
clusteralertgroup.management.cattle.io/event-alert             10h
clusteralertgroup.management.cattle.io/kube-components-alert   10h
clusteralertgroup.management.cattle.io/node-alert              10h
node.management.cattle.io/machine-zh82s   10h


 % kubectl  get ns local -o json > local.json
删除
     "spec": {
        "finalizers": [
            "kubernetes"
        ]
    },

curl -H "Content-Type:application/json" -X PUT --data-binary @local.json http://127.0.0.1:8001/api/v1/namespaces/local/finalize


% kubectl delete ns local --force --grace-period=0

 % kubectl edit ns local
 message: 'Some resources are remaining: clusteralertgroups.management.cattle.io
      has 5 resource instances, nodes.management.cattle.io has 1 resource instances'


kubectl api-resources -o wide --verbs=list --namespaced | xargs -n 1 kubectl get --show-kind --ignore-not-found -n local

# 位于名字空间中的资源
kubectl api-resources --namespaced=true

# 不在名字空间中的资源
kubectl api-resources --namespaced=false


% kubectl api-resources --namespaced=true -o wide
clusteralertgroups                                      management.cattle.io        true         ClusterAlertGroup                         [delete deletecollection get list patch create

nodes                                                   management.cattle.io        true         Node



 % kubectl get ClusterAlertGroup -n local
NAME                    AGE
cluster-scan-alert      10h
etcd-alert              10h
event-alert             10h
kube-components-alert   10h
node-alert              10h

% kubectl get node  -n local -o wide
NAME             STATUS   ROLES    AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE         KERNEL-VERSION      CONTAINER-RUNTIME
docker-desktop   Ready    <none>   10h   v1.19.7   192.168.65.4   <none>        Docker Desktop   4.19.121-linuxkit   docker://20.10.5


https://wener.me/notes/devops/kubernetes/k8s-faq/

删除 rancher 项目空间#
主要难点在于 get all 不会返回所有资源
可尝试 ketall
部分资源需要先 patch 才能删除


for ns in local p-66lfd ; do
  for error in app.project.cattle.io/cluster-alerting app.project.cattle.io/cluster-monitoring app.project.cattle.io/monitoring-operator app.project.cattle.io/project-monitoring clusteralertgroup.management.cattle.io/cluster-scan-alert clusteralertgroup.management.cattle.io/etcd-alert clusteralertgroup.management.cattle.io/event-alert clusteralertgroup.management.cattle.io/kube-components-alert clusteralertgroup.management.cattle.io/node-alert clusterroletemplatebinding.management.cattle.io/creator-cluster-owner clusterroletemplatebinding.management.cattle.io/u-b4qkhsnliz-admin node.management.cattle.io/machine-9sssc node.management.cattle.io/machine-ks6z6 node.management.cattle.io/machine-v4v89 project.management.cattle.io/p-cnj28 project.management.cattle.io/p-mbvfd projectalertgroup.management.cattle.io/projectalert-workload-alert projectalertrule.management.cattle.io/less-than-half-workload-available projectalertrule.management.cattle.io/memory-close-to-resource-limited projectroletemplatebinding.management.cattle.io/app-jdnmz projectroletemplatebinding.management.cattle.io/creator-project-owner projectroletemplatebinding.management.cattle.io/prtb-s6fhc projectroletemplatebinding.management.cattle.io/u-2gacgc4nfu-member projectroletemplatebinding.management.cattle.io/u-efxo6n6ndd-member  ; do
    for resource in `kubectl get -n $ns $error -o name` ; do
      kubectl patch -n $ns $resource -p '{"metadata": {"finalizers": []}}' --type='merge'
    done
  done
done

# 全局资源
for res in $(kubectl api-resources --namespaced=false --api-group management.cattle.io | cut -d ' ' -f 1); do
  echo "=== $res.management.cattle.io ==="
  kubectl get $res.management.cattle.io
done


# namespaced
groups="management.cattle.io catalog.cattle.io project.cattle.io"
for grp in $groups; do
for res in $(kubectl api-resources --namespaced=true --api-group $grp -o name); do
  echo "=== $res ==="
  kubectl get --all-namespaces $res
done
done


# 清除资源
cleargroup(){
  kubectl patch $1 -p '{"metadata":{"finalizers":[]}}' --type=merge $(kubectl get $1 -o jsonpath='{..metadata.name}')
  kubectl delete --all $1
}

cleargroup globalroles.management.cattle.io


 % kubectl delete --all clusteralertgroup.management.cattle.io   -n local --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
No resources found

% kubectl delete --all node.management.cattle.io  -n local --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
node.management.cattle.io "machine-zh82s" force deleted


