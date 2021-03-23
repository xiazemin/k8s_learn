openssl genrsa -out tls.key 2048

% openssl req -new -x509 -key tls.key -out tls.crt -subj /C=CN/ST=GuangDong/L=GuangZhou/O=DevOps/CN=tomcat.kubernetes.io -days 3650 ca0gu0@ca0gu0deMBP


% kubectl create secret tls tomcat-ingress-secret --cert=tls.crt --key=tls.key -n testing
secret/tomcat-ingress-secret created

% kubectl apply -f tomcat-ingress-tls.yamlWarning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Error from server (InternalError): error when creating "tomcat-ingress-tls.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": Post "https://ingress-nginx-controller-admission.ingress-nginx.svc:443/networking/v1beta1/ingresses?timeout=10s": x509: certificate signed by unknown authority

https://segmentfault.com/a/1190000023375829