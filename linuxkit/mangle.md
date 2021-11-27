# iptables -t mangle -nvL
Chain PREROUTING (policy ACCEPT 168M packets, 114G bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain INPUT (policy ACCEPT 167M packets, 113G bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain FORWARD (policy ACCEPT 1538K packets, 775M bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain OUTPUT (policy ACCEPT 93M packets, 34G bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain POSTROUTING (policy ACCEPT 94M packets, 34G bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-KUBELET-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-PROXY-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination