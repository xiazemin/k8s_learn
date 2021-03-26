https://hub.docker.com/r/rancher/rancher/tags?page=1&ordering=last_updated

http://docs.rancher.cn/


docker pull rancher/rancher:v2.5.8-rc1-linux-arm64

https://www.jianshu.com/p/40f4fbe1ec22

https://www.jianshu.com/p/adfeafc936f9?utm_campaign=hugo

 % docker pull registry.cn-hangzhou.aliyuncs.com/rancher/rancher:v2.4.2
 Digest: sha256:248ddca1169e8a4e06babd50e8105cbba0a326f86ec4de3e38d61e8909ffdb4e
Status: Downloaded newer image for registry.cn-hangzhou.aliyuncs.com/rancher/rancher:v2.4.2
registry.cn-hangzhou.aliyuncs.com/rancher/rancher:v2.4.2

 http://mirror.cnrancher.com/

  % docker inspect registry.cn-hangzhou.aliyuncs.com/rancher/rancher:v2.4.2

% mkdir -p  ~/docker_volume/rancher_home/rancher
% mkdir -p ~/docker_volume/rancher_home/auditlog


docker run -d --restart=unless-stopped -p 80:80 -p 443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher registry.cn-hangzhou.aliyuncs.com/rancher/rancher:v2.4.2  

WARNING: The requested image's platform (linux/amd64) does not match the detected host platform (linux/arm64/v8) and no specific platform was requested
ace7f0cb84c05e8bf43652a75bb6ccb293e4c25262bf430d4a00220ccfdb234b


https://www.jianshu.com/p/b4beb8b15bb8

docker pull rancher/rancher:v2.5.8-rc1

docker run -d --restart=unless-stopped -p 80:8088 -p 443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:v2.5.8-rc1

端口与nginx 冲突了

docker run -d --restart=unless-stopped --privileged=true  -p 8088:80 -p 8443:443 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:v2.5.8-rc1

必须是特权模式否则起不来

docker run -d --restart=unless-stopped --privileged=true -p 8443:443 -p 8088:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:v2.5.8-rc1

https://localhost:8443

https://blog.csdn.net/qq_43792385/article/details/104424848

 % kubectl apply -f https://localhost:8443/v3/import/tbbz5s24p7t9h9xzvk7svwnwfncprtdf96swgfxv8fkchv5lgkrnf7_c-tbkmc.yaml
Unable to connect to the server: x509: certificate signed by unknown authority

 %  curl --insecure -sfL https://localhost:8443/v3/import/tbbz5s24p7t9h9xzvk7svwnwfncprtdf96swgfxv8fkchv5lgkrnf7_c-tbkmc.yaml | kubectl apply -f -
clusterrole.rbac.authorization.k8s.io/proxy-clusterrole-kubeapiserver created
clusterrolebinding.rbac.authorization.k8s.io/proxy-role-binding-kubernetes-master created
namespace/cattle-system created
serviceaccount/cattle created
clusterrolebinding.rbac.authorization.k8s.io/cattle-admin-binding created
secret/cattle-credentials-fa0865b created
clusterrole.rbac.authorization.k8s.io/cattle-admin created
deployment.apps/cattle-cluster-agent created



 % docker inspect --format '{{ .NetworkSettings.IPAddress }}' rancher
172.17.0.3

https://www.jianshu.com/p/5fb3e1a998d6
https://zhuanlan.zhihu.com/p/27506799

 % kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user admin
clusterrolebinding.rbac.authorization.k8s.io/cluster-admin-binding created

 mkdir -p ~/docker_volume/rancher_home/volume


 http://docs.rancher.cn/docs/rancher2/installation_new/_index/

 docker logs -f rancher

  % helm repo add rancher-latest https://releases.rancher.com/server-charts/latest
"rancher-latest" has been added to your repositories



kubectl apply -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.12/deploy/manifests/00-crds.yaml

 % kubectl apply -f cert-manager.yaml
Warning: apiextensions.k8s.io/v1beta1 CustomResourceDefinition is deprecated in v1.16+, unavailable in v1.22+; use apiextensions.k8s.io/v1 CustomResourceDefinition
customresourcedefinition.apiextensions.k8s.io/certificaterequests.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/certificates.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/challenges.acme.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/clusterissuers.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/issuers.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/orders.acme.cert-manager.io created

apiextensions.k8s.io/v1beta1 => apiextensions.k8s.io/v1


error: error validating "cert-manager.yaml": error validating data: [ValidationError(CustomResourceDefinition.spec): unknown field "additionalPrinterColumns" in io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionSpec, ValidationError(CustomResourceDefinition.spec): unknown field "subresources" in io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionSpec, ValidationError(CustomResourceDefinition.spec): unknown field "validation" in io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionSpec, ValidationError(CustomResourceDefinition.spec): unknown field "version" in io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionSpec]; if you choose to ignore these errors, turn validation off with --validate=false

https://www.orchome.com/8203

https://github.com/jetstack/cert-manager

 kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml

 https://cert-manager.io/docs/installation/kubernetes/


 % kubectl apply -f cert-manager.crds.yaml
customresourcedefinition.apiextensions.k8s.io/certificaterequests.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/certificates.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/challenges.acme.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/clusterissuers.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/issuers.cert-manager.io created
customresourcedefinition.apiextensions.k8s.io/orders.acme.cert-manager.io created

% helm repo add jetstack https://charts.jetstack.io
"jetstack" has been added to your repositories

 % helm search  repo cert-manager
NAME                           	CHART VERSION	APP VERSION	DESCRIPTION
jetstack/cert-manager          	v1.2.0       	v1.2.0     	A Helm chart for cert-manager

 % helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.2.0 \
  --create-namespace

  cert-manager has been deployed successfully!

   % kubectl get pods --namespace cert-manager
NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-85f9bbcd97-k78s2              1/1     Running   0          90s
cert-manager-cainjector-74459fcc56-8ts82   1/1     Running   0          90s
cert-manager-webhook-57d97ccc67-hl4rc      1/1     Running   0          90s



kubectl create namespace cert-manager

kubectl label namespace cert-manager certmanager.k8s.io/disable-validation=true


helm install rancher-latest/rancher \
  --name rancher \
  --namespace cattle-system \
  --set hostname=rancher.my.org
Error: unknown flag: --name

https://stackoverflow.com/questions/57961162/helm-install-unknown-flag-name

helm install rancher-k8s rancher-latest/rancher  \
  --namespace cattle-system \
  --set hostname=rancher.my.org

  Rancher Server has been installed.

% kubectl -n cattle-system rollout status deploy/rancher-k8s
Waiting for deployment "rancher-k8s" rollout to finish: 0 of 3 updated replicas are available...
Waiting for deployment "rancher-k8s" rollout to finish: 1 of 3 updated replicas are available...
Waiting for deployment "rancher-k8s" rollout to finish: 2 of 3 updated replicas are available...
deployment "rancher-k8s" successfully rolled out

 % kubectl get pods -n cattle-system
NAME                                   READY   STATUS             RESTARTS   AGE
cattle-cluster-agent-685f5d455-dxp8m   0/1     CrashLoopBackOff   18         137m
helm-operation-4dh9m                   0/2     Pending            0          1s
helm-operation-znq5g                   2/2     Running            0          62s
rancher-k8s-6489f8689-86wcw            1/1     Running            0          5m46s
rancher-k8s-6489f8689-bj8px            1/1     Running            0          5m46s
rancher-k8s-6489f8689-htbxl            1/1     Running            0          5m46s

 % kubectl -n cattle-system get deploy rancher-k8s
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
rancher-k8s   3/3     3            3           7m48s

% kubectl describe deploy rancher-k8s  -n cattle-system
Name:                   rancher-k8s

http://docs.rancher.cn/docs/rancher2/installation_new/install-rancher-on-k8s/_index

% kubectl get svc rancher-k8s -n cattle-system
NAME          TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
rancher-k8s   ClusterIP   10.106.15.8   <none>        80/TCP,443/TCP   24m


 % kubectl get ns
NAME                                     STATUS   AGE
cattle-global-data                       Active   23m
cattle-global-nt                         Active   23m
cattle-system                            Active   159m
cert-manager                             Active   33m
cluster-fleet-local-local-1a3d67d0a899   Active   18m
default                                  Active   2d19h
fleet-clusters-system                    Active   18m
fleet-default                            Active   19m
fleet-local                              Active   18m
fleet-system                             Active   21m
ingress-nginx                            Active   8h
kube-node-lease                          Active   2d19h
kube-public                              Active   2d19h
kube-system                              Active   2d19h
kubernetes-dashboard                     Active   2d1h
local                                    Active   23m
p-ghxxw                                  Active   23m
p-vtq96                                  Active   23m
rancher-operator-system                  Active   23m


 % helm list  -n cattle-system
NAME           	NAMESPACE    	REVISION	UPDATED                                	STATUS  	CHART                        	APP VERSION
rancher-k8s    	cattle-system	1       	2021-03-24 17:56:49.843259 +0800 CST   	deployed	rancher-2.5.7                	v2.5.7
rancher-webhook	cattle-system	1       	2021-03-24 10:02:51.547798215 +0000 UTC	deployed	rancher-webhook-0.1.0-beta901	0.1.0-beta9

https://blog.csdn.net/oJinXuan1/article/details/85008262

% helm inspect values rancher-latest/rancher
# Additional Trusted CAs.
# Enable this flag and add your CA certs as a secret named tls-ca-additional in the namespace.
# See README.md for details.
additionalTrustedCAs: false

 % helm delete rancher-k8s -n cattle-system
Error: uninstallation completed with 1 error(s): unable to build kubernetes objects for delete: unable to recognize "": no matches for kind "Issuer" in version "cert-manager.io/v1beta1"

% helm delete rancher-webhook -n  cattle-system
release "rancher-webhook" uninstalled


% helm list  -n cattle-system
NAME	NAMESPACE	REVISION	UPDATED	STATUS	CHART	APP VERSION

% helm list  -A
NAME            	NAMESPACE              	REVISION	UPDATED                                	STATUS  	CHART                                                                             	APP VERSION
fleet           	fleet-system           	1       	2021-03-26 14:25:36.082813507 +0000 UTC	deployed	fleet-0.3.400                                                                     	0.3.4
fleet-agent     	fleet-system           	1       	2021-03-26 14:26:06.598600674 +0000 UTC	deployed	fleet-agent-v0.0.0+s-2879661f9d0e3cfaea1b12951bb30640413f45ac5ccec4579b5c261733160
fleet-crd       	fleet-system           	1       	2021-03-26 14:26:03.402779131 +0000 UTC	deployed	fleet-crd-0.3.400                                                                 	0.3.4
rancher-operator	rancher-operator-system	1       	2021-03-26 14:26:20.13096475 +0000 UTC 	deployed	rancher-operator-0.1.300                                                          	0.1.3


% helm delete fleet -n fleet-system
release "fleet" uninstalled

% helm delete fleet-agent -n fleet-system
release "fleet-agent" uninstalled

% helm delete fleet-crd -n fleet-system

 % helm delete rancher-operator -n      rancher-operator-system
release "rancher-operator" uninstalled

 % kubectl get all -n cattle-global-data
No resources found in cattle-global-data namespace.

https://stackoverflow.com/questions/55036464/how-to-deleteuninstall-helm-chart-on-specific-resource

 % kubectl delete ns cattle-global-data
namespace "cattle-global-data" deleted

% kubectl delete ns cattle-global-nt
namespace "cattle-global-nt" deleted

% kubectl delete ns cattle-system
namespace "cattle-system" deleted

% kubectl get all -n cluster-fleet-local-local-1a3d67d0a899
No resources found in cluster-fleet-local-local-1a3d67d0a899 namespace.

% kubectl delete ns cluster-fleet-local-local-1a3d67d0a899
namespace "cluster-fleet-local-local-1a3d67d0a899" deleted

https://blog.51cto.com/13760351/2494356

https://www.jianshu.com/p/5fb3e1a998d6
https://new.qq.com/omn/20210104/20210104A0BCZZ00.html
https://blog.51cto.com/13760351/2494356

https://zhuanlan.zhihu.com/p/109032970



