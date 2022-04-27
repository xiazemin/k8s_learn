Nginx、OpenRestry、Kong这三个项目紧密相连： Nginx是模块化设计的反向代理软件，C语言开发； OpenResty是以Nginx为核心的Web开发平台，可以解析执行Lua脚本（OpenResty与Lua的关系，类似于Jvm与Java，不过Java可以做的事情太多了，OpenResty主要用来做Web、API等）； Kong是一个OpenResty应用，是一个api gateway，具有API管理和请求代理的功能。

https://blog.csdn.net/qq_24509229/article/details/116171770

events和http指令位于main context，server位于http context，location位于server context：
OpenResty是一个集成了Nginx、LuaJIT和其它很多moudels的平台，用来托管完整的web应用——包含业务逻辑，而不单纯是静态文件服务器

Kong是一个OpenResty应用，用来管理api。
https://luarocks.org/

https://github.com/Kong/kong.git
认证插件：

Basic Auth
HMAC Auth
JWT Auth
Key Auth
LDAP Auth
OAuth 2.0 Auth
安全插件：

Bot Detection (机器人检测)
CORS (跨域请求)
IP Restriction (IP限制)
流控插件：

ACL (访问控制）
Rate Limiting （限速）
Request Size Limiting
Request Termination
Response Rate Limiting
微服务插件：

AWS Lambda
Azure Functions
Apache OpenWhisk
Serverless Functions
分析和监控插件：

Datadog
Prometheus
Zipkin
内容修改插件(Transformations)：

Correlation ID
Request Transformer
Response Transformer
日志插件：

File Log
HTTP Log
Loggly
StatsD
Syslog
TCP Log
UDP Log

Kong是一个Api网关，也是一个特性更丰富的反向代理，既然它有代理流量的功能，那么能不能直接成为Kubernetes的流量入口？使Kubernetes内部的服务都通过Kong发布。

Kong实现了一个Kubernetes Ingress Controller来做这件事。在Kubernetes中部署kong的方法见Kong CE or EE on Kubernetes。

https://blog.csdn.net/qq_24509229/article/details/116171770

https://docs.konghq.com/hub/

https://www.lijiaocn.com/%E9%A1%B9%E7%9B%AE/2018/09/30/integrate-kubernetes-with-kong.html
https://docs.konghq.com/install/kubernetes/


KongPlugin资源的定义：

apiVersion: configuration.konghq.com/v1
kind: KongPlugin
metadata:
  name: <名称>
  namespace: <命名空间>
  labels:
    global: "true"   #可选，如果设置，该插件为全局插件，应该使用双引号将true引起来
disabled: <boolean>  #可选，将该插件禁用
config:              #该插件的配置，插件的说明文档中能够查到响应的配置
    key: value
plugin: <插件名称>    #如key-auth，rate-limiting等


https://blog.csdn.net/twingao/article/details/104073337
https://blog.csdn.net/twingao/article/details/104073289
kongconsumers：Kong的用户，给不同的API用户提供不同的消费者身份。
kongcredentials：Kong用户的认证凭证。
kongingresses：定义代理行为规则，是对Ingress的补充配置。
kongplugins：插件的配置。
https://blog.csdn.net/twingao/article/details/104073276
https://blog.csdn.net/twingao/article/details/104073247
https://blog.csdn.net/twingao/article/details/104073231
https://blog.csdn.net/twingao/article/details/104073211
https://blog.csdn.net/twingao/article/details/104073159
https://blog.csdn.net/twingao/article/details/104073112
https://blog.csdn.net/twingao/article/details/104073090
https://blog.csdn.net/twingao/article/details/104073068

Kong的由来
Nginx是模块化设计的反向代理软件，C语言开发。
Lua是一个小巧的脚本语言，C语言开发，Lua脚本可以容易的被C/C++代码调用，也可以反过来调用C/C++的函数。
通过lua-nginx-module模块将luaJIT解释器集成进Nginx，使得Nginx可以执行Lua脚本，灵活但不宜用。
OpenResty是一个基于Nginx与Lua的高性能Web平台，其中内部集成了大量精良的Lua库，第三方模块以及大多数的依赖项。用于方便地搭建能够处理超高并发，扩展性极高的动态Web应用、Web服务和动态网关。
在Nginx上直接使用Lua开发难度比较大，OpenResty集成大量Lua库，第三方模块和所需的依赖项，这样使用Lua开发功能变得更加容易。
Kong是一个基于OpenResty开发的应用，可以认为是一个Web Server，更进一步是一个API Gateway，具有API管理和请求代理的功能。
可以举个例子：Nginx是OS，Lua是编程语言，OpenResty是开发平台或者框架，Kong是一个OpenResty之上的应用。

Kong特性
云原生(Cloud-Native)：Kong可以在Kubernetes或裸机上运行；
动态负载平衡(Dynamic Load Balancing)：跨多个上游服务的负载平衡业务。
基于哈希的负载平衡(Hash-based Load Balancing)：一致的散列/会话亲和的负载平衡。
断路器(Circuit-Breaker)：智能跟踪不健康的上游服务。
健康检查(Health Checks)：主动和被动监控您的上游服务。
服务发现(Service Discovery)：解决如Consul等第三方DNS解析器的SRV记录。
无服务器(Serverless)：从Kong中直接调用和保证AWS或OpenWhisk函数安全。
WebSockets：通过WebSockets与上游服务进行通信。
OAuth2.0：轻松的向API中添加OAuth2.0认证。
日志记录(Logging)：通过HTTP、TCP、UDP记录请求或者相应的日志，存储在磁盘中。
安全(Security)：ACL，Bot检测，IPs白名单/黑名单等。
系统日志(Syslog)：记录信息到系统日志。
SSL：为基础服务或API设置特定的SSL证书。
监视(Monitoring)：能够实时对关键负载和性能指标进行监控。
转发代理(Forward Proxy)：使端口连接到中间透明的HTTP代理。
认证(Authentications)：支持HMAC，JWT和BASIC方式进行认证等等。
速率限制(Rate-limiting)：基于多个变量的阻塞和节流请求。
转换(Transformations)：添加、删除或操作HTTP请求和响应。
缓存(Caching)：在代理层进行缓存和服务响应。
命令行工具(CLI)：能够通过命令行控制Kong的集群。
REST API：可以通过REST API灵活的操作Kong。
GEO复制：在不同的区域，配置总是最新的。
故障检测与恢复(Failure Detection & Recovery)：如果Cassandra节点失效，Kong并不会受影响。
群集(Clustering)：所有的Kong节点会自动加入群集，并更新各个节点上的配置。
可扩展性(Scalability)：通过添加节点，实现水平缩放。
性能(Performance)：通过缩放和使用Nigix，Kong能够轻松处理负载。
插件(Plugins)：基于插件的可扩展体系结构，能够方便的向Kong和API添加功能。

Kong以前只支持数据库方式，配置数据（路由规则，Service等）存放在数据库中，Kong各节点从数据库中同步配置数据以保持配置数据一致性。从1.1.0版本开始，Kong开始支持无数据方式，无数据库方式的配置数据存放在声明式配置文件中，Kong各个节点完全独立。

Kong有三种配置管理方式：

管理API方式：Kong的标准管理方式，Kong的配置数据都是通过调用Kong提供的Rest API进行管理，最终配置都存放在数据库中，Kong集群实例从数据库读取配置。
声明式配置文件Declarative Config：在无数据库方式下，Kong提供一种声明式配置文件，配置数据在文件中声明，然后通过文件或者Rest API的方式一次性加载到Kong中。
Ingress Controller：Kong在Kubernetes下支持Ingress Controller，Kong还定义了多个自定义资源CRDs，配置数据通过这些CRDs进行配置。