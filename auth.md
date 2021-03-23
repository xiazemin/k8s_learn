https://kubernetes.github.io/ingress-nginx/examples/auth/basic/

% htpasswd -c auth foo
New password:
Re-type new password:
Adding password for user foo

% kubectl create secret generic basic-auth --from-file=auth
secret/basic-auth created

% kubectl get secret basic-auth -o yaml
apiVersion: v1
data:
  auth: Zm9vOiRhcHIxJHg2YlJNUTBFJEhFVXhEVXJTWU9lQWlaYzdnVm5tZzEK
kind: Secret
metadata:
  creationTimestamp: "2021-03-23T06:30:46Z"
  managedFields:
  - apiVersion: v1
    fieldsType: FieldsV1
    fieldsV1:
      f:data:
        .: {}
        f:auth: {}
      f:type: {}
    manager: kubectl-create
    operation: Update
    time: "2021-03-23T06:30:46Z"
  name: basic-auth
  namespace: default
  resourceVersion: "83786"
  selfLink: /api/v1/namespaces/default/secrets/basic-auth
  uid: a5a2bf72-c7ed-47e8-aa93-59764ba0ab88
type: Opaque

 echo "
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: ingress-with-auth
  annotations:
    # type of authentication
    nginx.ingress.kubernetes.io/auth-type: basic
    # name of the secret that contains the user/password definitions
    nginx.ingress.kubernetes.io/auth-secret: basic-auth
    # message to display with an appropriate context why the authentication is required
    nginx.ingress.kubernetes.io/auth-realm: 'Authentication Required - foo'
spec:
  rules:
  - host: foo.bar.com
    http:
      paths:
      - path: /
        backend:
          serviceName: http-svc
          servicePort: 80
" | kubectl create -f -
Warning: networking.k8s.io/v1beta1 Ingress is deprecated in v1.19+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Error from server (InternalError): error when creating "STDIN": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1beta1/ingresses?timeout=10s": x509: certificate signed by unknown authority

https://www.cnblogs.com/leozhanggg/p/13189173.html

https://github.com/kubernetes/ingress-nginx/issues/5401

https://stackoverflow.com/questions/61365202/nginx-ingress-service-ingress-nginx-controller-admission-not-found



 % kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission
validatingwebhookconfiguration.admissionregistration.k8s.io "ingress-nginx-admission" deleted

问题解决


kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.44.0/deploy/static/provider/cloud/deploy.yaml


https://www.infoq.cn/article/it1bqeyocrl9igvos-xy