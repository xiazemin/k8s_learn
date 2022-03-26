service secrets configmap

% kubectl create ns inject
namespace/inject created

 % kubectl apply -f deploy.yaml -n inject
error: error validating "deploy.yaml": error validating data: apiVersion not set; if you choose to ignore these errors, turn validation off with --validate=false

% kubectl apply -f deploy.yaml -n inject
error: error validating "deploy.yaml": error validating data: [ValidationError(Deployment.spec): unknown field "containers" in io.k8s.api.apps.v1.DeploymentSpec, ValidationError(Deployment.spec.selector): unknown field "labels" in io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector, ValidationError(Deployment.spec): missing required field "template" in io.k8s.api.apps.v1.DeploymentSpec]; if you choose to ignore these errors, turn validation off with --validate=false

 % kubectl apply -f deploy.yaml -n inject
deployment.apps/inject created

% kubectl -n inject get pods            
NAME                    READY   STATUS      RESTARTS   AGE
inject-b7f5ff45-2x9j8   0/1     Completed   0          3s

 % docker pull nginx:latest
latest: Pulling from library/nginx
ae13dd578326: Pull complete 
6c0ee9353e13: Pull complete 
dca7733b187e: Pull complete 
352e5a6cac26: Pull complete 
9eaf108767c7: Pull complete 
be0c016df0be: Pull complete 
Digest: sha256:4ed64c2e0857ad21c38b98345ebb5edb01791a0a10b0e9e3d9ddde185cdbd31a
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest

% kubectl -n inject get pods
NAME                     READY   STATUS    RESTARTS   AGE
inject-556d8c558-7q4gq   1/1     Running   0          13s


 % istioctl kube-inject -f deploy.yaml |kubectl apply -f - -n inject 
deployment.apps/inject configured

 % kubectl get pods -n inject
NAME                      READY   STATUS    RESTARTS   AGE
inject-594c785fb6-tlhns   2/2     Running   0          30s


% istioctl kube-inject -f deploy.yaml > deploy-inject.yaml


containers:
    image: docker.io/istio/proxyv2:1.12.1
    name: istio-proxy
      - args:
        - proxy
        - sidecar


initContainers:
    image: docker.io/istio/proxyv2:1.12.1
    name: istio-init
      - args:
        - istio-iptables


 % kubectl exec -it -n inject inject-594c785fb6-tlhns -c nginx -- ifconfig
OCI runtime exec failed: exec failed: container_linux.go:380: starting container process caused: exec: "ifconfig": executable file not found in $PATH: unknown
command terminated with exit code 126


% kubectl exec -it -n inject inject-594c785fb6-tlhns -c nginx -- sh                     
# 

# curl www.baidu.com
<!DOCTYPE html>

# apt-get update
Get:1 http://deb.debian.org/debian bullseye InRelease [116 kB]
Get:2 http://security.debian.org/debian-security bullseye-security InRelease [44.1 kB]


# apt-get install net-tools
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
The following NEW packages will be installed:
  net-tools
0 upgraded, 1 newly installed, 0 to remove and 0 not upgraded.
Need to get 250 kB of archives.
After this operation, 1015 kB of additional disk space will be used.
Get:1 http://deb.debian.org/debian bullseye/main amd64 net-tools amd64 1.60+git20181103.0eebece-1 [250 kB]
Fetched 250 kB in 2s (156 kB/s)   
debconf: delaying package configuration, since apt-utils is not installed
Selecting previously unselected package net-tools.
(Reading database ... 7815 files and directories currently installed.)
Preparing to unpack .../net-tools_1.60+git20181103.0eebece-1_amd64.deb ...
Unpacking net-tools (1.60+git20181103.0eebece-1) ...
Setting up net-tools (1.60+git20181103.0eebece-1) ...


# ifconfig
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.1.4.109  netmask 255.255.0.0  broadcast 10.1.255.255
        ether c6:02:6c:0b:9f:56  txqueuelen 0  (Ethernet)
        RX packets 12308  bytes 10666188 (10.1 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 6747  bytes 6039085 (5.7 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 4940  bytes 17194716 (16.3 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 4940  bytes 17194716 (16.3 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0




 % kubectl exec -it -n inject inject-594c785fb6-tlhns -c istio-proxy -- ifconfig
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.1.4.109  netmask 255.255.0.0  broadcast 10.1.255.255
        ether c6:02:6c:0b:9f:56  txqueuelen 0  (Ethernet)
        RX packets 13291  bytes 10750250 (10.7 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 7505  bytes 9020989 (9.0 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 5525  bytes 22664975 (22.6 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 5525  bytes 22664975 (22.6 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0



 % kubectl exec -it -n inject inject-594c785fb6-tlhns -c istio-proxy -- route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.1.0.1        0.0.0.0         UG    0      0        0 eth0
10.1.0.0        0.0.0.0         255.255.0.0     U     0      0        0 eth0


% kubectl exec -it -n inject inject-594c785fb6-tlhns -c nginx -- route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.1.0.1        0.0.0.0         UG    0      0        0 eth0
10.1.0.0        0.0.0.0         255.255.0.0     U     0      0        0 eth0



% kubectl apply -f deploy.yaml 
deployment.apps/inject created


 % kubectl get pods
NAME                     READY   STATUS    RESTARTS   AGE
inject-556d8c558-l2b68   1/1     Running   0          15s




 % kubectl exec -it inject-556d8c558-l2b68 -- netstat -ntlp 
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      1/nginx: master pro 
tcp6       0      0 :::80                   :::*                    LISTEN      1/nginx: master pro 




 % kubectl -n inject exec -it inject-594c785fb6-tlhns  -- netstat -ntlp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      1/nginx: master pro 
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      -                   
tcp        0      0 127.0.0.1:15000         0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      -                   
tcp        0      0 127.0.0.1:15004         0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      -                   
tcp6       0      0 :::15020                :::*                    LISTEN      -                   
tcp6       0      0 :::80                   :::*                    LISTEN      1/nginx: master pro 




% kubectl -n inject exec -it inject-594c785fb6-tlhns  -c istio-proxy -- netstat -ntlp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 127.0.0.1:15000         0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 127.0.0.1:15004         0.0.0.0:*               LISTEN      1/pilot-agent       
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      17/envoy            
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      17/envoy            
tcp6       0      0 :::15020                :::*                    LISTEN      1/pilot-agent       
tcp6       0      0 :::80                   :::*                    LISTEN      -                 



% kubectl logs -f -n inject inject-594c785fb6-tlhns -c istio-init
2022-03-24T15:16:48.606265Z     info    Istio iptables environment:
ENVOY_PORT=
INBOUND_CAPTURE_PORT=
ISTIO_INBOUND_INTERCEPTION_MODE=
ISTIO_INBOUND_TPROXY_ROUTE_TABLE=
ISTIO_INBOUND_PORTS=
ISTIO_OUTBOUND_PORTS=
ISTIO_LOCAL_EXCLUDE_PORTS=
ISTIO_EXCLUDE_INTERFACES=
ISTIO_SERVICE_CIDR=
ISTIO_SERVICE_EXCLUDE_CIDR=
ISTIO_META_DNS_CAPTURE=
2022-03-24T15:16:48.606382Z     info    Istio iptables variables:
PROXY_PORT=15001
PROXY_INBOUND_CAPTURE_PORT=15006
PROXY_TUNNEL_PORT=15008
PROXY_UID=1337
PROXY_GID=1337
INBOUND_INTERCEPTION_MODE=REDIRECT
INBOUND_TPROXY_MARK=1337
INBOUND_TPROXY_ROUTE_TABLE=133
INBOUND_PORTS_INCLUDE=*
INBOUND_PORTS_EXCLUDE=15090,15021,15020
OUTBOUND_IP_RANGES_INCLUDE=*
OUTBOUND_IP_RANGES_EXCLUDE=
OUTBOUND_PORTS_INCLUDE=
OUTBOUND_PORTS_EXCLUDE=
KUBE_VIRT_INTERFACES=
ENABLE_INBOUND_IPV6=false
DNS_CAPTURE=false
CAPTURE_ALL_DNS=false
DNS_SERVERS=[],[]
OUTPUT_PATH=
NETWORK_NAMESPACE=
CNI_MODE=false
EXCLUDE_INTERFACES=

2022-03-24T15:16:48.608165Z     info    Writing following contents to rules file: /tmp/iptables-rules-1648135008606893849.txt3417690870
* nat
-N ISTIO_INBOUND
-N ISTIO_REDIRECT
-N ISTIO_IN_REDIRECT
-N ISTIO_OUTPUT
-A ISTIO_INBOUND -p tcp --dport 15008 -j RETURN
-A ISTIO_REDIRECT -p tcp -j REDIRECT --to-ports 15001
-A ISTIO_IN_REDIRECT -p tcp -j REDIRECT --to-ports 15006
-A PREROUTING -p tcp -j ISTIO_INBOUND
-A ISTIO_INBOUND -p tcp --dport 22 -j RETURN
-A ISTIO_INBOUND -p tcp --dport 15090 -j RETURN
-A ISTIO_INBOUND -p tcp --dport 15021 -j RETURN
-A ISTIO_INBOUND -p tcp --dport 15020 -j RETURN
-A ISTIO_INBOUND -p tcp -j ISTIO_IN_REDIRECT
-A OUTPUT -p tcp -j ISTIO_OUTPUT
-A ISTIO_OUTPUT -o lo -s 127.0.0.6/32 -j RETURN
-A ISTIO_OUTPUT -o lo ! -d 127.0.0.1/32 -m owner --uid-owner 1337 -j ISTIO_IN_REDIRECT
-A ISTIO_OUTPUT -o lo -m owner ! --uid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -m owner --uid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -o lo ! -d 127.0.0.1/32 -m owner --gid-owner 1337 -j ISTIO_IN_REDIRECT
-A ISTIO_OUTPUT -o lo -m owner ! --gid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -m owner --gid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -d 127.0.0.1/32 -j RETURN
-A ISTIO_OUTPUT -j ISTIO_REDIRECT
COMMIT
2022-03-24T15:16:48.608299Z     info    Running command: iptables-restore --noflush /tmp/iptables-rules-1648135008606893849.txt3417690870
2022-03-24T15:16:48.658720Z     info    Writing following contents to rules file: /tmp/ip6tables-rules-1648135008658589275.txt232986230

2022-03-24T15:16:48.658775Z     info    Running command: ip6tables-restore --noflush /tmp/ip6tables-rules-1648135008658589275.txt232986230
2022-03-24T15:16:48.663272Z     info    Running command: iptables-save 
2022-03-24T15:16:48.669538Z     info    Command output: 
# Generated by iptables-save v1.8.4 on Thu Mar 24 15:16:48 2022
*nat
:PREROUTING ACCEPT [0:0]
:INPUT ACCEPT [0:0]
:OUTPUT ACCEPT [0:0]
:POSTROUTING ACCEPT [0:0]
:ISTIO_INBOUND - [0:0]
:ISTIO_IN_REDIRECT - [0:0]
:ISTIO_OUTPUT - [0:0]
:ISTIO_REDIRECT - [0:0]
-A PREROUTING -p tcp -j ISTIO_INBOUND
-A OUTPUT -p tcp -j ISTIO_OUTPUT
-A ISTIO_INBOUND -p tcp -m tcp --dport 15008 -j RETURN
-A ISTIO_INBOUND -p tcp -m tcp --dport 22 -j RETURN
-A ISTIO_INBOUND -p tcp -m tcp --dport 15090 -j RETURN
-A ISTIO_INBOUND -p tcp -m tcp --dport 15021 -j RETURN
-A ISTIO_INBOUND -p tcp -m tcp --dport 15020 -j RETURN
-A ISTIO_INBOUND -p tcp -j ISTIO_IN_REDIRECT
-A ISTIO_IN_REDIRECT -p tcp -j REDIRECT --to-ports 15006
-A ISTIO_OUTPUT -s 127.0.0.6/32 -o lo -j RETURN
-A ISTIO_OUTPUT ! -d 127.0.0.1/32 -o lo -m owner --uid-owner 1337 -j ISTIO_IN_REDIRECT
-A ISTIO_OUTPUT -o lo -m owner ! --uid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -m owner --uid-owner 1337 -j RETURN
-A ISTIO_OUTPUT ! -d 127.0.0.1/32 -o lo -m owner --gid-owner 1337 -j ISTIO_IN_REDIRECT
-A ISTIO_OUTPUT -o lo -m owner ! --gid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -m owner --gid-owner 1337 -j RETURN
-A ISTIO_OUTPUT -d 127.0.0.1/32 -j RETURN
-A ISTIO_OUTPUT -j ISTIO_REDIRECT
-A ISTIO_REDIRECT -p tcp -j REDIRECT --to-ports 15001
COMMIT
# Completed on Thu Mar 24 15:16:48 2022



% docker ps |grep istio-proxy
b422025e8631   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up 1 second                       k8s_istio-proxy_istio-egressgateway-687f4db598-8dmmb_istio-system_cfb39b6a-3b04-4ae0-ba0c-334d0331c981_1
145bfe8c4017   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up Less than a second             k8s_istio-proxy_details-v1-79f774bdb9-dg7tp_bookinfo_489f4c5e-becf-475d-af3d-4d6efcbf4fec_1
bccc8708372e   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up 1 second                       k8s_istio-proxy_reviews-v3-84779c7bbc-gjm7n_bookinfo_54647799-3c8d-4cdd-8416-e5b6ce458b61_1
f6301cc2fb1f   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up 1 second                       k8s_istio-proxy_ratings-v1-b6994bb9-85mtv_bookinfo_955584c3-fabb-4e3c-9a95-f950c6ee1604_1
7ef85fca8084   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up Less than a second             k8s_istio-proxy_istio-ingressgateway-78f69bd5db-tdbtw_istio-system_a4da4019-7b05-4f94-8f19-07991f9a3b28_1
c7132042d226   4f3c74acb37a           "/usr/local/bin/pilo…"   3 seconds ago    Up 1 second                       k8s_istio-proxy_productpage-v1-6b746f74dc-k7lph_bookinfo_93e87998-d6db-4216-8460-c059f2dcf979_1
bc48e3f821a1   4f3c74acb37a           "/usr/local/bin/pilo…"   5 seconds ago    Up 2 seconds                      k8s_istio-proxy_reviews-v1-545db77b95-d79vb_bookinfo_1e4643c9-98a7-4ebb-ab61-eb2c9c3373df_1
2dc9b8f48263   4f3c74acb37a           "/usr/local/bin/pilo…"   5 seconds ago    Up 2 seconds                      k8s_istio-proxy_reviews-v2-7bf8c9648f-g8j2q_bookinfo_746e55ad-f7e5-4ca6-bbe2-f736b345b42c_1
37a07b3f0a88   4f3c74acb37a           "/usr/local/bin/pilo…"   42 minutes ago   Up 42 minutes                     k8s_istio-proxy_inject-594c785fb6-tlhns_inject_2eb35330-ce86-4ef6-bd3b-a61838cf882f_0
xiazemin@xiazemindeMBP istio-1.12.1 %
xiazemin@xiazemindeMBP istio-1.12.1 %
xiazemin@xiazemindeMBP istio-1.12.1 % docker ps |grep istio-proxy
b422025e8631   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_istio-egressgateway-687f4db598-8dmmb_istio-system_cfb39b6a-3b04-4ae0-ba0c-334d0331c981_1
145bfe8c4017   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_details-v1-79f774bdb9-dg7tp_bookinfo_489f4c5e-becf-475d-af3d-4d6efcbf4fec_1
bccc8708372e   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_reviews-v3-84779c7bbc-gjm7n_bookinfo_54647799-3c8d-4cdd-8416-e5b6ce458b61_1
f6301cc2fb1f   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_ratings-v1-b6994bb9-85mtv_bookinfo_955584c3-fabb-4e3c-9a95-f950c6ee1604_1
7ef85fca8084   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_istio-ingressgateway-78f69bd5db-tdbtw_istio-system_a4da4019-7b05-4f94-8f19-07991f9a3b28_1
c7132042d226   4f3c74acb37a           "/usr/local/bin/pilo…"   13 seconds ago   Up 10 seconds             k8s_istio-proxy_productpage-v1-6b746f74dc-k7lph_bookinfo_93e87998-d6db-4216-8460-c059f2dcf979_1
bc48e3f821a1   4f3c74acb37a           "/usr/local/bin/pilo…"   15 seconds ago   Up 11 seconds             k8s_istio-proxy_reviews-v1-545db77b95-d79vb_bookinfo_1e4643c9-98a7-4ebb-ab61-eb2c9c3373df_1
2dc9b8f48263   4f3c74acb37a           "/usr/local/bin/pilo…"   15 seconds ago   Up 11 seconds             k8s_istio-proxy_reviews-v2-7bf8c9648f-g8j2q_bookinfo_746e55ad-f7e5-4ca6-bbe2-f736b345b42c_1
37a07b3f0a88   4f3c74acb37a           "/usr/local/bin/pilo…"   42 minutes ago   Up 42 minutes             k8s_istio-proxy_inject-594c785fb6-tlhns_inject_2eb35330-ce86-4ef6-bd3b-a61838cf882f_0



 % docker exec -it --privileged b422025e8631 bash
istio-proxy@istio-egressgateway-687f4db598-8dmmb:/$

istio-proxy@istio-egressgateway-687f4db598-8dmmb:/$ iptables -nvL -t nat
Fatal: can't open lock file /run/xtables.lock: Read-only file system

$ sudo iptables -nvL -t nat
sudo: effective uid is not 0, is /usr/bin/sudo on a file system with the 'nosuid' option set or an NFS file system without root privileges?

https://superuser.com/questions/1580293/sudo-effective-uid-is-not-0-is-usr-bin-sudo-on-a-file-system-with-the-nosuid


https://blog.csdn.net/shida_csdn/article/details/79791897

 % docker exec -it --privileged --user root b422025e8631 bash
root@istio-egressgateway-687f4db598-8dmmb:/# iptables
iptables v1.8.4 (legacy): no command specified

/# sudo su root
root@istio-egressgateway-687f4db598-8dmmb:/# iptables -nvL -t nat
Fatal: can't open lock file /run/xtables.lock: Read-only file system

https://blog.csdn.net/yygydjkthh/article/details/50737237

 % docker exec -it --privileged --pid=host --user root b422025e8631 bash
unknown flag: --pid
See 'docker exec --help'.

https://github.com/docker/for-mac/issues/5547
https://stackoverflow.com/questions/49419092/disallow-egress-from-docker-containers-on-docker-for-mac

https://docker-docs.netlify.app/network/iptables/


% kubectl exec -n inject inject-594c785fb6-tlhns -c nginx -- ps -ef 
OCI runtime exec failed: exec failed: container_linux.go:380: starting container process caused: exec: "ps": executable file not found in $PATH: unknown
command terminated with exit code 126

 % kubectl exec -it -n inject inject-594c785fb6-tlhns -c nginx -- bash
root@inject-594c785fb6-tlhns:/# 
root@inject-594c785fb6-tlhns:/# apt-get install ps
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done


% kubectl exec -n inject inject-594c785fb6-tlhns -c nginx -- ps -ef
OCI runtime exec failed: exec failed: container_linux.go:380: starting container process caused: exec: "ps": executable file not found in $PATH: unknown
command terminated with exit code 126


安装procps包

root@inject-594c785fb6-tlhns:/# apt-get install procps
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
The following additional packages will be installed:
  libgpm2 libncurses6 libncursesw6 libprocps8 psmisc
Suggested packages:


% kubectl exec -n inject inject-594c785fb6-tlhns -c nginx -- ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 15:16 ?        00:00:00 nginx: master process nginx -g daemon off;
nginx       31     1  0 15:16 ?        00:00:00 nginx: worker process
nginx       32     1  0 15:16 ?        00:00:00 nginx: worker process
nginx       33     1  0 15:16 ?        00:00:00 nginx: worker process
root       584     0  0 16:29 ?        00:00:00 ps -ef


% kubectl exec -n inject inject-594c785fb6-tlhns -c istio-proxy -- ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
istio-p+     1     0  0 15:16 ?        00:00:08 /usr/local/bin/pilot-agent proxy sidecar --domain inject.svc.cluster.local --proxyLogLevel=warning --proxyComponentLogLevel=misc:error --log_output_level=default:info --concurrency 2
istio-p+    17     1  0 15:16 ?        00:00:26 /usr/local/bin/envoy -c etc/istio/proxy/envoy-rev0.json --restart-epoch 0 --drain-time-s 45 --drain-strategy immediate --parent-shutdown-time-s 60 --local-address-ip-version v4 --file-flush-interval-msec 1000 --disable-hot-restart --log-format %Y-%m-%dT%T.%fZ.%l.envoy %n.%v -l warning --component-log-level misc:error --concurrency 2
istio-p+    51     0  0 16:30 ?        00:00:00 ps -ef




 % kubectl exec -it -n istio-system istiod-76d66d9876-hgsmt -- sh
$
$
$ ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
istio-p+     1     0  0 08:16 ?        00:03:00 /usr/local/bin/pilot-discovery discovery --monitoringAddr=:15014 --log_output_level=default:info --domain cluster.local --keepaliveMaxSer
istio-p+    20     0  0 13:22 pts/0    00:00:00 sh
istio-p+    27    20  0 13:22 pts/0    00:00:00 ps -ef
$


% kubectl label ns inject istio-injection=enabled
namespace/inject labeled


 % kubectl -n inject exec -it inject-594c785fb6-tlhns -c istio-proxy -- netstat -ntlp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:15021           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      -
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:15090           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 127.0.0.1:15000         0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:15001           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 127.0.0.1:15004         0.0.0.0:*               LISTEN      1/pilot-agent
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      17/envoy
tcp        0      0 0.0.0.0:15006           0.0.0.0:*               LISTEN      17/envoy
tcp6       0      0 :::15020                :::*                    LISTEN      1/pilot-agent
tcp6       0      0 :::80                   :::*                    LISTEN      -




 % docker ps |grep istio-proxy
b422025e8631   4f3c74acb37a           "/usr/local/bin/pilo…"   22 hours ago         Up 22 hours                   k8s_istio-proxy_istio-egressgateway-687f4db598-8dmmb_istio-system_cfb39b6a-3b04-4ae0-ba0c-334d0331c981_1

% docker exec -it --privileged b422025e8631 sh

$ iptables -nvL -t nat
Fatal: can't open lock file /run/xtables.lock: Read-only file system
$ sudo su root
sudo: effective uid is not 0, is /usr/bin/sudo on a file system with the 'nosuid' option set or an NFS file system without root privileges?

https://blog.csdn.net/qingfeng14/article/details/52506074


https://blog.csdn.net/hello_percy/article/details/64124978

sudo works with a mechanism that is called setuid (Set User ID, or also called suid). If that bit is set on an executable file (like sudo), then the application is executed under the permissions of the user, who is the owner of that file (in case of sudo, the owner is the root user).


apt reinstall sudo
https://askubuntu.com/questions/625540/suddenly-cant-run-sudo

https://github.com/docker/for-linux/issues/1100
https://unix.stackexchange.com/questions/546822/why-does-sudo-fail-inside-docker-complaining-about-nosuid
https://stackoverflow.com/questions/67282511/suid-is-not-honoured-inside-docker-container
https://cn.2md-tuning-klub.com/708034-why-does-sudo-fail-inside-RIEXUY

