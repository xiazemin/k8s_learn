node节点的iptables是由kube-proxy生成的


kube-proxy只修改了filter和nat表，它对iptables的链进行了扩充，自定义了KUBE-SERVICES，KUBE-NODEPORTS，KUBE-POSTROUTING，KUBE-MARK-MASQ和KUBE-MARK-DROP五个链，并主要通过为 KUBE-SERVICES链（附着在PREROUTING和OUTPUT）增加rule来配制traffic routing 规则

https://blog.csdn.net/bijiarong8928/article/details/100964459


// the services chain
kubeServicesChain utiliptables.Chain = "KUBE-SERVICES"
 
// the external services chain
kubeExternalServicesChain utiliptables.Chain = "KUBE-EXTERNAL-SERVICES"
 
// the nodeports chain
kubeNodePortsChain utiliptables.Chain = "KUBE-NODEPORTS"
 
// the kubernetes postrouting chain
kubePostroutingChain utiliptables.Chain = "KUBE-POSTROUTING"
 
// the mark-for-masquerade chain
KubeMarkMasqChain utiliptables.Chain = "KUBE-MARK-MASQ"     /*对于未能匹配到跳转规则的traffic set mark 0x8000，有此标记的数据包会在filter表drop掉*/
 
// the mark-for-drop chain
KubeMarkDropChain utiliptables.Chain = "KUBE-MARK-DROP"    /*对于符合条件的包 set mark 0x4000, 有此标记的数据包会在KUBE-POSTROUTING chain中统一做MASQUERADE*/
 
// the kubernetes forward chain
kubeForwardChain utiliptables.Chain = "KUBE-FORWARD"



-A KUBE-MARK-DROP -j MARK -- set -xmark 0x8000/0x8000
-A KUBE-MARK-MASQ -j MARK -- set -xmark 0x4000/0x4000




对于KUBE-MARK-MASQ链中所有规则设置了kubernetes独有MARK标记，在KUBE-POSTROUTING链中对NODE节点上匹配kubernetes独有MARK标记的数据包，当报文离开node节点时进行SNAT，MASQUERADE源IP

-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -m mark --mark 0x4000/0x4000 -j MASQUERADE
而对于KUBE-MARK-DROP设置标记的报文则会在KUBE_FIREWALL中全部丢弃  

-A KUBE-FIREWALL -m comment --comment "kubernetes firewall for dropping marked packets" -m mark --mark 0x8000/0x8000 -j DROP



KUBE_SVC和KUBE-SEP

Kube-proxy接着对每个服务创建“KUBE-SVC-”链，并在nat表中将KUBE-SERVICES链中每个目标地址是service的数据包导入这个“KUBE-SVC-”链，如果endpoint尚未创建，KUBE-SVC-链中没有规则，任何incomming packets在规则匹配失败后会被KUBE-MARK-DROP。在iptables的filter中有如下处理，如果KUBE-SVC处理失败会通过KUBE_FIREWALL丢弃

Chain INPUT (policy ACCEPT 209 packets, 378K bytes)
  pkts bytes target     prot opt in     out     source               destination        
  540K 1370M KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
  540K 1370M KUBE-FIREWALL  all  --  *      *       0.0.0.0/0            0.0.0.0/0
　KUBE_FIREWALL内容如下，就是直接丢弃所有报文：

Chain KUBE-FIREWALL (2 references)
  pkts bytes target     prot opt in     out     source               destination        
     0     0 DROP       all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes firewall for dropping marked packets */ mark match 0x8000/0x8000


KUBE-SEP表示的是KUBE-SVC对应的endpoint，当接收到的 serviceInfo中包含endpoint信息时，为endpoint创建跳转规则，如上述的KUBE-SVC-HVYO5BWEF5HC7MD7有endpoint，其iptables规则如下：
-A KUBE-SVC-HVYO5BWEF5HC7MD7 -m comment --comment "oqton-backoffice/sonatype-nexus:" -j KUBE-SEP-ESZGVIJJ5GN2KKU



KUBE-SEP-ESZGVIJJ5GN2KKU中的处理为将经过该链的所有tcp报文，DNAT为container 内部暴露的访问方式172.20.5.141:8080。结合对KUBE-SVC的处理可可知，这种访问方式就是cluster IP的访问方式，即将目的IP是cluster IP且目的端口是service暴露的端口的报文DNAT为目的IP是container且目的端口是container暴露的端口的报文，

-A KUBE-SEP-ESZGVIJJ5GN2KKUR -p tcp -m comment --comment "oqton-backoffice/sonatype-nexus:" -m tcp -j DNAT --to-destination 172.20.5.141:8080　




如果service类型为nodePort，（从LB转发至node的数据包均属此类）那么将KUBE-NODEPORTS链中每个目的地址是NODE节点端口的数据包导入这个“KUBE-SVC-”链；KUBE-NODEPORTS必须位于KUBE-SERVICE链的最后一个，可以看到iptables在处理报文时会优先处理目的IP为cluster IP的报文，匹配失败之后再去使用NodePort方式。如下规则表明，NodePort方式下会将目的ip为node节点且端口为node节点暴露的端口的报文进行KUBE-SVC-HVYO5BWEF5HC7MD7处理，KUBE-SVC-HVYO5BWEF5HC7MD7中会对报文进行DNAT转换。因此Custer IP和NodePort方式的唯一不同点就是KUBE-SERVICE中是根据cluster IP还是根据node port进行匹配

-A KUBE-NODEPORTS -p tcp -m comment --comment "oqton-backoffice/sonatype-nexus:" -m tcp --dport 32257 -j KUBE-MARK-MASQ
-A KUBE-NODEPORTS -p tcp -m comment --comment "oqton-backoffice/sonatype-nexus:" -m tcp --dport 32257 -j KUBE-SVC-HVYO5BWEF5HC7MD7



如果服务用到了loadblance，此时报文是从LB inbound的，报文的outbound处理则是通过KUBE-FW实现outbound报文的负载均衡。如下对目的IP是50.1.1.1(LB公网IP)且目的端口是443(一般是https)的报文作了KUBE-FW-J4ENLV444DNEMLR3处理。
-A KUBE-SERVICES -d 50.1.1.1/32 -p tcp -m comment --comment "kube-system/nginx-ingress-lb:https loadbalancer IP" -m tcp --dport 443 -j KUBE-FW-J4ENLV444DNEMLR3




如下在KUBE-FW-J4ENLV444DNEMLR3中显示的是LB的3个endpoint(该endpoint可能是service)，使用比率对报文进行了负载均衡控制
Chain KUBE-SVC-J4ENLV444DNEMLR3 (3 references)
     10   600 KUBE-SEP-ZVUNFBS77WHMPNFT  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/nginx-ingress-lb:https */ statistic mode random probability 0.33332999982
     18  1080 KUBE-SEP-Y47C2UBHCAA5SP4C  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/nginx-ingress-lb:https */ statistic mode random probability 0.50000000000
     16   960 KUBE-SEP-QGNNICTBV4CXTTZM  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/nginx-ingress-lb:https */





而上述3条链对应的处理如下，可以看到上述的每条链都作了DNAT，将目的IP由LB公网IP转换为LB的container IP
-A KUBE-SEP-ZVUNFBS77WHMPNFT -s 172.20.1.231/32 -m comment --comment "kube-system/nginx-ingress-lb:https" -j KUBE-MARK-MASQ
-A KUBE-SEP-ZVUNFBS77WHMPNFT -p tcp -m comment --comment "kube-system/nginx-ingress-lb:https" -m tcp -j DNAT --to-destination 172.20.1.231:443
-A KUBE-SEP-Y47C2UBHCAA5SP4C -s 172.20.2.191/32 -m comment --comment "kube-system/nginx-ingress-lb:https" -j KUBE-MARK-MASQ
-A KUBE-SEP-Y47C2UBHCAA5SP4C -p tcp -m comment --comment "kube-system/nginx-ingress-lb:https" -m tcp -j DNAT --to-destination 172.20.2.191:443
-A KUBE-SEP-QGNNICTBV4CXTTZM -s 172.20.2.3/32 -m comment --comment "kube-system/nginx-ingress-lb:https" -j KUBE-MARK-MASQ
-A KUBE-SEP-QGNNICTBV4CXTTZM -p tcp -m comment --comment "kube-system/nginx-ingress-lb:https" -m tcp -j DNAT --to-destination 172.20.2.3:443


node节点上的iptables中有到达所有service的规则，service 的cluster IP并不是一个实际的IP，它的存在只是为了找出实际的endpoint地址，对达到cluster IP的报文都要进行DNAT为Pod IP(+port)，不同node上的报文实际上是通过POD IP传输的，cluster IP只是本node节点的一个概念，用于查找并DNAT，即目的地址为clutter IP的报文只是本node发送的，其他节点不会发送(也没有路由支持)，即默认下cluster ip仅支持本node节点的service访问，如果需要跨node节点访问，可以使用插件实现，如flannel，它将pod  ip进行了封装


https://www.jianshu.com/p/1be9b096a691



目前kubernetes的Service访问的方式为cluster ip或node port等（云的负载均衡）方式，其实如何能访问到service，就是iptables或ipvs的功能，


https://www.jianshu.com/p/67744d680286
https://www.jianshu.com/p/7cff00e253f4
clusterIP 访问方式
PREROUTING   KUBE-SERVICE   KUBE-SVC-XXX   KUBE-SEP-XXX 


nodePort 方式
非本机访问
PREROUTING   KUBE-SERVICE  KUBE-NODEPORTS  KUBE-SVC-XXX   KUBE-SEP-XXX 
本机访问
OUTPUT   KUBE-SERVICE  KUBE-NODEPORTS  KUBE-SVC-XXX   KUBE-SEP-XXX 


该服务的 nodePort 端口为 30070，其 iptables 访问规则和使用 clusterIP 方式访问有点类似，不过 nodePort 方式会比 clusterIP 的方式多走一条链 KUBE-NODEPORTS，其会在 KUBE-NODEPORTS 链设置 mark 标记并转发到 KUBE-SVC-5SB6FTEHND4GTL2W，nodeport 与 clusterIP 访问方式最后都是转发到了 KUBE-SVC-xxx 链。



kube-proxy ipvs 模式



底层数据结构：iptables 使用链表，ipvs 使用哈希表
负载均衡算法：iptables 只支持随机、轮询两种负载均衡算法而 ipvs 支持的多达 8 种；
操作工具：iptables 需要使用 iptables 命令行工作来定义规则，ipvs 需要使用 ipvsadm 来定义规则。


ipset 是 iptables 的一种扩展，在 iptables 中可以使用 -m set启用 ipset 模块，具体来说，ipvs 使用 ipset 来存储需要 NAT 或 masquared 时的 ip 和端口列表。在数据包过滤过程中，首先遍历 iptables 规则，在定义了使用 ipset 的条件下会跳转到 ipset 列表中进行匹配。



kube-proxy 在 ipvs 模式下自定义了八条链，分别为 KUBE-SERVICES、KUBE-FIREWALL、KUBE-POSTROUTING、KUBE-MARK-MASQ、KUBE-NODE-PORT、KUBE-MARK-DROP、KUBE-FORWARD、KUBE-LOAD-BALANCER，如下所示：


此外，由于 linux 内核原生的 ipvs 模式只支持 DNAT，不支持 SNAT，所以，在以下几种场景中 ipvs 仍需要依赖 iptables 规则：


clusterIP 访问方式
PREROUTING   KUBE-SERVICE KUBE-CLUSTER-IP  INPUT KUBE-FIREWALL POSTROUTING

nodePort 方式
PREROUTING   KUBE-SERVICE KUBE-NODE-PORT INPUT KUBE-FIREWALL POSTROUTING


https://github.com/kubernetes/kubernetes/blob/master/pkg/proxy/ipvs/README.md

https://github.com/kubernetes/kubernetes/issues/17470

https://github.com/kubernetes/kubernetes/issues/44063

https://github.com/cilium/k8s-iptables-diagram
