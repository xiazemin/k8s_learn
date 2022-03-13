% kubectl api-versions | grep admissionregistration.k8s.io
admissionregistration.k8s.io/v1


make build 
make build-image

Building the sidecar-injector binary...
go: downloading k8s.io/api v0.0.0-20180127130940-acf347b865f2
go: downloading k8s.io/apimachinery v0.0.0-20180126010752-19e3f5aa3adc
 => exporting to image                                                                                          1.2s
 => => exporting layers                                                                                         1.2s
 => => writing image sha256:317679564487576c0e31db98eed8d723738bbd159d899b498275ca598f03873a                    0.0s
 => => naming to docker.io/morvencao/sidecar-injector:v20220313-b326755                                         0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them


 % docker images |grep sidecar-injector
morvencao/sidecar-injector                                           v20220313-b326755                                       317679564487   About a minute ago   20.3MB

%  kubectl create ns sidecar-injector
namespace/sidecar-injector created


./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector

creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.mzvPWXYB 
Generating RSA private key, 2048 bit long modulus
.....+++
..............................................................+++
e is 65537 (0x10001)
error: unable to recognize "STDIN": no matches for kind "CertificateSigningRequest" in version "certificates.k8s.io/v1beta1"


 % cat deploy/mutatingwebhook.yaml | \
    deploy/webhook-patch-ca-bundle.sh > \
    deploy/mutatingwebhook-ca-bundle.yaml



% cat deploy/mutatingwebhook.yaml | \
    deploy/webhook-patch-ca-bundle.sh > \
    deploy/mutatingwebhook-ca-bundle.yaml
% kubectl create -f deploy/nginxconfigmap.yaml
configmap/nginx-configmap created
% kubectl create -f deploy/configmap.yaml
configmap/sidecar-injector-webhook-configmap created
% kubectl create -f deploy/deployment.yaml
deployment.apps/sidecar-injector-webhook-deployment created
% kubectl create -f deploy/service.yaml
service/sidecar-injector-webhook-svc created


% kubectl create -f deploy/mutatingwebhook-ca-bundle.yaml
error: unable to recognize "deploy/mutatingwebhook-ca-bundle.yaml": no matches for kind "MutatingWebhookConfiguration" in version "admissionregistration.k8s.io/v1beta1"


admissionregistration.k8s.io/v1beta1 换成 v1

% kubectl create -f deploy/mutatingwebhook-ca-bundle.yaml
error: error validating "deploy/mutatingwebhook-ca-bundle.yaml": error validating data: [ValidationError(MutatingWebhookConfiguration.webhooks[0]): missing required field "sideEffects" in io.k8s.api.admissionregistration.v1.MutatingWebhook, ValidationError(MutatingWebhookConfiguration.webhooks[0]): missing required field "admissionReviewVersions" in io.k8s.api.admissionregistration.v1.MutatingWebhook]; if you choose to ignore these errors, turn validation off with --validate=false


% kubectl create -f deploy/mutatingwebhook-ca-bundle.yaml --validate=false
The MutatingWebhookConfiguration "sidecar-injector-webhook-cfg" is invalid: 
* webhooks[0].sideEffects: Required value: must specify one of None, NoneOnDryRun
* webhooks[0].admissionReviewVersions: Required value: must specify one of v1, v1beta1




% kubectl -n sidecar-injector get pod
NAME                                                   READY   STATUS              RESTARTS   AGE
sidecar-injector-webhook-deployment-7689959d4b-8vhq4   0/1     ContainerCreating   0          64s

% kubectl -n sidecar-injector get deploy
NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
sidecar-injector-webhook-deployment   0/1     1            0           11m



% kubectl -n sidecar-injector describe pod sidecar-injector-webhook-deployment-7689959d4b-8vhq4
Name:           sidecar-injector-webhook-deployment-7689959d4b-8vhq4
Namespace:      sidecar-injector
Priority:       0
Events:
  Type     Reason       Age                  From               Message
  ----     ------       ----                 ----               -------
  Normal   Scheduled    11m                  default-scheduler  Successfully assigned sidecar-injector/sidecar-injector-webhook-deployment-7689959d4b-8vhq4 to docker-desktop
  Warning  FailedMount  9m39s                kubelet            Unable to attach or mount volumes: unmounted volumes=[webhook-certs], unattached volumes=[webhook-config kube-api-access-qng8z webhook-certs]: timed out waiting for the condition
  Warning  FailedMount  86s (x13 over 11m)   kubelet            MountVolume.SetUp failed for volume "webhook-certs" : secret "sidecar-injector-webhook-certs" not found
  Warning  FailedMount  33s (x4 over 7m20s)  kubelet            Unable to attach or mount volumes: unmounted volumes=[webhook-certs], unattached volumes=[webhook-certs webhook-config kube-api-access-qng8z]: timed out waiting for the condition



apiVersion: certificates.k8s.io/v1beta1
apiVersion: certificates.k8s.io/v1


 % ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.9Juy3c4P 
Generating RSA private key, 2048 bit long modulus
..........+++
...+++
e is 65537 (0x10001)
error: error validating "STDIN": error validating data: ValidationError(CertificateSigningRequest.spec): missing required field "signerName" in io.k8s.api.certificates.v1.CertificateSigningRequestSpec; if you choose to ignore these errors, turn validation off with --validate=false



 % ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.QGfOrpGo 
Generating RSA private key, 2048 bit long modulus
.............................+++
..................+++
e is 65537 (0x10001)
error: error parsing STDIN: error converting YAML to JSON: yaml: line 7: mapping values are not allowed in this context


% ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.cuwhJztZ 
Generating RSA private key, 2048 bit long modulus
.......................+++
.............................+++
e is 65537 (0x10001)
error: error parsing STDIN: error converting YAML to JSON: yaml: line 7: mapping values are not allowed in this context




% ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.ypUeXjQa 
Generating RSA private key, 2048 bit long modulus
.............................................................+++
....................................+++
e is 65537 (0x10001)
error: error parsing STDIN: error converting YAML to JSON: yaml: line 14: could not find expected ':'



% ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.2pu5Qg1d 
Generating RSA private key, 2048 bit long modulus
..................................................+++
............................................+++
e is 65537 (0x10001)
The CertificateSigningRequest "sidecar-injector-webhook-svc.sidecar-injector" is invalid: spec.signerName: Invalid value: "xiazemin": must be a fully qualified domain and path of the form 'example.com/signer-name'





% ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.Gg0rHtnf 
Generating RSA private key, 2048 bit long modulus
......................................................................................+++
........+++
e is 65537 (0x10001)
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.sidecar-injector created
NAME                                            AGE   SIGNERNAME                REQUESTOR            REQUESTEDDURATION   CONDITION
sidecar-injector-webhook-svc.sidecar-injector   0s    example.com/signer-name   docker-for-desktop   <none>              Pending
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.sidecar-injector approved
ERROR: After approving csr sidecar-injector-webhook-svc.sidecar-injector, the signed certificate did not appear on the resource. Giving up after 10 attempts.



kubernetes.io/kube-apiserver-client-kubelet

https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/

The kube-controller-manager ships with a built-in approver for certificates with a signerName of kubernetes.io/kube-apiserver-client-kubelet that delegates various permissions on CSRs for node credentials to authorization. The kube-controller-manager POSTs SubjectAccessReview resources to the API server in order to check authorization for certificate approval.




% ./deploy/webhook-create-signed-cert.sh \
    --service sidecar-injector-webhook-svc \
    --secret sidecar-injector-webhook-certs \
    --namespace sidecar-injector
creating certs in tmpdir /var/folders/2n/42n_1dfd6kjd6s3k7bt4cb3h0000gn/T/tmp.VJQAGYph 
Generating RSA private key, 2048 bit long modulus
.................................................................................................................+++
............................................+++
e is 65537 (0x10001)
certificatesigningrequest.certificates.k8s.io "sidecar-injector-webhook-svc.sidecar-injector" deleted
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.sidecar-injector created
NAME                                            AGE   SIGNERNAME                                    REQUESTOR            REQUESTEDDURATION   CONDITION
sidecar-injector-webhook-svc.sidecar-injector   0s    kubernetes.io/kube-apiserver-client-kubelet   docker-for-desktop   <none>              Pending
certificatesigningrequest.certificates.k8s.io/sidecar-injector-webhook-svc.sidecar-injector approved
ERROR: After approving csr sidecar-injector-webhook-svc.sidecar-injector, the signed certificate did not appear on the resource. Giving up after 10 attempts.


https://github.com/istio/istio/issues/4368

https://medium.com/ibm-cloud/diving-into-kubernetes-mutatingadmissionwebhook-6ef3c5695f74


sidecar-injector-webhook-svc.default

https://kubernetes.io/docs/tasks/tls/managing-tls-in-a-cluster/

https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/#signers

