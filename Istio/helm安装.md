为 Istio 组件，创建命名空间 istio-system :

$ kubectl create namespace istio-system

安装 Istio base chart，它包含了 Istio 控制平面用到的集群范围的资源：

$ helm install istio-base manifests/charts/base -n istio-system

安装 Istio discovery chart，它用于部署 istiod 服务：

$ helm install istiod manifests/charts/istio-control/istio-discovery \
    --set global.hub="docker.io/istio" \
    --set global.tag="1.12.1" \
    -n istio-system

(可选项) 安装 Istio 的入站网关 chart，它包含入站网关组件：

$ helm install istio-ingress manifests/charts/gateways/istio-ingress \
    --set global.hub="docker.io/istio" \
    --set global.tag="1.12.1" \
    -n istio-system

(可选项) 安装 Istio 的出站网关 chart，它包含了出站网关组件：

$ helm install istio-egress manifests/charts/gateways/istio-egress \
    --set global.hub="docker.io/istio" \
    --set global.tag="1.12.1" \
    -n istio-system



https://istio.io/latest/zh/docs/setup/install/helm/


 %  helm install istio-base manifests/charts/base -n istio-system
Error: INSTALLATION FAILED: failed to download "manifests/charts/base"

https://www.guojingyi.cn/899.html

https://blog.csdn.net/weixin_44782815/article/details/108830849

 % helm repo add stable http://mirror.azure.cn/kubernetes/charts/
"stable" already exists with the same configuration, skipping

 % helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "ingress-nginx" chart repository
...Successfully got an update from the "jetstack" chart repository
...Successfully got an update from the "stable" chart repository
...Successfully got an update from the "rancher-latest" chart repository
...Successfully got an update from the "apphub" chart repository
Update Complete. ⎈Happy Helming!⎈

% helm repo remove stable
"stable" has been removed from your repositories

 %  helm repo add stable http://mirror.azure.cn/kubernetes/charts/
"stable" has been added to your repositories


helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.9.2/charts/

 % helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.9.2/charts/
Error: looks like "https://storage.googleapis.com/istio-release/releases/1.9.2/charts/" is not a valid chart repository or cannot be reached: failed to fetch https://storage.googleapis.com/istio-release/releases/1.9.2/charts/index.yaml : 404 Not Found


https://istio.io/latest/docs/setup/install/helm/
 % helm repo add istio https://istio-release.storage.googleapis.com/charts
"istio" has been added to your repositories

helm repo add istio https://istio.io/charts
https://stackoverflow.com/questions/49311296/helm-repo-to-install-istio


% helm install istio-base istio/base -n istio-system
Error: INSTALLATION FAILED: rendered manifests contain a resource that already exists. Unable to continue with install: ServiceAccount "istio-reader-service-account" in namespace "istio-system" exists and cannot be imported into the current release: invalid ownership metadata; label validation error: missing key "app.kubernetes.io/managed-by": must be set to "Helm"; annotation validation error: missing key "meta.helm.sh/release-name": must be set to "istio-base"; annotation validation error: missing key "meta.helm.sh/release-namespace": must be set to "istio-system"

 % helm install istiod istio/istiod -n istio-system --wait
W0103 16:09:50.942939   45274 warnings.go:70] policy/v1beta1 PodDisruptionBudget is deprecated in v1.21+, unavailable in v1.25+; use policy/v1 PodDisruptionBudget
Error: INSTALLATION FAILED: rendered manifests contain a resource that already exists. Unable to continue with install: PodDisruptionBudget "istiod" in namespace "istio-system" exists and cannot be imported into the current release: invalid ownership metadata; label validation error: missing key "app.kubernetes.io/managed-by": must be set to "Helm"; annotation validation error: missing key "meta.helm.sh/release-name": must be set to "istiod"; annotation validation error: missing key "meta.helm.sh/release-namespace": must be set to "istio-system"

% helm status istiod -n istio-system
Error: release: not found


% istioctl install 
This will install the Istio 1.12.1 default profile with ["Istio core" "Istiod" "Ingress gateways"] components into the cluster. Proceed? (y/N) y
✔ Istio core installed                                                                                                                  
- Processing resources for Istiod. Waiting for Deployment/istio-system/istiod                                                           


% brew install kubernetes-helm
Updating Homebrew...

% istioctl analyze
Info [IST0102] (Namespace default) The namespace is not enabled for Istio injection. Run 'kubectl label namespace default istio-injection=enabled' to enable it, or 'kubectl label namespace default istio-injection=disabled' to explicitly mark it as not needing injection.
Info [IST0118] (Service default/apple-headless-service) Port name  (port: 5678, targetPort: 5678) doesn't follow the naming convention of Istio port.
Info [IST0118] (Service default/minio-service) Port name api (port: 9000, targetPort: 9000) doesn't follow the naming convention of Istio port.
Info [IST0118] (Service default/minio-service) Port name console (port: 9001, targetPort: 9001) doesn't follow the naming convention of Istio port.
Info [IST0118] (Service default/sidecar-injector-webhook-svc) Port name  (port: 7896, targetPort: 7896) doesn't follow the naming convention of Istio port.


 % istioctl  verify-install    
1 Istio control planes detected, checking --revision "default" only
error while fetching revision : control plane revision "" not found
1 Istio injectors detected
Error: Istio present but verify-install needs an IstioOperator or manifest for comparison. Supply flag --filename <yaml>

