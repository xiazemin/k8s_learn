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

