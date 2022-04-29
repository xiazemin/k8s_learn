helm常用命令
在helm的时候，你可以使用第三方开发的chart，也可以自己开发chart，以下是两种情况下使用的常见命令。更为详细的命令，可以安装好helm之后，使用helm help来查看，或查看官方文档。

使用第三方开发的chart
部署前

repo: add, list, remove, update, and index chart repositories
search: search for a keyword in charts
部署后

install: install a chart
list: list releases
status: display the status of the named release
upgrade: upgrade a release
rollback: roll back a release to a previous revision
uninstall: uninstall a release
自己开发chart
lint: examine a chart for possible issues
package: package a chart directory into a chart archive
push: push helm chart to chartmuseum
chart push: push helm chart to OCI repository
chart
chart可以说是helm里面最重要的概念了，关于chart也有很多内容需要掌握，在这里做一个列举。

chart开发：主要是指利用模板技术开发一个chart，会在后面做详细介绍；
chart hooks：在chart的生命周期中，提供一些hooks，方便进行一些前置或后置操作
在chart安装前，创建应用需要的Secret
在chart安装前，备份数据库
在chart卸载后，做一些清理工作
chart test：当你install了一个chart后，如何知道这个release是否运行正常呢？chart test提供了一种测试的方式，来验证你的应用是否正常运行，比如：
校验mysql应用能够正常连接并接受请求
校验services能够正常做load balance
library chart：一种以library形式存在的chart，可以在application chart之间进行共享避免重复逻辑；类似于编程语言中的public library；
chart校验：基于PKI、GnuPG等技术，对helm package进行签名，保证传输或发布过程中的安全性；
OCI（Open Container Initiative，容器发型规范）支持：helm 3引入（EXPERIMENTAL），能够将chart推送到支持OCI的仓库中，比如harbor, nexus，等，比如，将chart保存到harbor中：
保存chart：helm chart save kubeedge/ some-harbor-repo/kubeedge-cloud-chart:1.0.0
登录repo：helm registry login https://some-harbor-repo
推送chart：helm chart push some-harbor-repo/kubeedge-cloud-chart:1.0.0
高级特性
post rendering: 提供在helm install之前对manifests进行操作、配置的一种机制；一般结合kustomize使用。比如：
在install时插入sidecar，从而为deployment增加功能
不修改原chart的情况下，更改manifests的配置


添加仓库并找到redis chart
添加仓库：helm repo add stable https://charts.helm.sh/stable
查看已经添加的仓库：helm repo list
搜索仓库有哪些chart：helm search repo stable
更新仓库列表到本地：helm repo update
搜索redis：helm search repo redis
查看redis chart详情：helm show chart stable/redis
查看redis values（values：相当于chart的配置文件）：helm show values stable/redis

https://zhuanlan.zhihu.com/p/350328164
https://github.com/zkf1317/cloud-native-tutorial

helm repo add stable https://charts.helm.sh/stable

 % helm repo add elastic https://helm.elastic.co
"elastic" has been added to your repositories


https://artifacthub.io/packages/helm/microfunctions/kong

helm repo add bitnami https://charts.bitnami.com/bitnami

% helm search repo kong
NAME            CHART VERSION   APP VERSION     DESCRIPTION                                       
apphub/kong     0.27.0          1.3             The Cloud-Native Ingress and Service Mesh for A...
bitnami/kong    6.1.18          2.8.1           Kong is an open source Microservice API gateway...
stable/kong     0.36.7          1.4             DEPRECATED The Cloud-Native Ingress and API-man...

helm install kong bitnami/kong

% helm install kong bitnami/kong
Error: INSTALLATION FAILED: cannot re-use a name that is still in use


 % kubectl get pods |grep kong
kong-6486bffc75-ksb22   2/2     Running     0               75s
kong-6486bffc75-snxfz   2/2     Running     1 (2m36s ago)   17m
kong-migrate--1-h969v   0/1     Completed   0               17m
kong-postgresql-0       1/1     Running     0               17m

% helm  list
NAME    NAMESPACE       REVISION        UPDATED                                    STATUS  CHART           APP VERSION
kong    default         1               2022-04-28 21:48:24.723812 +0800 CST       failed  kong-6.1.18     2.8.1      


% helm uninstall kong           
release "kong" uninstalled

% helm install kong bitnami/kong
