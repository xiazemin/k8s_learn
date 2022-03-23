/ # ifconfig
cni0      Link encap:Ethernet  HWaddr B2:2A:59:02:9D:6C
          inet addr:10.1.0.1  Bcast:10.1.255.255  Mask:255.255.0.0
          inet6 addr: fe80::b02a:59ff:fe02:9d6c/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:121866 errors:0 dropped:0 overruns:0 frame:0
          TX packets:133745 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:12230308 (11.6 MiB)  TX bytes:107961619 (102.9 MiB)

docker0   Link encap:Ethernet  HWaddr 02:42:5A:D1:CB:E3
          inet addr:172.17.0.1  Bcast:172.17.255.255  Mask:255.255.0.0
          inet6 addr: fe80::42:5aff:fed1:cbe3/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:746 (746.0 B)

eth0      Link encap:Ethernet  HWaddr 02:50:00:00:00:01
          inet addr:192.168.65.3  Bcast:192.168.65.255  Mask:255.255.255.0
          inet6 addr: fe80::50:ff:fe00:1/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:20872 errors:0 dropped:0 overruns:0 frame:0
          TX packets:3155 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:29835130 (28.4 MiB)  TX bytes:218970 (213.8 KiB)

lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:570040 errors:0 dropped:0 overruns:0 frame:0
          TX packets:570040 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000
          RX bytes:288058370 (274.7 MiB)  TX bytes:288058370 (274.7 MiB)

services1 Link encap:Ethernet  HWaddr 52:89:A8:6C:0F:15
          inet addr:192.168.65.4  Bcast:0.0.0.0  Mask:255.255.255.255
          inet6 addr: fe80::5089:a8ff:fe6c:f15/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1607 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1617 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:163070 (159.2 KiB)  TX bytes:139001 (135.7 KiB)

veth05215c95 Link encap:Ethernet  HWaddr 6A:00:C4:12:36:7D
          inet6 addr: fe80::6800:c4ff:fe12:367d/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:3978 errors:0 dropped:0 overruns:0 frame:0
          TX packets:4119 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:8715778 (8.3 MiB)  TX bytes:335372 (327.5 KiB)

veth09cb50a Link encap:Ethernet  HWaddr 56:5A:20:32:83:59
          inet6 addr: fe80::545a:20ff:fe32:8359/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:14 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:1156 (1.1 KiB)

veth1aa90b96 Link encap:Ethernet  HWaddr 16:4D:19:E5:76:F4
          inet6 addr: fe80::144d:19ff:fee5:76f4/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:86 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4268 (4.1 KiB)

veth2789d3eb Link encap:Ethernet  HWaddr AA:17:C7:06:28:59
          inet6 addr: fe80::a817:c7ff:fe06:2859/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7664 errors:0 dropped:0 overruns:0 frame:0
          TX packets:8046 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1857046 (1.7 MiB)  TX bytes:5344927 (5.0 MiB)

veth3585213d Link encap:Ethernet  HWaddr 62:66:98:C0:08:CF
          inet6 addr: fe80::6066:98ff:fec0:8cf/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:5356 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7150 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:5149945 (4.9 MiB)  TX bytes:1448930 (1.3 MiB)

veth4478fddb Link encap:Ethernet  HWaddr 12:CB:F0:20:8A:02
          inet6 addr: fe80::10cb:f0ff:fe20:8a02/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:8466 errors:0 dropped:0 overruns:0 frame:0
          TX packets:8886 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1246215 (1.1 MiB)  TX bytes:5364548 (5.1 MiB)

veth4ce493c1 Link encap:Ethernet  HWaddr CE:7D:1B:D0:D3:6D
          inet6 addr: fe80::cc7d:1bff:fed0:d36d/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7164 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7540 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:968085 (945.3 KiB)  TX bytes:5120119 (4.8 MiB)

veth5250be2b Link encap:Ethernet  HWaddr DE:1E:7B:DD:BA:25
          inet6 addr: fe80::dc1e:7bff:fedd:ba25/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:18734 errors:0 dropped:0 overruns:0 frame:0
          TX packets:19950 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:2226367 (2.1 MiB)  TX bytes:71493781 (68.1 MiB)

veth5668ac85 Link encap:Ethernet  HWaddr 66:03:CE:B8:DF:E6
          inet6 addr: fe80::6403:ceff:feb8:dfe6/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:87 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4310 (4.2 KiB)

veth5c563d2b Link encap:Ethernet  HWaddr AE:C0:CF:8A:64:DF
          inet6 addr: fe80::acc0:cfff:fe8a:64df/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7995 errors:0 dropped:0 overruns:0 frame:0
          TX packets:8322 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:2569641 (2.4 MiB)  TX bytes:5215462 (4.9 MiB)

veth6b9acdd6 Link encap:Ethernet  HWaddr CE:FE:AA:7E:9A:10
          inet6 addr: fe80::ccfe:aaff:fe7e:9a10/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:9595 errors:0 dropped:0 overruns:0 frame:0
          TX packets:11649 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:5936411 (5.6 MiB)  TX bytes:2676186 (2.5 MiB)

veth6eb86be5 Link encap:Ethernet  HWaddr 72:76:21:7B:3B:98
          inet6 addr: fe80::7076:21ff:fe7b:3b98/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:5833 errors:0 dropped:0 overruns:0 frame:0
          TX packets:5990 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1020058 (996.1 KiB)  TX bytes:1805798 (1.7 MiB)

veth8b313c75 Link encap:Ethernet  HWaddr 62:0D:C0:62:39:43
          inet6 addr: fe80::600d:c0ff:fe62:3943/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:6979 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7315 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:995098 (971.7 KiB)  TX bytes:4976648 (4.7 MiB)

veth9b52fd14 Link encap:Ethernet  HWaddr CE:9D:28:AC:2B:AE
          inet6 addr: fe80::cc9d:28ff:feac:2bae/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:251 errors:0 dropped:0 overruns:0 frame:0
          TX packets:341 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:33150 (32.3 KiB)  TX bytes:190852 (186.3 KiB)

veth9e60ddbf Link encap:Ethernet  HWaddr 76:B2:CE:F8:37:5A
          inet6 addr: fe80::74b2:ceff:fef8:375a/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:4300 errors:0 dropped:0 overruns:0 frame:0
          TX packets:4634 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:890567 (869.6 KiB)  TX bytes:709160 (692.5 KiB)

vetha0a59e24 Link encap:Ethernet  HWaddr 6A:EB:5F:A0:6D:B1
          inet6 addr: fe80::68eb:5fff:fea0:6db1/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:97 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4758 (4.6 KiB)

vetha7760f41 Link encap:Ethernet  HWaddr AE:7F:8F:39:EA:82
          inet6 addr: fe80::ac7f:8fff:fe39:ea82/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7365 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7708 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1012080 (988.3 KiB)  TX bytes:5136221 (4.8 MiB)

vethaac8b33a Link encap:Ethernet  HWaddr C6:A8:CE:B1:73:A0
          inet6 addr: fe80::c4a8:ceff:feb1:73a0/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:90 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4436 (4.3 KiB)

vethae03c31f Link encap:Ethernet  HWaddr 16:8E:27:48:03:48
          inet6 addr: fe80::148e:27ff:fe48:348/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:759 errors:0 dropped:0 overruns:0 frame:0
          TX packets:982 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:252672 (246.7 KiB)  TX bytes:102444 (100.0 KiB)

vethafcac197 Link encap:Ethernet  HWaddr D6:3B:97:8A:95:C1
          inet6 addr: fe80::d43b:97ff:fe8a:95c1/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:4556 errors:0 dropped:0 overruns:0 frame:0
          TX packets:4894 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:927368 (905.6 KiB)  TX bytes:732011 (714.8 KiB)

vethb312dd11 Link encap:Ethernet  HWaddr BE:4C:5A:15:2B:50
          inet6 addr: fe80::bc4c:5aff:fe15:2b50/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:84 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4280 (4.1 KiB)

vethb40f0238 Link encap:Ethernet  HWaddr D2:4B:08:3E:40:1E
          inet6 addr: fe80::d04b:8ff:fe3e:401e/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:93 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:42 (42.0 B)  TX bytes:4590 (4.4 KiB)

vethbee26119 Link encap:Ethernet  HWaddr 56:CC:DA:2B:DB:D1
          inet6 addr: fe80::54cc:daff:fe2b:dbd1/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:11579 errors:0 dropped:0 overruns:0 frame:0
          TX packets:12549 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1792242 (1.7 MiB)  TX bytes:6999172 (6.6 MiB)

vethc2bc04a4 Link encap:Ethernet  HWaddr 72:DF:EC:C1:84:D7
          inet6 addr: fe80::70df:ecff:fec1:84d7/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1687 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1992 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:235177 (229.6 KiB)  TX bytes:182569 (178.2 KiB)

vethcbe5a408 Link encap:Ethernet  HWaddr 7A:D2:12:F1:4E:BA
          inet6 addr: fe80::78d2:12ff:fef1:4eba/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7642 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7949 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1732691 (1.6 MiB)  TX bytes:5065434 (4.8 MiB)

vethce0d77ba Link encap:Ethernet  HWaddr 66:79:B7:1F:A2:41
          inet6 addr: fe80::6479:b7ff:fe1f:a241/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:5294 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7100 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:5400350 (5.1 MiB)  TX bytes:1272919 (1.2 MiB)

vethdf94e690 Link encap:Ethernet  HWaddr E6:D1:4E:C2:F2:34
          inet6 addr: fe80::e4d1:4eff:fec2:f234/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7106 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7513 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:939869 (917.8 KiB)  TX bytes:4949628 (4.7 MiB)

vethdfed1836 Link encap:Ethernet  HWaddr 7A:B2:3B:BB:47:5E
          inet6 addr: fe80::78b2:3bff:febb:475e/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:9135 errors:0 dropped:0 overruns:0 frame:0
          TX packets:10294 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1896641 (1.8 MiB)  TX bytes:6613519 (6.3 MiB)

vethe1947ecf Link encap:Ethernet  HWaddr D2:50:8D:B4:4F:02
          inet6 addr: fe80::d050:8dff:feb4:4f02/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:7132 errors:0 dropped:0 overruns:0 frame:0
          TX packets:7520 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:941941 (919.8 KiB)  TX bytes:5127803 (4.8 MiB)

/



/ # traceroute 127.0.0.1
traceroute to 127.0.0.1 (127.0.0.1), 30 hops max, 46 byte packets
 1  localhost (127.0.0.1)  0.014 ms  0.052 ms  0.157 ms
/ #




/ # ping -c 3 127.0.0.1
PING 127.0.0.1 (127.0.0.1): 56 data bytes
64 bytes from 127.0.0.1: seq=0 ttl=64 time=2.548 ms
64 bytes from 127.0.0.1: seq=1 ttl=64 time=2.119 ms
64 bytes from 127.0.0.1: seq=2 ttl=64 time=0.317 ms

--- 127.0.0.1 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.317/1.661/2.548 ms

/ # ping -c 3 127.0.3.1
PING 127.0.3.1 (127.0.3.1): 56 data bytes
64 bytes from 127.0.3.1: seq=0 ttl=64 time=5.142 ms
64 bytes from 127.0.3.1: seq=1 ttl=64 time=0.184 ms
64 bytes from 127.0.3.1: seq=2 ttl=64 time=0.104 ms

--- 127.0.3.1 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.104/1.810/5.142 ms


