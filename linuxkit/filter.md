# iptables -t filter -nvL
Chain INPUT (policy ACCEPT 3386K packets, 2827M bytes)
 pkts bytes target     prot opt in     out     source               destination
    0     0 IN_WEB     tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:80
 167M  113G KUBE-NODEPORTS  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes health check service ports */
 307K   18M KUBE-EXTERNAL-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes externally-visible service portals */
 167M  113G KUBE-FIREWALL  all  --  *      *       0.0.0.0/0            0.0.0.0/0

Chain FORWARD (policy ACCEPT 175 packets, 12889 bytes)
 pkts bytes target     prot opt in     out     source               destination
1538K  775M KUBE-FORWARD  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding rules */
22724 1556K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes service portals */
22722 1555K KUBE-EXTERNAL-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes externally-visible service portals */
29608 8435K DOCKER-USER  all  --  *      *       0.0.0.0/0            0.0.0.0/0
29608 8435K DOCKER-ISOLATION-STAGE-1  all  --  *      *       0.0.0.0/0            0.0.0.0/0
 5142 6792K ACCEPT     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0            ctstate RELATED,ESTABLISHED
    0     0 DOCKER     all  --  *      docker0  0.0.0.0/0            0.0.0.0/0
11502  673K ACCEPT     all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0
    0     0 ACCEPT     all  --  docker0 docker0  0.0.0.0/0            0.0.0.0/0
    0     0 ACCEPT     all  --  *      br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0            ctstate RELATED,ESTABLISHED
    0     0 DOCKER     all  --  *      br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0
    0     0 ACCEPT     all  --  br-421a9e16d12b !br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0
    0     0 ACCEPT     all  --  br-421a9e16d12b br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0

Chain OUTPUT (policy ACCEPT 1567K packets, 526M bytes)
 pkts bytes target     prot opt in     out     source               destination
 484K   30M KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate NEW /* kubernetes service portals */
  93M   34G KUBE-FIREWALL  all  --  *      *       0.0.0.0/0            0.0.0.0/0

Chain DOCKER (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     tcp  --  !docker0 docker0  0.0.0.0/0            172.17.0.2           tcp dpt:443
    0     0 ACCEPT     tcp  --  !docker0 docker0  0.0.0.0/0            172.17.0.2           tcp dpt:80

Chain DOCKER-ISOLATION-STAGE-1 (1 references)
 pkts bytes target     prot opt in     out     source               destination
11502  673K DOCKER-ISOLATION-STAGE-2  all  --  docker0 !docker0  0.0.0.0/0            0.0.0.0/0
    0     0 DOCKER-ISOLATION-STAGE-2  all  --  br-421a9e16d12b !br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0
29608 8435K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

Chain DOCKER-ISOLATION-STAGE-2 (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 DROP       all  --  *      docker0  0.0.0.0/0            0.0.0.0/0
    0     0 DROP       all  --  *      br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0
11502  673K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

Chain DOCKER-USER (1 references)
 pkts bytes target     prot opt in     out     source               destination
29608 8435K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0

Chain IN_WEB (1 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-EXTERNAL-SERVICES (2 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-FIREWALL (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 DROP       all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes firewall for dropping marked packets */ mark match 0x8000/0x8000
    0     0 DROP       all  --  *      *      !127.0.0.0/8          127.0.0.0/8          /* block incoming localnet connections */ ! ctstate RELATED,ESTABLISHED,DNAT

Chain KUBE-FORWARD (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 DROP       all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate INVALID
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding rules */ mark match 0x4000/0x4000
20629 9087K ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding conntrack pod source rule */ ctstate RELATED,ESTABLISHED
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes forwarding conntrack pod destination rule */ ctstate RELATED,ESTABLISHED

Chain KUBE-KUBELET-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-NODEPORTS (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:https health check node port */ tcp dpt:30380
    0     0 ACCEPT     tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:http health check node port */ tcp dpt:30380

Chain KUBE-PROXY-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-SERVICES (2 references)
 pkts bytes target     prot opt in     out     source               destination
/ #