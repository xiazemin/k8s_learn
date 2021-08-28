https://www.cnblogs.com/rabbix/p/13941210.html

https://github.com/openresty/lua-resty-redis/blob/master/lib/resty/redis.lua

https://stackoverflow.com/questions/64815229/nginx-controller-kubernetes-need-to-change-host-header-within-ingress/64826818


kubectl apply -f ingress.yaml 

Error from server (InternalError): error when creating "ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": an error on the server ("") has prevented the request from succeeding


kubectl delete -f ../ingress-controller.yaml
kubectl apply -f ingress.yaml 
kubectl apply -f ../ingress-controller.yaml


%curl 127.0.0.1/apple
/apple%                                               

% curl -H "token:12344" 127.0.0.1/apple
<h1>系统开小差了</h1>

https://www.cnblogs.com/rongfengliang/p/10060467.html

https://www.huaweicloud.com/articles/11962254.html

kubectl delete -f ../../ingress/ingress-controller.yaml 
kubectl apply -f ../../ingress/access_by_lua_block/ingress.yaml
kubectl apply -f ../../ingress/ingress-controller.yaml 

curl -H "token:12344" 127.0.0.1/apple

https://www.cnblogs.com/linuxk/p/9706720.html

kubectl exec -n default -it ingress-nginx-controller-57648496fc-dbv26 -- /bin/bash

ERROR: Unable to lock database: Permission denied

https://www.coder.work/article/42725

https://stackoom.com/question/3Qqc3
https://www.devops.buzz/public/kubernetes/deployments/deployment-examples

https://www.huaweicloud.com/articles/11962254.html

% kubectl exec -n default -it ingress-nginx-controller-57648496fc-dbv26 --user=root -- /bin/bash
error: auth info "root" does not exist

ping: permission denied (are you root?)
docker exec -u 0 -it [container_id] sh

https://github.com/smebberson/docker-alpine/issues/65


% docker ps 
CONTAINER ID   IMAGE                    COMMAND                  CREATED          STATUS                PORTS  
f5a55b5d52c5   fa59b6fe51ab             "/usr/bin/dumb-init …"   18 minutes ago   Up 18 minutes                                                                                                k8s_controller_ingress-nginx-controller-57648496fc-dbv26_default_88ff9054-7fdd-4b9d-9804-c0ab453db81b_0
a2386b37775a   k8s.gcr.io/pause:3.4.1   "/pause"                 18 minutes ago   Up 18 minutes                                                                                                k8s_POD_ingress-nginx-controller-57648496fc-dbv26_default_88ff9054-7fdd-4b9d-9804-c0ab453db81b_0


docker exec -u 0 -it f5a55b5d52c5 /bin/bash
bash-5.1# apk add redis
fetch https://dl-cdn.alpinelinux.org/alpine/v3.13/main/x86_64/APKINDEX.tar.gz


274903771976:error:1416F086:SSL routines:tls_process_server_certificate:certificate verify failed:ssl/statem/statem_clnt.c:1913:
ERROR: https://dl-cdn.alpinelinux.org/alpine/v3.13/community: Permission denied

https://github.com/microsoft/vscode-remote-release/issues/5052

docker exec -it --user=0 --privileged f5a55b5d52c5 bash


 kubectl cp /opt/homebrew/bin/redis-cli ingress-nginx-controller-57648496fc-dbv26:/etc/nginx

 kubectl cp ./redis-cli ingress-nginx-controller-57648496fc-dbv26:/etc/nginx

 # ./redis-cli 
golang连接redis
PONG <nil>

% curl -H "token:12344" 127.0.0.1/apple
<h1>系统开小差了</h1>redis could not be resolved (3: Host not found)
https://atbug.com/nginx-dynamic-domain-parse-in-kubernetes/
service的短名称是解析不了的, 需要使用serviceName.namespace.svc.clusterName.

redis => redis.default.svc.docker-desktop

https://blog.csdn.net/u013928820/article/details/106164730/

通过访问服务名，依靠DNS解析，就是同一个namespace里的pod可以直接通过serviceName:port。不同的namespace里面，可以通过serviceName.namespace.svc.cluster.local
redis.default.svc.cluster.local

% curl -H "token:12344" 127.0.0.1/apple       
/apple%                                             

终于成功了

% ab -n 1000 -c 100 -H "token:12344" 127.0.0.1/apple   
Percentage of the requests served within a certain time (ms)
  50%     54
  66%     77
  75%     90
  80%     99
  90%    122
  95%    145
  98%    189
  99%    212
 100%    258 (longest request)

 % ab -n 1000 -c 200 -H "token:12344" 127.0.0.1/apple       
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
apr_socket_recv: Connection reset by peer (54)
Total of 1 requests completed
# ./redis-cli
bash: ./redis-cli: No such file or directory

# ls -al
lrwxrwxrwx    1 www-data www-data        35 Aug 28 11:23 redis-cli -> ../Cellar/redis/6.2.1/bin/redis-cli

 ls -al /opt/homebrew/Cellar/redis/6.2.1/bin/redis-cli
-r-xr-xr-x  1 xiazemin  wheel  304848  4 25 21:18 /opt/homebrew/Cellar/redis/6.2.1/bin/redis-cli


kubectl cp /opt/homebrew/Cellar/redis/6.2.1/bin/redis-cli ingress-nginx-controller-57648496fc-dbv26:/etc/nginx

# ./redis-cli 
bash: ./redis-cli: cannot execute binary file: Exec format error

https://github.com/crypt1d/redi.sh


https://github.com/SomajitDey/redis-client



# uname -a
Linux ingress-nginx-controller-57648496fc-dbv26 5.10.47-linuxkit #1 SMP PREEMPT Sat Jul 3 21:50:16 UTC 2021 x86_64 Linux

https://redis.io/clients#bash

% bash ../../ingress/access_by_lua_block/redis.cli.sh  -H "127.0.0.1" -P 30379 -p 123456 -g "key"
../../ingress/access_by_lua_block/redis.cli.sh: line 170: exec: {FD}: not found


# bash ./redis.sh -H 127.0.0.1 -P 6379 -p 123456 -g "key"
./redis.sh: connect: Connection refused

# bash ./redis.sh -H 127.0.0.1 -P 30379 -p 123456 -g "key"
./redis.sh: connect: Connection refused


# bash ./redis.sh -H "redis" -P 6379 -p 123456 -g "key"

https://learnku.com/articles/47718
