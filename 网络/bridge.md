bash-5.0# ip link add wangpo type bridge
bash-5.0#
bash-5.0# ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
6: wangpo: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 22:01:ca:58:3d:59 brd ff:ff:ff:ff:ff:ff


bash-5.0# ip link add wp2xmq type veth peer name xmq2wp
bash-5.0# ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
6: wangpo: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 22:01:ca:58:3d:59 brd ff:ff:ff:ff:ff:ff
7: xmq2wp@wp2xmq: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff
8: wp2xmq@xmq2wp: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether d2:98:28:54:86:aa brd ff:ff:ff:ff:ff:ff



bash-5.0# ip link set xmq2wp netns ximenqin
bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
7: xmq2wp@if8: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden


bash-5.0# ip link set wp2xmq master wangpo


bash-5.0# ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
6: wangpo: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether d2:98:28:54:86:aa brd ff:ff:ff:ff:ff:ff
8: wp2xmq@if7: <BROADCAST,MULTICAST> mtu 1500 qdisc noop master wangpo state DOWN mode DEFAULT group default qlen 1000
    link/ether d2:98:28:54:86:aa brd ff:ff:ff:ff:ff:ff link-netns ximenqin
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST> mtu 1500 master wangpo state disabled priority 32 cost 2








bash-5.0# ip netns exec ximenqin ip addr add dev xmq2wp 192.168.187.96/24
bash-5.0#
bash-5.0# ip netns exec ximenqin ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
    inet 192.168.188.96/24 scope global ximenqin
       valid_lft forever preferred_lft forever
    inet6 fe80::804f:1ff:fedd:cfd3/64 scope link
       valid_lft forever preferred_lft forever
7: xmq2wp@if8: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.96/24 scope global xmq2wp
       valid_lft forever preferred_lft forever
bash-5.0#
bash-5.0#
bash-5.0# ip netns exec ximenqin ip link set xmq2wp up
bash-5.0# ip netns exec ximenqin ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
    inet 192.168.188.96/24 scope global ximenqin
       valid_lft forever preferred_lft forever
    inet6 fe80::804f:1ff:fedd:cfd3/64 scope link
       valid_lft forever preferred_lft forever
7: xmq2wp@if8: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state LOWERLAYERDOWN group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.96/24 scope global xmq2wp
       valid_lft forever preferred_lft forever
bash-5.0#
bash-5.0# ip link set wp2xmq up
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state disabled priority 32 cost 2
bash-5.0# ip netns exec ximenqin ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
    inet 192.168.188.96/24 scope global ximenqin
       valid_lft forever preferred_lft forever
    inet6 fe80::804f:1ff:fedd:cfd3/64 scope link
       valid_lft forever preferred_lft forever
7: xmq2wp@if8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.96/24 scope global xmq2wp
       valid_lft forever preferred_lft forever
    inet6 fe80::bc27:e3ff:fe04:8513/64 scope link
       valid_lft forever preferred_lft forever




bash-5.0# ip link add wp2pjl type veth peer name pjl2wp
bash-5.0# ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
6: wangpo: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether d2:98:28:54:86:aa brd ff:ff:ff:ff:ff:ff
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master wangpo state UP mode DEFAULT group default qlen 1000
    link/ether d2:98:28:54:86:aa brd ff:ff:ff:ff:ff:ff link-netns ximenqin
9: pjl2wp@wp2pjl: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff
10: wp2pjl@pjl2wp: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether fa:c8:3e:62:73:4d brd ff:ff:ff:ff:ff:ff
bash-5.0# ip link set pjl2wp netns panjinlian
bash-5.0#
bash-5.0# ip netns exec panjinlian ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
9: pjl2wp@if10: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff link-netns eden
bash-5.0#
bash-5.0# ip link set wp2pjl master wangpo
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state disabled priority 32 cost 2
10: wp2pjl@if9: <BROADCAST,MULTICAST> mtu 1500 master wangpo state disabled priority 32 cost 2
bash-5.0#
bash-5.0#




bash-5.0# ip netns exec panjinlian ip addr add dev pjl2wp 192.168.187.69/24
bash-5.0#
bash-5.0# ip netns exec panjinlian ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
    inet 192.168.188.69/4 scope global panjinlian
       valid_lft forever preferred_lft forever
    inet6 fe80::5474:2eff:fe4e:4d48/64 scope link
       valid_lft forever preferred_lft forever
9: pjl2wp@if10: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.69/24 scope global pjl2wp
       valid_lft forever preferred_lft forever
bash-5.0# ip netns exec panjinlian ip link pjl2wp up
Command "pjl2wp" is unknown, try "ip link help".
bash-5.0# ip netns exec panjinlian ip link set pjl2wp up
bash-5.0# ip netns exec panjinlian ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
    inet 192.168.188.69/4 scope global panjinlian
       valid_lft forever preferred_lft forever
    inet6 fe80::5474:2eff:fe4e:4d48/64 scope link
       valid_lft forever preferred_lft forever
9: pjl2wp@if10: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state LOWERLAYERDOWN group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.69/24 scope global pjl2wp
       valid_lft forever preferred_lft forever
bash-5.0# ip link set wp2pjl up
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state disabled priority 32 cost 2
10: wp2pjl@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state disabled priority 32 cost 2



bash-5.0# ip netns exec ximenqin ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes

--- 192.168.187.69 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss




bash-5.0# ip netns exec ximenqin ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes

--- 192.168.187.69 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss
bash-5.0#
bash-5.0#
bash-5.0# ip netns exec ximenqin link
BusyBox v1.31.1 () multi-call binary.

Usage: link FILE LINK

Create hard LINK to FILE
bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
7: xmq2wp@if8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden
bash-5.0# ip netns exec ximenqin ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
    inet 192.168.188.96/24 scope global ximenqin
       valid_lft forever preferred_lft forever
    inet6 fe80::804f:1ff:fedd:cfd3/64 scope link
       valid_lft forever preferred_lft forever
7: xmq2wp@if8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether be:27:e3:04:85:13 brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.96/24 scope global xmq2wp
       valid_lft forever preferred_lft forever
    inet6 fe80::bc27:e3ff:fe04:8513/64 scope link
       valid_lft forever preferred_lft forever
bash-5.0#
bash-5.0# ip netns exec panjinlian ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
    inet 192.168.188.69/4 scope global panjinlian
       valid_lft forever preferred_lft forever
    inet6 fe80::5474:2eff:fe4e:4d48/64 scope link
       valid_lft forever preferred_lft forever
9: pjl2wp@if10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.69/24 scope global pjl2wp
       valid_lft forever preferred_lft forever
    inet6 fe80::1c28:80ff:fe97:658a/64 scope link
       valid_lft forever preferred_lft forever
bash-5.0



bash-5.0# ip link set wangpo up
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state forwarding priority 32 cost 2
10: wp2pjl@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state forwarding priority 32 cost 2
bash-5.0#



bash-5.0# ip link set wangpo up
bash-5.0# bridge link
8: wp2xmq@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state forwarding priority 32 cost 2
10: wp2pjl@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 master wangpo state forwarding priority 32 cost 2
bash-5.0#
bash-5.0#
bash-5.0# ip netns exec ximenqin ping -c 3 192.168.187.96
PING 192.168.187.96 (192.168.187.96): 56 data bytes
^C
--- 192.168.187.96 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss
bash-5.0# ip netns exec ximenqin ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes
64 bytes from 192.168.187.69: seq=0 ttl=64 time=19.218 ms
64 bytes from 192.168.187.69: seq=1 ttl=64 time=0.102 ms
64 bytes from 192.168.187.69: seq=2 ttl=64 time=0.083 ms

--- 192.168.187.69 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.083/6.467/19.218 ms
bash-5.0#
bash-5.0# ip netns exec panjinlian ping -c 192.168.187.69
ping: invalid number '192.168.187.69'
bash-5.0# ip netns exec panjinlian ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes

--- 192.168.187.69 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss
bash-5.0#


bash-5.0# ip netns exec panjinlian ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
    inet 192.168.188.69/4 scope global panjinlian
       valid_lft forever preferred_lft forever
    inet6 fe80::5474:2eff:fe4e:4d48/64 scope link
       valid_lft forever preferred_lft forever
9: pjl2wp@if10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 1e:28:80:97:65:8a brd ff:ff:ff:ff:ff:ff link-netns eden
    inet 192.168.187.69/24 scope global pjl2wp
       valid_lft forever preferred_lft forever
    inet6 fe80::1c28:80ff:fe97:658a/64 scope link
       valid_lft forever preferred_lft forever
bash-5.0#
bash-5.0#
bash-5.0# ip netns exec panjinlian ping -c 3 192.168.187.96
PING 192.168.187.96 (192.168.187.96): 56 data bytes
64 bytes from 192.168.187.96: seq=0 ttl=64 time=0.065 ms
64 bytes from 192.168.187.96: seq=1 ttl=64 time=0.081 ms
64 bytes from 192.168.187.96: seq=2 ttl=64 time=0.073 ms

--- 192.168.187.96 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.065/0.073/0.081 ms
bash-5.0# ip netns exec panjinlian ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes


--- 192.168.187.69 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss


bash-5.0# ip netns exec ximenqin ping -c 3 192.168.187.69
PING 192.168.187.69 (192.168.187.69): 56 data bytes
64 bytes from 192.168.187.69: seq=0 ttl=64 time=0.152 ms
64 bytes from 192.168.187.69: seq=1 ttl=64 time=0.071 ms
64 bytes from 192.168.187.69: seq=2 ttl=64 time=0.098 ms

--- 192.168.187.69 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.071/0.107/0.152 ms

