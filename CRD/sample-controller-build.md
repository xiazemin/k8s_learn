% go build -o sample-controller .


../../../../go/pkg/mod/sigs.k8s.io/json@v0.0.0-20211020170558-c049b76a60c6/internal/golang/encoding/json/encode.go:1255:18: sf.IsExported undefined (type reflect.StructField has no field or method IsExported)


https://githubmemory.com/repo/clarketm/json
https://githubmemory.com/repo/tinygo-org/tinygo/issues/2054


%  go build -o sample-controller .
% git checkout release-1.21 

问题解决


 % ./sample-controller -kubeconfig=$HOME/.kube/config
I1103 10:50:25.426210   36841 controller.go:115] Setting up event handlers
I1103 10:50:25.426376   36841 controller.go:156] Starting Foo controller
I1103 10:50:25.426381   36841 controller.go:159] Waiting for informer caches to sync

在mac 宿主机运行失败，mac 有一个虚拟机


https://blog.csdn.net/zhonglinzhang/article/details/86553744


http://www.asznl.com/post/43

https://www.qikqiak.com/k8strain/operator/crd/

https://cloud.tencent.com/developer/article/1717417

 % ./sample-controller -kubeconfig=$HOME/.kube/config
I1104 13:44:56.114207   58916 controller.go:115] Setting up event handlers
I1104 13:44:56.114648   58916 controller.go:156] Starting Foo controller
I1104 13:44:56.114654   58916 controller.go:159] Waiting for informer caches to sync
E1104 13:44:56.132686   58916 reflector.go:138] pkg/generated/informers/externalversions/factory.go:117: Failed to watch *v1alpha1.Foo: failed to list *v1alpha1.Foo: the server could not find the requested resource (get foos.samplecontroller.k8s.io)
E1104 13:44:57.427687   58916 reflector.go:138] pkg/generated/informers/externalversions/factory.go:117: Failed to watch *v1alpha1.Foo: failed to list *v1alpha1.Foo: the server



% kubectl create -f artifacts/examples/crd.yaml
Warning: apiextensions.k8s.io/v1beta1 CustomResourceDefinition is deprecated in v1.16+, unavailable in v1.22+; use apiextensions.k8s.io/v1 CustomResourceDefinition
customresourcedefinition.apiextensions.k8s.io/foos.samplecontroller.k8s.io created

% kubectl create -f artifacts/examples/example-foo.yaml
foo.samplecontroller.k8s.io/example-foo created


 % kubectl get deployments
NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
example-foo                0/1     1            0           14s


^@^@^@I1104 13:48:10.365898   58916 controller.go:228] Successfully synced 'default/example-foo'
I1104 13:48:10.365938   58916 event.go:291] "Event occurred" object="default/example-foo" kind="Foo" apiVersion="samplecontroller.k8s.io/v1alpha1" type="Normal" reason="Synced" message="Foo synced successfully"


 % kubectl get CustomResourceDefinition foos.samplecontroller.k8s.io
NAME                           CREATED AT
foos.samplecontroller.k8s.io   2021-11-04T05:46:53Z


% kubectl get Foo example-foo
NAME          AGE
example-foo   3m11s


% ETCDCTL_API=3 etcdctl --cacert=./etcd/ca.crt \
--cert=./etcd/peer.crt \
--key=./etcd/peer.key \
--endpoints=https://127.0.0.1:32389 \
get ""  --prefix --keys-only |grep foo
/registry/apiextensions.k8s.io/customresourcedefinitions/foos.samplecontroller.k8s.io
/registry/deployments/default/example-foo
/registry/events/default/example-foo-9bbb75dc8-v7dzv.16b441bc25c1f3ba
/registry/events/default/example-foo-9bbb75dc8-v7dzv.16b441bc4b1a9b8f
/registry/events/default/example-foo-9bbb75dc8-v7dzv.16b441cbaffd4dc2
/registry/events/default/example-foo-9bbb75dc8-v7dzv.16b441cbbc571207
/registry/events/default/example-foo-9bbb75dc8-v7dzv.16b441cbc264acc2
/registry/events/default/example-foo-9bbb75dc8.16b441bc25872175
/registry/events/default/example-foo.16b441bc243a6ed8
/registry/events/default/example-foo.16b441bc24a13b4a
/registry/pods/default/example-foo-9bbb75dc8-v7dzv
/registry/replicasets/default/example-foo-9bbb75dc8
/registry/samplecontroller.k8s.io/foos/default/example-foo


% ETCDCTL_API=3 etcdctl --cacert=./etcd/ca.crt \
--cert=./etcd/peer.crt \
--key=./etcd/peer.key \
--endpoints=https://127.0.0.1:32389 \
get /registry/pods/default/example-foo-9bbb75dc8-v7dzv -w json
{"header":{"cluster_id":4202879228857769416,"member_id":16554063148076462710,"revision":2058014,"raft_term":18},"kvs":[{"key":"L3JlZ2lzdHJ5L3BvZHMvZGVmYXVsdC9leGFtcGxlLWZvby05YmJiNzVkYzgtdjdkenY=","create_revision":2057101,"mod_revision":2057210,"version":4,"value":"azhzAAoJCgJ2MRIDUG9kEt4TCowMChtleGFtcGxlLWZvby05YmJiNzVkYzgtdjdkenYSFmV4YW1wbGUtZm9vLTliYmI3NWRjOC0aB2RlZmF1bHQiACokOTAxYzA5ODYtMTM0ZS00Y2I5LTliMjItYWI5N2QyMzlmZDkyMgA4AEIICOjnjYwGEABaDAoDYXBwEgVuZ2lueFoZCgpjb250cm9sbGVyEgtleGFtcGxlLWZvb1oeChFwb2QtdGVtcGxhdGUtaGFzaBIJOWJiYjc1ZGM4alYKClJlcGxpY2FTZXQaFWV4YW1wbGUtZm9vLTliYmI3NWRjOCIkYjhjZWQ5Y2YtMjViNS00NjhkLTg1MzYtODBlODRmZmIwMjY4KgdhcHBzL3YxMAE4AXoAigG5BQoXa3ViZS1jb250cm9sbGVyLW1hbmFnZXISBlVwZGF0ZRoCdjEiCAjo542MBhAAMghGaWVsZHNWMTr9BAr6BHsiZjptZXRhZGF0YSI6eyJmOmdlbmVyYXRlTmFtZSI6e30sImY6bGFiZWxzIjp7Ii4iOnt9LCJmOmFwcCI6e30sImY6Y29udHJvbGxlciI6e30sImY6cG9kLXRlbXBsYXRlLWhhc2giOnt9fSwiZjpvd25lclJlZmVyZW5jZXMiOnsiLiI6e30sIms6e1widWlkXCI6XCJiOGNlZDljZi0yNWI1LTQ2OGQtODUzNi04MGU4NGZmYjAyNjhcIn0iOnsiLiI6e30sImY6YXBpVmVyc2lvbiI6e30sImY6YmxvY2tPd25lckRlbGV0aW9uIjp7fSwiZjpjb250cm9sbGVyIjp7fSwiZjpraW5kIjp7fSwiZjpuYW1lIjp7fSwiZjp1aWQiOnt9fX19LCJmOnNwZWMiOnsiZjpjb250YWluZXJzIjp7Ims6e1wibmFtZVwiOlwibmdpbnhcIn0iOnsiLiI6e30sImY6aW1hZ2UiOnt9LCJmOmltYWdlUHVsbFBvbGljeSI6e30sImY6bmFtZSI6e30sImY6cmVzb3VyY2VzIjp7fSwiZjp0ZXJtaW5hdGlvbk1lc3NhZ2VQYXRoIjp7fSwiZjp0ZXJtaW5hdGlvbk1lc3NhZ2VQb2xpY3kiOnt9fX0sImY6ZG5zUG9saWN5Ijp7fSwiZjplbmFibGVTZXJ2aWNlTGlua3MiOnt9LCJmOnJlc3RhcnRQb2xpY3kiOnt9LCJmOnNjaGVkdWxlck5hbWUiOnt9LCJmOnNlY3VyaXR5Q29udGV4dCI6e30sImY6dGVybWluYXRpb25HcmFjZVBlcmlvZFNlY29uZHMiOnt9fX2KAbQECgdrdWJlbGV0EgZVcGRhdGUaAnYxIggIrOiNjAYQADIIRmllbGRzVjE6iAQKhQR7ImY6c3RhdHVzIjp7ImY6Y29uZGl0aW9ucyI6eyJrOntcInR5cGVcIjpcIkNvbnRhaW5lcnNSZWFkeVwifSI6eyIuIjp7fSwiZjpsYXN0UHJvYmVUaW1lIjp7fSwiZjpsYXN0VHJhbnNpdGlvblRpbWUiOnt9LCJmOnN0YXR1cyI6e30sImY6dHlwZSI6e319LCJrOntcInR5cGVcIjpcIkluaXRpYWxpemVkXCJ9Ijp7Ii4iOnt9LCJmOmxhc3RQcm9iZVRpbWUiOnt9LCJmOmxhc3RUcmFuc2l0aW9uVGltZSI6e30sImY6c3RhdHVzIjp7fSwiZjp0eXBlIjp7fX0sIms6e1widHlwZVwiOlwiUmVhZHlcIn0iOnsiLiI6e30sImY6bGFzdFByb2JlVGltZSI6e30sImY6bGFzdFRyYW5zaXRpb25UaW1lIjp7fSwiZjpzdGF0dXMiOnt9LCJmOnR5cGUiOnt9fX0sImY6Y29udGFpbmVyU3RhdHVzZXMiOnt9LCJmOmhvc3RJUCI6e30sImY6cGhhc2UiOnt9LCJmOnBvZElQIjp7fSwiZjpwb2RJUHMiOnsiLiI6e30sIms6e1wiaXBcIjpcIjEwLjEuMC4yMDNcIn0iOnsiLiI6e30sImY6aXAiOnt9fX0sImY6c3RhcnRUaW1lIjp7fX19Eo8ECoQBChVrdWJlLWFwaS1hY2Nlc3MtZGg0eHMSa9IBaAoOIgwKABCXHBoFdG9rZW4KKBomChIKEGt1YmUtcm9vdC1jYS5jcnQSEAoGY2EuY3J0EgZjYS5jcnQKKRInCiUKCW5hbWVzcGFjZRIYCgJ2MRISbWV0YWRhdGEubmFtZXNwYWNlEKQDEpUBCgVuZ2lueBIMbmdpbng6bGF0ZXN0KgBCAEpMChVrdWJlLWFwaS1hY2Nlc3MtZGg0eHMQARotL3Zhci9ydW4vc2VjcmV0cy9rdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50IgAyAGoUL2Rldi90ZXJtaW5hdGlvbi1sb2dyBkFsd2F5c4ABAIgBAJABAKIBBEZpbGUaBkFsd2F5cyAeMgxDbHVzdGVyRmlyc3RCB2RlZmF1bHRKB2RlZmF1bHRSDmRvY2tlci1kZXNrdG9wWABgAGgAcgCCAQCKAQCaARFkZWZhdWx0LXNjaGVkdWxlcrIBNgocbm9kZS5rdWJlcm5ldGVzLmlvL25vdC1yZWFkeRIGRXhpc3RzGgAiCU5vRXhlY3V0ZSisArIBOAoebm9kZS5rdWJlcm5ldGVzLmlvL3VucmVhY2hhYmxlEgZFeGlzdHMaACIJTm9FeGVjdXRlKKwCwgEAyAEA8AEB+gEUUHJlZW1wdExvd2VyUHJpb3JpdHkaugMKB1J1bm5pbmcSIwoLSW5pdGlhbGl6ZWQSBFRydWUaACIICOjnjYwGEAAqADIAEh0KBVJlYWR5EgRUcnVlGgAiCAis6I2MBhAAKgAyABInCg9Db250YWluZXJzUmVhZHkSBFRydWUaACIICKzojYwGEAAqADIAEiQKDFBvZFNjaGVkdWxlZBIEVHJ1ZRoAIggI6OeNjAYQACoAMgAaACIAKgwxOTIuMTY4LjY1LjQyCjEwLjEuMC4yMDM6CAjo542MBhAAQtcBCgVuZ2lueBIMEgoKCAis6I2MBhAAGgAgASgAMgxuZ2lueDpsYXRlc3Q6X2RvY2tlci1wdWxsYWJsZTovL25naW54QHNoYTI1Njo2NDRhNzA1MTZhMjYwMDRjOTdkMGQ4NWM3ZmUxZDBjM2E2N2VhOGFiN2RkZjRhZmYxOTNkOWYzMDE2NzBjZjM2Qklkb2NrZXI6Ly85NWM3MjFiMzZhOTdiMzcwZjI4ZWZhN2U2MzllMjdkMjdjZDIyNmE5Zjc3NzIwYTQ1ODM0YWE3ZDI1N2RlNGZjSAFKCkJlc3RFZmZvcnRaAGIMCgoxMC4xLjAuMjAzGgAiAA=="}],"count":1}


https://www.huweihuang.com/kubernetes-notes/etcd/etcdctl-v3.html