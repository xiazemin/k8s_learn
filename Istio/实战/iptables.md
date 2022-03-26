% docker run -it --privileged --pid=host ubuntu nsenter -t 1 -m -u -n -i sh

/ # iptables -nvL -t nat
Chain PREROUTING (policy ACCEPT 334 packets, 20073 bytes)
 pkts bytes target     prot opt in     out     source               destination
25659 2145K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
 2352  154K desktop    all  --  *      *       0.0.0.0/0            0.0.0.0/0
  226 13560 DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain INPUT (policy ACCEPT 150 packets, 9000 bytes)
 pkts bytes target     prot opt in     out     source               destination

Chain OUTPUT (policy ACCEPT 10660 packets, 645K bytes)
 pkts bytes target     prot opt in     out     source               destination
 136K 8203K KUBE-SERVICES  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes service portals */
30542 1833K DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain POSTROUTING (policy ACCEPT 10787 packets, 652K bytes)
 pkts bytes target     prot opt in     out     source               destination
 160K   10M KUBE-POSTROUTING  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* kubernetes postrouting rules */
    0     0 MASQUERADE  all  --  *      docker0  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match src-type LOCAL
    0     0 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0
    0     0 CNI-53e12e307d10cdbf76f0b5f3  all  --  *      *       10.1.4.63            0.0.0.0/0            /* name: "default" id: "4f7bca5dcba79bbcb2bb331b32bcc8eceeaeff662d2a400aaca02af4f022269d" */
    0     0 CNI-6b632f3e65065b9d4953df6e  all  --  *      *       10.1.4.69            0.0.0.0/0            /* name: "default" id: "680cb521ee93d98fc324b446642a8e18864bdbb3c4e65a6f0bdef1c5eb0a3f9a" */
  209 18263 CNI-c908956a107e9d9db83a6a70  all  --  *      *       10.1.4.73            0.0.0.0/0            /* name: "default" id: "44c024c2417c7f31a6b3e6d151e19bfc06e6cc3688631689390bf2f05b29ee4f" */
  202 17695 CNI-78ece3a0f0583aa66e03cb1f  all  --  *      *       10.1.4.79            0.0.0.0/0            /* name: "default" id: "feaac7512b6eeead1ffc7a11f501a54c122ce06fbfb288ebd93c1b3a3669e61d" */
    0     0 CNI-5fbcf167bf44def32ac0d583  all  --  *      *       10.1.4.87            0.0.0.0/0            /* name: "default" id: "e27106f180d666433139a4d7aa4625de35607a20b1d981b223cc89b1ad7d0fae" */
    0     0 CNI-2915e11a20b6350aa5e5ec79  all  --  *      *       10.1.4.91            0.0.0.0/0            /* name: "default" id: "9a8b6f61eba02ec4dd46d76b1d9b4c7cd65e886a39787f180d3c1a4792d9a371" */
  847 78952 CNI-3733a3c57758d082e2a2f9c8  all  --  *      *       10.1.4.92            0.0.0.0/0            /* name: "default" id: "024b4d066a6f83c1da711aaf663300591c4b27fd7f38ab1f5a8ec06422a2e5de" */
  811 76684 CNI-e7e80df250509dc798c3ad65  all  --  *      *       10.1.4.93            0.0.0.0/0            /* name: "default" id: "9bfb6341eca48ba043e168c4807d5009fa48f8d7f790395e7107354c6347d71f" */