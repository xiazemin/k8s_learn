/ # iptables -t nat -nvL
Chain PREROUTING (policy ACCEPT 194 packets, 11981 bytes)
 pkts bytes target     prot opt in     out     source               destination
23979 1651K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
12998  833K desktop    all  --  *      *       0.0.0.0/0            0.0.0.0/0
    4  5023 DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain INPUT (policy ACCEPT 27 packets, 1620 bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain OUTPUT (policy ACCEPT 7134 packets, 436K bytes)
 pkts bytes target     prot opt in     out     source               destination
 471K   28M KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
 306K   18M DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain POSTROUTING (policy ACCEPT 7154 packets, 438K bytes)
 pkts bytes target     prot opt in     out     source               destination
 493K   30M KUBE-POSTROUTING  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */
  101  6060 MASQUERADE  all  --  *      docker0  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match src-type LOCAL
 9599  576K MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0
    1    40 MASQUERADE  all  --  *      br-421a9e16d12b  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match src-type LOCAL
    0     0 MASQUERADE  all  --  *      !br-421a9e16d12b  172.18.0.0/16        0.0.0.0/0
    1    60 CNI-0d314787d804206fd0db3f98  all  --  *      *       10.1.0.214           0.0.0.0/0            /* name: "default" id: "d56bc6d2512b0f78cebc9ad1ff17c1fc725ad61a952db34a7980e2d0d783d108" */
    0     0 CNI-48f8ec86c3b0a658588868a6  all  --  *      *       10.1.0.215           0.0.0.0/0            /* name: "default" id: "af82f2fdcea925dc1710fe6d0c9c276509f8e6cd7a794cd3743ea843b677b7f5" */
 1036 73821 CNI-91b2cd657fb8b0fedb8c77ad  all  --  *      *       10.1.0.216           0.0.0.0/0            /* name: "default" id: "c657423898aef5e5327acd484ef00d00242175702ac40cca490a6f25a1242bfe" */
    0     0 CNI-fab4888526a9bec71834f286  all  --  *      *       10.1.0.217           0.0.0.0/0            /* name: "default" id: "778a2bc9fb4e6afba59cc04cefd72e1b0d88124b7d6e6b007197811ac3eec346" */
 1028 73250 CNI-137db40c6a5fbc453fbbcfce  all  --  *      *       10.1.0.218           0.0.0.0/0            /* name: "default" id: "ec757b45cd049459142a9347627310fdba91799d5250309e7f2e5dac5cfd13b0" */
   31  1860 CNI-6563068de8207c093cbc345d  all  --  *      *       10.1.0.220           0.0.0.0/0            /* name: "default" id: "293e15775fa73def3d7470b91c3feb05bb7438638ddf20743efd8d9d4e895de3" */
    0     0 CNI-40508a9bf71b10c255802fb8  all  --  *      *       10.1.0.221           0.0.0.0/0            /* name: "default" id: "02c1e14ac96f8e807f79c6d18609e6345746f37d8e269d2cfaa3b9e4ed6aca2d" */
    0     0 CNI-e4d1efb036294232e4444b5f  all  --  *      *       10.1.0.222           0.0.0.0/0            /* name: "default" id: "71b423acc07836a756a2d34d31db93d5319d1e7fd75d206a6926631003eda89d" */
   18  1236 CNI-e3e4ea5629d0c5627784ddd4  all  --  *      *       10.1.0.223           0.0.0.0/0            /* name: "default" id: "f474c66ddaa93f99db2ec563dc6d77f69c863b25ed00b7cde35eb5c3b058cc8a" */
10623  805K CNI-e93ed372dfdce36419bee3f1  all  --  *      *       10.1.0.224           0.0.0.0/0            /* name: "default" id: "6856a2970e865ce1463eee892de5a5f14be32c2495f4a030cb3fb55034e3282f" */
    0     0 MASQUERADE  tcp  --  *      *       172.17.0.2           172.17.0.2           tcp dpt:443
    0     0 MASQUERADE  tcp  --  *      *       172.17.0.2           172.17.0.2           tcp dpt:80

Chain CNI-0d314787d804206fd0db3f98 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "d56bc6d2512b0f78cebc9ad1ff17c1fc725ad61a952db34a7980e2d0d783d108" */
    1    60 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "d56bc6d2512b0f78cebc9ad1ff17c1fc725ad61a952db34a7980e2d0d783d108" */

Chain CNI-137db40c6a5fbc453fbbcfce (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "ec757b45cd049459142a9347627310fdba91799d5250309e7f2e5dac5cfd13b0" */
 1028 73250 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "ec757b45cd049459142a9347627310fdba91799d5250309e7f2e5dac5cfd13b0" */

Chain CNI-40508a9bf71b10c255802fb8 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "02c1e14ac96f8e807f79c6d18609e6345746f37d8e269d2cfaa3b9e4ed6aca2d" */
    0     0 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "02c1e14ac96f8e807f79c6d18609e6345746f37d8e269d2cfaa3b9e4ed6aca2d" */

Chain CNI-48f8ec86c3b0a658588868a6 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "af82f2fdcea925dc1710fe6d0c9c276509f8e6cd7a794cd3743ea843b677b7f5" */
    0     0 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "af82f2fdcea925dc1710fe6d0c9c276509f8e6cd7a794cd3743ea843b677b7f5" */

Chain CNI-6563068de8207c093cbc345d (1 references)
 pkts bytes target     prot opt in     out     source               destination
   30  1800 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "293e15775fa73def3d7470b91c3feb05bb7438638ddf20743efd8d9d4e895de3" */
    1    60 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "293e15775fa73def3d7470b91c3feb05bb7438638ddf20743efd8d9d4e895de3" */

Chain CNI-91b2cd657fb8b0fedb8c77ad (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "c657423898aef5e5327acd484ef00d00242175702ac40cca490a6f25a1242bfe" */
 1036 73821 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "c657423898aef5e5327acd484ef00d00242175702ac40cca490a6f25a1242bfe" */

Chain CNI-e3e4ea5629d0c5627784ddd4 (1 references)
 pkts bytes target     prot opt in     out     source               destination
   16  1112 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "f474c66ddaa93f99db2ec563dc6d77f69c863b25ed00b7cde35eb5c3b058cc8a" */
    2   124 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "f474c66ddaa93f99db2ec563dc6d77f69c863b25ed00b7cde35eb5c3b058cc8a" */

Chain CNI-e4d1efb036294232e4444b5f (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "71b423acc07836a756a2d34d31db93d5319d1e7fd75d206a6926631003eda89d" */
    0     0 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "71b423acc07836a756a2d34d31db93d5319d1e7fd75d206a6926631003eda89d" */

Chain CNI-e93ed372dfdce36419bee3f1 (1 references)
 pkts bytes target     prot opt in     out     source               destination
 9382  722K ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "6856a2970e865ce1463eee892de5a5f14be32c2495f4a030cb3fb55034e3282f" */
 1241 83762 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "6856a2970e865ce1463eee892de5a5f14be32c2495f4a030cb3fb55034e3282f" */

Chain CNI-fab4888526a9bec71834f286 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            10.1.0.0/16          /* name: "default" id: "778a2bc9fb4e6afba59cc04cefd72e1b0d88124b7d6e6b007197811ac3eec346" */
    0     0 MASQUERADE  all  --  *      *       0.0.0.0/0           !224.0.0.0/4          /* name: "default" id: "778a2bc9fb4e6afba59cc04cefd72e1b0d88124b7d6e6b007197811ac3eec346" */

Chain DOCKER (2 references)
 pkts bytes target     prot opt in     out     source               destination
   23  1380 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:7443 to:172.17.0.2:443
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:7071 to:172.17.0.2:80

Chain KUBE-KUBELET-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-MARK-DROP (0 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            MARK or 0x8000

Chain KUBE-MARK-MASQ (26 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            MARK or 0x4000

Chain KUBE-NODEPORTS (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:api */ tcp dpt:30000
    0     0 KUBE-SVC-SXW22BMJJ7T3N2OP  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:api */ tcp dpt:30000
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       127.0.0.0/8          0.0.0.0/0            /* default/ingress-nginx-controller:https */ tcp dpt:32342
    0     0 KUBE-XLB-Q7CDIBSFDYNOJNFE  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:https */ tcp dpt:32342
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/etcd-svc-docker-desktop-xzm */ tcp dpt:32389
    0     0 KUBE-SVC-EDBHCP4VQID7F5J2  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/etcd-svc-docker-desktop-xzm */ tcp dpt:32389
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/redis:tcp */ tcp dpt:30379
    0     0 KUBE-SVC-CKFHGLZY3HDORVFT  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/redis:tcp */ tcp dpt:30379
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/apple-service */ tcp dpt:30080
    0     0 KUBE-SVC-Y4TE457BRBWMNDKG  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/apple-service */ tcp dpt:30080
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       127.0.0.0/8          0.0.0.0/0            /* default/ingress-nginx-controller:http */ tcp dpt:30701
    0     0 KUBE-XLB-D7TXZ2ONB4DT7BQA  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:http */ tcp dpt:30701
    0     0 KUBE-MARK-MASQ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:console */ tcp dpt:30001
    0     0 KUBE-SVC-ED7LY7V3PRCUB6IJ  tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:console */ tcp dpt:30001

Chain KUBE-POSTROUTING (1 references)
 pkts bytes target     prot opt in     out     source               destination
 7476  459K RETURN     all  --  *      *       0.0.0.0/0            0.0.0.0/0            mark match ! 0x4000/0x4000
    0     0 MARK       all  --  *      *       0.0.0.0/0            0.0.0.0/0            MARK xor 0x4000
    0     0 MASQUERADE  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service traffic requiring SNAT */ random-fully

Chain KUBE-PROXY-CANARY (0 references)
 pkts bytes target     prot opt in     out     source               destination

Chain KUBE-SEP-3JXLGUMKBPWA3G7K (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.218           0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */ tcp to:10.1.0.218:53

Chain KUBE-SEP-4XKQ3BKJH4UM3LNT (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.220           0.0.0.0/0            /* default/ingress-nginx-controller-admission:https-webhook */
   20  1200 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller-admission:https-webhook */ tcp to:10.1.0.220:8443

Chain KUBE-SEP-5EF3GOGMU4HLZ57C (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.216           0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */ tcp to:10.1.0.216:53

Chain KUBE-SEP-5VPUENENYDAFRS6E (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.220           0.0.0.0/0            /* default/ingress-nginx-controller:http */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:http */ tcp to:10.1.0.220:80

Chain KUBE-SEP-FEVJFI3I56DO4VH7 (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.220           0.0.0.0/0            /* default/ingress-nginx-controller:https */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:https */ tcp to:10.1.0.220:443

Chain KUBE-SEP-FLOJ2NBWC6RHDXIG (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       192.168.65.4         0.0.0.0/0            /* kube-system/etcd-svc-docker-desktop-xzm */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/etcd-svc-docker-desktop-xzm */ tcp to:192.168.65.4:2379

Chain KUBE-SEP-FYCASJGNWKVYLA7B (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       192.168.65.4         0.0.0.0/0            /* default/kubernetes:https */
   27  1620 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/kubernetes:https */ tcp to:192.168.65.4:6443

Chain KUBE-SEP-GXFFDQL3JO5LZDO7 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.218           0.0.0.0/0            /* kube-system/kube-dns:dns */
   61  4638 DNAT       udp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns */ udp to:10.1.0.218:53

Chain KUBE-SEP-JPYJVE4BPHWKYBDW (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.216           0.0.0.0/0            /* kube-system/kube-dns:dns */
   67  5090 DNAT       udp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns */ udp to:10.1.0.216:53

Chain KUBE-SEP-KWOQFGYL6KRGKJ23 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.218           0.0.0.0/0            /* kube-system/kube-dns:metrics */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:metrics */ tcp to:10.1.0.218:9153

Chain KUBE-SEP-N6ACUILWO7XJGY5F (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.217           0.0.0.0/0            /* default/apple-service */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/apple-service */ tcp to:10.1.0.217:5678

Chain KUBE-SEP-NGE6IH2X2N4M4XEF (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.224           0.0.0.0/0            /* cattle-system/cattle-cluster-agent:http */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cattle-system/cattle-cluster-agent:http */ tcp to:10.1.0.224:80

Chain KUBE-SEP-NSIQZ33RUVLUUTPS (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.224           0.0.0.0/0            /* cattle-system/cattle-cluster-agent:https-internal */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cattle-system/cattle-cluster-agent:https-internal */ tcp to:10.1.0.224:444

Chain KUBE-SEP-NVZBSVXMHBDHFW7S (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.223           0.0.0.0/0            /* default/minio-service:console */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:console */ tcp to:10.1.0.223:9001

Chain KUBE-SEP-OYXC66MH76XMJWAI (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.223           0.0.0.0/0            /* default/minio-service:api */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:api */ tcp to:10.1.0.223:9000

Chain KUBE-SEP-QCJUKRQWA7676OUJ (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.216           0.0.0.0/0            /* kube-system/kube-dns:metrics */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:metrics */ tcp to:10.1.0.216:9153

Chain KUBE-SEP-RAT4S652HRQXMAAD (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       10.1.0.222           0.0.0.0/0            /* default/redis:tcp */
    0     0 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/redis:tcp */ tcp to:10.1.0.222:6379

Chain KUBE-SERVICES (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SVC-RXZQBFX6IWO22WWW  tcp  --  *      *       0.0.0.0/0            10.101.247.125       /* cattle-system/cattle-cluster-agent:http cluster IP */ tcp dpt:80
    0     0 KUBE-SVC-SXW22BMJJ7T3N2OP  tcp  --  *      *       0.0.0.0/0            10.99.62.36          /* default/minio-service:api cluster IP */ tcp dpt:9000
    0     0 KUBE-SVC-ERIFXISQEP7F7OF4  tcp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:dns-tcp cluster IP */ tcp dpt:53
    0     0 KUBE-SVC-Q7CDIBSFDYNOJNFE  tcp  --  *      *       0.0.0.0/0            10.108.133.209       /* default/ingress-nginx-controller:https cluster IP */ tcp dpt:443
    0     0 KUBE-SVC-DISNXZXWEI7GIGLU  tcp  --  *      *       0.0.0.0/0            10.101.247.125       /* cattle-system/cattle-cluster-agent:https-internal cluster IP */ tcp dpt:443
    0     0 KUBE-SVC-EDBHCP4VQID7F5J2  tcp  --  *      *       0.0.0.0/0            10.111.136.178       /* kube-system/etcd-svc-docker-desktop-xzm cluster IP */ tcp dpt:2379
    0     0 KUBE-SVC-CKFHGLZY3HDORVFT  tcp  --  *      *       0.0.0.0/0            10.105.105.140       /* default/redis:tcp cluster IP */ tcp dpt:6379
  128  9728 KUBE-SVC-TCOU7JCQXEZGVUNU  udp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:dns cluster IP */ udp dpt:53
    0     0 KUBE-SVC-Y4TE457BRBWMNDKG  tcp  --  *      *       0.0.0.0/0            10.105.42.239        /* default/apple-service cluster IP */ tcp dpt:5678
    0     0 KUBE-SVC-D7TXZ2ONB4DT7BQA  tcp  --  *      *       0.0.0.0/0            10.108.133.209       /* default/ingress-nginx-controller:http cluster IP */ tcp dpt:80
   27  1620 KUBE-SVC-NPX46M4PTMTKRN6Y  tcp  --  *      *       0.0.0.0/0            10.96.0.1            /* default/kubernetes:https cluster IP */ tcp dpt:443
   20  1200 KUBE-SVC-XUD33RTORZBRAEIL  tcp  --  *      *       0.0.0.0/0            10.99.126.26         /* default/ingress-nginx-controller-admission:https-webhook cluster IP */ tcp dpt:443
    0     0 KUBE-SVC-ED7LY7V3PRCUB6IJ  tcp  --  *      *       0.0.0.0/0            10.99.62.36          /* default/minio-service:console cluster IP */ tcp dpt:9001
    0     0 KUBE-SVC-JD5MR3NA4I4DYORP  tcp  --  *      *       0.0.0.0/0            10.96.0.10           /* kube-system/kube-dns:metrics cluster IP */ tcp dpt:9153
 4617  277K KUBE-NODEPORTS  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service nodeports; NOTE: this must be the last rule in this chain */ ADDRTYPE match dst-type LOCAL

Chain KUBE-SVC-CKFHGLZY3HDORVFT (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-RAT4S652HRQXMAAD  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/redis:tcp */

Chain KUBE-SVC-D7TXZ2ONB4DT7BQA (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-5VPUENENYDAFRS6E  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:http */

Chain KUBE-SVC-DISNXZXWEI7GIGLU (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-NSIQZ33RUVLUUTPS  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cattle-system/cattle-cluster-agent:https-internal */

Chain KUBE-SVC-ED7LY7V3PRCUB6IJ (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-NVZBSVXMHBDHFW7S  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:console */

Chain KUBE-SVC-EDBHCP4VQID7F5J2 (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-FLOJ2NBWC6RHDXIG  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/etcd-svc-docker-desktop-xzm */

Chain KUBE-SVC-ERIFXISQEP7F7OF4 (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-5EF3GOGMU4HLZ57C  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */ statistic mode random probability 0.50000000000
    0     0 KUBE-SEP-3JXLGUMKBPWA3G7K  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns-tcp */

Chain KUBE-SVC-JD5MR3NA4I4DYORP (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-QCJUKRQWA7676OUJ  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:metrics */ statistic mode random probability 0.50000000000
    0     0 KUBE-SEP-KWOQFGYL6KRGKJ23  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:metrics */

Chain KUBE-SVC-NPX46M4PTMTKRN6Y (1 references)
 pkts bytes target     prot opt in     out     source               destination
   27  1620 KUBE-SEP-FYCASJGNWKVYLA7B  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/kubernetes:https */

Chain KUBE-SVC-Q7CDIBSFDYNOJNFE (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-FEVJFI3I56DO4VH7  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller:https */

Chain KUBE-SVC-RXZQBFX6IWO22WWW (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-NGE6IH2X2N4M4XEF  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cattle-system/cattle-cluster-agent:http */

Chain KUBE-SVC-SXW22BMJJ7T3N2OP (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-OYXC66MH76XMJWAI  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/minio-service:api */

Chain KUBE-SVC-TCOU7JCQXEZGVUNU (1 references)
 pkts bytes target     prot opt in     out     source               destination
   67  5090 KUBE-SEP-JPYJVE4BPHWKYBDW  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns */ statistic mode random probability 0.50000000000
   61  4638 KUBE-SEP-GXFFDQL3JO5LZDO7  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kube-system/kube-dns:dns */

Chain KUBE-SVC-XUD33RTORZBRAEIL (1 references)
 pkts bytes target     prot opt in     out     source               destination
   20  1200 KUBE-SEP-4XKQ3BKJH4UM3LNT  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/ingress-nginx-controller-admission:https-webhook */

Chain KUBE-SVC-Y4TE457BRBWMNDKG (2 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-SEP-N6ACUILWO7XJGY5F  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* default/apple-service */

Chain KUBE-XLB-D7TXZ2ONB4DT7BQA (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* masquerade LOCAL traffic for default/ingress-nginx-controller:http LB IP */ ADDRTYPE match src-type LOCAL
    0     0 KUBE-SVC-D7TXZ2ONB4DT7BQA  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* route LOCAL traffic for default/ingress-nginx-controller:http LB IP to service chain */ ADDRTYPE match src-type LOCAL
    0     0 KUBE-SEP-5VPUENENYDAFRS6E  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* Balancing rule 0 for default/ingress-nginx-controller:http */

Chain KUBE-XLB-Q7CDIBSFDYNOJNFE (1 references)
 pkts bytes target     prot opt in     out     source               destination
    0     0 KUBE-MARK-MASQ  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* masquerade LOCAL traffic for default/ingress-nginx-controller:https LB IP */ ADDRTYPE match src-type LOCAL
    0     0 KUBE-SVC-Q7CDIBSFDYNOJNFE  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* route LOCAL traffic for default/ingress-nginx-controller:https LB IP to service chain */ ADDRTYPE match src-type LOCAL
    0     0 KUBE-SEP-FEVJFI3I56DO4VH7  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* Balancing rule 0 for default/ingress-nginx-controller:https */

Chain desktop (1 references)
 pkts bytes target     prot opt in     out     source               destination
    2   955 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   955 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   204 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   921 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    5  2196 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    6  2768 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    5  1512 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    5  1511 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   204 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    1    68 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3  1023 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    2   136 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
   54  4041 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
   16  1775 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   995 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    4  1071 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   994 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   995 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   995 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    3   995 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    0     0 RETURN     tcp  --  *      *       0.0.0.0/0            172.17.0.0/16
    0     0 RETURN     tcp  --  *      *       0.0.0.0/0            172.18.0.0/16
    0     0 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.5
    6   360 RETURN     tcp  --  *      *       0.0.0.0/0            192.168.65.0/24
    0     0 RETURN     tcp  --  *      *       0.0.0.0/0            127.0.0.0/8
   30  1800 RETURN     tcp  --  *      *       0.0.0.0/0            10.1.0.0/16
    5   300 RETURN     tcp  --  *      *       0.0.0.0/0            10.96.0.0/12
/ # ?
