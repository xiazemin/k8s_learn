% go build -o sample-controller .


../../../../go/pkg/mod/sigs.k8s.io/json@v0.0.0-20211020170558-c049b76a60c6/internal/golang/encoding/json/encode.go:1255:18: sf.IsExported undefined (type reflect.StructField has no field or method IsExported)


https://githubmemory.com/repo/clarketm/json
https://githubmemory.com/repo/tinygo-org/tinygo/issues/2054


%  go build -o sample-controller .
% git checkout release-1.21 

问题解决


 % ./sample-controller -kubeconfig=$HOME/.kube/config
I1103 10:50:25.426210   36841 controller.go:115] Setting up event handlers
I1103 10:50:25.426376   36841 controller.go:156] Starting Foo controller
I1103 10:50:25.426381   36841 controller.go:159] Waiting for informer caches to sync

在mac 宿主机运行失败，mac 有一个虚拟机


https://blog.csdn.net/zhonglinzhang/article/details/86553744


http://www.asznl.com/post/43

https://www.qikqiak.com/k8strain/operator/crd/

https://cloud.tencent.com/developer/article/1717417







