https://zhuanlan.zhihu.com/p/69065079
docker hub是一个现成的从国内不需要代理就能访问，也能拉取gcr镜像的地方。
docker hub提供了auto build的功能，可以与一个代码仓库绑定，通过向代码仓库提交Dockerfile自动触发docker hub的build镜像操作。


只需要在Dockerfile中写一行FROM gcr.io/google_containers/...并提交的代码库即可。

1 创建 github repository
2,创建 docker hub repository
创建一个名叫kube-proxy-amd64的repo，repo名称将会是镜像名称；
切到Builds页签下绑定到github，配置auto build
填写auto build配置
save即可，然后到github提交代码触发自动构建；

TIPS

一个repo只能放一种镜像，可以有不同tag，如：
push username/nginx:1.14，username/kube-proxy-amd64:1.7两个镜像会在docker hub的username用户下产生nginx和kube-proxy-amd64两个repo；而push username/folder/nginx:1.14，加一个folder段会push失败。
自动构建规则可使用正则匹配

https://docs.docker.com/docker-hub/
需要升级到高级版才行
