https://kubernetes.io/zh/docs/concepts/services-networking/service/
https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/namespaces/


写上namespace 否则是default
% kubectl apply -f sample/apple.cattle-system.yaml
pod/apple-app created
service/apple-service created

 % kubectl get Ingress  -n cattle-system

 Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME                  CLASS    HOSTS       ADDRESS     PORTS   AGE
rancher-k8s-ingress   <none>   localhost   localhost   80      37h

 kubectl edit Ingress rancher-k8s-ingress  -n cattle-system


或者，你也可以使用 kubectl edit 更新资源：

kubectl edit deployment/my-nginx
这相当于首先 get 资源，在文本编辑器中编辑它，然后用更新的版本 apply 资源：

kubectl get deployment my-nginx -o yaml > /tmp/nginx.yaml

https://kubernetes.io/zh/docs/concepts/cluster-administration/manage-deployment/


apiVersion: extensions/v1beta1
apiVersion: networking.k8s.io/v1

      - backend:
          serviceName: apple-service
          servicePort: 5678
        path: /cattle/apple
        pathType: Prefix

Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
error: ingresses.extensions "rancher-k8s-ingress" is invalid
A copy of your changes has been stored to "/var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-6uad1.yaml"
error: At least one of apiVersion, kind and name was changed


 % mv /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-6uad1.yaml sample/apple.cattle-system.ingress.yaml

 % kubectl delete -f sample/ingress-rancher-k8s.yaml
ingress.networking.k8s.io "rancher-k8s-ingress" deleted

% kubectl apply -f sample/apple.cattle-system.ingress.yaml
Error from server (InternalError): error when creating "sample/apple.cattle-system.ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1beta1/ingresses?timeout=10s": dial tcp 10.111.205.18:443: connect: connection refused


先干掉rancher
      - backend:
          service:
            name: rancher-k8s
            port:
              number: 80
        path: /rancher
        pathType: Prefix


 % kubectl apply -f sample/apple.cattle-system.ingress.yaml
Error from server (InternalError): error when creating "sample/apple.cattle-system.ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1beta1/ingresses?timeout=10s": dial tcp 10.111.205.18:443: connect: connection refused

https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.34.1/deploy/static/provider/cloud/deploy.yaml

https://stackoverflow.com/questions/65360431/internal-error-occurred-failed-calling-webhook-validate-nginx-ingress-kubernet

https://zhuanlan.zhihu.com/p/258729984

https://blog.csdn.net/cds992/article/details/106246616/


干掉

# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
# ingresses.extensions "rancher-k8s-ingress" was not valid:
# * : Invalid value: "The edited file failed validation": [ValidationError(Ingress.spec.rules[0].http.paths[0].backend): unknown field "serviceName" in io.k8s.api.networking.v1.IngressBackend, ValidationError(Ingress.spec.rules[0].http.paths[0].backend): unknown field "servicePort" in io.k8s.api.networking.v1.IngressBackend, ValidationError(Ingress.spec.rules[0].http.paths[1].backend): unknown field "serviceName" in io.k8s.api.networking.v1.IngressBackend, ValidationError(Ingress.spec.rules[0].http.paths[1].backend): unknown field "servicePort" in io.k8s.api.networking.v1.IngressBackend]
#

  creationTimestamp: "2021-03-24T12:45:35Z"
  generation: 2

  resourceVersion: "230863"
  uid: 0c6d830a-1342-4f0d-a185-3a33ad7a141c
  
    field.cattle.io/publicEndpoints: '[{"addresses":[""],"port":80,"protocol":"HTTP","serviceName":"cattle-system:rancher-k8s","ingressName":"cattle-system:rancher-k8s-ingress","hostname":"localhost","path":"/rancher","allNodes":false}]'

  selfLink: /apis/extensions/v1beta1/namespaces/cattle-system/ingresses/rancher-k8s-ingress


      - backend:
          service:
            name: rancher-k8s
            port:
              number: 80
        path: /rancher
        pathType: Prefix

    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"networking.k8s.io/v1","kind":"Ingress","metadata":{"annotations":{"ingress.kubernetes.io/rewrite-target":"/","nginx.ingress.kubernetes.io/rancherk8s":"http://rancher-k8s.cattle-system.svc.cluster.local:80"},"name":"rancher-k8s-ingress","namespace":"cattle-system"},"spec":{"rules":[{"host":"localhost","http":{"paths":[{"backend":{"service":{"name":"rancher-k8s","port":{"number":80}}},"path":"/rancher","pathType":"Prefix"}]}}]}}
    nginx.ingress.kubernetes.io/rancherk8s: http://rancher-k8s.cattle-system.svc.cluster.local:80


% kubectl get ns -o yaml > sample/apple.my-ns.ns.yaml
% kubectl apply -f sample/apple.my-ns.ns.yaml
namespace/my-ns created

% kubectl apply -f sample/apple.my-ns.yaml
pod/apple-app created
service/apple-service created

% kubectl apply -f sample/apple.my-ns.ingress.yaml
Error from server (InternalError): error when creating "sample/apple.my-ns.ingress.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1beta1/ingresses?timeout=10s": service "ingress-nginx-controller-admission" not found


https://stackoverflow.com/questions/61616203/nginx-ingress-controller-failed-calling-webhook



validate.nginx.ingress.kubernetes.io
    matchPolicy: Equivalent
    rules:
      - apiGroups:
          - networking.k8s.io
        apiVersions:
          - v1beta1
        operations:
          - CREATE
          - UPDATE
        resources:
          - ingresses
    failurePolicy: Fail
    sideEffects: None
    admissionReviewVersions:
      - v1
      - v1beta1
    clientConfig:
      service:
        namespace: ingress-nginx
        name: ingress-nginx-controller-admission
        path: /networking/v1beta1/ingresses


 % kubectl get -A ValidatingWebhookConfiguration
NAME                      WEBHOOKS   AGE
cert-manager-webhook      1          41h
ingress-nginx-admission   1          2d1h
rancher.cattle.io         2          41h


 % kubectl get -A ValidatingWebhookConfiguration -o yaml > ValidatingWebhookConfiguration.all.yaml


 % kubectl delete ValidatingWebhookConfiguration -A
error: resource(s) were provided, but no name, label selector, or --all flag specified


% kubectl delete ValidatingWebhookConfiguration --all
validatingwebhookconfiguration.admissionregistration.k8s.io "cert-manager-webhook" deleted
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted
validatingwebhookconfiguration.admissionregistration.k8s.io "rancher.cattle.io" deleted


% kubectl apply -f sample/apple.my-ns.ingress.yaml
ingress.networking.k8s.io/rancher-k8s-ingress created

% kubectl get Ingress -n my-ns
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
NAME                  CLASS    HOSTS       ADDRESS   PORTS   AGE
rancher-k8s-ingress   <none>   localhost             80      34s


% kubectl describe Ingress rancher-k8s-ingress -n my-ns
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Name:             rancher-k8s-ingress
Namespace:        my-ns
Address:
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host        Path  Backends
  ----        ----  --------
  localhost
              /my-ns/apple   apple-service:5678 (<none>)
Annotations:  ingress.kubernetes.io/rewrite-target: /
Events:       <none>

% curl http://localhost/my-ns/apple
curl: (7) Failed to connect to localhost port 80: Connection refused
