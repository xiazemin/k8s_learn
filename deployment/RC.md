https://blog.csdn.net/yjk13703623757/article/details/53746273/

https://blog.elkfun.com/?p=1665

https://www.qieseo.com/168006.html

https://www.cnadn.net/post/2335.htm

RC只支持基于等式的selector（env=dev或environment!=qa）但Replica Set还支持新的、基于集合的selector（version in (v1.0, v2.0)或env notin (dev, qa)），这对复杂的运维管理带来很大方便。
使用Deployment升级Pod，只需要定义Pod的最终状态，k8s会为你执行必要的操作，虽然能够使用命令 # kubectl rolling-update 完成升级，但它是在客户端与服务端多次交互控制RC完成的，所以REST API中并没有rolling-update的接口，这为定制自己的管理系统带来了一些麻烦。
Deployment拥有更加灵活强大的升级、回滚功能。


 % kubectl apply -f deployment/ReplicationController.yaml
replicationcontroller/apple-rc created

% kubectl get rc  
NAME       DESIRED   CURRENT   READY   AGE
apple-rc   3         3         3       32s

 % kubectl get pod |grep apple-rctl
apple-rctl-2hdpt                            1/1     Running     0          29s
apple-rctl-wrmwp                            1/1     Running     0          15s