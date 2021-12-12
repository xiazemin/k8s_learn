% kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml 
error: error validating "deploy/ValidatingWebhookConfiguration.yaml": error validating data: [ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "clientConfig" in io.k8s.api.admissionregistration.v1.ValidatingWebhook, ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "sideEffects" in io.k8s.api.admissionregistration.v1.ValidatingWebhook, ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "admissionReviewVersions" in io.k8s.api.admissionregistration.v1.ValidatingWebhook]; if you choose to ignore these errors, turn validation off with --validate=false


kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
error: error validating "deploy/ValidatingWebhookConfiguration.yaml": error validating data: [ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "sideEffects" in io.k8s.api.admissionregistration.v1.ValidatingWebhook, ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "admissionReviewVersions" in io.k8s.api.admissionregistration.v1.ValidatingWebhook]; if you choose to ignore these errors, turn validation off with --validate=false


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
error: error validating "deploy/ValidatingWebhookConfiguration.yaml": error validating data: ValidationError(ValidatingWebhookConfiguration.webhooks[0]): missing required field "sideEffects" in io.k8s.api.admissionregistration.v1.ValidatingWebhook; if you choose to ignore these errors, turn validation off with --validate=false


 


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
error: error when retrieving current configuration of:
Resource: "admissionregistration.k8s.io/v1, Resource=validatingwebhookconfigurations", GroupVersionKind: "admissionregistration.k8s.io/v1, Kind=ValidatingWebhookConfiguration"
Name: "", Namespace: ""
from server for: "deploy/ValidatingWebhookConfiguration.yaml": resource name may not be empty


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
The ValidatingWebhookConfiguration "admission-webhook-example" is invalid: 
* webhooks[0].name: Invalid value: "admission-webhook-example-deployment-5bc4865848-v9fvb:7896/validate": a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')
* webhooks[0].rules[0].operations[0]: Unsupported value: "GET": supported values: "*", "CONNECT", "CREATE", "DELETE", "UPDATE"


 


 


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
The ValidatingWebhookConfiguration "admission-webhook-example" is invalid: 
* webhooks[0].name: Invalid value: "admission-webhook-example": should be a domain with at least three segments separated by dots
* webhooks[0].rules[0].operations[0]: Unsupported value: "GET": supported values: "*", "CONNECT", "CREATE", "DELETE", "UPDATE"


 


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
The ValidatingWebhookConfiguration "admission-webhook-example" is invalid: 
* webhooks[0].name: Invalid value: "admission-webhook-example.default": should be a domain with at least three segments separated by dots
* webhooks[0].rules[0].operations[0]: Unsupported value: "GET": supported values: "*", "CONNECT", "CREATE", "DELETE", "UPDATE"


 


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
The ValidatingWebhookConfiguration "admission-webhook-example" is invalid: webhooks[0].rules[0].operations[0]: Unsupported value: "GET": supported values: "*", "CONNECT", "CREATE", "DELETE", "UPDATE"


 kubectl apply -f deploy/ValidatingWebhookConfiguration.yaml
validatingwebhookconfiguration.admissionregistration.k8s.io/admission-webhook-example created



% kubectl describe ValidatingWebhookConfiguration admission-webhook-example
Name:         admission-webhook-example
Namespace:    
Labels:       <none>
Annotations:  <none>
API Version:  admissionregistration.k8s.io/v1
Kind:         ValidatingWebhookConfiguration
Metadata:
  Creation Timestamp:  2021-12-12T03:01:55Z
  Generation:          1
  Managed Fields:
    API Version:  admissionregistration.k8s.io/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:metadata:
        f:annotations:
          .:
          f:kubectl.kubernetes.io/last-applied-configuration:
      f:webhooks:
        .:
        k:{"name":"admission-webhook.example.default"}:
          .:
          f:admissionReviewVersions:
          f:clientConfig:
            .:
            f:url:
          f:failurePolicy:
          f:matchPolicy:
          f:name:
          f:namespaceSelector:
          f:objectSelector:
          f:rules:
          f:sideEffects:
          f:timeoutSeconds:
    Manager:         kubectl-client-side-apply
    Operation:       Update
    Time:            2021-12-12T03:01:55Z
  Resource Version:  4028164
  UID:               fc052e5c-2383-4390-9ca8-bcfa85889bfb
Webhooks:
  Admission Review Versions:
    v1
    v1beta1
  Client Config:
    URL:           https://admission-webhook-example-deployment-5bc4865848-v9fvb:7896/validate
  Failure Policy:  Fail
  Match Policy:    Equivalent
  Name:            admission-webhook.example.default
  Namespace Selector:
  Object Selector:
  Rules:
    API Groups:
      *
    API Versions:
      *
    Operations:
      UPDATE
    Resources:
      */status
    Scope:          *
  Side Effects:     None
  Timeout Seconds:  1
Events:             <none>

https://kubernetes.io/zh/docs/reference/access-authn-authz/extensible-admission-controllers/#webhook-configuration


k8s把pod 暴露为服务
kubectl expose deployment admission-webhook-example-deployment --type=NodePort
error: couldn't find port via --port flag or introspection
See 'kubectl expose -h' for help and examples


% kubectl apply -f deploy/service.yaml 
The Service "admission-webhook-example" is invalid: spec.ports[0].nodePort: Invalid value: 37896: provided port is not in the valid range. The range of valid ports is 30000-32767

% kubectl apply -f deploy/service.yaml
service/admission-webhook-example created

% telnet 127.0.0.1 30896
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Connection closed by foreign host.


  selector:
    app: admission-webhook-example

% telnet 127.0.0.1 30896                        
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

% curl 127.0.0.1:30896                          
Client sent an HTTP request to an HTTPS server.


 % kubectl edit deployment example-foo 
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example:7896/validate?timeout=1s": dial tcp: lookup admission-webhook-example on 192.168.65.5:53: no such host
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-9cd4q.yaml` to try this update again.


 % kubectl logs -f admission-webhook-example-deployment-5bc4865848-v9fvb
I1212 02:43:35.063846       1 main.go:50] Server started
2021/12/12 03:18:19 http: TLS handshake error from 192.168.65.3:58566: tls: first record does not look like a TLS handshake


https://admission-webhook-example.default:7896/validate


% kubectl edit deployment example-foo                        
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example.default:7896/validate?timeout=1s": dial tcp: lookup admission-webhook-example.default on 192.168.65.5:53: no such host
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-qnlov.yaml` to try this update again.

 % kubectl exec -it dnsutils sh
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
# 
# nslookup admission-webhook-example.default 
Server:         10.96.0.10
Address:        10.96.0.10#53

Name:   admission-webhook-example.default.svc.cluster.local
Address: 10.111.163.92



 % kubectl edit deployment example-foo                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example.default.svc.cluster.local:7896/validate?timeout=1s": context deadline exceeded
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-qrf84.yaml` to try this update again.

 % kubectl edit deployment example-foo                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example.default.svc.cluster.local:30896/validate?timeout=1s": context deadline exceeded
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-fcemi.yaml` to try this update again.



 % curl https://127.0.0.1:30896
curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.



超时时间改成10


 % kubectl edit deployment example-foo                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example.default.svc.cluster.local:7896/validate?timeout=10s": dial tcp: lookup admission-webhook-example.default.svc.cluster.local on 192.168.65.5:53: read udp 192.168.65.4:57839->192.168.65.5:53: i/o timeout
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-0uw3h.yaml` to try this update again.


   url: "https://docker.for.mac.host.internal:30896/validate" 


% kubectl edit deployment example-foo                        
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://docker.for.mac.host.internal:30896/validate?timeout=10s": x509: certificate is valid for sidecar-injector-webhook-svc, sidecar-injector-webhook-svc.default, sidecar-injector-webhook-svc.default.svc, not docker.for.mac.host.internal
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-46c27.yaml` to try this update again.

问题原因
我们的kubernetes的apiserver-advertise-address是一个内网IP，默认情况下，kubernetes自建的CA会为apiserver签发一个证书，证书的默认可访问的是内网IP、kubernetes、kubernetes.default kubernetes.default.svc、kubernetes.default.svc.cluster.local，不包含设备的外网IP。所以直接通过admin.conf去访问kubernetes是不可能的。

https://blog.csdn.net/DANTE54/article/details/105297228

    url: "https://kubernetes.default.svc.cluster.local:30896/validate" 



     % kubectl edit deployment example-foo                                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://kubernetes.default.svc.cluster.local:30896/validate?timeout=10s": dial tcp: lookup kubernetes.default.svc.cluster.local on 192.168.65.5:53: read udp 192.168.65.4:65346->192.168.65.5:53: i/o timeout
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-68d14.yaml` to try this update again.


% kubectl edit deployment example-foo                                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://admission-webhook-example.default.svc:7896/validate?timeout=10s": x509: certificate is valid for sidecar-injector-webhook-svc, sidecar-injector-webhook-svc.default, sidecar-injector-webhook-svc.default.svc, not admission-webhook-example.default.svc
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-rnnxb.yaml` to try this update again.


修改service名字
metadata:
  name: sidecar-injector-webhook-svc

% kubectl edit deployment example-foo                        
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": Post "https://sidecar-injector-webhook-svc.default.svc:7896/validate?timeout=10s": x509: certificate signed by unknown authority
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-kdcd1.yaml` to try this update again.


caBundle: "Ci0tLS0tQk...<base64-encoded PEM bundle containing the CA that signed the webhook's serving certificate>...tLS0K"
    

    caBundle: MIIDMzCCAhsCAQAwMzExMC8GA1UEAwwoc2lkZWNhci1pbmplY3Rvci13ZWJob29rLXN2Yy5kZWZhdWx0LnN2YzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKOw2eVcCYhm4Pq/WUPJTGjMvzV64MACctjyEXO25DrF3NAHeExb+7jdwV7RJDbFFDuJz2P/x8DFGvLfv2lMlEmKPQZVhAkdFwrU9eWtXf6BJFytrqxrZv9L2PFGsAToSCnd0/VFkh+bHDF9gFhyt2nxGFskV6ZzExpK+nBjM1k2eo5SHsQQbFi95nV+NMvIw6rFcEMfjt4Tr/VfSd81AMHQkrNfx+mhg29FKYzUfW3nodwkdAHVNiB5csqzBPNCOkEc8ghrWq8ei3lxcVryHi28CW7KzhGK1aWWo90EVp+vAUUesZaeWCF7SJmmwm5WX/GlrgYcbR6SzrL1rGA1BzsCAwEAAaCBujCBtwYJKoZIhvcNAQkOMYGpMIGmMAkGA1UdEwQCMAAwCwYDVR0PBAQDAgXgMBMGA1UdJQQMMAoGCCsGAQUFBwMBMHcGA1UdEQRwMG6CHHNpZGVjYXItaW5qZWN0b3Itd2ViaG9vay1zdmOCJHNpZGVjYXItaW5qZWN0b3Itd2ViaG9vay1zdmMuZGVmYXVsdIIoc2lkZWNhci1pbmplY3Rvci13ZWJob29rLXN2Yy5kZWZhdWx0LnN2YzANBgkqhkiG9w0BAQsFAAOCAQEAYX69adcqE9VM/p0Vcc5TWtlNzCTmwMK+gU/+DjNIbv6VltbnjYnIV0u62rhyA2PF5WZSqTu3e3UhV8oL5HrF1PyWbhbECUVEtQsWb8q+lwl4Uu1vFw6+dO5yMe/j12vNhxRRnhIKGJlsdetclz2Q9MMsIBA+hxZH3f728+Q4OyVkZ6qTPmlnJLvJbcBYjbGfWffRItXlIzrmI3/IOi7Mf8YwpiY85RYJ7zxneNsJXZhC949T943mGflSHfdHNNhVMGhtUEPnEA8LrJjlOkfTJVllxkM0zPFYVOtdwotjxy74sHgQd43y3/zYX5z+yuaaj9tGKWgbFaoNrTKdiqXfmQ==


ca-bundle.crt 文件存储了各大证书颁发证的根证书交叉文件。curl 访问https网站时，会比对这个文件里的根证书。如果这个文件过老，那就是有新的根证书未加入到这个文件里。导致curl无法正常访问https网站。

