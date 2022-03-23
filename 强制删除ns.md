kubectl get ns istio-operator  -o json | jq 'del(.spec.finalizers)' |curl -v -H "Content-Type: application/json" -X PUT --data-binary @- http://127.0.0.1:8080/api/v1/namespaces/istio-operator/finalize


https://www.cnblogs.com/elnino/p/11104084.html


kubectl delete ns knative-serving --force --grace-period=0


kubectl proxy

kubectl get ns knative-serving -o json | jq 'del(.spec.finalizers)' |curl -v -H "Content-Type: application/json" -X PUT --data-binary @- http://127.0.0.1:8080/api/v1/namespaces/knative-serving/finalize

https://www.jianshu.com/p/00351f9eb37b