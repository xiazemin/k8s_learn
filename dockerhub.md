首先看一下自己登陆没有：
vim ~/.docker/config.json
没有登陆的话使用如下指令登陆
% docker login
Authenticating with existing credentials...
Login Succeeded

https://blog.csdn.net/weixin_42072543/article/details/89288270

 % docker push apple:v1
The push refers to repository [docker.io/library/apple]
5f70bf18a086: Preparing
37d1f76ddcd2: Preparing
denied: requested access to the resource is denied

% docker tag apple:v1 xiazemin/k8s_learm:0.1.0
 % docker push xiazemin/k8s_learm:0.1.0
The push refers to repository [docker.io/xiazemin/k8s_learm]
5f70bf18a086: Pushing  1.024kB
37d1f76ddcd2: Pushing [==================================================>]  6.091MB

https://hub.docker.com/repository/docker/xiazemin/k8s_learm/tags?page=1&ordering=last_updated