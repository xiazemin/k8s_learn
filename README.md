https://github.com/AliyunContainerService/k8s-for-docker-desktop

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



 curl -o ingress-nginx.yaml https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-0.32.0/deploy/static/provider/cloud/deploy.yaml

kubectl apply -f ingress-nginx.yaml


 % kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx
NAMESPACE       NAME                                       READY   STATUS              RESTARTS   AGE
ingress-nginx   ingress-nginx-admission-create-hszqp       0/1     Completed           0          49s
ingress-nginx   ingress-nginx-admission-patch-q87l4        0/1     Completed           0          48s
ingress-nginx   ingress-nginx-controller-98f46f89d-g27tk   0/1     ContainerCreating   0          59s


brew install helm

# Add helm repo
helm repo add stable http://mirror.azure.cn/kubernetes/charts/

# Update charts repo
 % helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈Happy Helming!⎈


https://cloud.tencent.com/document/product/457/42946

% kubectl describe pod ingress-nginx-controller-98f46f89d-g27tk -n ingress-nginx

Events:
  Type     Reason       Age                  From               Message
  ----     ------       ----                 ----               -------
  Normal   Scheduled    13m                  default-scheduler  Successfully assigned ingress-nginx/ingress-nginx-controller-98f46f89d-g27tk to docker-desktop
  Warning  FailedMount  13m (x6 over 13m)    kubelet            MountVolume.SetUp failed for volume "webhook-cert" : secret "ingress-nginx-admission" not found
  Normal   Pulling      13m                  kubelet            Pulling image "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0"
  Warning  Failed       45s                  kubelet            Failed to pull image "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0": rpc error: code = Unknown desc = context canceled
  Warning  Failed       45s                  kubelet            Error: ErrImagePull
  Normal   BackOff      44s                  kubelet            Back-off pulling image "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0"
  Warning  Failed       44s                  kubelet            Error: ImagePullBackOff
  Normal   Pulling      32s (x2 over 6m31s)  kubelet            Pulling image "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0"
  Normal   Pulled       1s                   kubelet            Successfully pulled image "quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0" in 31.139592431


  docker pull quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.32.0


% kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx
NAMESPACE       NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx   ingress-nginx-admission-create-hszqp       0/1     Completed   0          18m
ingress-nginx   ingress-nginx-admission-patch-q87l4        0/1     Completed   0          18m
ingress-nginx   ingress-nginx-controller-98f46f89d-g27tk   1/1     Running     0          18m

https://stackoverflow.com/questions/51234378/why-do-pods-with-completed-status-still-show-up-in-kubctl-get-pods


%kubectl delete -f ingress-nginx.yaml

Error from server (Forbidden): error when creating "ingress-nginx.yaml": serviceaccounts "ingress-nginx-admission" is forbidden: unable to create new content in namespace ingress-nginx because it is being terminated

https://blog.51cto.com/liujingyu/2531898

%  kubectl delete -f ingress-nginx.yaml --force

%kubectl apply -f ingress-nginx.yaml


https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html

 % cp -r k8s-for-docker-desktop/sample sample

% kubectl create -f sample/apple.yaml
pod/apple-app created
service/apple-service created
% kubectl create -f sample/banana.yaml
pod/banana-app created
service/banana-service created
% kubectl create -f sample/ingress.yaml
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
ingress.extensions/example-ingress created

% kubectl get Ingress -o wide
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME              CLASS    HOSTS   ADDRESS     PORTS   AGE
example-ingress   <none>   *       localhost   80      49s

 % curl -kL http://localhost/apple
<html>
<head><title>503 Service Temporarily Unavailable</title></head>
<body>

% kubectl get pods
NAME         READY   STATUS             RESTARTS   AGE
apple-app    0/1     CrashLoopBackOff   8          32m
banana-app   0/1     CrashLoopBackOff   8          32m

% kubectl describe pod apple-app
  Warning  BackOff    1s (x3 over 20s)  kubelet            Back-off restarting failed container

% kubectl logs apple-app
runtime: failed to create new OS thread (have 2 already; errno=22)
fatal error: newosproc

runtime stack:
runtime.throw(0x6955a0, 0x9)
	/usr/local/go/src/runtime/panic.go:596 +0x95
runtime.newosproc(0xc420024000, 0xc420034000)
	/usr/local/go/src/runtime/os_linux.go:163 +0x18c
runtime.newm(0x6a4348, 0x0)
	/usr/local/go/src/runtime/proc.go:1628 +0x137
runtime.main.func1()
	/usr/local/go/src/runtime/proc.go:126 +0x36
runtime.systemstack(0x7cc200)
	/usr/local/go/src/runtime/asm_amd64.s:327 +0x79
runtime.mstart()
	/usr/local/go/src/runtime/proc.go:1132

goroutine 1 [running]:
runtime.systemstack_switch()
	/usr/local/go/src/runtime/asm_amd64.s:281 fp=0xc420020788 sp=0xc420020780
runtime.main()
	/usr/local/go/src/runtime/proc.go:127 +0x6c fp=0xc4200207e0 sp=0xc420020788
runtime.goexit()

https://github.com/golang/go/issues/19163

增加核数为6

https://blog.csdn.net/sqhren626232/article/details/101013390

%kubectl delete -f  sample/apple.yaml

spec:
  containers:
    - name: banana-app
      image: hashicorp/http-echo
      args:
        - "-text=banana"
      command: [ "/bin/bash", "-ce", "tail -f /dev/null" ]

% kubectl create -f sample/apple.yaml
pod/apple-app created
service/apple-service created



% kubectl describe pod apple-app
 Error: failed to start container "apple-app": Error response from daemon: OCI runtime create failed: container_linux.go:367: starting container process caused: exec: "/bin/bash": stat /bin/bash: no such file or directory: unknown

