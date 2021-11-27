 % kubectl get  configmap kube-proxy -n kube-system -o yaml
apiVersion: v1
data:
  config.conf: |-
    apiVersion: kubeproxy.config.k8s.io/v1alpha1
    bindAddress: 0.0.0.0
    bindAddressHardFail: false
    clientConnection:
      acceptContentTypes: ""
      burst: 10
      contentType: application/vnd.kubernetes.protobuf
      kubeconfig: /var/lib/kube-proxy/kubeconfig.conf
      qps: 5
    clusterCIDR: ""
    configSyncPeriod: 15m0s
    conntrack:
      maxPerCore: 0
      min: 0
      tcpCloseWaitTimeout: 1h0m0s
      tcpEstablishedTimeout: 24h0m0s
    detectLocalMode: ""
    enableProfiling: false
    healthzBindAddress: 0.0.0.0:10256
    hostnameOverride: ""
    iptables:
      masqueradeAll: false
      masqueradeBit: 14
      minSyncPeriod: 0s
      syncPeriod: 30s
    ipvs:
      excludeCIDRs: null
      minSyncPeriod: 0s
      scheduler: ""
      strictARP: false
      syncPeriod: 30s
      tcpFinTimeout: 0s
      tcpTimeout: 0s
      udpTimeout: 0s
    kind: KubeProxyConfiguration
    metricsBindAddress: 127.0.0.1:10249
    mode: ""
    nodePortAddresses: null
    oomScoreAdj: -999
    portRange: ""
    showHiddenMetricsForVersion: ""
    udpIdleTimeout: 250ms
    winkernel:
      enableDSR: false
      networkName: ""
      sourceVip: ""
  kubeconfig.conf: |-
    apiVersion: v1
    kind: Config
    clusters:
    - cluster:
        certificate-authority: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
        server: https://vm.docker.internal:6443
      name: default
    contexts:
    - context:
        cluster: default
        namespace: default
        user: default
      name: default
    current-context: default
    users:
    - name: default
      user:
        tokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token
kind: ConfigMap
metadata:
  creationTimestamp: "2021-08-23T03:19:00Z"
  labels:
    app: kube-proxy
  name: kube-proxy
  namespace: kube-system
  resourceVersion: "288"
  uid: a11e8cb0-d15b-4283-ba20-9f8f7231cf60





  https://www.cnblogs.com/plefan/p/14966487.html

   % kubectl get svc
NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
apple-service                        NodePort       10.105.42.239    <none>        5678:30080/TCP                  94d



/ # iptables -t nat -L KUBE-SERVICES |egrep apple-service
KUBE-SVC-Y4TE457BRBWMNDKG  tcp  --  anywhere             10.105.42.239        /* default/apple-service cluster IP */ tcp dpt:5678
/ # iptables -t filter -L KUBE-SERVICES |egrep apple-service



/ # iptables -t nat -L KUBE-SERVICES
Chain KUBE-SERVICES (2 references)
target     prot opt source               destination
KUBE-SVC-DISNXZXWEI7GIGLU  tcp  --  anywhere             10.101.247.125       /* cattle-system/cattle-cluster-agent:https-internal cluster IP */ tcp dpt:https
KUBE-SVC-EDBHCP4VQID7F5J2  tcp  --  anywhere             10.111.136.178       /* kube-system/etcd-svc-docker-desktop-xzm cluster IP */ tcp dpt:2379
KUBE-SVC-Q7CDIBSFDYNOJNFE  tcp  --  anywhere             10.108.133.209       /* default/ingress-nginx-controller:https cluster IP */ tcp dpt:https
KUBE-SVC-CKFHGLZY3HDORVFT  tcp  --  anywhere             10.105.105.140       /* default/redis:tcp cluster IP */ tcp dpt:6379
KUBE-SVC-D7TXZ2ONB4DT7BQA  tcp  --  anywhere             10.108.133.209       /* default/ingress-nginx-controller:http cluster IP */ tcp dpt:http
KUBE-SVC-NPX46M4PTMTKRN6Y  tcp  --  anywhere             10.96.0.1            /* default/kubernetes:https cluster IP */ tcp dpt:https
KUBE-SVC-XUD33RTORZBRAEIL  tcp  --  anywhere             10.99.126.26         /* default/ingress-nginx-controller-admission:https-webhook cluster IP */ tcp dpt:https
KUBE-SVC-ED7LY7V3PRCUB6IJ  tcp  --  anywhere             10.99.62.36          /* default/minio-service:console cluster IP */ tcp dpt:9001
KUBE-SVC-JD5MR3NA4I4DYORP  tcp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:metrics cluster IP */ tcp dpt:9153
KUBE-SVC-TCOU7JCQXEZGVUNU  udp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:dns cluster IP */ udp dpt:domain
KUBE-SVC-Y4TE457BRBWMNDKG  tcp  --  anywhere             10.105.42.239        /* default/apple-service cluster IP */ tcp dpt:5678
KUBE-SVC-SXW22BMJJ7T3N2OP  tcp  --  anywhere             10.99.62.36          /* default/minio-service:api cluster IP */ tcp dpt:9000
KUBE-SVC-ERIFXISQEP7F7OF4  tcp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:dns-tcp cluster IP */ tcp dpt:domain
KUBE-SVC-RXZQBFX6IWO22WWW  tcp  --  anywhere             10.101.247.125       /* cattle-system/cattle-cluster-agent:http cluster IP */ tcp dpt:http
KUBE-NODEPORTS  all  --  anywhere             anywhere             /* kubernetes service nodeports; NOTE: this must be the last rule in this chain */ ADDRTYPE match dst-type LOCAL



/ #  iptables -t nat -L KUBE-SVC-Y4TE457BRBWMNDKG
Chain KUBE-SVC-Y4TE457BRBWMNDKG (2 references)
target     prot opt source               destination
KUBE-SEP-N6ACUILWO7XJGY5F  all  --  anywhere             anywhere             /* default/apple-service */



/ #  iptables -t nat -L KUBE-SEP-N6ACUILWO7XJGY5F
Chain KUBE-SEP-N6ACUILWO7XJGY5F (1 references)
target     prot opt source               destination
KUBE-MARK-MASQ  all  --  10.1.0.217           anywhere             /* default/apple-service */
DNAT       tcp  --  anywhere             anywhere             /* default/apple-service */ tcp to:10.1.0.217:5678



/ #  iptables -t nat -L KUBE-MARK-MASQ
Chain KUBE-MARK-MASQ (26 references)
target     prot opt source               destination
MARK       all  --  anywhere             anywhere             MARK or 0x4000





~ % kubectl get svc -n kube-system
NAME                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
etcd-svc-docker-desktop-xzm   NodePort    10.111.136.178   <none>        2379:32389/TCP           23d
kube-dns                      ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   96d



/ # iptables -t nat -L KUBE-SERVICES |grep kube-dns
KUBE-SVC-JD5MR3NA4I4DYORP  tcp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:metrics cluster IP */ tcp dpt:9153
KUBE-SVC-TCOU7JCQXEZGVUNU  udp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:dns cluster IP */ udp dpt:domain
KUBE-SVC-ERIFXISQEP7F7OF4  tcp  --  anywhere             10.96.0.10           /* kube-system/kube-dns:dns-tcp cluster IP */ tcp dpt:domain




/ # iptables -t nat -L KUBE-SVC-TCOU7JCQXEZGVUNU
Chain KUBE-SVC-TCOU7JCQXEZGVUNU (1 references)
target     prot opt source               destination
KUBE-SEP-JPYJVE4BPHWKYBDW  all  --  anywhere             anywhere             /* kube-system/kube-dns:dns */ statistic mode random probability 0.50000000000
KUBE-SEP-GXFFDQL3JO5LZDO7  all  --  anywhere             anywhere             /* kube-system/kube-dns:dns */
/ #
/ #
/ #
/ #
/ # iptables -t nat -L KUBE-SEP-JPYJVE4BPHWKYBDW
Chain KUBE-SEP-JPYJVE4BPHWKYBDW (1 references)
target     prot opt source               destination
KUBE-MARK-MASQ  all  --  10.1.0.216           anywhere             /* kube-system/kube-dns:dns */
DNAT       udp  --  anywhere             anywhere             /* kube-system/kube-dns:dns */ udp to:10.1.0.216:53
/ #
/ #
/ #
/ # iptables -t nat -L KUBE-SEP-GXFFDQL3JO5LZDO7
Chain KUBE-SEP-GXFFDQL3JO5LZDO7 (1 references)
target     prot opt source               destination
KUBE-MARK-MASQ  all  --  10.1.0.218           anywhere             /* kube-system/kube-dns:dns */
DNAT       udp  --  anywhere             anywhere             /* kube-system/kube-dns:dns */ udp to:10.1.0.218:53




