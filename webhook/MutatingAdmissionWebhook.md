https://blog.hdls.me/15564491070483.html

Kubernetes 集群管理员可以使用 webhooks 来创建额外的资源更改及验证准入插件，这些准入插件可以通过 apiserver 的准入链来工作，而不需要重新编译 apiserver。这使得开发者可以对于很多动作都可以自定义准入逻辑，比如对任何资源的创建、更新、删除，给开发者提供了很大的自由和灵活度。可以使用的应用数量巨大。一些常见的使用常见包括：

在创建资源之前做些更改。Istio 是个非常典型的例子，在目标 pods 中注入 Envoy sidecar 容器，来实现流量管理和规则执行。
自动配置 StorageClass。监听 PersistentVolumeClaim 资源，并按照事先定好的规则自动的为之增添对应的 StorageClass。使用者无需关心 StorageClass 的创建。
验证复杂的自定义资源。确保只有其被定义后且所有的依赖项都创建好并可用，自定义资源才可以创建。
namespace 的限制。在多租户系统中，避免资源在预先保留的 namespace 中被创建。

Webhooks 和 Initializers 之间的区别是什么呢？

Webhooks 可以应用于更多操作，包括对于资源 "增删改" 的 "mutate" 和 "admit"；然而 Initializers 不可以对 "删" 资源进行 "admit"。
Webhooks 在创建资源前不允许查询；然而 Initializers 可以监听未初始化的资源，通过参数 ?includeUninitialized=true 来实现。
由于 Initializers 会把 "预创建" 状态也持久化到 etcd，因此会引入高延迟且给 etcd 带来负担，尤其在 apiserver 升级或失败时；然而 Webhooks 消耗的内存和计算资源更少。
Webhooks 比 Initializers 对失败的保障更强大。Webhooks 的配置中可以配置失败策略，用以避免资源在创建的时候被 hang 住。然而 Initializers 在尝试创建资源的时候可能会 block 住所有的资源。

https://medium.com/ibm-cloud/kubernetes-initializers-deep-dive-and-tutorial-3bc416e4e13e

https://github.com/morvencao/kube-mutating-webhook-tutorial

https://github.com/kubernetes/kubernetes/tree/release-1.9/test/images/webhook

