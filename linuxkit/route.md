/ # route -e
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
default         192.168.65.1    0.0.0.0         UG        0 0          0 eth0
10.1.0.0        *               255.255.0.0     U         0 0          0 cni0
127.0.0.0       *               255.0.0.0       U         0 0          0 lo
172.17.0.0      *               255.255.0.0     U         0 0          0 docker0
172.18.0.0      *               255.255.0.0     U         0 0          0 br-421a9e16d12b
192.168.65.0    *               255.255.255.0   U         0 0          0 eth0
192.168.65.5    *               255.255.255.255 UH        0 0          0 services1


/ # ifconfig
br-421a9e16d12b Link encap:Ethernet  HWaddr 02:42:42:DF:C5:C0
          inet addr:172.18.0.1  Bcast:172.18.255.255  Mask:255.255.0.0
          UP BROADCAST MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

cni0      Link encap:Ethernet  HWaddr 06:BE:94:CC:A5:73
          inet addr:10.1.0.1  Bcast:10.1.255.255  Mask:255.255.0.0
          inet6 addr: fe80::4be:94ff:fecc:a573/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:5819654 errors:0 dropped:0 overruns:0 frame:0
          TX packets:6295138 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:1144109299 (1.0 GiB)  TX bytes:3366178196 (3.1 GiB)

docker0   Link encap:Ethernet  HWaddr 02:42:DF:31:8E:2F
          inet addr:172.17.0.1  Bcast:172.17.255.255  Mask:255.255.0.0
          inet6 addr: fe80::42:dfff:fe31:8e2f/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:691476 errors:0 dropped:0 overruns:0 frame:0
          TX packets:779683 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:128132016 (122.1 MiB)  TX bytes:633453568 (604.1 MiB)

eth0      Link encap:Ethernet  HWaddr 02:50:00:00:00:01
          inet addr:192.168.65.3  Bcast:192.168.65.255  Mask:255.255.255.0
          inet6 addr: fe80::50:ff:fe00:1/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:136938715 errors:0 dropped:0 overruns:0 frame:0
          TX packets:64538173 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:116477521499 (108.4 GiB)  TX bytes:24675357566 (22.9 GiB)

lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:44428429 errors:0 dropped:0 overruns:0 frame:0
          TX packets:44428429 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:11475400211 (10.6 GiB)  TX bytes:11475400211 (10.6 GiB)

services1 Link encap:Ethernet  HWaddr E2:84:AF:68:F0:01
          inet addr:192.168.65.4  Bcast:0.0.0.0  Mask:255.255.255.255
          inet6 addr: fe80::e084:afff:fe68:f001/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:34966 errors:0 dropped:0 overruns:0 frame:0
          TX packets:34988 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:3050733 (2.9 MiB)  TX bytes:2346563 (2.2 MiB)

veth03c6deb0 Link encap:Ethernet  HWaddr 92:13:6C:47:4B:B0
          inet6 addr: fe80::9013:6cff:fe47:4bb0/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:88128 errors:0 dropped:0 overruns:0 frame:0
          TX packets:168551 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:6148932 (5.8 MiB)  TX bytes:11460870 (10.9 MiB)

veth0a198a4 Link encap:Ethernet  HWaddr 12:F7:A2:9A:C5:C2
          inet6 addr: fe80::10f7:a2ff:fe9a:c5c2/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:27375 errors:0 dropped:0 overruns:0 frame:0
          TX packets:35127 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:5565821 (5.3 MiB)  TX bytes:29774830 (28.3 MiB)

veth257e7f5f Link encap:Ethernet  HWaddr 42:13:9C:A7:AE:70
          inet6 addr: fe80::4013:9cff:fea7:ae70/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:168 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:11436 (11.1 KiB)

veth30b0175d Link encap:Ethernet  HWaddr 06:D7:E9:91:D8:89
          inet6 addr: fe80::4d7:e9ff:fe91:d889/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:24889 errors:0 dropped:0 overruns:0 frame:0
          TX packets:24046 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1968183 (1.8 MiB)  TX bytes:7730483 (7.3 MiB)

veth46b3b56 Link encap:Ethernet  HWaddr 86:06:23:D7:3A:B6
          inet6 addr: fe80::8406:23ff:fed7:3ab6/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:23 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:1682 (1.6 KiB)

veth74ed2a5d Link encap:Ethernet  HWaddr EA:59:1D:C5:5D:3A
          inet6 addr: fe80::e859:1dff:fec5:5d3a/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:716454 errors:0 dropped:0 overruns:0 frame:0
          TX packets:751359 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:136830804 (130.4 MiB)  TX bytes:236760595 (225.7 MiB)

veth75d8f631 Link encap:Ethernet  HWaddr 8A:E3:63:8D:C6:75
          inet6 addr: fe80::88e3:63ff:fe8d:c675/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:344 errors:0 dropped:0 overruns:0 frame:0
          TX packets:564 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:25360 (24.7 KiB)  TX bytes:46582 (45.4 KiB)

veth8c1c4559 Link encap:Ethernet  HWaddr 1E:9D:BF:00:90:FA
          inet6 addr: fe80::1c9d:bfff:fe00:90fa/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:3604748 errors:0 dropped:0 overruns:0 frame:0
          TX packets:3881589 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:939962963 (896.4 MiB)  TX bytes:2577254578 (2.3 GiB)

vetha09700d3 Link encap:Ethernet  HWaddr 4E:69:97:3F:04:C2
          inet6 addr: fe80::4c69:97ff:fe3f:4c2/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:67 errors:0 dropped:0 overruns:0 frame:0
          TX packets:226 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:5334 (5.2 KiB)  TX bytes:36155 (35.3 KiB)

vetha4d3962e Link encap:Ethernet  HWaddr 4A:5F:C8:C7:09:8B
          inet6 addr: fe80::485f:c8ff:fec7:98b/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:344777 errors:0 dropped:0 overruns:0 frame:0
          TX packets:355018 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:30831378 (29.4 MiB)  TX bytes:95211792 (90.8 MiB)

vetha7f6e8fa Link encap:Ethernet  HWaddr 9A:0C:03:16:5C:40
          inet6 addr: fe80::980c:3ff:fe16:5c40/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:344981 errors:0 dropped:0 overruns:0 frame:0
          TX packets:354983 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:30836727 (29.4 MiB)  TX bytes:95282607 (90.8 MiB)

vethd54e91dd Link encap:Ethernet  HWaddr E2:F0:99:CF:76:E2
          inet6 addr: fe80::e0f0:99ff:fecf:76e2/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:734985 errors:0 dropped:0 overruns:0 frame:0
          TX packets:799979 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:84101465 (80.2 MiB)  TX bytes:347620205 (331.5 MiB)

/ # ifconfig lo
lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:44435075 errors:0 dropped:0 overruns:0 frame:0
          TX packets:44435075 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:11477194856 (10.6 GiB)  TX bytes:11477194856 (10.6 GiB)




/ # nslookup docker.for.mac.host.internal
Server:		192.168.65.5
Address:	192.168.65.5:53

Non-authoritative answer:

Non-authoritative answer:
Name:	docker.for.mac.host.internal
Address: 192.168.65.2




The host has a changing IP address (or none if you have no network access). We recommend that you connect to the special DNS name host.docker.internal which resolves to the internal IP address used by the host. This is for development purpose and will not work in a production environment outside of Docker Desktop for Mac.

https://docs.docker.com/desktop/mac/networking/

/ # cat /etc/resolv.conf
# DNS requests are forwarded to the host. DHCP DNS options are ignored.
nameserver 192.168.65.5



/ # nslookup gateway.docker.internal
Server:		192.168.65.5
Address:	192.168.65.5:53

Non-authoritative answer:
Name:	gateway.docker.internal
Address: 192.168.65.2

Non-authoritative answer:



The way docker does DNS by default is fundamentally broken. It should always use dnsmasq and the host's DNS configuration by default (and not use 8.8.8.8). Each container should (by default) resolve DNS by querying the host, which should forward the request to its own resolver, and provide resolution for all of .docker.internal, including host.docker.internal (which should be configurable to allow containers named host). All other accessible containers should also resolve in .docker.internal

https://github.com/moby/libnetwork/pull/2348

entrypoint.sh 


function fix_linux_internal_host() {
DOCKER_INTERNAL_HOST="host.docker.internal"
if ! grep $DOCKER_INTERNAL_HOST /etc/hosts > /dev/null ; then
DOCKER_INTERNAL_IP='/sbin/ip route | awk '/default/ { print $3 }' | awk '!seen[$0]++'
echo -e "$DOCKER_INTERNAL_IP\t$DOCKER_INTERNAL_HOST" | tee -a /etc/hosts > /dev/null
echo "Added $DOCKER_INTERNAL_HOST to hosts /etc/hosts"
fi
}


/ # cat /etc/hosts
127.0.0.1       localhost
::1     localhost ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters


/ # ps -e | grep com.docker.vpnkit
31031 root      0:00 grep com.docker.vpnkit



 Today this works on Desktop through a magic proxy called vpnkit running on OSX host ( ps -e | grep com.docker.vpnkit ) for more info which allows you a way to reach the OSX host via host.docker.internal (the DNS proxy in vpnkit sends this A record to the VM)
By adding a SRV Record in the DNS Server in the dockerd daemon that maps host.docker.internal to the Gateway-IP, containers running in a native Linux environment will be able to reach the Linux Host , but on MAC, containers will be sending traffic to the Linux VM and not the OSX host

If users were able to insert A records into dockerd's DNS server through cmdline/config options and so custom mappings such as host.docker.internal could be mapped to the Host IP / or interface, this functionality could be achieved (Note: this would not work for default docker0 bridge, since for this case, the host's resolv.conf is mounted so the nameserver does not point to dockerd). I will bring up this feature with the team .


https://docs.docker.com/engine/reference/commandline/run/#add-entries-to-container-hosts-file---add-host



 % kubectl get svc kube-dns -n kube-system  -o yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    prometheus.io/port: "9153"
    prometheus.io/scrape: "true"
  creationTimestamp: "2021-08-23T03:19:00Z"
  labels:
    k8s-app: kube-dns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: CoreDNS
  name: kube-dns
  namespace: kube-system
  resourceVersion: "284"
  uid: 02a8a10b-6aeb-4fc0-aea2-1c7f161598f5
spec:
  clusterIP: 10.96.0.10
  clusterIPs:
  - 10.96.0.10
  ipFamilies:
  - IPv4
  ipFamilyPolicy: SingleStack
  ports:
  - name: dns
    port: 53
    protocol: UDP
    targetPort: 53
  - name: dns-tcp
    port: 53
    protocol: TCP
    targetPort: 53
  - name: metrics
    port: 9153
    protocol: TCP
    targetPort: 9153
  selector:
    k8s-app: kube-dns
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}


   % kubectl -n kube-system  get pod |grep dns
coredns-558bd4d5db-qc6px                 1/1     Running   17         97d
coredns-558bd4d5db-wzcgb                 1/1     Running   17         97d


