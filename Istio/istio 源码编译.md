https://www.servicemesher.com/blog/istio-deepin-part1-framework-and-environment/

编译镜像
# make init # 初始化，检查目录结构、Go版本号、初始化环境变量、检查vendor等
make docker # 对各组件（istioctl、mixer、pilot、istio-auth等）进行二进制包编译、测试、镜像编译
make push # 推送镜像到dockerhub
# 其他指令
make pilot  docker.pilot # 编译pilot组件和镜像
make app  docker.app # 编译app组件和镜像
make proxy  docker.proxy # 编译proxy组件和镜像
make proxy_init  docker.proxy_init # 编译proxy_init组件和镜像
make proxy_debug  docker.proxy_debug # 编译proxy_debug组件和镜像
make sidecar_injector  docker.sidecar_injector # 编译sidecar_injector组件和镜像
make proxyv2  docker.proxyv2 # 编译proxyv2组件和镜像

make push.docker.pilot # 推送pilot镜像到dockerhub，其他组件类似


 % make pilot  docker.pilot
Unable to find image 'gcr.io/istio-testing/build-tools:master-2021-12-21T23-55-44' locally
docker: Error response from daemon: Get "https://gcr.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers).
See 'docker run --help'.
make: *** [pilot] Error 125

% make init
Unable to find image 'gcr.io/istio-testing/build-tools:master-2021-12-21T23-55-44' locally
docker: Error response from daemon: Get "https://gcr.io/v2/": context deadline exceeded.
See 'docker run --help'.
make: *** [init] Error 125


https://github.com/istio/release-builder

https://github.com/istio/test-infra/blob/master/prow/config/jobs/release-builder.yaml

 % git submodule add https://github.com/istio/release-builder
Cloning into '/Users/xiazemin/source/k8s_learn/Istio/release-builder'...

 %cd release-builder 
 %mkdir -p /tmp/istio-release; go run main.go build --manifest example/manifest.yaml; go run main.go validate --release /tmp/istio-release/out


 vi Istio/release-builder/example/manifest.yaml
 % mkdir -p ./istio-release; go run main.go build --manifest example/manifest.yaml; go run main.go validate --release ./istio-release/out

./tools/docker --save
go: GOPATH entry is relative; must be absolute path: "istio-release/work".
For more details see: 'go help gopath'
make: *** [dockerx.save] Error 2
Error: failed to build: failed to build Docker: failed to create istio docker archives: exit status 2
exit status 1
2022-01-08T04:29:43.352128Z     info    test temporary dir at /tmp/release-test439917684
panic: failed to read manifest file: open istio-release/out/manifest.yaml: no such file or directory

goroutine 1 [running]:
istio.io/release-builder/pkg/validate.NewReleaseInfo(0x16f71733b, 0x13, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, ...)


% cp example/manifest.yaml istio-release/out 



用默认的路径编译
mkdir -p /tmp/istio-release; go run main.go build --manifest example/manifest.yaml; go run main.go validate --release /tmp/istio-release/out


make: *** [dockerx.save] Error 2
Error: failed to build: failed to build Docker: failed to create istio docker archives: exit status 2
exit status 1




% vi Makefile.core.mk 

#export GOPROXY ?= https://proxy.golang.org
#export GOSUMDB ?= sum.golang.org
export GOPROXY = https://goproxy.cn


https://www.jianshu.com/p/501b8698038e



https://aijishu.com/a/1060000000016469

https://www.tetrate.io/blog/envoy-1-16-%E7%9A%84%E6%9C%80%E6%96%B0%E6%B6%88%E6%81%AF%EF%BC%9A%E6%94%AF%E6%8C%81-arm64-%E4%BB%A5%E5%8F%8A%E5%85%B6%E4%BB%96%E5%8A%9F%E8%83%BD%EF%BC%81/?lang=zh-hans

https://issueexplorer.com/issue/istio/istio/35275
istio社区目前还不支持ARM64镜像，使用默认的istioctl工具无法自动化部署istio服务。

如需使用，建议参考https://aijishu.com/a/1060000000016469 进行镜像的手动编译、打包和运行；

另外也可以使用https://hub.docker.com/search?q=istio&type=image&architecture=arm64 raspbernetes的镜像进行运行尝试。


https://github.com/alibaba/nacos/issues/4309




https://hub.docker.com/search?q=istio&type=image&architecture=arm64

https://hub.docker.com/r/istioarm64/build-tools

https://hub.docker.com/r/istiojfh/pilot


% docker pull istiojfh/pilot:1.8.1-arm64
1.8.1-arm64: Pulling from istiojfh/pilot
04da93b342eb: Pull complete 
b235194751de: Pull complete 
606a67bb8db9: Pull complete 
ed8ae9a93463: Pull complete 
b233ccb7052d: Pull complete 
7f4d9d2c73b5: Pull complete 
Digest: sha256:917eb59b1e6fda4c84a5631f0bf29fc9b79a804ce0bf9b00a429d2635ff6c523
Status: Downloaded newer image for istiojfh/pilot:1.8.1-arm64
docker.io/istiojfh/pilot:1.8.1-arm64



Istio/istiod.yaml
  image: docker.io/istio/pilot:1.12.0

  image: docker.io/istiojfh/pilot:1.8.1-arm64

 % kubectl apply -f ../istiod.yaml 
deployment.apps/istiod created




https://hub.docker.com/r/querycapistio/build-tools

docker pull querycapistio/build-tools:release-1.12-latest
d1ba108e2816: Pull complete 
Digest: sha256:520397dc865b950c5954effeb5bb9a1b88b97ce8b9fbf44361c63c9089811dd5
Status: Downloaded newer image for querycapistio/build-tools:release-1.12-latest
docker.io/querycapistio/build-tools:release-1.12-latest


% docker tag docker.io/querycapistio/build-tools:release-1.12-latest gcr.io/istio-testing/build-tools:release-1.8-2020-11-12T22-53-11


% make build                                                            fatal: not a git repository: /work/../../.git/modules/Istio/istio
fatal: not a git repository: /work/../../.git/modules/Istio/istio
Makefile.core.mk:191: *** "TAG cannot be empty".  Stop.
make: *** [build] Error 2

 % make docker    
fatal: not a git repository: /work/../../.git/modules/Istio/istio
fatal: not a git repository: /work/../../.git/modules/Istio/istio
Makefile.core.mk:179: *** "TAG cannot be empty".  Stop.
make: *** [docker] Error 2


docker tag docker.io/querycapistio/build-tools:release-1.12-latest gcr.io/istio-testing/build-tools:master-2021-12-21T23-55-44


https://zhuanlan.zhihu.com/p/36574132


GOPATH="/Users/xiazemin/go"
% echo $GOPATH                      
/Users/xiazemin/go
% export ISTIO=$GOPATH/src/istio.io
% mkdir -p $ISTIO
cd $ISTIO

% cp -r ~/source/k8s_learn/Istio/istio .
% cd istio 

% vi .env

% git clone https://github.com/Istio/istio 

% make init
Unable to find image 'gcr.io/istio-testing/build-tools:master-2022-01-07T14-34-14' locally
^Cmake: *** [init] Error 255

docker tag docker.io/querycapistio/build-tools:release-1.12-latest gcr.io/istio-testing/build-tools:master-2022-01-07T14-34-14


% make initISTIO_OUT=/work/out/darwin_amd64 ISTIO_BIN=/gobin GOOS_LOCAL=darwin bin/retry.sh SSL_ERROR_SYSCALL bin/init.sh
Skipping envoy debug. Set DEBUG_IMAGE to download.
/work/out/linux_amd64/release /work
Downloading envoy: https://storage.googleapis.com/istio-build/proxy/envoy-alpha-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.tar.gz to /work/out/linux_amd64/release/envoy-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a


 % cp -r out/darwin_amd64  out/darwin_arm64

 % vi tools/docker-copy.sh 
TARGET_ARCH=${TARGET_ARCH:-arm64}

% export TARGET_ARCH=arm64 
% make init                
ISTIO_OUT=/work/out/darwin_arm64 ISTIO_BIN=/gobin GOOS_LOCAL=darwin bin/retry.sh SSL_ERROR_SYSCALL bin/init.sh
Skipping envoy debug. Set DEBUG_IMAGE to download.
/work/out/linux_arm64/release /work
Downloading envoy: https://storage.googleapis.com/istio-build/proxy/envoy-alpha-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.tar.gz to /work/out/linux_arm64/release/envoy-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a

real    0m37.638s
user    0m1.026s
sys     0m3.065s
Copying /work/out/linux_arm64/release/envoy-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a to /work/out/linux_arm64/release/envoy
/work
/work/out/linux_arm64/release /work
Downloading envoy: https://storage.googleapis.com/istio-build/proxy/envoy-centos-alpha-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.tar.gz to /work/out/linux_arm64/release/envoy-centos-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a

real    0m33.033s
user    0m1.124s
sys     0m2.811s
Copying /work/out/linux_arm64/release/envoy-centos-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a to /work/out/linux_arm64/release/envoy-centos
/work
/work/out/darwin_arm64/release /work
Downloading envoy: https://github.com/istio/proxy/releases/download/1.0.2/istio-proxy-1.0.2-macos.tar.gz to /work/out/darwin_arm64/release/envoy-1.0.2

real    0m14.291s
user    0m0.434s
sys     0m0.978s
Copying /work/out/darwin_arm64/release/envoy-1.0.2 to /work/out/darwin_arm64/release/envoy
/work
/work/out/linux_arm64/release /work
Downloading WebAssembly file: https://storage.googleapis.com/istio-build/proxy/stats-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.wasm to /work/out/linux_arm64/release/stats-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.wasm

real    0m1.723s
user    0m0.043s
sys     0m0.082s
/work
/work/out/linux_arm64/release /work
Downloading WebAssembly file: https://storage.googleapis.com/istio-build/proxy/stats-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.compiled.wasm to /work/out/linux_arm64/release/stats-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.compiled.wasm

real    0m1.944s
user    0m0.045s
sys     0m0.075s
/work
/work/out/linux_arm64/release /work
Downloading WebAssembly file: https://storage.googleapis.com/istio-build/proxy/metadata_exchange-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.wasm to /work/out/linux_arm64/release/metadata_exchange-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.wasm

real    0m1.432s
user    0m0.029s
sys     0m0.068s
/work
/work/out/linux_arm64/release /work
Downloading WebAssembly file: https://storage.googleapis.com/istio-build/proxy/metadata_exchange-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.compiled.wasm to /work/out/linux_arm64/release/metadata_exchange-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a.compiled.wasm

real    0m2.525s
user    0m0.054s
sys     0m0.142s
/work
Copying /work/out/darwin_arm64/release/envoy-1.0.2 to /work/out/darwin_arm64/envoy
Copying /work/out/linux_arm64/release/envoy-centos-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a to /work/out/linux_arm64/envoy-centos
Copying /work/out/linux_arm64/release/envoy-52a86aa341d9b5d4d8b5127277f81bbd0c4db16a to /work/out/linux_arm64/envoy
touch /work/out/darwin_arm64/istio_is_init
mkdir -p /work/out/darwin_arm64/logs
mkdir -p /work/out/darwin_arm64/release



编译镜像
# make init # 初始化，检查目录结构、Go版本号、初始化环境变量、检查vendor等
make docker # 对各组件（istioctl、mixer、pilot、istio-auth等）进行二进制包编译、测试、镜像编译
make push # 推送镜像到dockerhub

# 其他指令
make pilot  docker.pilot # 编译pilot组件和镜像
make app  docker.app # 编译app组件和镜像
make proxy  docker.proxy # 编译proxy组件和镜像
make proxy_init  docker.proxy_init # 编译proxy_init组件和镜像
make proxy_debug  docker.proxy_debug # 编译proxy_debug组件和镜像
make sidecar_injector  docker.sidecar_injector # 编译sidecar_injector组件和镜像
make proxyv2  docker.proxyv2 # 编译proxyv2组件和镜像

make push.docker.pilot # 推送pilot镜像到dockerhub，其他组件类似


% make pilot  docker.pilot


 % docker exec -it gcr.io/istio-testing/build-tools:master-2022-01-07T14-34-14 sh
Error: No such container: gcr.io/istio-testing/build-tools:master-2022-01-07T14-34-14


% ps aux |grep make


 % docker run -it gcr.io/istio-testing/build-tools:master-2022-01-07T14-34-14 sh
 没有反应，运行不起来 应该是 darwin arm64


% curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.12.1 TARGET_ARCH=arm64 sh -
https://istio.io/latest/docs/setup/getting-started/



 % istioctl   verify-install 
3 Istio control planes detected, checking --revision "default" only
error while fetching revision : control plane revision "" not found
1 Istio injectors detected
Error: Istio present but verify-install needs an IstioOperator or manifest for comparison. Supply flag --filename <yaml>

https://github.com/istio/istio/commit/032be16c2947d9d9658dc7d8c9b369b7fea7fd43


% istioctl manifest apply --set values.global.mtls.enabled=true 
Run the command with the --force flag if you want to ignore the validation error and proceed.
Error: generate config: unknown field "mtls" in v1alpha1.GlobalConfig


./istioctl manifest apply --set values.global.mtls.enabled=true
./istioctl manifest generate --set values.global.mtls.enabled=true > i1.5mtls.yaml
./istioctl verify-install --enableVerbose=false --filename i1.5mtls.yaml

https://github.com/istio/istio/issues/20000



% istioctl verify-install -f istio-1.12.1/manifest.yaml          
Checked 0 custom resource definitions
Checked 0 Istio Deployments
! No Istio installation found: unable to decode "istio-1.12.1/manifest.yaml": Object 'Kind' is missing in '{"dashboards":{"istio-extension-dashboard":13277,"istio-mesh-dashboard":7639,"istio-performance-dashboard":11829,"istio-service-dashboard":7636,"istio-workload-dashboard":7630,"pilot-dashboard":7645},"dependencies":{"api":{"sha":"1a632586cbd49d7b151affa135e69bcd8da2db18"},"client-go":{"goversionenabled":true,"sha":"6cbf560fe24ab852d213eaa5ef8adc0b182893ae"},"envoy":{"sha":"ea23f47b27464794980c05ab290a3b73d801405e"},"gogo-genproto":{"sha":"5eda25c962701b2c662124de51756bbba2b6afed"},"istio":{"sha":"88902a51acfb0383809608ccff169319560f768c"},"pkg":{"sha":"57f93bc2eff758f4daf523988aaf89b789af55b7"},"proxy":{"sha":"e6f45abcf874983fbff384459d70b28c072f68b5"},"test-infra":{"sha":"cd5b05895243226cce64411bac33bce72f2fd29d"},"tools":{"sha":"8aa7dcd19fa4ea57628f5ad57b0303acc7228d93"}},"docker":"docker.io/istio","version":"1.12.1"}'
Error: no Istio installation found

 % istioctl  analyze
Error [IST0145] (Gateway default/bookinfo-gateway) Conflict with gateways knative-serving/knative-ingress-gateway (workload selector istio=ingressgateway, port 80, hosts *).


% kubectl -n knative-serving get deploy |awk '{print $1}' |grep -v NAME |xargs kubectl -n knative-serving  delete deploy

% kubectl -n knative-serving get all |awk '{print $1}' |grep -v NAME |xargs kubectl -n knative-serving  delete 



 % istioctl  analyze                                                     Error [IST0145] (Gateway default/bookinfo-gateway) Conflict with gateways knative-serving/knative-ingress-gateway (workload selector istio=ingressgateway, port 80, hosts *).


 % kubectl get gateway -o wide                         
NAME               AGE
bookinfo-gateway   4d21h


% kubectl delete gateway bookinfo-gateway
gateway.networking.istio.io "bookinfo-gateway" deleted



 % istioctl  analyze                                   
Error [IST0101] (VirtualService default/bookinfo) Referenced gateway not found: "bookinfo-gateway"



 % kubectl get VirtualService             
NAME       GATEWAYS               HOSTS   AGE
bookinfo   ["bookinfo-gateway"]   ["*"]   4d21h


 % kubectl delete VirtualService bookinfo
virtualservice.networking.istio.io "bookinfo" deleted



检查都安装了什么
kubectl -n istio-system get IstioOperator installed-state -o yaml > installed-state.yaml

% kubectl -n istio-system get IstioOperator installed-state -o yaml
Error from server (NotFound): istiooperators.install.istio.io "installed-state" not found



% kubectl delete istiooperators.install.istio.io -n istio-system example-istiocontrolplane
Error from server (NotFound): istiooperators.install.istio.io "example-istiocontrolplane" not found


% istioctl operator remove
Operator controller is not installed in istio-operator namespace (no Deployment detected).
Aborting, use --force to override.


istioctl manifest generate | kubectl delete -f -
kubectl delete ns istio-system --grace-period=0 --force


istioctl install --set profile=demo -y

Istio core encountered an error: failed to update resource with server-side apply for obj ServiceAccount/istio-system/istio-reader-service-account: serviceaccounts "istio-reader-service-account" is forbidden: unable to create new content in namespace istio-system because it is being terminated


% kubectl edit ns istio-system  
namespace/istio-system edited

删除
spec:
  finalizers:
  - kubernetes


% kubectl get ns istio-system -o yaml > istio-system.yaml 


 % kubectl apply -f istio-system.yaml                      
Warning: Detected changes to resource istio-system which is currently being deleted.
namespace/istio-system configured


 % kubectl delete -f istio-system.yaml
namespace "istio-system" deleted

https://support.huaweicloud.com/intl/en-us/cce_faq/cce_faq_00277.html

https://www.ibm.com/docs/en/cloud-private/3.2.0?topic=console-namespace-is-stuck-in-terminating-state



 % kubectl get ns istio-system -o json > istio-system.json 

 % kubectl proxy      

 curl -k -H "Content-Type: application/json" -X PUT --data-binary @./source/k8s_learn/Istio/istio-system.json  http://127.0.0.1:8001/api/v1/namespaces/istio-system/finalize






Istio/istio-1.12.1/manifests/charts/istio-control/istio-discovery/files/gen-istio.yaml
image: "gcr.io/istio-testing/pilot:latest"

Istio/istio-1.12.1/manifests/charts/gateways/istio-ingress/templates/deployment.yaml
Istio/istio-1.12.1/manifests/charts/gateways/istio-egress/templates/deployment.yaml

 image: "{{ .Values.global.hub }}/{{ .Values.global.proxy.image | default "proxyv2" }}:{{ .Values.global.tag }}"


 proxyv2

  querycapistio/proxyv2:1.12.1

image: "gcr.io/istio-testing/pilot:latest" 

  istiojfh/pilot:1.8.1-arm64




