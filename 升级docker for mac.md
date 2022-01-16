1,升级到最新版本的docker for mac k8s启动失败解决方案：
% cd k8s-docker-desktop-for-mac
% sh load_images.sh 
重启

https://github.com/AliyunContainerService/k8s-for-docker-desktop
当前 master 分支已经在 Docker for Mac/Windows 4.3.0 (包含 Docker CE 20.10.11 和 Kubernetes 1.22.4) 版本测试通过

本机安装的是 Docker CE 20.10.12 k8s 1.22.5

% git checkout -b v1.22.5
切换到一个新分支 'v1.22.5'

k8s-docker-desktop-for-mac/images
v1.22.4 => v1.22.5

 % sh load_images.sh 

 Error response from daemon: manifest for gotok8s/kube-proxy:v1.22.5 not found: manifest unknown: manifest unknown
Error response from daemon: No such image: gotok8s/kube-proxy:v1.22.5
  
  https://registry.hub.docker.com/r/gotok8s/kube-scheduler/tags
  => v1.23.1

 % sh load_images.sh 
 重启

 https://www.wenjiangs.com/doc/o06xwu0g20
/Applications/Docker.app/Contents/MacOS/com.docker.diagnose gather 
time="2022-01-16T11:58:02+08:00" level=warning msg="/sysrq failed: Post \"http://unix/sysrq\": context deadline exceeded" type=unixsock


https://docs.docker.com/desktop/mac/troubleshoot/#self-diagnose-tool
 /Applications/Docker.app/Contents/MacOS/com.docker.diagnose check

% pred='process matches ".*(ocker|vpnkit).*" || (process in {"taskgated-helper", "launchservicesd", "kernel"} && eventMessage contains[c] "docker")'

% /usr/bin/log stream --style syslog --level=debug --color=always --predicate "$pred"

https://registry.hub.docker.com/r/gotok8s/coredns/tags
gotok8s/coredns:v1.8.4  =》 gotok8s/coredns:v1.8.6

gotok8s/pause:3.5 =》 gotok8s/pause:3.6

gotok8s/etcd:3.5.0-0  =》 gotok8s/etcd:3.5.1-0 



cat ~/Library/Containers/com.docker.docker/Data/log/vm/dockerd.log

6:06:58:47.903][I] time="2022-01-16T06:58:47.903436000Z" level=warning msg="Error getting v2 registry: Get \"https://k8s.gcr.io/v2/\": dial tcp 74.125.204.82:443: i/o timeout"
[016:06:58:47.904][I] time="2022-01-16T06:58:47.903676300Z" level=info msg="Attempting next endpoint for pull after error: Get \"https://k8s.gcr.io/v2/\": dial tcp 74.125.204.82:443: i/o timeout"
[016:06:58:47.914][I] time="2022-01-16T06:58:47.913605200Z" level=error msg="Handler for POST /v1.40/images/create returned error: Get \"https://k8s.gcr.io/v2/\": dial tcp 74.125.204.82:443: i/o timeout"
[016:06:59:04.954][I] time="2022-01-16T06:59:04.953483700Z" level=warning msg="Error getting v2 registry: Get \"https://k8s.gcr.io/v2/\": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"
[016:06:59:04.954][I] time="2022-01-16T06:59:04.953603900Z" level=info msg="Attempting next endpoint for pull after error: Get \"https://k8s.gcr.io/v2/\": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"
[016:06:59:04.962][I] time="2022-01-16T06:59:04.961834500Z" level=error msg="Handler for POST /v1.40/images/create returned error: Get \"https://k8s.gcr.io/v2/\": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"



cat ~/Library/Containers/com.docker.docker/Data/log/vm/containerd.log
rm ~/Library/Containers/com.docker.docker/Data/log/vm/dockerd.log

https://docs.docker.com/config/daemon/#read-the-logs

rm ~/Library/Containers/com.docker.docker/Data/log/vm/*



[016:07:09:53.926][I] time="2022-01-16T07:09:53.924353448Z" level=warning msg="Error getting v2 registry: Get \"https://k8s.gcr.io/v2/\": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"


修改成 k8s.gcr.io/kube-proxy:v1.22.5=gotok8s/kube-proxy:v1.23.1


进去发现镜像有问题
goroutine 148 [syscall]:

syscall.Syscall6(0xe8, 0xe, 0xc000e2fbec, 0x7, 0xffffffffffffffff, 0x0, 0x0)

/usr/local/go/src/syscall/asm_linux_amd64.s:43 +0x5

k8s.io/kubernetes/vendor/golang.org/x/sys/unix.EpollWait(0x9f1cf0abbceca3d5, {0xc000e2fbec, 0xa8aac008f7ef639c, 0x73679bfc3eb14da9}, 0x208fb1b158c13294)

/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/vendor/golang.org/x/sys/unix/zsyscall_linux_amd64.go:77 +0x58

k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify.(*fdPoller).wait(0xc000d8e000)

/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify/inotify_poller.go:86 +0x7d

k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify.(*Watcher).readEvents(0xc00067e1e0)

/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify/inotify.go:192 +0x2b0

created by k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify.NewWatcher

/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/vendor/github.com/fsnotify/fsnotify/inotify.go:59 +0x1c7


改成
k8s.gcr.io/kube-scheduler:v1.22.5=gotok8s/kube-scheduler:v1.22.4
% sh load_images.sh


F0116 07:46:25.108668       1 controllermanager.go:222] error building controller context: failed to wait for apiserver being healthy: timed out waiting for the condition: failed to get apiserver /healthz status: Get "https://192.168.65.4:6443/healthz?timeout=32s": dial tcp 192.168.65.4:6443: connect: connection refused

goroutine 1 [running]:

k8s.io/kubernetes/vendor/k8s.io/klog/v2.stacks(0xc000130001, 0xc0001b6500, 0x142, 0x251)

/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes/vendor/k8s.io/klog/v2/klog.go:1026 +0xb9

k8s.io/kubernetes/vendor/k8s.io/klog/v2.(*loggingT).output(0x7515640, 0xc000000003, 0x0, 0x0, 0xc0005bbc00, 0x0, 0x6108721, 0x14, 0xde, 0x0)

/workspace/src/k8s


清理了，重新部署，ok