
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


 % kubectl get services
NAME             TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
apple-service    ClusterIP   10.100.104.71    <none>        5678/TCP   14m
banana-service   ClusterIP   10.109.109.102   <none>        5678/TCP   13m
kubernetes       ClusterIP   10.96.0.1        <none>        443/TCP    26h


 % kubectl describe Ingress example-ingress
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Name:             example-ingress
Namespace:        default
Address:
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host        Path  Backends
  ----        ----  --------
  *
              /apple    apple-service:5678 (10.1.0.11:5678)
              /banana   banana-service:5678 (10.1.0.12:5678)
Annotations:  ingress.kubernetes.io/rewrite-target: /
Events:
  Type    Reason  Age    From                      Message
  ----    ------  ----   ----                      -------
  Normal  CREATE  5m52s  nginx-ingress-controller  Ingress default/example-ingress

https://stackoverflow.com/questions/66557860/default-http-backend80-error-endpoints-default-http-backend-not-found

https://github.com/nginxinc/kubernetes-ingress/issues/966

https://stackoverflow.com/questions/63558461/endpoints-default-http-backend-not-found-in-ingress-resource

Create default-http-backend service in kube-system namespace and error will be gone.


 % kubectl get pod -n ingress-nginx
NAME                                       READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-v9bn7       0/1     Completed   0          34m
ingress-nginx-admission-patch-55hln        0/1     Completed   0          34m
ingress-nginx-controller-98f46f89d-vlrkd   1/1     Running     0          35m

% kubectl logs ingress-nginx-admission-create-v9bn7 -n  ingress-nginx
{"err":"secrets \"ingress-nginx-admission\" not found","level":"info","msg":"no secret found","source":"k8s/k8s.go:101","time":"2021-03-22T15:02:38Z"}
{"level":"info","msg":"creating new secret","source":"cmd/create.go:23","time":"2021-03-22T15:02:38Z"}

https://kubernetes.github.io/ingress-nginx/troubleshooting/

https://www.jianshu.com/p/131bf15235b1


https://github.com/kubernetes/ingress-nginx/blob/master/docs/deploy/index.md

 % curl -o ingress-nginx.yaml https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.44.0/deploy/static/provider/cloud/deploy.yaml


https://matthewpalmer.net/kubernetes-app-developer/articles/kubernetes-ingress-guide-nginx-example.html

 https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml

  % kubectl get pods -n ingress-nginx \
  -l app.kubernetes.io/name=ingress-nginx --watch

kubectl apply -f ingress-nginx.yaml

TopologySpreadConstraints:[]core.TopologySpreadConstraint(nil)}}: field is immutable

% kubectl delete -f ingress-nginx.yaml

% kubectl apply -f ingress-nginx.v2.yaml
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
configmap/ingress-nginx-controller created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
service/ingress-nginx-controller-admission created
service/ingress-nginx-controller created
deployment.apps/ingress-nginx-controller created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
serviceaccount/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created


https://github.com/AliyunContainerService/k8s-for-docker-desktop

 % docker pull k8s.gcr.io/ingress-nginx/controller:v0.44.0
Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)

 %  kubectl get pods -n ingress-nginx
NAME                                        READY   STATUS             RESTARTS   AGE
ingress-nginx-admission-create-mklnk        0/1     Completed          0          8m48s
ingress-nginx-admission-patch-qthqj         0/1     Completed          0          8m47s
ingress-nginx-controller-7fc74cf778-tq5gw   0/1     ImagePullBackOff   0          8m48s


 Warning  Failed       7m55s (x4 over 10m)     kubelet            Failed to pull image "k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a": rpc error: code = Unknown desc = Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
  Normal   BackOff      5m39s (x12 over 9m59s)  kubelet            Back-off pulling image "k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a"
  Warning  Failed       41s (x33 over 9m59s)    kubelet            Error: ImagePullBackOff
xiazemin@bogon MyBlogSrc % docker pull k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a
Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)


