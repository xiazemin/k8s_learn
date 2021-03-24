docker pull golang:alpine
docker pull ubuntu  x
https://www.jianshu.com/p/0aa535e681f5?t=1500627922686
https://hub.docker.com/_/centos?tab=tags&page=1&ordering=last_updated

% docker pull centos:centos7.9.2009

https://blog.csdn.net/u010771890/article/details/102727814

% go build -o chat/c chat/client/main.go 
% go build -o chat/s chat/server/main.go

 ./chat/s :12048

./chat/c 127.0.0.1:12048

docker build -t chatroomserver:v1 .

 => ERROR [build-env 4/4] RUN go build -v -o /go/src/app/server server/main.go                                                                           0.5s
------
 > [build-env 4/4] RUN go build -v -o /go/src/app/server server/main.go:
#10 0.433 package server/main.go is not in GOROOT (/usr/local/go/src/server/main.go)
------
executor failed running [/bin/sh -c go build -v -o /go/src/app/server server/main.go]: exit code: 1

https://zhuanlan.zhihu.com/p/269115851

 % cd server/
 % go mod init server
go: creating new go.mod: module server

docker run -it -p 0.0.0.0:10248:10248 chatroomserver:v1  //默认监听10248

docker run -it -p 0.0.0.0:10248:10248 chatroomserver:v1 0.0.0.0:yourport  //也可以自己设置监听port

docker tag chatroomserver:v1 yourresposity/chatroomserver:v1  //将image重新命名，增加用户名作为前缀

docker push yourresposity/chatroomserver //上传入仓库

kubectl apply -f chatroomserverservice.yaml  

kubectl apply -f chatroomserver.yaml

https://www.jianshu.com/p/41dc6aada2ef

https://blog.csdn.net/weixin_41806245/article/details/93745532

https://stackoverflow.com/questions/64286761/package-models-is-not-in-goroot-when-build-docker

https://stackoverflow.com/questions/65095510/docker-build-from-parent-directory-is-giving-error-package-is-not-in-goroot


 => ERROR [build-env 4/4] RUN go build -v -o /go/src/app/server main.go                                                                                  0.5s
------
 > [build-env 4/4] RUN go build -v -o /go/src/app/server main.go:
#10 0.476 $GOPATH/go.mod exists but should not

 % rm go.mod