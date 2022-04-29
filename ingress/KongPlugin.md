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

创建KongPlugin资源后，还需要将该资源与Kubernetes中的Ingress、Service或KongConsumer资源关联
https://blog.csdn.net/twingao/article/details/104073337

Kong官方预置大量的插件，可以在Kong Hub查看官方插件，其中有三类插件，官方开发的开源插件；第三方开发的开源插件；适用于Kong Enterprise的收费插件。其中官方开发的开源插件有35个。

https://blog.csdn.net/twingao/article/details/104073425


Kong开源了大量的开源插件，当这些开源插件不能满足我们的需求，就需要修改这些开源插件或者自定义插件。Kong提供了方便地自定义插件机制，用户可以开发自己的定制插件。自定义插件可以和Kong进行深层次的集成，如使用数据库表结构，或者扩展Admin API。

如果插件实现了所有可选模块，则其目录结构如下所示：

complete-plugin
├── api.lua
├── daos.lua
├── handler.lua
├── migrations
│   ├── init.lua
│   └── 000_base_complete_plugin.lua
└── schema.lua

插件的每个模块对应不同的功能，其中handler.lua和schema.lua是两个必须模块。

Module Name	Required	Description
api.lua	No	插件需要向Admin API暴露接口时使用。
daos.lua	No	数据层相关，当插件需要访问数据库是配置。
handler.lua	Yes	插件的主要逻辑，Kong在不同阶段执行其对应的handler。
migrations/*.lua	No	插件依赖的数据库表结构，启用了daos.lua时需要定义。
schema.lua	Yes	插件的配置参数定义，主要用于Kong参数验证。

https://blog.csdn.net/twingao/article/details/104073442

https://github.com/Kong/kong/tree/master/kong/plugins/key-auth

如何在Kong for Kubernetes加载自定义插件。
https://blog.csdn.net/twingao/article/details/104073451

jianshu.com/p/ca70c26a853d

https://www.jianshu.com/p/27dda1e4b4ce