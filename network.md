Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress


cp ingress-nginx.v4.yaml ingress-nginx.v7.yaml

 % kubectl delete -f ingress-nginx.v4.yaml  --force

 % kubectl apply -f ingress-nginx.v7.yaml
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
configmap/ingress-nginx-controller created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
service/ingress-nginx-controller-admission created
service/ingress-nginx-controller created
deployment.apps/ingress-nginx-controller created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
serviceaccount/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created

% kubectl delete -f sample/apple.my-ns.ingress.yaml
ingress.networking.k8s.io "rancher-k8s-ingress" deleted

% kubectl delete ValidatingWebhookConfiguration --all
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted

% kubectl apply -f sample/apple.my-ns.ingress.yaml
ingress.networking.k8s.io/rancher-k8s-ingress created

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

 % curl http://127.0.0.1/my-ns/apple
curl: (7) Failed to connect to 127.0.0.1 port 80: Connection refused