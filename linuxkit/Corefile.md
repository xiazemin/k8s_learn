https://coredns.io/2017/07/23/corefile-explained/

To explain more, let take a look at this “Corefile”:

ZONE:[PORT] {
    [PLUGIN]...
}


The ZONE is root zone ., the PLUGIN is chaos. The chaos plugin does not have any properties, but it does take an argument: CoreDNS-001. This text is returned on a CH class query: dig CH txt version.bind @localhost

. {
   chaos CoreDNS-001
}
If CoreDNS can’t find a Corefile to load is loads the following builtin one that loads the whoami plugin:

. {
    whoami
}

https://coredns.io/manual/configuration/

https://github.com/coredns/coredns.io/tree/master/content/manual

docker pull coredns

docker start -p 53:53 -v /Corefile:/Corefile -v /zones:/zones -d --name coredns coredns

https://segmentfault.com/a/1190000022179401

https://docs.ucloud.cn/uk8s/administercluster/custom_dns_service

https://haojianxun.github.io/docs/%E5%AE%89%E8%A3%85%E9%85%8D%E7%BD%AE/coredns%E5%AE%89%E8%A3%85%E5%92%8C%E9%85%8D%E7%BD%AE/

如何配置外部dns
有些服务不在kubernetes内部，在内部环境内需要通过dns去访问,名称后缀为carey.com

carey:53 {
        errors
        cache 30
        proxy . 10.150.0.1
    }


https://www.jianshu.com/p/f7f45d435a80

/etc/coredns/Corefile

.:53 {
  # 绑定interface ip
  bind 127.0.0.1
  # 先走本机的hosts
  # https://coredns.io/plugins/hosts/
  hosts {
    # 自定义sms.service search.service 的解析
    # 因为解析的域名少我们这里直接用hosts插件即可完成需求
    # 如果有大量自定义域名解析那么建议用file插件使用 符合RFC 1035规范的DNS解析配置文件
    10.6.6.2 sms.service
    10.6.6.3 search.service
    # ttl
    ttl 60
    # 重载hosts配置
    reload 1m
    # 继续执行
    fallthrough
  }
  # file enables serving zone data from an RFC 1035-style master file.
  # https://coredns.io/plugins/file/
  # file service.signed service
  # 最后所有的都转发到系统配置的上游dns服务器去解析
  forward . /etc/resolv.conf
  # 缓存时间ttl
  cache 120
  # 自动加载配置文件的间隔时间
  reload 6s
  # 输出日志
  log
  # 输出错误
  errors
}
运行CoreDNS

coredns -conf ./Corefile

https://www.iamle.com/archives/2679.html

http://declanfeng.com/docs/linuxsoftware/linuxsoftware-1c9sudpod28sp

https://kubernetes.io/zh/docs/tasks/administer-cluster/coredns/

https://jimmysong.io/kubernetes-handbook/practice/coredns.html
https://support.huaweicloud.com/usermanual-cce/cce_01_0162.html
https://blog.csdn.net/hguisu/article/details/95060802


添加全局自定义域名解析
可以为 coredns 配置 hosts 来实现为 kubernetes 集群添加全局的自定义域名解析:

编辑 coredns 配置:

kubectl -n kube-system edit configmap coredns

https://imroc.cc/k8s/trick/customize-dns-resolution/