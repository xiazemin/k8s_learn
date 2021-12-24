https://www.cnblogs.com/zjq-blogs/p/14067613.html

 % kubectl apply -f ingress.yaml
Error from server (InternalError): error when creating "ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": an error on the server ("") has prevented the request from succeeding


% kubectl get validatingwebhookconfigurations ingress-nginx-admission  -o yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"admissionregistration.k8s.io/v1","kind":"ValidatingWebhookConfiguration","metadata":{"annotations":{},"labels":{"app.kubernetes.io/component":"admission-webhook","app.kubernetes.io/instance":"ingress-nginx","app.kubernetes.io/managed-by":"Helm","app.kubernetes.io/name":"ingress-nginx","app.kubernetes.io/version":"0.44.0","helm.sh/chart":"ingress-nginx-3.23.0"},"name":"ingress-nginx-admission"},"webhooks":[{"admissionReviewVersions":["v1"],"clientConfig":{"service":{"name":"ingress-nginx-controller-admission","namespace":"default","path":"/networking/v1/ingresses"}},"failurePolicy":"Fail","matchPolicy":"Equivalent","name":"validate.nginx.ingress.kubernetes.io","rules":[{"apiGroups":["networking.k8s.io"],"apiVersions":["v1"],"operations":["CREATE","UPDATE"],"resources":["ingresses"]}],"sideEffects":"None"}]}
  creationTimestamp: "2021-08-28T12:52:06Z"
  generation: 2
  labels:
    app.kubernetes.io/component: admission-webhook
    app.kubernetes.io/instance: ingress-nginx
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/version: 0.44.0
    helm.sh/chart: ingress-nginx-3.23.0
  name: ingress-nginx-admission
  resourceVersion: "233225"
  uid: 5b55615b-0c11-4157-a26e-2ba6c2cf601d
webhooks:
- admissionReviewVersions:
  - v1
  clientConfig:
    caBundle: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJkVENDQVJ1Z0F3SUJBZ0lRR1JadXplNytzbnAvRGs0akZ5cTh3VEFLQmdncWhrak9QUVFEQWpBUE1RMHcKQ3dZRFZRUUtFd1J1YVd3eE1DQVhEVEl4TURneU5UQTJNVEF5TmxvWUR6SXhNakV3T0RBeE1EWXhNREkyV2pBUApNUTB3Q3dZRFZRUUtFd1J1YVd3eE1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRTFoc05wK243CnRtZ2lBNlQzY1pFNVBPSE00WUJjcHlJS1V4UkhHeWVheFhNVGEzUDhCWnBDQWh4YkZRYVZWZjVrSkVBY2FwK04KQ0ZOKy8yR2dVZmROenFOWE1GVXdEZ1lEVlIwUEFRSC9CQVFEQWdJRU1CTUdBMVVkSlFRTU1Bb0dDQ3NHQVFVRgpCd01CTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkhyOG9CaGtpSWIzblArSU5ieDZNUmFJClRMWkhNQW9HQ0NxR1NNNDlCQU1DQTBnQU1FVUNJSFFiY215ZGtUT1ltODFZKytZYVl6OHVtbUgzTEt3bUFaVFIKTmhnUTZaUXlBaUVBMEJQOFR2cjE0YXBaR1dUTDFsanY0MTEwdVArWWxvcWpuRis1NFV0TVJRTT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    service:
      name: ingress-nginx-controller-admission
      namespace: default
      path: /networking/v1/ingresses
      port: 443
  failurePolicy: Fail
  matchPolicy: Equivalent
  name: validate.nginx.ingress.kubernetes.io
  namespaceSelector: {}
  objectSelector: {}
  rules:
  - apiGroups:
    - networking.k8s.io
    apiVersions:
    - v1
    operations:
    - CREATE
    - UPDATE
    resources:
    - ingresses
    scope: '*'
  sideEffects: None
  timeoutSeconds: 10


https://blog.csdn.net/qq_35396734/article/details/120247811

https://github.com/kubernetes/kubernetes/issues/71647



kubectl proxy 让外部网络访问K8S service的ClusterIP

% kubectl proxy --port=8009
Starting to serve on 127.0.0.1:8009

 % curl 127.0.0.1:8009
{
  "paths": [
    "/.well-known/openid-configuration",
    "/api",
    "/api/v1",
    "/apis",
    "/apis/",
    "/apis/admissionregistration.k8s.io",
    "/apis/admissionregistration.k8s.io/v1",
    "/apis/admissionregistration.k8s.io/v1beta1",
    "/apis/apiextensions.k8s.io",
    "/apis/apiextensions.k8s.io/v1",
    "/apis/apiextensions.k8s.io/v1beta1",
    "/apis/apiregistration.k8s.io",
    "/apis/apiregistration.k8s.io/v1",
    "/apis/apiregistration.k8s.io/v1beta1",
    "/apis/apps",
    "/apis/apps/v1",
    "/apis/authentication.k8s.io",
    "/apis/authentication.k8s.io/v1",
    "/apis/authentication.k8s.io/v1beta1",
    "/apis/authorization.k8s.io",
    "/apis/authorization.k8s.io/v1",
    "/apis/authorization.k8s.io/v1beta1",
    "/apis/autoscaling",
    "/apis/autoscaling/v1",
    "/apis/autoscaling/v2beta1",
    "/apis/autoscaling/v2beta2",
    "/apis/batch",
    "/apis/batch/v1",
    "/apis/batch/v1beta1",
    "/apis/catalog.cattle.io",
    "/apis/catalog.cattle.io/v1",
    "/apis/certificates.k8s.io",
    "/apis/certificates.k8s.io/v1",
    "/apis/certificates.k8s.io/v1beta1",
    "/apis/coordination.k8s.io",
    "/apis/coordination.k8s.io/v1",
    "/apis/coordination.k8s.io/v1beta1",
    "/apis/custom.metrics.k8s.io",
    "/apis/custom.metrics.k8s.io/v1beta1",
    "/apis/discovery.k8s.io",
    "/apis/discovery.k8s.io/v1",
    "/apis/discovery.k8s.io/v1beta1",
    "/apis/events.k8s.io",
    "/apis/events.k8s.io/v1",
    "/apis/events.k8s.io/v1beta1",
    "/apis/extensions",
    "/apis/extensions/v1beta1",
    "/apis/flowcontrol.apiserver.k8s.io",
    "/apis/flowcontrol.apiserver.k8s.io/v1beta1",
    "/apis/management.cattle.io",
    "/apis/management.cattle.io/v3",
    "/apis/networking.k8s.io",
    "/apis/networking.k8s.io/v1",
    "/apis/networking.k8s.io/v1beta1",
    "/apis/node.k8s.io",
    "/apis/node.k8s.io/v1",
    "/apis/node.k8s.io/v1beta1",
    "/apis/policy",
    "/apis/policy/v1",
    "/apis/policy/v1beta1",
    "/apis/rbac.authorization.k8s.io",
    "/apis/rbac.authorization.k8s.io/v1",
    "/apis/rbac.authorization.k8s.io/v1beta1",
    "/apis/samplecontroller.k8s.io",
    "/apis/samplecontroller.k8s.io/v1alpha1",
    "/apis/scheduling.k8s.io",
    "/apis/scheduling.k8s.io/v1",
    "/apis/scheduling.k8s.io/v1beta1",
    "/apis/storage.k8s.io",
    "/apis/storage.k8s.io/v1",
    "/apis/storage.k8s.io/v1beta1",
    "/apis/ui.cattle.io",
    "/apis/ui.cattle.io/v1",
    "/healthz",
    "/healthz/autoregister-completion",
    "/healthz/etcd",
    "/healthz/log",
    "/healthz/ping",
    "/healthz/poststarthook/aggregator-reload-proxy-client-cert",
    "/healthz/poststarthook/apiservice-openapi-controller",
    "/healthz/poststarthook/apiservice-registration-controller",
    "/healthz/poststarthook/apiservice-status-available-controller",
    "/healthz/poststarthook/bootstrap-controller",
    "/healthz/poststarthook/crd-informer-synced",
    "/healthz/poststarthook/generic-apiserver-start-informers",
    "/healthz/poststarthook/kube-apiserver-autoregistration",
    "/healthz/poststarthook/priority-and-fairness-config-consumer",
    "/healthz/poststarthook/priority-and-fairness-config-producer",
    "/healthz/poststarthook/priority-and-fairness-filter",
    "/healthz/poststarthook/rbac/bootstrap-roles",
    "/healthz/poststarthook/scheduling/bootstrap-system-priority-classes",
    "/healthz/poststarthook/start-apiextensions-controllers",
    "/healthz/poststarthook/start-apiextensions-informers",
    "/healthz/poststarthook/start-cluster-authentication-info-controller",
    "/healthz/poststarthook/start-kube-aggregator-informers",
    "/healthz/poststarthook/start-kube-apiserver-admission-initializer",
    "/livez",
    "/livez/autoregister-completion",
    "/livez/etcd",
    "/livez/log",
    "/livez/ping",
    "/livez/poststarthook/aggregator-reload-proxy-client-cert",
    "/livez/poststarthook/apiservice-openapi-controller",
    "/livez/poststarthook/apiservice-registration-controller",
    "/livez/poststarthook/apiservice-status-available-controller",
    "/livez/poststarthook/bootstrap-controller",
    "/livez/poststarthook/crd-informer-synced",
    "/livez/poststarthook/generic-apiserver-start-informers",
    "/livez/poststarthook/kube-apiserver-autoregistration",
    "/livez/poststarthook/priority-and-fairness-config-consumer",
    "/livez/poststarthook/priority-and-fairness-config-producer",
    "/livez/poststarthook/priority-and-fairness-filter",
    "/livez/poststarthook/rbac/bootstrap-roles",
    "/livez/poststarthook/scheduling/bootstrap-system-priority-classes",
    "/livez/poststarthook/start-apiextensions-controllers",
    "/livez/poststarthook/start-apiextensions-informers",
    "/livez/poststarthook/start-cluster-authentication-info-controller",
    "/livez/poststarthook/start-kube-aggregator-informers",
    "/livez/poststarthook/start-kube-apiserver-admission-initializer",
    "/logs",
    "/metrics",
    "/openapi/v2",
    "/openid/v1/jwks",
    "/readyz",
    "/readyz/autoregister-completion",
    "/readyz/etcd",
    "/readyz/informer-sync",
    "/readyz/log",
    "/readyz/ping",
    "/readyz/poststarthook/aggregator-reload-proxy-client-cert",
    "/readyz/poststarthook/apiservice-openapi-controller",
    "/readyz/poststarthook/apiservice-registration-controller",
    "/readyz/poststarthook/apiservice-status-available-controller",
    "/readyz/poststarthook/bootstrap-controller",
    "/readyz/poststarthook/crd-informer-synced",
    "/readyz/poststarthook/generic-apiserver-start-informers",
    "/readyz/poststarthook/kube-apiserver-autoregistration",
    "/readyz/poststarthook/priority-and-fairness-config-consumer",
    "/readyz/poststarthook/priority-and-fairness-config-producer",
    "/readyz/poststarthook/priority-and-fairness-filter",
    "/readyz/poststarthook/rbac/bootstrap-roles",
    "/readyz/poststarthook/scheduling/bootstrap-system-priority-classes",
    "/readyz/poststarthook/start-apiextensions-controllers",
    "/readyz/poststarthook/start-apiextensions-informers",
    "/readyz/poststarthook/start-cluster-authentication-info-controller",
    "/readyz/poststarthook/start-kube-aggregator-informers",
    "/readyz/poststarthook/start-kube-apiserver-admission-initializer",
    "/readyz/shutdown",
    "/version"
  ]
}%

 % docker tag matrix matrix:0.0.1

 % kubectl set image deployment/matrix-deployment  matrix-deployment=matrix:0.0.1 
deployment.apps/matrix-deployment image updated



 % curl 127.0.0.1:8009/metrics |head -n 10
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0# HELP aggregator_openapi_v2_regeneration_count [ALPHA] Counter of OpenAPI v2 spec regeneration count broken down by causing APIService name and reason.
# TYPE aggregator_openapi_v2_regeneration_count counter
aggregator_openapi_v2_regeneration_count{apiservice="*",reason="startup"} 0
aggregator_openapi_v2_regeneration_count{apiservice="k8s_internal_local_delegation_chain_0000000002",reason="update"} 0
aggregator_openapi_v2_regeneration_count{apiservice="v1beta1.custom.metrics.k8s.io",reason="add"} 0
aggregator_openapi_v2_regeneration_count{apiservice="v1beta1.custom.metrics.k8s.io",reason="delete"} 0
# HELP aggregator_openapi_v2_regeneration_duration [ALPHA] Gauge of OpenAPI v2 spec regeneration duration in seconds.
# TYPE aggregator_openapi_v2_regeneration_duration gauge
aggregator_openapi_v2_regeneration_duration{reason="add"} 0.835911584
aggregator_openapi_v2_regeneration_duration{reason="delete"} 0.841896292
100 65536    0 65536    0     0  2461k      0 --:--:-- --:--:-- --:--:-- 2461k
curl: (23) Failed writing body (0 != 16384)


https://blog.csdn.net/zwqjoy/article/details/87865283


 % curl http://127.0.0.1:12345/httpInc
httpInc%



 % curl http://127.0.0.1:8009/api/v1/namespaces/default
{
  "kind": "Namespace",
  "apiVersion": "v1",
 "metadata": {
    "name": "default",
    "uid": "3b84ec33-6346-4b28-9261-2757ef35c85b",
    "resourceVersion": "2648253",
    "creationTimestamp": "2021-08-23T03:18:59Z",
    "labels": {
      "field.cattle.io/projectId": "p-ww28c",
      "kubernetes.io/metadata.name": "default"
    },
    "annotations": {
      "cattle.io/status": "{\"Conditions\":[{\"Type\":\"ResourceQuotaInit\",\"Status\":\"True\",\"Message\":\"\",\"LastUpdateTime\":\"2021-11-19T06:56:10Z\"},{\"Type\":\"InitialRolesPopulated\",\"Status\":\"True\",\"Message\":\"\",\"LastUpdateTime\":\"2021-11-19T06:56:16Z\"}]}",
      "field.cattle.io/projectId": "c-m-r7xmtmss:p-ww28c",
      "lifecycle.cattle.io/create.namespace-auth": "true"
    },
    "finalizers": [
      "controller.cattle.io/namespace-auth"
    ],
    "managedFields": [
      {
        "manager": "kube-apiserver",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-08-23T03:18:59Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:labels":{".":{},"f:kubernetes.io/metadata.name":{}}}}
      },
      {
        "manager": "rancher",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-11-19T06:56:10Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:cattle.io/status":{},"f:field.cattle.io/projectId":{},"f:lifecycle.cattle.io/create.namespace-auth":{}},"f:finalizers":{".":{},"v:\"controller.cattle.io/namespace-auth\"":{}}}}
      },
      {
        "manager": "agent",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-11-19T06:56:20Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:labels":{"f:field.cattle.io/projectId":{}}}}
      }
    ]
  },
  "spec": {
    "finalizers": [
      "kubernetes"
    ]
  },
  "status": {
    "phase": "Active"
  }
}


curl http://[k8s-master]:8009/api/v1/namespaces/[namespace-name]/services/[service-name]/proxy


 % curl http://127.0.0.1:8009/api/v1/namespaces/default/services


% curl http://127.0.0.1:8009/api/v1/namespaces/default/services/matrix-deployment
{
  "kind": "Service",
  "apiVersion": "v1",
  "metadata": {
    "name": "matrix-deployment",
    "namespace": "default",
    "uid": "ee4cc453-980c-4e33-a3c6-e7175d151eae",
    "resourceVersion": "4402175",
    "creationTimestamp": "2021-12-24T03:42:34Z",
    "labels": {
      "app": "matrix-deployment"
    },
    "annotations": {
      "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"matrix-deployment\"},\"name\":\"matrix-deployment\",\"namespace\":\"default\"},\"spec\":{\"ports\":[{\"name\":\"http\",\"port\":12345,\"protocol\":\"TCP\",\"targetPort\":12345}],\"selector\":{\"app\":\"matrix-deployment\"},\"type\":\"ClusterIP\"}}\n"
    },
    "managedFields": [
      {
        "manager": "kubectl-client-side-apply",
        "operation": "Update",
        "apiVersion": "v1",
        "time": "2021-12-24T03:42:34Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {"f:metadata":{"f:annotations":{".":{},"f:kubectl.kubernetes.io/last-applied-configuration":{}},"f:labels":{".":{},"f:app":{}}},"f:spec":{"f:ports":{".":{},"k:{\"port\":12345,\"protocol\":\"TCP\"}":{".":{},"f:name":{},"f:port":{},"f:protocol":{},"f:targetPort":{}}},"f:selector":{".":{},"f:app":{}},"f:sessionAffinity":{},"f:type":{}}}
      }
    ]
  },
  "spec": {
    "ports": [
      {
        "name": "http",
        "protocol": "TCP",
        "port": 12345,
        "targetPort": 12345
      }
    ],
    "selector": {
      "app": "matrix-deployment"
    },
    "clusterIP": "10.96.185.219",
    "clusterIPs": [
      "10.96.185.219"
    ],
    "type": "ClusterIP",
    "sessionAffinity": "None",
    "ipFamilies": [
      "IPv4"
    ],
    "ipFamilyPolicy": "SingleStack"
  },
  "status": {
    "loadBalancer": {

    }
  }
}



% kubectl get pod |grep nginx                           
ingress-nginx-admission-create-vjn92                    0/1     Completed          0          117d
ingress-nginx-admission-patch-wlq6p                     0/1     Completed          0          117d
ingress-nginx-controller-57648496fc-84wl8               1/1     Running            38         117d

 % kubectl logs -f ingress-nginx-controller-57648496fc-84wl8

 E1224 03:47:42.495895      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-with-auth in namespace default"
W1224 03:48:17.091904      13 controller.go:906] Error obtaining Endpoints for Service "default/apple-service": no object matching key "default/apple-service" in local store
W1224 03:48:20.426021      13 controller.go:906] Error obtaining Endpoints for Service "default/apple-service": no object matching key "default/apple-service" in local store
E1224 03:48:35.929136      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-with-auth in namespace default"
E1224 03:49:42.512480      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-with-auth in namespace default"
E1224 03:51:42.534704      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-with-auth in namespace default"
E1224 03:52:48.862859      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-matrix in namespace default"
E1224 03:53:42.562522      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-with-auth in namespace default"
W1224 03:53:53.834938      13 controller.go:906] Error obtaining Endpoints for Service "default/apple-service": no object matching key "default/apple-service" in local store
W1224 03:53:57.167650      13 controller.go:906] Error obtaining Endpoints for Service "default/apple-service": no object matching key "default/apple-service" in local store





E1224 03:57:26.987047      13 server.go:84] "failed to process webhook request" err="rejecting admission review because the request does not contain an Ingress resource but networking.k8s.io/v1, Kind=Ingress with name ingress-matrix in namespace default"


kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission


https://github.com/kubernetes/minikube/issues/11121


 % kubectl get ValidatingWebhookConfiguration ingress-nginx-admission  -o yaml > ingress-nginx-admission.yaml 


  % kubectl delete  ValidatingWebhookConfiguration ingress-nginx-admission 
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted


% kubectl apply -f ingress.yaml                                          
ingress.networking.k8s.io/ingress-matrix created

% kubectl apply -f  ingress-nginx-admission.yaml 
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created



W1224 04:13:42.976147      13 controller.go:994] Service "default/matrix-deployment" does not have any active Endpoint.
^@^@^@^@^@^@

% kubectl get endpoints |grep matrix-deployment
matrix-deployment                     


https://kubernetes.io/zh/docs/concepts/services-networking/service/

% kubectl get pod |grep matrix-deployment
matrix-deployment-5ccd76bb66-lnxbm                      0/1     CrashLoopBackOff   11         34m


 % kubectl describe pod matrix-deployment-5ccd76bb66-lnxbm 
Name:         matrix-deployment-5ccd76bb66-lnxbm
Namespace:    default
Priority:     0
Node:         docker-desktop/192.168.65.4
Start Time:   Fri, 24 Dec 2021 11:48:17 +0800
Labels:       app=matrix-deployment
              pod-template-hash=5ccd76bb66
              version=stable
Annotations:  sidecar.istio.io/inject: false
Status:       Running
IP:           10.1.27.47
IPs:
  IP:           10.1.27.47
Controlled By:  ReplicaSet/matrix-deployment-5ccd76bb66
Containers:
  matrix-deployment:
    Container ID:   docker://bc78f2b5cc769cc3b5bd94120bfcfea1c46e795d11f07ec6f28b3cdd3d9ca7af
    Image:          matrix:0.0.3
    Image ID:       docker://sha256:f7ef9ed85535e1c4dace11ca60f5275cd1412d523f8fce2808bfe016eea8910d
    Port:           12345/TCP
    Host Port:      0/TCP
    State:          Waiting
      Reason:       CrashLoopBackOff
    Last State:     Terminated
      Reason:       Error
      Exit Code:    1
      Started:      Fri, 24 Dec 2021 12:19:39 +0800
      Finished:     Fri, 24 Dec 2021 12:19:39 +0800
    Ready:          False
    Restart Count:  11
    Limits:
      cpu:     300m
      memory:  500Mi
    Requests:
      cpu:        100m
      memory:     100Mi
    Environment:  <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-d4frg (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  kube-api-access-d4frg:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age                    From               Message
  ----     ------     ----                   ----               -------
  Normal   Scheduled  34m                    default-scheduler  Successfully assigned default/matrix-deployment-5ccd76bb66-lnxbm to docker-desktop
  Normal   Pulled     33m (x5 over 34m)      kubelet            Container image "matrix:0.0.3" already present on machine
  Normal   Created    33m (x5 over 34m)      kubelet            Created container matrix-deployment
  Normal   Started    33m (x5 over 34m)      kubelet            Started container matrix-deployment
  Warning  BackOff    4m49s (x143 over 34m)  kubelet            Back-off restarting failed container



CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o matrix .



% kubectl get ingress ingress-with-auth  -o yaml >ingress-with-auth.yaml


 % kubectl delete  ingress ingress-with-auth
ingress.networking.k8s.io "ingress-with-auth" deleted


 % curl http://localhost:80/httpInc
httpInc%                            



 % ab -n 200 http://localhost:80/httpInc 
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Finished 200 requests


