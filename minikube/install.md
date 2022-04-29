wget https://github.com/kubernetes/minikube/releases/download/v1.7.3/minikube-linux-amd64

./minikube-linux-amd64 start --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --vm-driver=none --cpus=2 --memory=4096 
✨  Using the none driver based on user configuration
⚠️  The 'none' driver does not respect the --cpus flag
⚠️  The 'none' driver does not respect the --memory flag
✅  正在使用镜像存储库 registry.cn-hangzhou.aliyuncs.com/google_containers
🤹  Running on localhost (CPUs=2, Memory=1826MB, Disk=40252MB) ...
ℹ️   OS release is CentOS Linux 8 (Core)
🐳  正在 Docker 20.10.5 中准备 Kubernetes v1.17.3…
    ▪ kubelet.resolv-conf=/run/systemd/resolve/resolv.conf
💾  正在下载 kubectl v1.17.3
💾  正在下载 kubelet v1.17.3
💾  正在下载 kubeadm v1.17.3
^@^@🚀  正在启动 Kubernetes ...
^@^@

^@^@^@^@🌟  Enabling addons: default-storageclass, storage-provisioner
🤹  开始配置本地主机环境...

⚠️  “none”驱动程序提供有限的隔离功能，并且可能会降低系统安全性和可靠性。
⚠️  如需了解详情，请参阅：
👉  https://minikube.sigs.k8s.io/docs/reference/drivers/none/

⚠️  kubectl 和 minikube 配置将存储在 /root 中
⚠️  如需以您自己的用户身份使用 kubectl 或 minikube 命令，您可能需要重新定位该命令。例如，如需覆盖您的自定义设置，请运行：

    ▪ sudo mv /root/.kube /root/.minikube $HOME
    ▪ sudo chown -R $USER $HOME/.kube $HOME/.minikube

💡  此操作还可通过设置环境变量 CHANGE_MINIKUBE_NONE_USER=true 自动完成
⌛  等待集群上线...
^@🏄  完成！kubectl 已经配置至 "minikube"
💡  为获得最佳结果，请安装 kubectl：https://kubernetes.io/docs/tasks/tools/install-kubectl/

https://github.com/kubernetes/minikube

https://www.cnblogs.com/harmful-chan/p/12731014.html?ivk_sa=1024320u

https://kubernetes.io/docs/tasks/tools/

curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl 

chmod +x ./kubectl
mv ./kubectl /usr/bin/kubectl


]# kubectl get ns
NAME              STATUS   AGE
default           Active   7m42s
kube-node-lease   Active   7m44s
kube-public       Active   7m44s
kube-system       Active   7m44s


# kubectl get nodes
NAME             STATUS   ROLES    AGE     VERSION
vm-4-15-centos   Ready    master   7m50s   v1.17.3

 mv software/minikube-linux-amd64 /usr/local/bin/minikube

 https://www.cnblogs.com/skywalk-cnb/p/14663722.html
 