docker run -it --privileged --pid=host alpine:latest nsenter -t 1 -m -u -n -i sh

# iptables -t raw -nvL
Chain PREROUTING (policy ACCEPT 169M packets, 114G bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain OUTPUT (policy ACCEPT 93M packets, 34G bytes)
 pkts bytes target     prot opt in     out     source               destination

 