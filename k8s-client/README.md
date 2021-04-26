vi ~/.kube/config

current-context: docker-desktop

 kubectl version
Client Version: version.Info{Major:"1", Minor:"19", GitVersion:"v1.19.7", GitCommit:"1dd5338295409edcfff11505e7bb246f0d325d15", GitTreeState:"clean", BuildDate:"2021-01-13T13:23:52Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"19", GitVersion:"v1.19.7", GitCommit:"1dd5338295409edcfff11505e7bb246f0d325d15", GitTreeState:"clean", BuildDate:"2021-01-13T13:15:20Z", GoVersion:"go1.15.5", Compiler:"gc", Platform:"linux/arm64"}


go mod init k8s_learn
go: creating new go.mod: module k8s_learn

https://github.com/kubernetes/client-go

go get -u -v k8s.io/client-go@kubernetes-1.19.7

//https://www.cnblogs.com/LC161616/p/12046863.html

https://www.cnblogs.com/jokerjason/p/11898428.html

% go run k8s-client/exp2/Nodes.go
[0] docker-desktop

https://blog.csdn.net/niyuelin1990/article/details/79076805
https://blog.csdn.net/qq_37950254/article/details/89603207

https://zhuanlan.zhihu.com/p/165970638

https://blog.csdn.net/huwh_/article/details/78821805

https://www.jianshu.com/p/48946a317d81



% go mod tidy       
go: k8s.io/api@v0.19.7: Get "https://goproxy.cn/k8s.io/api/@v/v0.19.7.mod": dial tcp: lookup goproxy.cn: no such host

% export GOPROXY=https://goproxy.io

% go mod tidy 

% go run k8s-client/exp5/listpods.go
kube-system
[10.1.0.85 10.1.0.84 192.168.65.4 192.168.65.4 192.168.65.4 192.168.65.4 192.168.65.4 10.1.0.86 10.1.0.87]
