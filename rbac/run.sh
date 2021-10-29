% kubectl get pods
NAME                                        READY   STATUS      RESTARTS   AGE
apple-app                                   1/1     Running     11         64d
ingress-nginx-admission-create-vjn92        0/1     Completed   0          61d
ingress-nginx-admission-patch-wlq6p         0/1     Completed   0          61d
ingress-nginx-controller-57648496fc-84wl8   1/1     Running     18         61d
minio-deployment-55bf5bff5d-cvq7v           1/1     Running     13         59d
redis-f9f74787-tq6tw                        1/1     Running     13         60d


% kubectl apply -f rbac/role.yaml 
role.rbac.authorization.k8s.io/pod-reader created


% kubectl apply -f rbac/cluster-role.yaml
clusterrole.rbac.authorization.k8s.io/secret-reader created

 % kubectl apply -f rbac/role-binding.yaml 
rolebinding.rbac.authorization.k8s.io/read-pods created



%kubectl create role 
创建 Role 对象，定义在某一名字空间中的权限。例如:

创建名称为 "pod-reader" 的 Role 对象，允许用户对 Pods 执行 get、watch 和 list 操作：

kubectl create role pod-reader --verb=get --verb=list --verb=watch --resource=pods

%kubectl create clusterrole
创建名称为 "pod-reader" 的 ClusterRole对象，允许用户对 Pods 对象执行 get、watch和list` 操作：

kubectl create clusterrole pod-reader --verb=get,list,watch --resource=pods


%kubectl create rolebinding
在名字空间 "acme" 中，将名为 admin 的 ClusterRole 中的权限授予名称 "bob" 的用户:

kubectl create rolebinding bob-admin-binding --clusterrole=admin --user=bob --namespace=acme




% kubectl get role
NAME                      CREATED AT
ingress-nginx             2021-08-28T12:52:06Z
ingress-nginx-admission   2021-08-28T12:52:06Z
pod-reader                2021-10-29T02:03:05Z



% kubectl get rolebinding
NAME                      ROLE                           AGE
ingress-nginx             Role/ingress-nginx             61d
ingress-nginx-admission   Role/ingress-nginx-admission   61d
read-pods                 Role/pod-reader                7m26s




 % kubectl get clusterrole       
NAME                                                                   CREATED AT
admin                                                                  2021-08-23T03:18:58Z
cluster-admin                                                          2021-08-23T03:18:58Z
edit                                                                   2021-08-23T03:18:58Z
ingress-nginx                                                          2021-08-28T12:52:06Z
ingress-nginx-admission                                                2021-08-28T12:52:06Z
kubeadm:get-nodes                                                      2021-08-23T03:19:00Z
secret-reader                                                          2021-10-29T02:03:17Z
system:aggregate-to-admin                                              2021-08-23T03:18:58Z
system:aggregate-to-edit                                               2021-08-23T03:18:58Z
system:aggregate-to-view                                               2021-08-23T03:18:58Z
system:auth-delegator                                                  2021-08-23T03:18:58Z
system:basic-user                                                      2021-08-23T03:18:58Z
system:certificates.k8s.io:certificatesigningrequests:nodeclient       2021-08-23T03:18:58Z
system:certificates.k8s.io:certificatesigningrequests:selfnodeclient   2021-08-23T03:18:58Z
system:certificates.k8s.io:kube-apiserver-client-approver              2021-08-23T03:18:58Z
system:certificates.k8s.io:kube-apiserver-client-kubelet-approver      2021-08-23T03:18:58Z
system:certificates.k8s.io:kubelet-serving-approver                    2021-08-23T03:18:58Z
system:certificates.k8s.io:legacy-unknown-approver                     2021-08-23T03:18:58Z
system:controller:attachdetach-controller                              2021-08-23T03:18:58Z
system:controller:certificate-controller                               2021-08-23T03:18:58Z
system:controller:clusterrole-aggregation-controller                   2021-08-23T03:18:58Z
system:controller:cronjob-controller                                   2021-08-23T03:18:58Z
system:controller:daemon-set-controller                                2021-08-23T03:18:58Z
system:controller:deployment-controller                                2021-08-23T03:18:58Z
system:controller:disruption-controller                                2021-08-23T03:18:58Z
system:controller:endpoint-controller                                  2021-08-23T03:18:58Z
system:controller:endpointslice-controller                             2021-08-23T03:18:58Z
system:controller:endpointslicemirroring-controller                    2021-08-23T03:18:58Z
system:controller:ephemeral-volume-controller                          2021-08-23T03:18:58Z
system:controller:expand-controller                                    2021-08-23T03:18:58Z
system:controller:generic-garbage-collector                            2021-08-23T03:18:58Z
system:controller:horizontal-pod-autoscaler                            2021-08-23T03:18:58Z
system:controller:job-controller                                       2021-08-23T03:18:58Z
system:controller:namespace-controller                                 2021-08-23T03:18:58Z
system:controller:node-controller                                      2021-08-23T03:18:58Z
system:controller:persistent-volume-binder                             2021-08-23T03:18:58Z
system:controller:pod-garbage-collector                                2021-08-23T03:18:58Z
system:controller:pv-protection-controller                             2021-08-23T03:18:58Z
system:controller:pvc-protection-controller                            2021-08-23T03:18:58Z
system:controller:replicaset-controller                                2021-08-23T03:18:58Z
system:controller:replication-controller                               2021-08-23T03:18:58Z
system:controller:resourcequota-controller                             2021-08-23T03:18:58Z
system:controller:root-ca-cert-publisher                               2021-08-23T03:18:58Z
system:controller:route-controller                                     2021-08-23T03:18:58Z
system:controller:service-account-controller                           2021-08-23T03:18:58Z
system:controller:service-controller                                   2021-08-23T03:18:58Z
system:controller:statefulset-controller                               2021-08-23T03:18:58Z
system:controller:ttl-after-finished-controller                        2021-08-23T03:18:58Z
system:controller:ttl-controller                                       2021-08-23T03:18:58Z
system:coredns                                                         2021-08-23T03:19:00Z
system:discovery                                                       2021-08-23T03:18:58Z
system:heapster                                                        2021-08-23T03:18:58Z
system:kube-aggregator                                                 2021-08-23T03:18:58Z
system:kube-controller-manager                                         2021-08-23T03:18:58Z
system:kube-dns                                                        2021-08-23T03:18:58Z
system:kube-scheduler                                                  2021-08-23T03:18:58Z
system:kubelet-api-admin                                               2021-08-23T03:18:58Z
system:monitoring                                                      2021-08-23T03:18:58Z
system:node                                                            2021-08-23T03:18:58Z
system:node-bootstrapper                                               2021-08-23T03:18:58Z
system:node-problem-detector                                           2021-08-23T03:18:58Z
system:node-proxier                                                    2021-08-23T03:18:58Z
system:persistent-volume-provisioner                                   2021-08-23T03:18:58Z
system:public-info-viewer                                              2021-08-23T03:18:58Z
system:service-account-issuer-discovery                                2021-08-23T03:18:58Z
system:volume-scheduler                                                2021-08-23T03:18:58Z
view                                                                   2021-08-23T03:18:58Z
vpnkit-controller                                                      2021-08-23T03:19:39Z



% kubectl get clusterrolebinding
NAME                                                   ROLE                                                                               AGE
cluster-admin                                          ClusterRole/cluster-admin                                                          66d
docker-for-desktop-binding                             ClusterRole/cluster-admin                                                          66d
ingress-nginx                                          ClusterRole/ingress-nginx                                                          61d
ingress-nginx-admission                                ClusterRole/ingress-nginx-admission                                                61d
kubeadm:get-nodes                                      ClusterRole/kubeadm:get-nodes                                                      66d
kubeadm:kubelet-bootstrap                              ClusterRole/system:node-bootstrapper                                               66d
kubeadm:node-autoapprove-bootstrap                     ClusterRole/system:certificates.k8s.io:certificatesigningrequests:nodeclient       66d
kubeadm:node-autoapprove-certificate-rotation          ClusterRole/system:certificates.k8s.io:certificatesigningrequests:selfnodeclient   66d
kubeadm:node-proxier                                   ClusterRole/system:node-proxier                                                    66d
storage-provisioner                                    ClusterRole/system:persistent-volume-provisioner                                   66d
system:basic-user                                      ClusterRole/system:basic-user                                                      66d
system:controller:attachdetach-controller              ClusterRole/system:controller:attachdetach-controller                              66d
system:controller:certificate-controller               ClusterRole/system:controller:certificate-controller                               66d
system:controller:clusterrole-aggregation-controller   ClusterRole/system:controller:clusterrole-aggregation-controller                   66d
system:controller:cronjob-controller                   ClusterRole/system:controller:cronjob-controller                                   66d
system:controller:daemon-set-controller                ClusterRole/system:controller:daemon-set-controller                                66d
system:controller:deployment-controller                ClusterRole/system:controller:deployment-controller                                66d
system:controller:disruption-controller                ClusterRole/system:controller:disruption-controller                                66d
system:controller:endpoint-controller                  ClusterRole/system:controller:endpoint-controller                                  66d
system:controller:endpointslice-controller             ClusterRole/system:controller:endpointslice-controller                             66d
system:controller:endpointslicemirroring-controller    ClusterRole/system:controller:endpointslicemirroring-controller                    66d
system:controller:ephemeral-volume-controller          ClusterRole/system:controller:ephemeral-volume-controller                          66d
system:controller:expand-controller                    ClusterRole/system:controller:expand-controller                                    66d
system:controller:generic-garbage-collector            ClusterRole/system:controller:generic-garbage-collector                            66d
system:controller:horizontal-pod-autoscaler            ClusterRole/system:controller:horizontal-pod-autoscaler                            66d
system:controller:job-controller                       ClusterRole/system:controller:job-controller                                       66d
system:controller:namespace-controller                 ClusterRole/system:controller:namespace-controller                                 66d
system:controller:node-controller                      ClusterRole/system:controller:node-controller                                      66d
system:controller:persistent-volume-binder             ClusterRole/system:controller:persistent-volume-binder                             66d
system:controller:pod-garbage-collector                ClusterRole/system:controller:pod-garbage-collector                                66d
system:controller:pv-protection-controller             ClusterRole/system:controller:pv-protection-controller                             66d
system:controller:pvc-protection-controller            ClusterRole/system:controller:pvc-protection-controller                            66d
system:controller:replicaset-controller                ClusterRole/system:controller:replicaset-controller                                66d
system:controller:replication-controller               ClusterRole/system:controller:replication-controller                               66d
system:controller:resourcequota-controller             ClusterRole/system:controller:resourcequota-controller                             66d
system:controller:root-ca-cert-publisher               ClusterRole/system:controller:root-ca-cert-publisher                               66d
system:controller:route-controller                     ClusterRole/system:controller:route-controller                                     66d
system:controller:service-account-controller           ClusterRole/system:controller:service-account-controller                           66d
system:controller:service-controller                   ClusterRole/system:controller:service-controller                                   66d
system:controller:statefulset-controller               ClusterRole/system:controller:statefulset-controller                               66d
system:controller:ttl-after-finished-controller        ClusterRole/system:controller:ttl-after-finished-controller                        66d
system:controller:ttl-controller                       ClusterRole/system:controller:ttl-controller                                       66d
system:coredns                                         ClusterRole/system:coredns                                                         66d
system:discovery                                       ClusterRole/system:discovery                                                       66d
system:kube-controller-manager                         ClusterRole/system:kube-controller-manager                                         66d
system:kube-dns                                        ClusterRole/system:kube-dns                                                        66d
system:kube-scheduler                                  ClusterRole/system:kube-scheduler                                                  66d
system:monitoring                                      ClusterRole/system:monitoring                                                      66d
system:node                                            ClusterRole/system:node                                                            66d
system:node-proxier                                    ClusterRole/system:node-proxier                                                    66d
system:public-info-viewer                              ClusterRole/system:public-info-viewer                                              66d
system:service-account-issuer-discovery                ClusterRole/system:service-account-issuer-discovery                                66d
system:volume-scheduler                                ClusterRole/system:volume-scheduler                                                66d
vpnkit-controller                                      ClusterRole/vpnkit-controller                                                      66d




% kubectl get clusterrolebinding docker-for-desktop-binding  
NAME                         ROLE                        AGE
docker-for-desktop-binding   ClusterRole/cluster-admin   66d


% kubectl get clusterrole cluster-admin 
NAME            CREATED AT
cluster-admin   2021-08-23T03:18:58Z


% kubectl get ServiceAccount
NAME                      SECRETS   AGE
default                   1         66d
ingress-nginx             1         61d
ingress-nginx-admission   1         61d

