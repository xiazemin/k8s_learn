Ingress：对集群中服务的外部访问进行管理的 API 对象，可以理解为kubernetes中一种定义从集群外部到集群内服务的HTTP/HTTPS 路由的资源类型；
Ingress Controller：负责通过负载均衡器（如nginx/haproxy）来实现 Ingress中定义的路由，通常以pod形式运行；
IngressClass：一种定义某个具体Ingress Controller的资源类型，kubernetes 1.18-1.21中支持，Ingress通过IngressClass来指定使用哪一种Ingrss Controller。


为什么需要使用Ingress暴露服务？Kubernetes还支持通过NodePort、LoadBalancer向集群外部暴露内部服务，与这两种方式相比，ingress可以通过单一IP为多个服务提供访问，可以通过主机名和路径将流量路由到特定服务，使得服务访问更加灵活，同时还可以用来提供https服务。

Nginx Ingress Controller即是Nginx公司使用nginx实现的一种Ingress Controller。（注意区别于kubernetes社区开发的ingress-nginx，两者均是Ingress Controller，功能略有差异，Nginx Ingress Controller额外定义了两种资源类型VirtualServer和VirtualServerRoute，实现了一些Ingress不支持的的功能。


https://github.com/nginxinc/kubernetes-ingress/
https://github.com/kubernetes/ingress-nginx

nginx ingress controller 程序已经将ingress定义的路由规则转化为相应的nginx配置文件，因此我们访问nginx ingress controller才能实现相应的路由功能。从上面的nginx配置文件也可看出， nginx ingress 控制器并非直接将请求转发给相应服务，而是通过http请求头部信息确定客户端尝试访问哪个服务，通过与该服务关联的Endpoint对象查看pod IP，并将客户端的请求转发给其中的一个pod。

https://blog.csdn.net/sinat_32582203/article/details/119939537
https://www.nginx.com/blog/guide-to-choosing-ingress-controller-part-4-nginx-ingress-controller-options/

nginx.ingress.kubernetes.io/backend-protocol: "HTTPS"
nginx.ingress.kubernetes.io/backend-protocol: "GRPC"

https://blog.csdn.net/longtds/article/details/116189814


kong如何更新nginx相关的配置

https://blog.csdn.net/oqqYuan1234567890/article/details/105087348
https://github.com/Kong/kubernetes-ingress-controller

https://it.baiked.com/kubernetes/2477.html