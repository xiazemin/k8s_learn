Kong的核心实体概念。

Client：指下游客户端。
Service：指上游服务。
Route：定义匹配客户端请求的规则，每个路由都与一个服务相关联，而服务可能有多个与之相关联的路由。
Plugin：它是在代理生命周期中运行的业务逻辑，插件可以生效于全局或者特定的路由和服务。
Consumer：消费者表示服务的消费者或者使用者，可能是一个用户，也可能是一个应用。


https://blog.csdn.net/twingao/article/details/104073211

代理协议
Kong支持http/https、tcp/tls和grpc/grpcs协议的代理。

路由匹配

请求必须包括所有配置属性（and）

请求中必须至少匹配一个属性值（or）

{undefined
“hosts”: [“example.com”, “foo-service.com”],
“paths”: ["/foo", “/bar”],
“methods”: [“GET”]
}

https://blog.csdn.net/twingao/article/details/104073231


Kong提供了两种负载均衡方案：

基于DNS的负载均衡
当使用基于DNS的负载均衡时，上游服务的注册是在Kong之外完成，而Kong只接收来自DNS服务器的负载均衡域名解析。使用包含主机名（而不是IP地址）URL定义的每个API将自动使用基于DNS的负载均衡。

Ring-Balancer环形均衡器
使用环形均衡器时，上游服务的添加和删除将由Kong处理，不需要进行DNS域名解析（当然如果target配置为域名也需要解析域名，但不需要DNS实现负载均衡）。Kong将扮演服务注册中心的角色。使用upstream和target配置服务。

加权循环算法：默认情况下环形均衡器将使用加权循环的方案。
散列算法：以none、consumer、IP或者header为输入的散列算法。none就是加权循环算法。

https://blog.csdn.net/twingao/article/details/104073247


Kong无数据库介绍
Kong在1.1.0版本（2019-03-27）开始支持无数据库模式，只需要将kong.conf中的配置项database = off即可启用无数据库方式。

无数据库采用声明方式定义实体，所有的实体都配置在其中。一旦该文件加载到Kong中，它将替换整个配置。当需要增量更改时，将对声明性配置文件进行更改，然后将其全部重新加载。
由于没有中心数据库进行协调，多个Kong节点是完全彼此独立。
由于配置实体的唯一方式是通过声明性配置，管理API的CRUD操作只支持只读方法，即GET方法，不支持POST、DELETE等方法。
无数据库加载实体声明的方式：

在kong.conf配置文件中配置declarative_config = /etc/kong/kong.yaml
调用管理API，一次性加载声明配置。
由于没有数据库，有些Kong的插件将不再能够使用，有些插件能够部分使用。

兼容的插件

bot-detection
correlation-id
cors
datadog
file-log
http-log
tcp-log
udp-log
syslog
ip-restriction
prometheus
zipkin
request-transformer
response-transformer
request-termination

部分兼容的插件

acl
basic-auth
hmac-auth
jwt
key-auth
rate-limiting
response-ratelimiting
key-auth

不兼容的插件

oauth2

https://blog.csdn.net/twingao/article/details/104073276


