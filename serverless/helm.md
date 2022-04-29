https://artifacthub.io/packages/helm/softonic/knative-serving

https://knative.dev/docs/

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.4.0/serving-crds.yaml

customresourcedefinition.apiextensions.k8s.io/certificates.networking.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/configurations.serving.knative.dev created
customresourcedefinition.apiextensions.k8s.io/clusterdomainclaims.networking.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/domainmappings.serving.knative.dev created
customresourcedefinition.apiextensions.k8s.io/ingresses.networking.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/metrics.autoscaling.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/podautoscalers.autoscaling.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/revisions.serving.knative.dev created
customresourcedefinition.apiextensions.k8s.io/routes.serving.knative.dev created
customresourcedefinition.apiextensions.k8s.io/serverlessservices.networking.internal.knative.dev created
customresourcedefinition.apiextensions.k8s.io/services.serving.knative.dev created
customresourcedefinition.apiextensions.k8s.io/images.caching.internal.knative.dev created

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.4.0/serving-core.yaml

namespace/knative-serving unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-aggregated-addressable-resolver unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-addressable-resolver unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-admin unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-edit unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-view unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-core unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-podspecable-binding unchanged
serviceaccount/controller unchanged
clusterrole.rbac.authorization.k8s.io/knative-serving-admin unchanged
clusterrolebinding.rbac.authorization.k8s.io/knative-serving-controller-admin unchanged
clusterrolebinding.rbac.authorization.k8s.io/knative-serving-controller-addressable-resolver unchanged
customresourcedefinition.apiextensions.k8s.io/images.caching.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/certificates.networking.internal.knative.dev unchanged
^@^@^@^@^@^@^@customresourcedefinition.apiextensions.k8s.io/configurations.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/clusterdomainclaims.networking.internal.knative.dev unchanged
^@customresourcedefinition.apiextensions.k8s.io/domainmappings.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/ingresses.networking.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/metrics.autoscaling.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/podautoscalers.autoscaling.internal.knative.dev unchanged
^@customresourcedefinition.apiextensions.k8s.io/revisions.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/routes.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/serverlessservices.networking.internal.knative.dev unchanged
^@^@^@^@^@^@^@customresourcedefinition.apiextensions.k8s.io/services.serving.knative.dev unchanged
image.caching.internal.knative.dev/queue-proxy unchanged
configmap/config-autoscaler unchanged
configmap/config-defaults unchanged
configmap/config-deployment unchanged
configmap/config-domain unchanged
configmap/config-features unchanged
configmap/config-gc unchanged
configmap/config-leader-election unchanged
configmap/config-logging unchanged
configmap/config-network unchanged
configmap/config-observability unchanged
configmap/config-tracing unchanged
horizontalpodautoscaler.autoscaling/activator unchanged
poddisruptionbudget.policy/activator-pdb unchanged
deployment.apps/activator configured
service/activator-service unchanged
deployment.apps/autoscaler configured
service/autoscaler unchanged
deployment.apps/controller configured
service/controller unchanged
deployment.apps/domain-mapping unchanged
deployment.apps/domainmapping-webhook unchanged
service/domainmapping-webhook unchanged
horizontalpodautoscaler.autoscaling/webhook created
poddisruptionbudget.policy/webhook-pdb created
deployment.apps/webhook created
service/webhook created
validatingwebhookconfiguration.admissionregistration.k8s.io/config.webhook.serving.knative.dev created
mutatingwebhookconfiguration.admissionregistration.k8s.io/webhook.serving.knative.dev created
mutatingwebhookconfiguration.admissionregistration.k8s.io/webhook.domainmapping.serving.knative.dev created
secret/domainmapping-webhook-certs created
validatingwebhookconfiguration.admissionregistration.k8s.io/validation.webhook.domainmapping.serving.knative.dev created
validatingwebhookconfiguration.admissionregistration.k8s.io/validation.webhook.serving.knative.dev created
secret/webhook-certs created

kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.4.0/kourier.yaml
namespace/kourier-system created
configmap/kourier-bootstrap created
configmap/config-kourier created
serviceaccount/net-kourier created
clusterrole.rbac.authorization.k8s.io/net-kourier created
clusterrolebinding.rbac.authorization.k8s.io/net-kourier created
deployment.apps/net-kourier-controller created
service/net-kourier-controller created
deployment.apps/3scale-kourier-gateway created
service/kourier created
service/kourier-internal created


kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'

kubectl --namespace kourier-system get service kourier

https://blog.csdn.net/Small_StarOne/article/details/103815250

  Normal   Scheduled               <unknown>             default-scheduler        Successfully assigned knative-serving/activator-5f9b99c7b6-k4bdt to vm-4-15-centos
  Warning  FailedCreatePodSandBox  12m                   kubelet, vm-4-15-centos  Failed to create pod sandbox: rpc error: code = DeadlineExceeded desc = context deadline exceeded
  Normal   SandboxChanged          12m                   kubelet, vm-4-15-centos  Pod sandbox changed, it will be killed and re-created.
  Normal   Created                 10m (x2 over 11m)     kubelet, vm-4-15-centos  Created container activator
  Normal   Started                 10m (x2 over 11m)     kubelet, vm-4-15-centos  Started container activator
  Warning  Unhealthy               9m43s (x11 over 10m)  kubelet, vm-4-15-centos  Readiness probe failed: Get http://172.17.0.17:8012/: dial tcp 172.17.0.17:8012: connect: connection refused
  Normal   Pulled                  4m6s (x5 over 11m)    kubelet, vm-4-15-centos  Container image "gcr.io/knative-releases/knative.dev/serving/cmd/activator:latest" already present on machine
  Warning  Unhealthy               44s (x23 over 10m)    kubelet, vm-4-15-centos  Liveness probe failed: Get http://172.17.0.17:8012/: dial tcp 172.17.0.17:8012: connect: connection refused

  内存不足

https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#install-a-networking-layer

https://knative.dev/docs/install/yaml-install/eventing/install-eventing-with-yaml/


unable to recognize "https://github.com/knative/serving/releases/download/knative-v1.4.0/serving-core.yaml": no matches for kind "PodDisruptionBudget" in version "policy/v1"
unable to recognize "https://github.com/knative/serving/releases/download/knative-v1.4.0/serving-core.yaml": no matches for kind "PodDisruptionBudget" in version "policy/v1"


https://blog.csdn.net/cd_yourheart/article/details/107463650



# kubectl apply -f serving-core.yaml
unable to recognize "serving-core.yaml": no matches for kind "PodDisruptionBudget" in version "policy/v1"
unable to recognize "serving-core.yaml": no matches for kind "PodDisruptionBudget" in version "policy/v1"

# kubectl api-versions |grep policy
policy/v1beta1
