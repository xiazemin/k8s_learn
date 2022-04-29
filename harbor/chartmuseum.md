一、启用harbor的Chart仓库服务

# ./install.sh --with-chartmuseum

启用后，默认创建的项目就带有helm charts功能了。

插件地址: https://github.com/chartmuseum/helm-push 
安装命令：
helm plugin install https://github.com/chartmuseum/helm-push

也可以下载离线包
解压到对应用户的环境目录下 /root/.local/share/helm/plugins/helm-push/

三、添加repo

helm repo add  --username admin --password xxxx  harbor https://harbor.xxx.com/chartrepo/library
/chartrepo/是必须带的

推送
helm push mysql-1.4.0.tgz --username=admin --password=xxxxx https://harbor.xxx.com/chartrepo/library


安装
helm install web --version 1.4.0 myrepo/demo

如果下载失败先更新repo
helm repo update

https://www.cnblogs.com/opsdemo/p/15039252.html
https://blog.csdn.net/catoop/article/details/121848279


https://github.com/goharbor/harbor-helm

Chartmuseum		
chartmuseum.enabled	Enable chartmusuem to store chart	true
默认是允许的


% helm package .
Successfully packaged chart and saved it to: /Users/xiazemin/source/k8s_learn/helm/mychart/mychart-0.1.0.tgz

helm repo index .
  index.yaml  

索引文件是一个名为 index.yaml 的 Yaml 文件。它包含有关 Chart 包的一些元数据，包括 Chart.yaml 文件的内容。有效的 Chart Repository 必须具有索引文件。索引文件包含有关 Chart Repository 中每个 Chart 的信息。可以通过 helm repo index 命令将本地目录下的 Chart 生成索引文件。


使用 gh-pages 分支作为 Chart Repository，例如：https://USERNAME.github.io/REPONAME。可以访问 Milvus Chart 来演示 Chart Repository , 效果如 https://milvus-io.github.io/milvus-helm/


添加到 Helm Hub 仓库
要添加至 Helm Hub , 需 fork 该 helm/hub 项目，并根据以下两点信息来创建新的 Pull Request：

在 config/repo-values.yaml 添加你的 Chart Repository，并使用短名称和该 Repository 的基本URL。
在 repos.yaml 文件中，添加 Repository 的联系信息。当需要联系这些 Repository 管理者时，此联系信息是必要的。如果改 Repository 由团队或公司所有，我们建议你使用邮件列表。

https://blog.csdn.net/jeffzhesi/article/details/106565173


helm repo add  --username admin --password Harbor12345 harbor https://124.220.185.35:30003/chartrepo/library

Error: looks like "https://124.220.185.35:30003/chartrepo/library" is not a valid chart repository or cannot be reached: Get "https://124.220.185.35:30003/chartrepo/library/index.yaml": x509: cannot validate certificate for 124.220.185.35 because it doesn't contain any IP SANs

helm cm-push --ca-file /etc/docker/certs.d/core.harbor.domain/ca.crt  mychart-0.1.0.tgz harbor

https://blog.csdn.net/qq_37837432/article/details/124130202


https://blog.csdn.net/liuzehn/article/details/120805900


1.使用FQDN创建ssl证书：

sudo openssl req -subj '/CN=124.220.185.35:30003/' -x509 -days 3650 -batch -nodes -newkey rsa:2048 -keyout private/helm.key -out private/helm.crt

Generating a RSA private key
...................................................................................................................................+++++
...................................+++++
writing new private key to 'private/helm.key'
-----


https://www.thinbug.com/q/42116783

https://www.it610.com/article/1188193745498120192.htm


helm repo add harbor https://124.220.185.35:30003/chartrepo/library --cert-file private/helm.crt  --username admin --password Harbor12345 

Error: looks like "https://124.220.185.35:30003/chartrepo/library" is not a valid chart repository or cannot be reached: Get "https://124.220.185.35:30003/chartrepo/library/index.yaml": dial tcp 124.220.185.35:30003: connect: network is unreachable

https://blog.csdn.net/began2014/article/details/100950034/

% helm repo index .  --url https://124.220.185.35:30003 

% helm repo add harbor https://124.220.185.35:30003/chartrepo/library --username admin --password Harbor12345 --insecure-skip-tls-verify 
"harbor" has been added to your repositories

https://blog.csdn.net/weixin_42758299/article/details/122965560


helm push mychart-0.1.0.tgz harbor

Error: this feature has been marked as experimental and is not enabled by default. Please set HELM_EXPERIMENTAL_OCI=1 in your environment to use this feature


helm push，将本地charts加入harbor charts

一种方式是直接上传charts文件夹
helm push seldon-mab harbor-test-helm
seldon-mab是charts目录
harbor-test-helm是harbor charts repo名称。
另一种是将charts package文件包push
helm push seldon-core-operator-1.5.1.tgz harbor-test-helm

https://www.jianshu.com/p/d5aabe1cd9e4


% export HELM_EXPERIMENTAL_OCI=1

% helm push mychart-0.1.0.tgz harbor
Error: scheme prefix missing from remote (e.g. "oci://")

https://blog.csdn.net/lzyjzb/article/details/110382677


https://stackoverflow.com/questions/71145976/how-to-push-a-helm-chart-to-harbor-using-helm-cli-v3-7-2-and-harbor-2-4-0-d4affc

% helm cm-push mychart-0.1.0.tgz harbor
Error: unknown command "cm-push" for "helm"
Run 'helm --help' for usage.


https://github.com/chartmuseum/helm-push

helm plugin install https://github.com/chartmuseum/helm-push
Error: Unable to get repository: Cloning into '/Users/xiazemin/Library/Caches/helm/plugins/https-github.com-chartmuseum-helm-push'...
fatal: unable to access 'https://github.com/chartmuseum/helm-push/': Recv failure: Connection reset by peer
: exit status 128

Downloading and installing helm-push v0.10.2 ...
https://github.com/chartmuseum/helm-push/releases/download/v0.10.2/helm-push_0.10.2_darwin_amd64.tar.gz
Installed plugin: cm-push

helm cm-push --help


% helm cm-push mychart-0.1.0.tgz harbor
Pushing mychart-0.1.0.tgz to harbor...
Error: Post "https://124.220.185.35:30003/api/chartrepo/library/charts": x509: cannot validate certificate for 124.220.185.35 because it doesn't contain any IP SANs
Usage:
  helm cm-push [flags]

% helm cm-push mychart-0.1.0.tgz harbor  --insecure
Pushing mychart-0.1.0.tgz to harbor...
Done.

https://github.com/chartmuseum/helm-push/
