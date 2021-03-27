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