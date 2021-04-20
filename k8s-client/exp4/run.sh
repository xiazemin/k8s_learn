% kubectl get deployment
No resources found in default namespace.

% kubectl get pod  
No resources found in default namespace.

% kubectl get svc
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
banana-service   ClusterIP   10.111.66.167   <none>        5678/TCP   26d
kubernetes       ClusterIP   10.96.0.1       <none>        443/TCP    29d

% kubectl get endpoints
NAME             ENDPOINTS           AGE
banana-service   <none>              26d
kubernetes       192.168.65.4:6443   29d

 % go run k8s-client/exp4/pod.go 
You must specify the deployment name.