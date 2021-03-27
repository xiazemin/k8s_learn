https://kubernetes.github.io/ingress-nginx/user-guide/ingress-path-matching/

https://blog.51cto.com/weifan/2451983

https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/ingress-path-matching.md

https://www.qedev.com/linux/87370.html


 % docker build . -t apple:v2
[+] Building 9.4s (12/12) FINISHED

 % docker image rm  apple:v1
Untagged: apple:v1

% kubectl delete -f sample/apple.my-ns.yaml
pod "apple-app" deleted
service "apple-service" deleted

% kubectl apply -f sample/apple.my-ns.yaml


 % curl http://localhost/apple
/apple%

% curl http://localhost/apple/banana
/apple/banana%

% curl http://localhost/apple/banana/1
404 page not found

 % curl http://localhost/apple1/banana
 <head><title>404 Not Found</title></head>

 % cp sample/apple.my-ns.ingress.yaml sample/apple.my-ns.ingress.rewrite.yaml
% kubectl delete -f sample/apple.my-ns.ingress.yaml
ingress.networking.k8s.io "rancher-k8s-ingress" deleted

% kubectl apply -f sample/apple.my-ns.ingress.rewrite.yaml
ingress.networking.k8s.io/rancher-k8s-ingress created

没有成功

 % kubectl apply -f sample/apple.my-ns.ingress.regx.yaml
% curl http://localhost/apple1/banana
/apple1/banana



% curl http://localhost/apple1/banana
/apple1/banana%

% curl http://localhost/apple1/banana/abc
/apple1/banana/abc%

% curl http://localhost/apple1/banana/abc/def
/apple1/banana/abc/def%

% curl http://localhost/apple1/banana/abc/def/hijk
404 page not found

% curl http://localhost/apple1/apple/banana/abc
404 page not found

% curl http://localhost/apple/banana
/apple/banana
