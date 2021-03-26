https://jimmysong.io/kubernetes-handbook/guide/connecting-to-applications-port-forward.html

% kubectl get service -n my-ns
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
apple-service   ClusterIP   10.105.14.198   <none>        5678/TCP   79s

% kubectl get pod -n my-ns
NAME        READY   STATUS    RESTARTS   AGE
apple-app   1/1     Running   0          2m11s

% kubectl port-forward apple-app -n my-ns 5678:5678
Forwarding from 127.0.0.1:5678 -> 5678
Forwarding from [::1]:5678 -> 5678
Handling connection for 5678

 % curl http://127.0.0.1:5678/apple
apple


% kubectl logs ingress-nginx-controller-84599cfff-l77nt -n ingress-nginx

 13 event.go:282] Event(v1.ObjectReference{Kind:"Pod", Namespace:"ingress-nginx", Name:"ingress-nginx-controller-84599cfff-l77nt", UID:"63d4056b-dedb-4c62-b71d-fd2b7f46c59d", APIVersion:"v1", ResourceVersion:"464784", FieldPath:""}): type: 'Normal' reason: 'RELOAD' NGINX reload triggered due to a change in configuration
2021/03/26 16:03:22 [alert] 50#50: failed to load the 'resty.core' module (https://github.com/openresty/lua-resty-core); ensure you are using an OpenResty release from https://openresty.org/en/download.html (reason: error loading module 'resty.core.shdict' from file '/usr/local/lib/lua/resty/core/shdict.lua':
	not enough memory) in /etc/nginx/nginx.conf:1321
I0326 16:03:38.676616      13 status.go:281] "updating Ingress status" namespace="my-ns" ingress="rancher-k8s-ingress" currentValue=[] newValue=[{IP: Hostname:localhost Ports:[]}]
W0326 16:03:38.698230      13 controller.go:994] Service "default/apple-service" does not have any active Endpoint.
W0326 16:03:38.698351      13 controller.go:994] Service "default/banana-service" does not have any active Endpoint.


nginx: lua atpanic: Lua VM crashed, reason: not enough memory
nginx: lua atpanic: Lua VM crashed, reason: not enough memory
2021/03/26 16:04:03 [emerg] 1167#1167: io_setup() failed (38: Function not implemented)


 % kubectl get Ingress
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME              CLASS    HOSTS   ADDRESS     PORTS   AGE
example-ingress   <none>   *       localhost   80      2d3h


 % kubectl delete -f sample/ingress.v1.yaml
ingress.networking.k8s.io "example-ingress" deleted



192.168.64.6 - - [26/Mar/2021:16:20:31 +0000] "GET /apple HTTP/1.1" 200 5 "-" "curl/7.64.1" 78 0.037 [my-ns-apple-service-5678] [] 10.1.0.255:5678 5 0.037 200 7f898983f9770ef39a25001d18cbf499


 % curl  http://127.0.0.1/apple
<html>
<head><title>404 Not Found</title></head>

% curl  http://localhost/apple
apple


              number: 5678
        path: /apples_bana
        pathType: Prefix 
status:
  loadBalancer:
    ingress:
    - hostname: localhost

端口号和hostname 有用，path 没有 用，
curl  http://localhost/apple
apple



https://jimmysong.io/kubernetes-handbook/cloud-native/setup-kubernetes-with-rancher-and-aliyun.html

https://www.qikqiak.com/k8s-book/docs/40.ingress.html