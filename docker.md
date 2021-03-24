
https://printempw.github.io/build-x86-docker-images-on-an-m1-macs/


启用实验性功能
Docker 的 buildx 还是实验性功能，需要在 Docker Desktop 设置中开启，具体位于 Preferences > Experimental Features > Enable CLI experimental features。

新建 builder 实例
Docker 默认的 builder 不支持同时指定多个架构，所以要新建一个：

docker buildx create --use --name m1_builder


构建多架构 Docker 镜像
使用 buildx 构建：

docker buildx build \
  --platform linux/amd64,linux/arm64 -t chatroomserver:v1 .

[+] Building 0.0s (0/0)
error: multiple platforms feature is currently not supported for docker driver. Please switch to a different driver (eg. "docker buildx create --use")

 % docker buildx create --use --name m1_builder
m1_builder

% docker buildx inspect --bootstrap
[+] Building 46.7s (0/1)
 => [internal] booting buildkit                                                                                                                                      46.7s
 => => pulling image moby/buildkit:buildx-stable-1                                                                                                                   44.9s
 => => creating container buildx_buildkit_m1_builder0                                                                                                                 1.8s
INFO: 2021/03/24 10:33:14 parsed scheme: "unix"
INFO: 2021/03/24 10:33:14 scheme "unix" not registered, fallback to default scheme
INFO: 2021/03/24 10:33:14 ccResolverWrapper: sending update to cc: {[{unix:///run/buildkit/buildkitd.sock  <nil> 0 <nil>}] <nil> <nil>}
INFO: 2021/03/24 10:33:14 ClientConn switching balancer to "pick_first"
INFO: 2021/03/24 10:33:14 Channel switches to new LB policy "pick_first"
[+] Building 46.8s (1/1) FINISHED
 => [internal] booting buildkit                                                                                                                                      46.8s
 => => pulling image moby/buildkit:buildx-stable-1                                                                                                                   44.9s
 => => creating container buildx_buildkit_m1_builder0                                                                                                                 1.9s
INFO: 2021/03/24 10:33:14 Subchannel Connectivity change to READY
INFO: 2021/03/24 10:33:14 pickfirstBalancer: HandleSubConnStateChange: 0xc000256830, {READY <nil>}
INFO: 2021/03/24 10:33:14 Channel Connectivity change to READY
INFO: 2021/03/24 10:33:14 parsed scheme: "unix"
INFO: 2021/03/24 10:33:14 scheme "unix" not registered, fallback to default scheme
INFO: 2021/03/24 10:33:14 ccResolverWrapper: sending update to cc: {[{unix:///run/buildkit/buildkitd.sock  <nil> 0 <nil>}] <nil> <nil>}
INFO: 2021/03/24 10:33:14 ClientConn switching balancer to "pick_first"
INFO: 2021/03/24 10:33:14 Channel switches to new LB policy "pick_first"
INFO: 2021/03/24 10:33:14 Subchannel Connectivity change to CONNECTING
INFO: 2021/03/24 10:33:14 pickfirstBalancer: HandleSubConnStateChange: 0xc00008cb40, {CONNECTING <nil>}
INFO: 2021/03/24 10:33:14 Channel Connectivity change to CONNECTING
INFO: 2021/03/24 10:33:14 Subchannel picks a new address "unix:///run/buildkit/buildkitd.sock" to connect
INFO: 2021/03/24 10:33:14 Subchannel Connectivity change to READY
INFO: 2021/03/24 10:33:14 pickfirstBalancer: HandleSubConnStateChange: 0xc00008cb40, {READY <nil>}
INFO: 2021/03/24 10:33:14 Channel Connectivity change to READY
Name:   m1_builder
Driver: docker-container

Nodes:
Name:      m1_builder0
Endpoint:  unix:///var/run/docker.sock
Status:    running
Platforms: linux/arm64, linux/amd64, linux/riscv64, linux/ppc64le, linux/s390x, linux/386, linux/arm/v7, linux/arm/v6


docker buildx build \
  --platform linux/amd64,linux/arm64 -t chatroomserver:v1 .


=> [linux/amd64 build-env 3/4] WORKDIR /go/src/app                                                                                                                   0.7s
 => ERROR [linux/amd64 build-env 4/4] RUN go build -v -o /go/src/app/server main.go                                                                                   1.2s
 => CANCELED [linux/arm64 build-env 2/4] ADD .  .                                                                                                                     0.0s
 => CACHED [linux/arm64 build-env 3/4] WORKDIR /go/src/app                                                                                                            0.0s
 => CACHED [linux/arm64 build-env 4/4] RUN go build -v -o /go/src/app/server main.go                                                                                  0.0s
 => CANCELED [linux/arm64 stage-1 2/3] COPY --from=build-env /go/src/app/server /app/server                                                                           0.0s
------
 > [linux/amd64 build-env 4/4] RUN go build -v -o /go/src/app/server main.go:
#12 0.754 $GOPATH/go.mod exists but should not


1. COPY指令

COPY指令能够将构建命令所在的主机本地的文件或目录，复制到镜像文件系统。

exec格式用法（推荐）：
COPY ["<src>",... "<dest>"]，推荐，特别适合路径中带有空格的情况

shell格式用法：
COPY <src>... <dest>

 

2. ADD指令
ADD指令不仅能够将构建命令所在的主机本地的文件或目录，而且能够将远程URL所对应的文件或目录，作为资源复制到镜像文件系统。
所以，可以认为ADD是增强版的COPY，支持将远程URL的资源加入到镜像的文件系统。
 

exec格式用法（推荐）：
ADD ["<src>",... "<dest>"]，特别适合路径中带有空格的情况

shell格式用法：
ADD <src>... <dest>



=> ERROR [linux/amd64 build-env 4/5] WORKDIR /go/src/app

 failed to solve: rpc error: code = Unknown desc = mkdir /tmp/buildkit-mount126733939/go/src/app: not a directory


把 WORKDIR /go/src/app 改成WORKDIR /go/src/app/ 注意要有/

error: failed to solve: rpc error: code = Unknown desc = executor failed running [/bin/sh -c go build -v -o /go/src/app/server server/main.go]: exit code: 1


RUN go build -v -o /go/src/app/server server/main.go 改成
UN go build -v -o /go/src/app/server main.go 注意这里需要用workdir 下的相对路径


https://segmentfault.com/a/1190000021166703?utm_source=tag-newest
https://github.com/docker/buildx

 => CANCELED [linux/arm64 build-env 2/5] ADD  ./server/main.go /go/src/app/  太慢了


 https://blog.csdn.net/nklinsirui/article/details/80967677

https://hub.docker.com/r/arm64v8/centos/
 docker pull arm64v8/centos:7
 7: Pulling from arm64v8/centos
6717b8ec66cd: Pull complete
Digest: sha256:43964203bf5d7fe38c6fca6166ac89e4c095e2b0c0a28f6c7c678a1348ddc7fa
Status: Downloaded newer image for arm64v8/centos:7
docker.io/arm64v8/centos:7

 % docker pull golang:alpine
alpine: Pulling from library/golang
159e5727ea61: Pull complete
89e1f0424fba: Pull complete
4e458f4c6c66: Pull complete
ec4e3c46f4b2: Pull complete
577341a4c2b5: Pull complete
Digest: sha256:49b4eac11640066bc72c74b70202478b7d431c7d8918e0973d6e4aeb8b3129d2
Status: Downloaded newer image for golang:alpine
docker.io/library/golang:alpine


https://blog.csdn.net/freewebsys/article/details/79224625


 [linux/arm64 build-env 1/5] FROM docker.io/library/golang:alpine@sha25  2032.2s
 => => sha256:e6948a1906ae692419137d9db26a066c5d9e4fd87

 https://hub.docker.com/_/golang

 https://hub.docker.com/r/arm64v8/alpine/

 docker pull arm64v8/alpine

 FROM golang:alpine AS build-env  => FROM arm64v8/alpine AS build-env

% docker buildx build \
  --platform linux/amd64,linux/arm64 -t apple:v1 .

 => ERROR [linux/arm64 build-env 5/5] RUN go build -v -o /go/src/app/server main.go                                                                                   1.0s
 => ERROR [linux/amd64 build-env 5/5] RUN go build -v -o /go/src/app/server main.go
 #16 0.881 /bin/sh: go: not found

 没有装golang
 docker pull golang:latest

  FROM golang:alpine AS build-env  =>  FROM golang:latest AS build-env

  https://github.com/docker-library/golang

docker pull golang:latest@e87ba3a72191
invalid reference format

ls ~/Library/Containers/com.docker.docker/Data/
 % docker pull golang:latest
 https://www.jianshu.com/p/f5505e64a4a4

 docker pull arm64v8/golang
 
