https://kubernetes.github.io/ingress-nginx/examples/auth/external-auth/

% curl -i https://httpbin.org/basic-auth/user/passwd         
HTTP/2 401 
date: Thu, 26 Aug 2021 02:40:51 GMT
content-length: 0
server: gunicorn/19.9.0
www-authenticate: Basic realm="Fake Realm"
access-control-allow-origin: *
access-control-allow-credentials: true

% curl -i https://user:passwd@httpbin.org/basic-auth/user/passwd
HTTP/2 200 
date: Thu, 26 Aug 2021 02:41:04 GMT
content-type: application/json
content-length: 47
server: gunicorn/19.9.0
access-control-allow-origin: *
access-control-allow-credentials: true

{
  "authenticated": true, 
  "user": "user"
}

% kubectl apply -f ingress-external-auth.yaml 
Error from server (InternalError): error when creating "ingress-external-auth.yaml": Internal error occurred: failed calling webhook "validate.nginx.ingress.kubernetes.io": an error on the server ("") has prevented the request from succeeding

kubectl delete -f ../../ingress-controller.yaml
kubectl delete -f ../basicAuth/ingress-with-auth.yaml
kubectl apply -f ingress-external-auth.yaml 
kubectl apply -f ../../ingress-controller.yaml


curl 127.0.0.1/apple                     


<html>
<head><title>401 Authorization Required</title></head>
<body>
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx</center>
</body>
</html>

% curl user:passwd@127.0.0.1/apple
/apple%                                          