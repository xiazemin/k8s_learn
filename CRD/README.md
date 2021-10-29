https://kubernetes.io/zh/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/
https://kubernetes.io/zh/docs/concepts/extend-kubernetes/api-extension/custom-resources/
https://www.jianshu.com/p/cc7eea6dd1fb
Kubernetes目前常使用CRD+Controller的方式扩展API，官方提供了CRD代码的自动生成器code-generator
https://juejin.cn/post/6844903783298777096

默认的生成脚本在code-generator下的generate-groups.sh，如我想生成./generate-groups.sh all github.com/nevermore/project/pkg/client github.com/nevermore/project/pkg/apis "foo:v1 bar:v1beta1"，则需要在$GOPATH/src/新建好相应的路径，mkdir -p $GOPATH/src/github.com/nevermore/project/pkg/client，mkdir -p $GOPATH/src/github.com/nevermore/project/pkg/apis/foo/v1，进入到v1路径下，新建三个文件touch doc.go types.go regsiter.go，修改每个文件开头为package v1；同理配置apis/bar/v1beta1相应的文件。
最终生成相应的clientset、listers、informers。

