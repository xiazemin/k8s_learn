Sidecar 注入
简单来说，Sidecar 注入会将额外容器的配置添加到 Pod 模板中。Istio 服务网格目前所需的容器有：

istio-init init 容器用于设置 iptables 规则，以便将入站/出站流量通过 sidecar 代理。初始化容器与应用程序容器在以下方面有所不同：

它在启动应用容器之前运行，并一直运行直至完成。
如果有多个初始化容器，则每个容器都应在启动下一个容器之前成功完成。
因此，您可以看到，对于不需要成为实际应用容器一部分的设置或初始化作业来说，这种容器是多么的完美。在这种情况下，istio-init 就是这样做并设置了 iptables 规则。

istio-proxy 这个容器是真正的 sidecar 代理（基于 Envoy）。

手动注入
在手动注入方法中，可以使用 istioctl 修改容器模板并添加前面提到的两个容器的配置。不论是手动注入还是自动注入，Istio 都从 istio-sidecar-injector 和的 istio 两个 Configmap 对象中获取配置。

configmap 包含了 istio-init 初始化容器和 istio-proxy 代理容器的配置。该配置包括容器镜像的名称以及拦截模式，权限要求等参数。

从安全的角度来看，重要的是要注意 istio-init 需要 NET_ADMIN 权限来修改 pod 命名空间中的 iptables，如果 istio-proxy 是 TPROXY 模式，也需要这一权限。由于该仅限于 pod 的命名空间，因此应该没有问题。但是，我们注意到最近的 open-shift 版本可能会出现一些问题，因此需要一种解决方法。本文结尾处提到了一个这样的选择。

要修改当前的 Pod 模板以进行 sidecar 注入，您可以：

$ istioctl kube-inject -f demo-red.yaml | kubectl apply -f -

或者

要使用修改后的 Configmap 或本地 Configmap：

https://istio.io/latest/zh/blog/2019/data-plane-setup/