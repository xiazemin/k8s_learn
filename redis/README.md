https://www.cnblogs.com/kevingrace/p/14412134.html

https://www.jianshu.com/p/a34790c730cf

https://segmentfault.com/a/1190000039196137

https://www.gotkx.com/?p=76



https://kubernetes.io/zh/docs/tutorials/configuration/configure-redis-using-configmap/
https://www.cnblogs.com/zisefeizhu/p/14282299.html
https://juejin.cn/post/6844903806719754254

http://www.mydlq.club/article/76/
https://www.jianshu.com/p/0fc0e1b869c6
https://www.codenong.com/cs109763447/

https://blog.csdn.net/xujiamin0022016/article/details/109763447
https://www.gl.sh.cn/2021/03/26/k8s_bu_shu_dan_jie_dian_redis.html

https://kuboard.cn/learning/k8s-practice/ocp/redis.html

https://blog.51cto.com/u_15100527/2616832

 % kubectl apply -f redis-config.yaml
configmap/redis-config created
% kubectl apply -f redis-deploy.yaml 
deployment.apps/redis created

kubectl apply -f redis-service.yaml 
service/redis created

 % docker run -it  --rm  redis:latest /bin/bash
root@948f69ef03a3:/data# 

 % kubectl get svc -o wide |grep redis 
redis                                NodePort       10.108.154.209   <none>        6379:30379/TCP               8s     app=redis


 %  redis-cli -h 127.0.0.1 -p 30379 -a 123456
Warning: Using a password with '-a' or '-u' option on the command line interface may not be safe.
127.0.0.1:30379> 

kubectl describe ing  ingress-with-auth 

kubectl logs ingress-nginx-controller-57648496fc-qn2lz  -n <namespace> 
https://www.cxyzjd.com/article/weixin_33788244/90227960

2021/08/28 07:08:44 [emerg] 75#75: io_setup() failed (38: Function not implemented)
192.168.65.3 - - [28/Aug/2021:07:08:56 +0000] "GET /apple HTTP/1.1" 500 39 "-" "curl/7.64.1" 91 0.009 [default-apple-service-5678] [] - - - - 1ebb5fd894d7bb64ff7c835c4012358a

