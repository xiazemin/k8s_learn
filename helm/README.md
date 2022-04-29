helm chart 是一种描述如何部署应用到 kubernetes 中的文档格式。
helm 项目提供了命令行工具 helm 来进行部署包的管理，并且支持接入 chart 仓库

helm repo add <repo-name> <repo-url>来添加开发环境的 chart 仓库。
helm search repo <repo-name>可以浏览当前仓库中有的 chart。

%helm search repo kong
NAME       	CHART VERSION	APP VERSION	DESCRIPTION
apphub/kong	0.27.0       	1.3        	The Cloud-Native Ingress and Service Mesh for A...
stable/kong	0.36.7       	1.4        	DEPRECATED The Cloud-Native Ingress and API-man...


helm create <charts-name>可以创建一个初始 chart
% helm create mychart
Creating mychart
% tree
.
|____Chart.yaml
|____charts
|____.helmignore
|____templates
| |____deployment.yaml
| |____NOTES.txt
| |____ingress.yaml
| |____tests
| | |____test-connection.yaml
| |____service.yaml
| |____hpa.yaml
| |____serviceaccount.yaml
| |_____helpers.tpl
|____values.yaml

说明：
charts目录： [可选]，该目录中放置当前Chart依赖的其它Chart
Chart.yaml：用于描述Chart的基本信息，包括名称版本等
templates目录： 部署文件模版目录
values.yaml文件： 为templates目录中的yaml文件提供变量


除了 Chart.yaml、values.yaml 是必须要有的，其他都是可选的。
实际使用时，我们无疑是需要 templates/ 里的模板的，比如 deployment.yaml、service.yaml，模板从 values.yaml 中获取值来渲染得到最终的部署相关资源描述文件。

Chart.yaml
Chart.yaml 中主要是放一些概要信息，比如 chart 的名称、版本、维护者、依赖（即子 chart）。
在 Helm 3 中，不再需要 requirements.yaml 来描述 dependencies 。
在声明了 dependencies 的 chart 中执行helm dependency update或helm dependency build将会自动生成一个 Chart.lock 文件，且如果设置了依赖项的 repository，会到仓库中查找并打包为 .tgz 文件下载到 charts/ 路径下。

values.yaml
values.yaml 中放置一些需要传递给模板的值，在模板文件中你将可以通过 {undefined{.Values.xxx.xxx}} 来进行引用。

覆盖子 chart 配置
一个可用服务通常涉及到多个应用，比如业务应用会依赖 redis 之类中间件，这就是 Chart.yaml 中设置 dependencies 的用意。当我们依赖于别人提供的 chart，但在部署时又需要对其中的一些配置进行调整，这时候就可以通过在父 Chart.yaml 中设置能对应到子 chart 中 values.yaml 中的配置的配置项，来达到覆盖子 chart 中配置值的目的。

在 mychart/charts/mysubchart/values.yaml 中存在一个配置项 dessert。
dessert: cake
可以通过在 mychart/values.yaml 中设置另一个值来覆盖。指明子 chart 名称，然后是 values.yaml 中的路径即可。
mysubchart:
  dessert: ice-cream

https://blog.csdn.net/alading2009/article/details/119820869

https://artifacthub.io/packages/search?ts_query_web=nginx&sort=relevance&page=1


1语法检查
helm lint --strict /root/yml/my-hello-world/
==> Linting /root/yml/my-hello-world/
[INFO] Chart.yaml: icon is recommended
1 chart(s) linted, no failures

2 查看最终yaml文件
helm install --dry-run --debug /root/yml/my-hello-world

3,打包
# helm package ./my-hello-world/
Successfully packaged chart and saved it to: /tmp/my-hello-world-0.1.0.tgz

4,利用chart包安装
# helm install  my-hello-world -n test  /root/yml/my-hello-world/my-hello-world-0.1.0.tgz

说明：
my-hello-world为创建的Release名
test 为安装到的namespace。

5,利用chart目录安装
语法：
helm install release_name -n namespace_name chart_dir

6,利用helm仓库安装
语法：
helm install release_name -n namespace_name repot_name/chart_name

https://artifacthub.io/

7,使用helm get mainfest查看安装好的chart


https://www.hellodemos.com/hello-helm-intermediate/helm-intermediate-getting-started-zh-cn.html

https://help.aliyun.com/document_detail/128090.html

https://cloud.tencent.com/developer/article/1898297


https://github.com/artifacthub/hub

3.更新仓库信息
helm repo update

4.删除仓库
使用 remove 命令可以删除一个或多个仓库:

helm repo remove test-tepo

使用 helm 查看已安装的 chart 列表:
# helm list

需要在更新的时候在命令行显式地指定变量名称和值来进行更新，命令如下:
helm upgrade --set service.externalPort=8090 my-tomcat test-repo/tomcat

卸载 charts
卸载 chart 可以使用 uninstall 命令:
helm uninstall my-tomcat

打包 chart
假如我们修改 chart 后，需要将其进行打包，可以使用如下命令:

helm package /opt/helm/work/tomcat

拉取 chart
我们可以直接将仓库中的 chart 拉取到本地并进行一些修改，拉取的命令如下:

helm pull test-repo/tomcat --untar --untardir /opt/helm/work
上面的命令中指定了 --untar 来指定下载 chart 包后进行解压缩， --untardir 命令可以指定解压缩的目

https://zhuanlan.zhihu.com/p/431501827
https://www.jianshu.com/p/b809cffd9c36

Kubernetes (简称 k8s) 是一个能够部署和管理容器的平台。然而，在 k8s 里还没有抽象到“应用”这一层概念。一个应用往往由多个 k8s 资源 (Deployment、Service、ConfigMap）组成。所以，我们需要一个工具在 k8s 之上来部署和管理一个应用所包含的资源（K8s API Resource），这就是 Helm 所做的事情。


https://developer.aliyun.com/article/709186
https://community.anaplan.com/t5/App-Hub/ct-p/APPHUB


helm三大概念
chart：chart就是helm package，包含了一个k8s app应用运行起来的所有要素，比如service, deployment, configmap, serviceaccount, rbac, 等，这些要素都是以template文件的形式存在，再结合values文件，最终渲染出能够被k8s执行的yaml文件；
repository：仓库是charts的集合，方便进行分享和分发。下面是官网仓库和阿里云仓库的地址，大家可以进去看看，感受一下；
https://artifacthub.io/
https://developer.aliyun.com/hub
release：release是helm chart在kubernetes的一个运行实例，你可以用不同的release name多次安装同一个chart，比如：当集群中需要多个redis实例，你可以使用不同的配置文件安装redis chart。

