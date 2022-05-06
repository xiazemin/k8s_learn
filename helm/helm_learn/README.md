https://blog.csdn.net/bbwangj/article/details/81087911

brew install helm

Chart Install 过程：

Helm从指定的目录或者tgz文件中解析出Chart结构信息
Helm将指定的Chart结构和Values信息通过gRPC传递给Tiller
Tiller根据Chart和Values生成一个Release
Tiller将Release发送给Kubernetes用于生成Release
Chart Update过程：

Helm从指定的目录或者tgz文件中解析出Chart结构信息
Helm将要更新的Release的名称和Chart结构，Values信息传递给Tiller
Tiller生成Release并更新指定名称的Release的History
Tiller将Release发送给Kubernetes用于更新Release
Chart Rollback过程：

Helm将要回滚的Release的名称传递给Tiller
Tiller根据Release的名称查找History
Tiller从History中获取上一个Release
Tiller将上一个Release发送给Kubernetes用于替换当前Release


curl https://raw.githubusercontent.com/helm/helm/master/scripts/get > get_helm.sh

Tiller 是以 Deployment 方式部署在 Kubernetes 集群中的，只需使用以下指令便可简单的完成安装。
$ helm init

由于 Helm 默认会去 storage.googleapis.com 拉取镜像

helm init --client-only --stable-repo-url https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts/
helm repo add incubator https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
helm repo update


# 创建服务端
helm init --service-account tiller --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.9.1  --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
 
# 创建TLS认证服务端，参考地址：https://github.com/gjmzj/kubeasz/blob/master/docs/guide/helm.md
helm init --service-account tiller --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.9.1 --tiller-tls-cert /etc/kubernetes/ssl/tiller001.pem --tiller-tls-key /etc/kubernetes/ssl/tiller001-key.pem --tls-ca-cert /etc/kubernetes/ssl/ca.pem --tiller-namespace kube-system --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
在 Kubernetes 中安装 Tiller 服务，因为官方的镜像因为某些原因无法拉取，使用-i指定自己的镜像，可选镜像：registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.9.1（阿里云），该镜像的版本与helm客户端的版本相同，使用helm version可查看helm客户端版本。

因为 Helm 的服务端 Tiller 是一个部署在 Kubernetes 中 Kube-System Namespace 下 的 Deployment，它会去连接 Kube-Api 在 Kubernetes 里创建和删除应用。

