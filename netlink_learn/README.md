https://blog.csdn.net/ooobama/article/details/98496887

运行过Docker Hub的Docker镜像的话，会发现其中一些容器时需要挂载/var/run/docker.sock文件。这个文件是什么呢？为什么有些容器需要使用它？简单地说，它是Docker守护进程(Docker daemon)默认监听的Unix域套接字(Unix domain socket)，容器中的进程可以通过它与Docker守护进程进行通信。


不妨看一下 [Portainer](http://portainer.io/)，它提供了图形化界面用于管理Docker主机和Swarm集群。如果使用Portainer管理本地Docker主机的话，需要绑定/var/run/docker.sock:

1
docker run -d -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock portainer/portainer

安装Docker之后，Docker守护进程会监听Unix域套接字：/var/run/docker.sock。这一点可以通过Docker daemon的配置选项看出来(在ubuntu上执行cat /etc/default/docker )：

1
-H unix:///var/run/docker.sock
注: 监听网络TCP套接字或者其他套接字需要配置相应的-H选项。

Docker engine API v1.27(最新版)定义的所有HTTP接口都可以通过/var/run/docker.sock调用。

运行容器

使用Portainer的UI，可以轻松创建容器。实际上，HTTP请求是通过docker.sock发送给Docker守护进程的。可以通过curl创建容器来说明这一点。使用HTTP接口运行容器需要两个步骤，先创建容器，然后启动容器。

1. 创建nginx容器

curl命令通过Unix套接字发送{“Image”:”nginx”}到Docker守护进程的/containers/create接口，这个将会基于Nginx镜像创建容器并返回容器的ID。

1
curl -XPOST --unix-socket /var/run/docker.sock -d ‘{“Image”:”nginx”}’ -H ‘Content-Type: application/json’ http://localhost/containers/create
输出返回了容器ID:

1
{“Id”:”fcb65c6147efb862d5ea3a2ef20e793c52f0fafa3eb04e4292cb4784c5777d65",”Warnings”:null}
2. 启动nginx容器

使用返回的容器ID，调用/containers/<ID>/start接口，即可启动新创建的容器。

1
curl -XPOST --unix-socket /var/run/docker.sock http://localhost/containers/fcb6...7d65/start
查看已启动的容器:

docker ls
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
fcb65c6147ef nginx “nginx -g ‘daemon …” 5 minutes ago Up 5 seconds 80/tcp, 443/tcp ecstatic_kirch
...
可知，使用docker.sock运行容器其实非常简单。

Docker守护进程的事件流

Docker的API提供了/events接口，可以用于获取Docker守护进程产生的所有事件流。负载均衡组件(load balancer)组件可以通过它获

https://www.cnblogs.com/fundebug/p/6723464.html

https://docs.docker.com/engine/api/v1.27/#operation/ContainerAttachWebsocket

https://stackoverflow.com/questions/38532483/where-is-var-lib-docker-on-mac-os-x


Is actually at:

/Volumes/{DISKNAME}/var/run/docker.sock
If you run this, it should prove it, as long as your running VirtualBox 5.2.8 or later and the share for /Volumes is setup to be auto-mounted and permanent AND you generated the default docker-machine while on that version of Virtualbox:

#!/bin/bash
docker run -d --restart unless-stopped -p 9000:9000 \
-v /var/run/docker.sock:/var/run/docker.sock portainer/portainer \
--no-auth

