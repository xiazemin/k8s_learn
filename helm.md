http://www.coderdocument.com/docs/helm/v2/index.html

https://github.com/helm/charts/tree/master/stable/nginx-ingress


https://github.com/kubernetes/ingress-nginx

https://www.imooc.com/article/291355

https://zhuanlan.zhihu.com/p/165651732


 % helm repo add apphub https://apphub.aliyuncs.com

% helm search repo  ingress-nginx
No results found
% helm search hub ingress-nginx
URL                                               	CHART VERSION	APP VERSION	DESCRIPTION
https://hub.helm.sh/charts/ingress-nginx/ingres...	3.24.0       	0.44.0     	Ingress controller for Kubernetes using NGINX a...
https://hub.helm.sh/charts/wener/ingress-nginx    	3.24.0       	0.44.0     	Ingress controller for Kubernetes using NGINX a...

helm repo add ingresshub  https://hub.helm.sh/charts/wener/ingress-nginx 
Error: looks like "https://hub.helm.sh/charts/wener/ingress-nginx" is not a valid chart repository or cannot be reached: error converting YAML to JSON: yaml: line 2: mapping values are not allowed in this context

https://artifacthub.io/packages/helm/wener/ingress-nginx

https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx


 % helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
"ingress-nginx" has been added to your repositories
% helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "ingress-nginx" chart repository
...Successfully got an update from the "apphub" chart repository
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈Happy Helming!⎈

 % helm search repo ingress-nginx
NAME                       	CHART VERSION	APP VERSION	DESCRIPTION
ingress-nginx/ingress-nginx	3.24.0       	0.44.0     	Ingress controller for Kubernetes using NGINX a...

 % helm install ingress-nginx-0440  ingress-nginx/ingress-nginx
NAME: ingress-nginx-0440
LAST DEPLOYED: Tue Mar 23 11:04:22 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The ingress-nginx controller has been installed.
It may take a few minutes for the LoadBalancer IP to be available.
You can watch the status by running 'kubectl --namespace default get services -o wide -w ingress-nginx-0440-controller'

An example Ingress that makes use of the controller:

  apiVersion: networking.k8s.io/v1beta1
  kind: Ingress
  metadata:
    annotations:
      kubernetes.io/ingress.class: nginx
    name: example
    namespace: foo
  spec:
    rules:
      - host: www.example.com
        http:
          paths:
            - backend:
                serviceName: exampleService
                servicePort: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
        - hosts:
            - www.example.com
          secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls

% helm list
NAME                	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART               	APP VERSION
ingress-nginx-0.44.0	default  	1       	2021-03-23 11:03:42.864706 +0800 CST	failed  	ingress-nginx-3.24.0	0.44.0
ingress-nginx-0440  	default  	1       	2021-03-23 11:04:22.154773 +0800 CST	deployed	ingress-nginx-3.24.0	0.44.0

 % helm uninstall ingress-nginx-0.44.0
release "ingress-nginx-0.44.0" uninstalled

% kubectl --namespace default get services -o wide -w ingress-nginx-0440-controller
NAME                            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE   SELECTOR
ingress-nginx-0440-controller   LoadBalancer   10.109.121.93   localhost     80:31061/TCP,443:31444/TCP   9m    app.kubernetes.io/component=controller,app.kubernetes.io/instance=ingress-nginx-0440,app.kubernetes.io/name=ingress-nginx

 % kubectl logs ingress-nginx-0440-controller-f58b5bcc8-9qf8b
Error from server (BadRequest): container "controller" in pod "ingress-nginx-0440-controller-f58b5bcc8-9qf8b" is waiting to start: trying and failing to pull image


% kubectl describe pod  ingress-nginx-0440-controller-f58b5bcc8-9qf8b
Name:         ingress-nginx-0440-controller-f58b5bcc8-9qf8b
Namespace:    default
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Tue, 23 Mar 2021 11:04:27 +0800
Labels:       app.kubernetes.io/component=controller
              app.kubernetes.io/instance=ingress-nginx-0440
              app.kubernetes.io/name=ingress-nginx
              pod-template-hash=f58b5bcc8
Annotations:  <none>
Status:       Pending
IP:           10.1.0.60
IPs:
  IP:           10.1.0.60
Controlled By:  ReplicaSet/ingress-nginx-0440-controller-f58b5bcc8
Containers:
  controller:
    Container ID:
    Image:         k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a
    Image ID:
    Ports:         80/TCP, 443/TCP, 8443/TCP
    Host Ports:    0/TCP, 0/TCP, 0/TCP
    Args:
      /nginx-ingress-controller
      --publish-service=$(POD_NAMESPACE)/ingress-nginx-0440-controller
      --election-id=ingress-controller-leader
      --ingress-class=nginx
      --configmap=$(POD_NAMESPACE)/ingress-nginx-0440-controller
      --validating-webhook=:8443
      --validating-webhook-certificate=/usr/local/certificates/cert
      --validating-webhook-key=/usr/local/certificates/key
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:      100m
      memory:   90Mi
    Liveness:   http-get http://:10254/healthz delay=10s timeout=1s period=10s #success=1 #failure=5
    Readiness:  http-get http://:10254/healthz delay=10s timeout=1s period=10s #success=1 #failure=3
    Environment:
      POD_NAME:       ingress-nginx-0440-controller-f58b5bcc8-9qf8b (v1:metadata.name)
      POD_NAMESPACE:  default (v1:metadata.namespace)
      LD_PRELOAD:     /usr/local/lib/libmimalloc.so
    Mounts:
      /usr/local/certificates/ from webhook-cert (ro)
      /var/run/secrets/kubernetes.io/serviceaccount from ingress-nginx-0440-token-krwz6 (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  webhook-cert:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  ingress-nginx-0440-admission
    Optional:    false
  ingress-nginx-0440-token-krwz6:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  ingress-nginx-0440-token-krwz6
    Optional:    false
QoS Class:       Burstable
Node-Selectors:  kubernetes.io/os=linux
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age                  From               Message
  ----     ------     ----                 ----               -------
  Normal   Scheduled  14m                  default-scheduler  Successfully assigned default/ingress-nginx-0440-controller-f58b5bcc8-9qf8b to docker-desktop
  Normal   Pulling    12m (x4 over 14m)    kubelet            Pulling image "k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a"
  Warning  Failed     11m (x4 over 13m)    kubelet            Failed to pull image "k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a": rpc error: code = Unknown desc = Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
  Warning  Failed     11m (x4 over 13m)    kubelet            Error: ErrImagePull
  Normal   BackOff    9m3s (x15 over 13m)  kubelet            Back-off pulling image "k8s.gcr.io/ingress-nginx/controller:v0.44.0@sha256:3dd0fac48073beaca2d67a78c746c7593f9c575168a17139a9955a82c63c4b9a"
  Warning  Failed     4m5s (x36 over 13m)  kubelet            Error: ImagePullBackOff

  % helm delete ingress-nginx-0440
release "ingress-nginx-0440" uninstalled


http://www.coderdocument.com/docs/helm/v2/index.html

