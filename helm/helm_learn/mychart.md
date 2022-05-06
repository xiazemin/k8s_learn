% helm create mychart 
Creating mychart


% cd mychart
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


Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes


Templates 目录下 YAML 文件模板(go template语法)填充的值默认都是在 values.yaml 里定义的，比如在 deployment.yaml 中定义的容器镜像:

% cat mychart/values.yaml|grep repository
  repository: nginx

  以上变量值是在 create chart 的时候就自动生成的默认值，你可以根据实际情况进行修改。

编写应用的介绍信息
打开 mychart/Chart.yaml：

编写应用具体部署信息
编辑 mychart/values.yaml，它默认会在 Kubernetes 部署一个 Nginx。下面是 mychart 应用的 values.yaml 文件的内容：


 helm lint mychart/     #检查依赖和模版配置是否正确


 % helm lint mychart
==> Linting mychart
[ERROR] Chart.yaml: unable to parse YAML
        error converting YAML to JSON: yaml: line 27: could not find expected ':'
[ERROR] templates/: cannot load Chart.yaml: error converting YAML to JSON: yaml: line 27: could not find expected ':'
[ERROR] : unable to load chart
        cannot load Chart.yaml: error converting YAML to JSON: yaml: line 27: could not find expected ':'

Error: 1 chart(s) linted, 1 chart(s) failed

% helm lint mychart
==> Linting mychart
[INFO] Chart.yaml: icon is recommended

1 chart(s) linted, 0 chart(s) failed


% helm package mychart  //将应用打包
Successfully packaged chart and saved it to: /Users/xiazemin/go/src/github.com/xiazemin/helm_learn/mychart-0.1.0.tgz


mychart 目录会被打包为一个 mychart-0.1.0.tgz 格式的压缩包，该压缩包会被放到当前目录下。

如果你想看到更详细的输出，可以加上 --debug 参数来查看打包的输出。

离线部署
注意: ~/.kube/config不存在的情况下要用 helm --kubeconfig 指定配置文件


# 方式一
$ helm install demo-test ./mychart

# 可根据不同的配置来install，默认是values.yaml
# helm install demo-test ./mychart -f ./mychart/values-prod.yaml

# 方式二
$ helm install demo-test ./mychart-0.1.0.tgz

$ helm list

# 升级
# $ helm upgrade demo-test ./mychart-0.2.0.tgz

$ helm uninstall demo-test


% helm install demo-test ./mychart
NAME: demo-test
LAST DEPLOYED: Sun Apr  4 16:53:14 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=mychart,app.kubernetes.io/instance=demo-test" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT

% helm list
NAME            NAMESPACE       REVISION        UPDATED                                 STATUS          CHART         APP VERSION
demo-test       default         1               2021-04-04 16:53:14.890414 +0800 CST    deployed        mychart-0.1.0 1.16.0     

将应用发布到 Repository
harbor1.6+ 支持存储 helm charts，这里使用 helm 安装 harbor

这里为了简化测试操作，我关闭了数据卷的挂载并使用的是 NodePort 方式进行访问。

https://github.com/goharbor/harbor/

https://blog.csdn.net/zyl290760647/article/details/83752877

按时安装harbor ,这里为了简化测试操作，我关闭了数据卷的挂载并使用的是 NodePort 方式进行访问。

helm -n test install harbor goharbor/harbor --set persistence.enabled=false --set expose.type=nodePort --set expose.tls.enabled=false --set externalURL=http://192.168.10.196:30002

https://blog.csdn.net/zhwyj1019/article/details/97490222


$ helm repo add goharbor https://helm.goharbor.io

$ helm repo update

# 查看harbor chart的各个版本
$ helm search repo harbor -l

# --version选择chart的版本
$ helm install harbor goharbor/harbor --set persistence.enabled=false --set expose.type=nodePort --set expose.tls.enabled=false --set externalURL=http://192.168.4.82:30002$ kubectl get svc


% helm repo add goharbor https://helm.goharbor.io
"goharbor" has been added to your repositories

% helm search repo harbor -l
NAME            CHART VERSION   APP VERSION     DESCRIPTION                                       
bitnami/harbor  9.8.3           2.2.1           Harbor is an an open source trusted cloud nativ...


% helm install harbor goharbor/harbor --set persistence.enabled=false --set expose.type=nodePort --set expose.tls.enabled=false --set externalURL=http://localhost:30002   

Error: cannot re-use a name that is still in use


https://www.cnblogs.com/qiyebao/p/13389621.html



