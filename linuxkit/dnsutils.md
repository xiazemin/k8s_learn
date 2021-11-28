https://kubernetes.io/zh/docs/tasks/administer-cluster/dns-debugging-resolution/

apiVersion: v1
kind: Pod
metadata:
  name: dnsutils
  namespace: default
spec:
  containers:
  - name: dnsutils
    image: k8s.gcr.io/e2e-test-images/jessie-dnsutils:1.3
    command:
      - sleep
      - "3600"
    imagePullPolicy: IfNotPresent
  restartPolicy: Always



  kubectl exec -i -t dnsutils -- nslookup kubernetes.default



  https://hub.docker.com/r/tutum/dnsutils/tags?page=1&ordering=last_updated

  % docker pull tutum/dnsutils:latest
latest: Pulling from tutum/dnsutils


% kubectl apply -f linuxkit/dnsutil.pod.yaml 

pod/dnsutils created


%   kubectl exec -i -t dnsutils -- nslookup kubernetes.default
Server:         10.96.0.10
Address:        10.96.0.10#53

Name:   kubernetes.default.svc.cluster.local
Address: 10.96.0.1

%   kubectl exec -i -t dnsutils -- nslookup apple-service     
Server:         10.96.0.10
Address:        10.96.0.10#53

Name:   apple-service.default.svc.cluster.local
Address: 10.105.42.239



 % kubectl  -n kube-system exec -it coredns-558bd4d5db-qc6px -- sh
OCI runtime exec failed: exec failed: container_linux.go:380: starting container process caused: exec: "sh": executable file not found in $PATH: unknown
command terminated with exit code 126


% 
xiazemin@xiazemindeMacBook-Pro k8s_learn %   kubectl exec -i -t dnsutils -- dig apple-service 

; <<>> DiG 9.9.5-3ubuntu0.2-Ubuntu <<>> apple-service
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 22314
;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;apple-service.                 IN      A

;; Query time: 234 msec
;; SERVER: 10.96.0.10#53(10.96.0.10)
;; WHEN: Sun Nov 28 10:27:36 UTC 2021
;; MSG SIZE  rcvd: 42

https://blog.csdn.net/chuhe163/article/details/113487353
https://blog.csdn.net/weixin_34032621/article/details/91394935


%   kubectl exec -i -t dnsutils -- dig +trace apple-service A

; <<>> DiG 9.9.5-3ubuntu0.2-Ubuntu <<>> +trace apple-service A
;; global options: +cmd
;; Received 28 bytes from 10.96.0.10#53(10.96.0.10) in 8 ms


kubectl exec -i -t dnsutils -- nslookup kube-dns.kube-system 
Server:         10.96.0.10
Address:        10.96.0.10#53

Name:   kube-dns.kube-system.svc.cluster.local
Address: 10.96.0.10


%   kubectl exec -i -t dnsutils -- dig kube-dns.kube-system.svc.cluster.local

; <<>> DiG 9.9.5-3ubuntu0.2-Ubuntu <<>> kube-dns.kube-system.svc.cluster.local
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 17754
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;kube-dns.kube-system.svc.cluster.local.        IN A

;; ANSWER SECTION:
kube-dns.kube-system.svc.cluster.local. 6 IN A  10.96.0.10

;; Query time: 7 msec
;; SERVER: 10.96.0.10#53(10.96.0.10)
;; WHEN: Sun Nov 28 10:32:34 UTC 2021
;; MSG SIZE  rcvd: 121

