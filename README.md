https://github.com/AliyunContainerService/k8s-for-docker-desktop


 % kubectl get nodes
NAME             STATUS   ROLES    AGE   VERSION
docker-desktop   Ready    master   11h   v1.19.7

 % kubectl config use-context docker-desktop
Switched to context "docker-desktop".

 % kubectl cluster-info
Kubernetes master is running at https://kubernetes.docker.internal:6443
KubeDNS is running at https://kubernetes.docker.internal:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

%  kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.4/aio/deploy/recommended.yaml
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created

https://kubernetes.io/zh/docs/tasks/extend-kubernetes/http-proxy-access-api/
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard

% kubectl get pods -n kubernetes-dashboard
NAME                                         READY   STATUS             RESTARTS   AGE
dashboard-metrics-scraper-7b59f7d4df-b6qz9   1/1     Running            0          30m
kubernetes-dashboard-665f4c5ff-tp5sm         0/1     CrashLoopBackOff   10         30m

CrashLoopBackOff 告诉我们，Kubernetes 正在尽力启动这个 Pod，但是一个或多个容器已经挂了，或者正被删除。
https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/troubleshooting-kubeadm/

deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created

% kubectl get pod -n kubernetes-dashboard
NAME                                         READY   STATUS             RESTARTS   AGE
dashboard-metrics-scraper-7b59f7d4df-b6qz9   1/1     Running            0          6m54s
kubernetes-dashboard-665f4c5ff-tp5sm         0/1     CrashLoopBackOff   5          6m54s

 % kubectl proxy
Starting to serve on 127.0.0.1:8001
% TOKEN=$(kubectl -n kube-system describe secret default| awk '$1=="token:"{print $2}')
kubectl config set-credentials docker-for-desktop --token="${TOKEN}"
echo $TOKEN
User "docker-for-desktop" set.
eyJhbGc

no endpoints available for service "https:kubernetes-dashboard:"
https://github.com/kubernetes/dashboard/issues/3322
 %  kubectl -n kubernetes-dashboard get svc -o wide
NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE   SELECTOR
dashboard-metrics-scraper   ClusterIP   10.108.121.130   <none>        8000/TCP   14m   k8s-app=dashboard-metrics-scraper
kubernetes-dashboard        ClusterIP   10.105.145.27    <none>        443/TCP    14m   k8s-app=kubernetes-dashboard

 % kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   11h
  
 % kubectl get namespace
NAME                   STATUS   AGE
default                Active   12h
kube-node-lease        Active   12h
kube-public            Active   12h
kube-system            Active   12h
kubernetes-dashboard   Active   22m

% kubectl get pods -n kube-public
No resources found in kube-public namespace.

% kubectl logs kubernetes-dashboard-665f4c5ff-tp5sm --namespace=kubernetes-dashboard
 2021/03/22 02:41:24 Initializing csrf token from kubernetes-dashboard-csrf secret
panic: an error on the server ("") has prevented the request from succeeding (get secrets kubernetes-dashboard-csrf)

goroutine 1 [running]:
github.com/kubernetes/dashboard/src/app/backend/client/csrf.(*csrfTokenManager).init(0x4000395040)
	/home/runner/work/dashboard/dashboard/src/app/backend/client/csrf/manager.go:41 +0x350
 
https://docs.aws.amazon.com/eks/latest/userguide/dashboard-tutorial.html
https://github.com/kubernetes/kubernetes/issues/29540

kubectl 最简单的登陆方式是在rancher里copy kubeconfig 文件，当然也可以手动配置，更灵活，可以支持多个集群
https://kubernetes.io/zh/docs/tasks/access-application-cluster/configure-access-multiple-clusters/
https://kubernetes.io/zh/docs/concepts/configuration/organize-cluster-access-kubeconfig/

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-set-credentials-em-

kubectl 可以根据不同context 链接 不同 服务，注意配置认证的时候，token方式不能配置用户名，否则报
error: more than one authentication method found for xx; found [token basicAuth], only one is allowed
error: you cannot specify more than one authentication method at the same time: --token, --username/--password

 kubectl config set-credentials xx --token=yy


https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.4/aio/deploy/recommended.yaml

% kubectl get pods -n kubernetes-dashboard
NAME READY STATUS RESTARTS AGE
dashboard-metrics-scraper-7b59f7d4df-b6qz9 1/1 Running 0 59m
kubernetes-dashboard-665f4c5ff-tp5sm 0/1 CrashLoopBackOff 15 59m

% kubectl logs kubernetes-dashboard-665f4c5ff-tp5sm --namespace=kubernetes-dashboard
2021/03/22 03:03:12 Starting overwatch
2021/03/22 03:03:12 Using namespace: kubernetes-dashboard
2021/03/22 03:03:12 Using in-cluster config to connect to apiserver
2021/03/22 03:03:12 Using secret token for csrf signing
2021/03/22 03:03:12 Initializing csrf token from kubernetes-dashboard-csrf secret
panic: an error on the server ("") has prevented the request from succeeding (get secrets kubernetes-dashboard-csrf)

goroutine 1 [running]:
github.com/kubernetes/dashboard/src/app/backend/client/csrf.(*csrfTokenManager).init(0x400038c800)
/home/runner/work/dashboard/dashboard/src/app/backend/client/csrf/manager.go:41 +0x350
github.com/kubernetes/dashboard/src/app/backend/client/csrf.NewCsrfTokenManager(...)
/home/runner/work/dashboard/dashboard/src/app/backend/client/csrf/manager.go:66
github.com/kubernetes/dashboard/src/app/backend/client.(*clientManager).initCSRFKey(0x4000201680)
/home/runner/work/dashboard/dashboard/src/app/backend/client/manager.go:502 +0xb0
github.com/kubernetes/dashboard/src/app/backend/client.(*clientManager).init(0x4000201680)
/home/runner/work/dashboard/dashboard/src/app/backend/client/manager.go:470 +0x40
github.com/kubernetes/dashboard/src/app/backend/client.NewClientManager(...)
/home/runner/work/dashboard/dashboard/src/app/backend/client/manager.go:551
main.main()
/home/runner/work/dashboard/dashboard/src/app/backend/dashboard.go:105 +0x1dc


https://blog.csdn.net/dengkuo19860718/article/details/101447560
% kubectl get deployment -n  kubernetes-dashboard
NAME                        READY   UP-TO-DATE   AVAILABLE   AGE
dashboard-metrics-scraper   1/1     1            1           4h46m
kubernetes-dashboard        0/1     1            0           4h46m

 %  kubectl rollout history deployment kubernetes-dashboard  -n  kubernetes-dashboard
deployment.apps/kubernetes-dashboard
REVISION  CHANGE-CAUSE
1         <none>


https://stackoverflow.com/questions/64533410/kubernetes-dashboard-an-error-on-the-server-unknown-has-prevented-the-reques%E3%80%81

把ClusterRoleBinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: kubernetes-dashboard

改成
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin

The ClusterRoleBinding "kubernetes-dashboard" is invalid: roleRef: Invalid value: rbac.RoleRef{APIGroup:"rbac.authorization.k8s.io", Kind:"ClusterRole", Name:"cluster-admin"}: cannot change roleRef

增加
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
   name: kubernetes-dashboard
   labels:
       k8s-app: kubernetes-dashboard
roleRef:
   apiGroup: rbac.authorization.k8s.io
   kind: ClusterRole
   name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile

---
  
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1alpha1
metadata:
  name: admin
rules:
  - apiGroups: ["*"]
    resources: ["*"]
    verbs: ["*"]
    nonResourceURLs: ["*"] 

Warning: rbac.authorization.k8s.io/v1beta1 ClusterRoleBinding is deprecated in v1.17+, unavailable in v1.22+; use rbac.authorization.k8s.io/v1 ClusterRoleBinding

unable to recognize "kubernetes-dashboard.yaml": no matches for kind "ClusterRole" in version "rbac.authorization.k8s.io/v1alpha1"
Error from server (Invalid): error when applying patch:
{"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"rbac.authorization.k8s.io/v1beta1\",\"kind\":\"ClusterRoleBinding\",\"metadata\":{\"annotations\":{},\"labels\":{\"k8s-app\":\"kubernetes-dashboard\"},\"name\":\"kubernetes-dashboard\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"cluster-admin\"},\"subjects\":[{\"kind\":\"ServiceAccount\",\"name\":\"kubernetes-dashboard\",\"namespace\":\"kube-system\"}]}\n"},"labels":{"k8s-app":"kubernetes-dashboard"}},"roleRef":{"name":"cluster-admin"},"subjects":[{"kind":"ServiceAccount","name":"kubernetes-dashboard","namespace":"kube-system"}]}
to:
Resource: "rbac.authorization.k8s.io/v1beta1, Resource=clusterrolebindings", GroupVersionKind: "rbac.authorization.k8s.io/v1beta1, Kind=ClusterRoleBinding"
Name: "kubernetes-dashboard", Namespace: ""
for: "kubernetes-dashboard.yaml": ClusterRoleBinding.rbac.authorization.k8s.io "kubernetes-dashboard" is invalid: roleRef: Invalid value: rbac.RoleRef{APIGroup:"rbac.authorization.k8s.io", Kind:"ClusterRole", Name:"cluster-admin"}: cannot change roleRef


https://stackoverflow.com/questions/60901631/dashboard-not-running
kubectl -n kubernetes-dashboard get all
kubectl -n kubernetes-dashboard describe svc kubernetes-dashboard
kubectl describe pod kubernetes-dashboard-665f4c5ff-tp5sm -n  kubernetes-dashboard
Node-Selectors:  kubernetes.io/os=linux

kubectl exec -it -n kubernetes-dashboard  kubernetes-dashboard-665f4c5ff-tp5sm sh


% kubectl get nodes --show-labels -n  kubernetes-dashboard
NAME             STATUS   ROLES    AGE   VERSION   LABELS
docker-desktop   Ready    master   17h   v1.19.7   beta.kubernetes.io/arch=arm64,beta.kubernetes.io/os=linux,kubernetes.io/arch=arm64,kubernetes.io/hostname=docker-desktop,kubernetes.io/os=linux,node-role.kubernetes.io/master=

% kubectl get nodes --show-labels -n kube-system
NAME             STATUS   ROLES    AGE   VERSION   LABELS
docker-desktop   Ready    master   17h   v1.19.7   beta.kubernetes.io/arch=arm64,beta.kubernetes.io/os=linux,kubernetes.io/arch=arm64,kubernetes.io/hostname=docker-desktop,kubernetes.io/os=linux,node-role.kubernetes.io/master=

% go run getos.go 
darwin
https://kubernetes.io/zh/docs/reference/kubernetes-api/labels-annotations-taints/#kubernetes-io-os

 kubernetes.io/os=linux  改成 darwin

 也不对，因为运行在虚拟机器，所以linux 没有问题

 %  kubectl rollout history deployment kubernetes-dashboard  -n  kubernetes-dashboard
deployment.apps/kubernetes-dashboard
REVISION  CHANGE-CAUSE
1         <none>
2         <none>


 % kubectl rollout undo deployment kubernetes-dashboard --to-revision=1

--record 的作用是将当前命令记录到 revision 记录中，这样我们就可以知道每个 revison 对应的是哪个配置文件。通过 kubectl rollout history deployment httpd 查看 revison 历史记录


https://segmentfault.com/a/1190000020675199
% kubectl get pod kubernetes-dashboard-665f4c5ff-tp5sm -n kubernetes-dashboard -o yaml | kubectl replace --force -f -
pod "kubernetes-dashboard-665f4c5ff-tp5sm" deleted
pod/kubernetes-dashboard-665f4c5ff-tp5sm replaced

% kubectl delete pod  dashboard-metrics-scraper-854fd5544b-gjbvh -n  kubernetes-dashboard --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
pod "dashboard-metrics-scraper-854fd5544b-gjbvh" force deleted


 % kubectl delete -f kubernetes-dashboard.v3.yaml
namespace "kubernetes-dashboard" deleted
serviceaccount "kubernetes-dashboard" deleted

https://blog.csdn.net/chenleiking/article/details/80197975

 % kubectl get pods  -n  kubernetes-dashboard -o wide

 已经全部删除了

  % kubectl apply -f kubernetes-dashboard.yaml --record


% kubectl get pods  -n  kubernetes-dashboard -o wide
NAME                                         READY   STATUS              RESTARTS   AGE   IP          NODE             NOMINATED NODE   READINESS GATES
dashboard-metrics-scraper-7b59f7d4df-2vwmx   1/1     Running             0          76s   10.1.0.11   docker-desktop   <none>           <none>
kubernetes-dashboard-665f4c5ff-qz4h9         0/1     ContainerCreating   0          76s   <none>      docker-desktop   <none>           <none>

应该是资源不够，加大内存问题解决

% kubectl get pods  -n  kubernetes-dashboard -o wide
NAME                                         READY   STATUS    RESTARTS   AGE    IP          NODE             NOMINATED NODE   READINESS GATES
dashboard-metrics-scraper-7b59f7d4df-2vwmx   1/1     Running   1          6m8s   10.1.0.17   docker-desktop   <none>           <none>
kubernetes-dashboard-665f4c5ff-qz4h9         1/1     Running   0          6m8s   10.1.0.16   docker-desktop   <none>           <none>


TOKEN=$(kubectl -n kube-system describe secret default| awk '$1=="token:"{print $2}')
kubectl config set-credentials docker-for-desktop --token="${TOKEN}"
echo $TOKEN


http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/overview?namespace=default


