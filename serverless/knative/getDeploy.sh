kubectl -n knative-serving get deploy |awk '{print $1}' |grep -v NAME |xargs -I{} sh -c "kubectl -n knative-serving get deploy {} -o yaml > knative-serving/{}.yaml"

kubectl -n knative-eventing get deploy |awk '{print $1}' |grep -v NAME |xargs -I{} sh -c "kubectl -n knative-eventing get deploy {} -o yaml > knative-eventing/{}.yaml"