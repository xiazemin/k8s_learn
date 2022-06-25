% helm init --client-only --stable-repo-url https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts/
Error: unknown command "init" for "helm"

Did you mean this?
	lint

Run 'helm --help' for usage.

https://blog.csdn.net/weixin_44782815/article/details/108830849

https://stackoverflow.com/questions/63164194/what-helm-repository-should-i-add

helm repo add stable https://kubernetes-charts.storage.googleapis.com

https://github.com/helm/charts

https://www.cnblogs.com/LiuQizhong/p/11883307.html

https://www.pianshen.com/article/8650309196/

从kubernetes 1.6开始默认开启RBAC。这是Kubernetes安全性/企业可用的一个重要特性。但是在RBAC开启的情况下管理及配置Tiller变的非常复杂。为了简化helm的尝试成本我们给出了一个不需要关注安全规则的默认配置。但是，这会导致一些用户意外获得了他们并不需要的权限。并且，管理员/SRE需要学习很多额外的知识才能将Tiller部署的到关注安全的生产环境的多租户K8S集群中并使其正常工作。

在了解了社区成员通常的使用场景后，我们发现Tiller的发布管理系统不需要依靠集群内的Operator来维护状态或充当Helm发布信息的中央枢纽。相反，我们可以简单地从Kubernetes API服务器中获取信息，渲染Charts客户端，并在Kubernetes中存储安装记录。

移除掉Tiller大大简化了hlem的安全模型实现方式。Helm3现在可以支持所有的kubernetes认证及鉴权等全部安全特性。Helm和本地的kubeconfig flie中的配置使用一致的权限。管理员可以按照自己认为合适的粒度来管理用户权限。

https://blog.csdn.net/zzh_gaoxingjiuhao/article/details/104182596

https://github.com/rancher/helm3-charts

https://www.linuxba.com/archives/8340

https://www.cnblogs.com/qiyebao/p/13389621.html