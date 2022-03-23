/ # ip netns add ximenqin
/ # ip netns add panjinlian
/ #
/ # ip netns ls
panjinlian
ximenqin
services (id: 0)


/ # ip link add ximenqin type veth peer name panjinlian

/ # ip link delete   panjinlian
/ # ip link delete   ximenqin
Cannot find device "ximenqin"
/ # ip link |grep ximenqin
RTNETLINK answers: Invalid argument
RTNETLINK answers: Invalid argument

bash-5.0# ip netns add ximenqin
bash-5.0# ip netns add panjinlian
bash-5.0# ip netns ls
panjinlian
ximenqin
eden
services


bash-5.0# ip link add ximenqin type veth peer name panjinlian



# ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@ximenqin: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff
5: ximenqin@panjinlian: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff


bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::



bash-5.0# ip link set ximenqin netns ximenqin
bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns eden



bash-5.0# ip link set panjinlian netns panjinlian
bash-5.0# ip netns exec panjinlian ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin



bash-5.0# iplink
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN qlen 1000
    link/tunnel6 00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00 brd 00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00
bash-5.0#




bash-5.0# ip netns exec ximenqin ip addr add dev ximenqin 192.168.188.96/24
bash-5.0#
bash-5.0# ip netns exec ximenqin ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
    inet 192.168.188.96/24 scope global ximenqin
       valid_lft forever preferred_lft forever
bash-5.0#
bash-5.0# ip netns exec panjinlian ip addr add dev panjinlian 192.168.188.69/4
bash-5.0#
bash-5.0#
bash-5.0# ip netns exec panjinlian ip addr
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
4: panjinlian@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 56:74:2e:4e:4d:48 brd ff:ff:ff:ff:ff:ff link-netns ximenqin
    inet 192.168.188.69/4 scope global panjinlian
       valid_lft forever preferred_lft forever




0# ip netns exec ximenqin ip link set ximenqin up
bash-5.0#
bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state LOWERLAYERDOWN mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
bash-5.0#
bash-5.0# ip netns exec panjinlian ip link set panjinlian up
bash-5.0#
bash-5.0# ip netns exec ximenqin ip link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/tunnel6 :: brd ::
5: ximenqin@if4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 82:4f:01:dd:cf:d3 brd ff:ff:ff:ff:ff:ff link-netns panjinlian
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
bash-5.0#



bash-5.0# ip netns exec ximenqin ping  -c 3 192.168.188.69
PING 192.168.188.69 (192.168.188.69): 56 data bytes
64 bytes from 192.168.188.69: seq=0 ttl=64 time=0.770 ms
64 bytes from 192.168.188.69: seq=1 ttl=64 time=0.057 ms
64 bytes from 192.168.188.69: seq=2 ttl=64 time=0.080 ms

--- 192.168.188.69 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.057/0.302/0.770 ms



