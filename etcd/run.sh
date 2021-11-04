kubectl get  pod etcd-docker-desktop  -n kube-system -o yaml >etcd/etcd-docker-desktop.pod.yaml

kubectl get svc kube-dns -n kube-system -o yaml > etcd/kube-dns.svc.yaml 

% kubectl apply -f etcd/etcd-svc-docker-desktop-xzm.svc.yaml 
service/etcd-svc-docker-desktop-xzm created

 % kubectl get svc etcd-svc-docker-desktop-xzm -n kube-system 
NAME                          TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
etcd-svc-docker-desktop-xzm   NodePort   10.105.187.254   <none>        2379:32379/TCP   35s


ETCDCTL_API=3 etcdctl --cacert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/ca.crt \
--cert=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.crt \
--key=/Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key \
--endpoints=https://127.0.0.1:32379 \
get /registry/namespaces -w=json | jq .


% docker cp 9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e:/run/config/pki/etcd/ca.crt etcd 
% docker cp 9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e:/run/config/pki/etcd/peer.crt etcd
% docker cp 9d2f3cf9d7bc9a933a216f6560eb22a4c62d59555bd449161ce8a58afce29d4e:/run/config/pki/etcd/peer.key etcd 

ETCDCTL_API=3 etcdctl --cacert=./etcd/ca.crt \
--cert=./etcd/peer.crt \
--key=./etcd/peer.key \
--endpoints=https://127.0.0.1:32379 \
get /registry/namespaces -w=json | jq .


 vimdiff ./etcd/peer.key /Users/xiazemin/Library/Group\ Containers/group.com.docker/pki/etcd/peer.key


#https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#expose

kubectl -n kube-system expose pod etcd-docker-desktop --port=2379 --type=NodePort --name=etcd-svc-docker-desktop-xzm

service/etcd-svc-docker-desktop-xzm exposed


https://blog.csdn.net/xinghun_4/article/details/50492041

https://stackoverflow.com/questions/52522570/how-to-expose-a-kubernetes-service-on-a-specific-nodeport/52523675


kubectl -n kube-system get svc etcd-svc-docker-desktop-xzm -o yaml > etcd/etcd-svc-docker-desktop-xzm_expose.svc.yaml

kubectl apply -f etcd/etcd-svc-docker-desktop-xzm_expose.svc.yaml 

% telnet 127.0.0.1 32379


% ETCDCTL_API=3 etcdctl --cacert=./etcd/ca.crt \
--cert=./etcd/peer.crt \
--key=./etcd/peer.key \
--endpoints=https://127.0.0.1:32379 \
get /registry/namespaces --prefix -w=json
{"level":"warn","ts":"2021-11-04T13:35:09.507+0800","logger":"etcd-client","caller":"v3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0x140004be380/#initially=[https://127.0.0.1:32379]","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = context deadline exceeded"}
Error: context deadline exceeded

端口冲突了
 % lsof |grep 32379
com.docke 71118 xiazemin   69u     IPv6  0xb7573feed65f41d         0t0                 TCP localhost:32379->localhost:62688 (CLOSE_WAIT)
com.docke 71118 xiazemin   89u     IPv6  0xb7573fef567a0bd         0t0                 TCP *:32379 (LISTEN)

% lsof |grep 32389



% ETCDCTL_API=3 etcdctl --cacert=./etcd/ca.crt \
--cert=./etcd/peer.crt \
--key=./etcd/peer.key \
--endpoints=https://127.0.0.1:32389 \
get /registry/namespaces --prefix -w=json

{"header":{"cluster_id":4202879228857769416,"member_id":16554063148076462710,"revision":2056701,"raft_term":18},"kvs":[{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMvZGVmYXVsdA==","create_revision":208,"mod_revision":208,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEoYCCusBCgdkZWZhdWx0EgAaACIAKiQzYjg0ZWMzMy02MzQ2LTRiMjgtOTI2MS0yNzU3ZWYzNWM4NWIyADgAQggIo6eMiQYQAFomChtrdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUSB2RlZmF1bHR6AIoBewoOa3ViZS1hcGlzZXJ2ZXISBlVwZGF0ZRoCdjEiCAijp4yJBhAAMghGaWVsZHNWMTpJCkd7ImY6bWV0YWRhdGEiOnsiZjpsYWJlbHMiOnsiLiI6e30sImY6a3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lIjp7fX19fRIMCgprdWJlcm5ldGVzGggKBkFjdGl2ZRoAIgA="},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1ub2RlLWxlYXNl","create_revision":52,"mod_revision":52,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEpYCCvsBCg9rdWJlLW5vZGUtbGVhc2USABoAIgAqJDQ5OTMxMzYxLTkzMzYtNGIwZC04MDQxLWQ3YzRiZDc2ZmIyNjIAOABCCAihp4yJBhAAWi4KG2t1YmVybmV0ZXMuaW8vbWV0YWRhdGEubmFtZRIPa3ViZS1ub2RlLWxlYXNlegCKAXsKDmt1YmUtYXBpc2VydmVyEgZVcGRhdGUaAnYxIggIoaeMiQYQADIIRmllbGRzVjE6SQpHeyJmOm1ldGFkYXRhIjp7ImY6bGFiZWxzIjp7Ii4iOnt9LCJmOmt1YmVybmV0ZXMuaW8vbWV0YWRhdGEubmFtZSI6e319fX0SDAoKa3ViZXJuZXRlcxoICgZBY3RpdmUaACIA"},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1wdWJsaWM=","create_revision":45,"mod_revision":45,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEo4CCvMBCgtrdWJlLXB1YmxpYxIAGgAiACokNTIxM2Q4NzAtNWRmNC00MzhlLWI1ZDgtMzc4OTA0ZmI5ZGQ5MgA4AEIICKGnjIkGEABaKgoba3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lEgtrdWJlLXB1YmxpY3oAigF7Cg5rdWJlLWFwaXNlcnZlchIGVXBkYXRlGgJ2MSIICKGnjIkGEAAyCEZpZWxkc1YxOkkKR3siZjptZXRhZGF0YSI6eyJmOmxhYmVscyI6eyIuIjp7fSwiZjprdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUiOnt9fX19EgwKCmt1YmVybmV0ZXMaCAoGQWN0aXZlGgAiAA=="},{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMva3ViZS1zeXN0ZW0=","create_revision":13,"mod_revision":13,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEo4CCvMBCgtrdWJlLXN5c3RlbRIAGgAiACokZWJhZDc5YzMtNmI3OC00OGQwLWI1ZGItZmI4YzNiYjk1MWVhMgA4AEIICKGnjIkGEABaKgoba3ViZXJuZXRlcy5pby9tZXRhZGF0YS5uYW1lEgtrdWJlLXN5c3RlbXoAigF7Cg5rdWJlLWFwaXNlcnZlchIGVXBkYXRlGgJ2MSIICKGnjIkGEAAyCEZpZWxkc1YxOkkKR3siZjptZXRhZGF0YSI6eyJmOmxhYmVscyI6eyIuIjp7fSwiZjprdWJlcm5ldGVzLmlvL21ldGFkYXRhLm5hbWUiOnt9fX19EgwKCmt1YmVybmV0ZXMaCAoGQWN0aXZlGgAiAA=="}],"count":4}

