https://www.namecheap.com/support/knowledgebase/article.aspx/9393/69/where-do-i-find-ssl-ca-bundle/

CA bundle is a file that contains root and intermediate certificates. The end-entity certificate along with a CA bundle constitutes the certificate chain.

The chain is required to improve compatibility of the certificates with web browsers and other kind of clients so that browsers recognize your certificate and no security warnings appear.


https://stackoverflow.com/questions/60633405/use-secret-for-validatingwebhookconfiguration

% sh deploy/ca-k8s.sh 
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1ekNDQWMrZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeE1EZ3lNekF5TWpFek1Gb1hEVE14TURneU1UQXlNakV6TUZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTGNwCmVIejZSMkp4TElwdmFyak03cDVNNExVcitLa1h4UjRlMmJyNHRZVmRNSnQxUS9SNHFKNE85UFlSbE5WcXNRNnYKSUJjQkRTMHRIN3ZvK21pR0g3dFVYZHFnUVFYYlg4S0JhV0tsbG9Da2pSK245Qkl6MkNwZE81Z3pKbzVzd05NNwoyZEx1d0E5OUY2TU4vLzdvWDNhSW5DeVVGY0ltRGt3cVI0SGhGWExIbHFweWlXWE1QekRrdU5ERHhxQktIM3ViCldFRTZiWGEwbU5WQTdlT1lpOWViRG5yQ2lkQlo2d1ZNYkpvOHpqQmdJSVJuMGY0Z3pXelBhNGNoN1kra3oySS8KUUU2TUczYlZyaWY0Y2hvRStWV1l2MEpCOUJRakZENThBcGdHdVBuNXNlVjVoU0VKNjdnVkJnUlJ5YUFBbjJLZApmTkp4bkxIeU1sbkNvLzhCQm0wQ0F3RUFBYU5DTUVBd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZPNmg4bUFVbEovZjRZSzNkZ0t2U2FFYWFBb0xNQTBHQ1NxR1NJYjMKRFFFQkN3VUFBNElCQVFDYWpmRXdzU3JPYXJkb1FPT3I0UWxqRmZENUJIb2p0OVM5bWREN2YwYlFFQ0dQZE91VQpLSkdqMEZsdklNejlPYWFVZlVHTVhESjlFTlZXTDg3STFjeENKOGtxVDZCVjVQUG9sNTZuSEVZL0pEZUdpR3hCCnltcG1SWjFYZ1lIeXBUTy9HYXBJNzZ1TXJEK1MwaXJhU29LdlZ5cnJqT0dvdXlGeldKdWs4Vk1jTldGOGkvTnoKbnFnNExoVyt5aTRtWWNYdFZJd1Q5RXVtQ0l4U3FBSUhDbHBaelVQQUlENVRlM0FiMFdiK1E5bmQvYVZOWjY3bQpXeFEvSUx2VXFkN2ZPdGk0WG5HV0RpazNEbmN0ejBIVk8xUXh0NzFKQit0QWdMNzBRcW1vcmZYK1lxQ3FhQUV0CmdVeEJhQnRKT1NWRW8zTDM3QlpKNW01dXdNNzJ5WlhxamxWTAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==

https://cuisongliu.github.io/2020/07/kubernetes/admission-webhook/
https://cuisongliu.github.io/2020/07/kubernetes/admission-webhook/#%e8%ae%be%e7%bd%aeca%e8%af%81%e4%b9%a6%e7%bb%99%e5%af%b9%e5%ba%94%e7%9a%84webhook%e8%b5%84%e6%ba%90


 % kubectl edit deployment example-foo                        
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": webhook response was absent
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-cmb9b.yaml` to try this update again.


 % kubectl logs -f  admission-webhook-example-deployment-7bbc77f89c-kv7xb
I1212 05:06:20.385425       1 webhook.go:52] Ready to write reponse ...
I1212 05:06:20.428713       1 webhook.go:52] Ready to write reponse ...
I1212 05:06:20.476906       1 webhook.go:52] Ready to write reponse ...
I1212 05:06:20.523446       1 webhook.go:52] Ready to write reponse ...


 % kubectl edit deployment example-foo                         
error: deployments.apps "example-foo" could not be patched: Internal error occurred: failed calling webhook "admission-webhook.example.default": webhook response was absent
You can run `kubectl replace -f /var/folders/hb/1q06vpss7_d1fw3ymxb1ztbr0000gn/T/kubectl-edit-e773k.yaml` to try this update again.


 % kubectl logs -f admission-webhook-example-deployment-7b996d5c64-z4fx6

I1212 05:37:17.506784       1 webhook.go:177] kube-scheduler%!(EXTRA types.UID=ea35d340-8ccb-4ca6-8817-af3874fb5021)
I1212 05:37:17.506805       1 webhook.go:56] Ready to write reponse ..



https://github.com/kubernetes/apiserver/blob/master/pkg/admission/plugin/webhook/request/admissionreview.go

func VerifyAdmissionResponse(uid types.UID, mutating bool, review runtime.Object) (*AdmissionResponse, error) {
case *admissionv1beta1.AdmissionReview:
if r.Response == nil {
    return nil, fmt.Errorf("webhook response was absent")
}


https://github.com/kubernetes/apiserver/blob/623d4c094b8511d0d746b6f3a8475ea8297c2946/pkg/admission/plugin/webhook/validating/dispatcher.go#L251

func (d *validatingDispatcher) callHook(ctx context.Context, h *v1.ValidatingWebhook, invocation *generic.WebhookInvocation, attr *generic.VersionedAttributes) error {
	if attr.Attributes.IsDryRun() {
uid, request, response, err := webhookrequest.CreateAdmissionObjects(attr, invocation)
result, err := webhookrequest.VerifyAdmissionResponse(uid, false, response)


https://github.com/kubernetes/apiserver/blob/623d4c094b8511d0d746b6f3a8475ea8297c2946/pkg/admission/plugin/webhook/request/admissionreview.go#L133

func CreateAdmissionObjects(versionedAttributes *generic.VersionedAttributes, invocation *generic.WebhookInvocation) (uid types.UID, request, response runtime.Object, err error) {
	for _, version := range invocation.Webhook.GetAdmissionReviewVersions() {
		switch version {



返回值类型改成
v1beta1.AdmissionReview{}


 %  kubectl edit deployment example-foo 
deployment.apps/example-foo edited


 %  kubectl edit deployment example-foo 
deployment.apps/example-foo edited

 % kubectl logs admission-webhook-example-deployment-7b996d5c64-ddvgp |grep example-foo
I1212 06:17:13.094363       1 webhook.go:185] example-foo%!(EXTRA types.UID=a88d558d-8d1d-4237-b52f-a0500a3bb651, string=admission.k8s.io/v1, string=example-foo, v1beta1.Operation=UPDATE, string=docker-for-desktop, map[string]v1.ExtraValue=map[], []string=[system:masters system:authenticated])




