https://www.cxyzjd.com/article/weixin_42072280/112857621

从最早的参考 K8S 官方的控制器代码并手动复制，到用 client-gen 生成框架代码，到现在使用 kubebuilder ，已经非常方便我们完成 CRD/Controller，甚至 Operator 的开发（当然 Operator 的开发也有专用的 operator-sdk开源框架）。

go mod init my.domain

## 这两个复杂的命令可以通过kubebuilder --help查看
# 初始化
kubebuilder init --domain example.com --license apache2 --owner "The Kubernetes authors" 
# 创建CRD api
kubebuilder create api --group webapp --version v1 --kind Frigate 
# 安装CRD
make install 
# 启动controller(本地)
make run

https://juejin.cn/post/6844903735815061511

https://www.cnblogs.com/alisystemsoftware/p/11580202.html 
