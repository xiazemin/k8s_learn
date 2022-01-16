https://github.com/knative/docs

https://knative.dev/docs/getting-started/

The Knative CLI (kn) provides a quick and easy interface for creating Knative resources, such as Knative Services and Event Sources, without the need to create or modify YAML files directly.

https://github.com/knative/client/releases
https://github.com/knative/client/releases/tag/knative-v1.1.0

% curl -O https://github.com/knative/client/releases/download/knative-v1.1.0/kn-darwin-amd64


% chmod +x kn-darwin-amd64 
% ./kn-darwin-amd64 version
./kn-darwin-amd64: line 1: syntax error near unexpected token `<'
./kn-darwin-amd64: line 1: `<html><body>You are being <a href="https://objects.githubusercontent.com/github-production-release-asset-2e65be/161563145/6e755cc9-ce12-41e0-96e5-e493e509909d?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20220114%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20220114T153506Z&amp;X-Amz-Expires=300&amp;X-Amz-Signature=dd5f4496d813ceb5b243a5aa7373923a96567ab7b4a177ac4123cffb562c154d&amp;X-Amz-SignedHeaders=host&amp;actor_id=0&amp;key_id=0&amp;repo_id=161563145&amp;response-content-disposition=attachment%3B%20filename%3Dkn-darwin-amd64&amp;response-content-type=application%2Foctet-stream">redirected</a>.</body></html>'

% curl -O https://github.com/knative-sandbox/kn-plugin-quickstart/releases/download/knative-v1.1.0/kn-quickstart-darwin-amd64


% chmod +x kn-quickstart-darwin-amd64 
% ./kn-quickstart-darwin-amd64 
./kn-quickstart-darwin-amd64: line 1: syntax error near unexpected token `<'
./kn-quickstart-darwin-amd64: line 1: `<html><body>You are being <a href="https://objects.githubusercontent.com/github-production-release-asset-2e65be/372742238/6c2367ab-9265-4075-b05f-20b3af963974?X-Amz-Algorithm=AWS4-HMAC-SHA256&amp;X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20220114%2Fus-east-1%2Fs3%2Faws4_request&amp;X-Amz-Date=20220114T153743Z&amp;X-Amz-Expires=300&amp;X-Amz-Signature=6fe77fc11c0b90e7e6d3a577d68ec8c5386aca58fba28c395def4be45d879ff5&amp;X-Amz-SignedHeaders=host&amp;actor_id=0&amp;key_id=0&amp;repo_id=372742238&amp;response-content-disposition=attachment%3B%20filename%3Dkn-quickstart-darwin-amd64&amp;response-content-type=application%2Foctet-stream">redirected</a>.</body></html>'


重定向了，需要单独下载
 % mv ~/Downloads/kn-quickstart-darwin-amd64 .

 %  ./kn-quickstart-darwin-amd64 version
Version:      v1.1.0
Build Date:   2021-12-14 15:56:28
Git Revision: b3fcecb

% mv ~/Downloads/kn-darwin-amd64 .
% chmod +x kn-darwin-amd64

 % ./kn-darwin-amd64 version
Version:      v1.1.0
Build Date:   2021-12-14 13:59:14
Git Revision: 530841f1
Supported APIs:
* Serving
  - serving.knative.dev/v1 (knative-serving v0.28.0)
* Eventing
  - sources.knative.dev/v1 (knative-eventing v0.28.0)
  - eventing.knative.dev/v1 (knative-eventing v0.28.0)


The quickstart plugin completes the following functions:

Checks if you have the selected Kubernetes instance installed, and creates a cluster called knative.
Installs Knative Serving with Kourier as the default networking layer, and nip.io as the DNS.
Installs Knative Eventing and creates an in-memory Broker and Channel implementation.



% ./kn-quickstart-darwin-amd64 kind                
Running Knative Quickstart using Kind
✅ Checking dependencies...
Error: creating cluster: kind version: kind version: exec: "kind": executable file not found in $PATH
Usage:
  kn-quickstart kind [flags]

Flags:
  -h, --help          help for kind
  -n, --name string   kind cluster name to be used by kn-quickstart (default knative) (default "knative")

creating cluster: kind version: kind version: exec: "kind": executable file not found in $PATH


% cp ./kn-darwin-amd64 /usr/local/bin/kn
% cp ./kn-quickstart-darwin-amd64 /usr/local/bin/kn-quickstart


https://knative.dev/docs/install/
https://knative.dev/docs/install/serving/install-serving-with-yaml/
https://knative.dev/docs/install/eventing/install-eventing-with-yaml/


https://www.cnblogs.com/mathli/p/11006397.html
https://zhuanlan.zhihu.com/p/141551023

To install the Knative Serving component:
Install the required custom resources by running the command:
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.1.0/serving-crds.yaml
Install the core components of Knative Serving by running the command:
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.1.0/serving-core.yaml


% kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.1.0/serving-crds.yaml
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


 % kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.1.0/serving-core.yaml
namespace/knative-serving created
clusterrole.rbac.authorization.k8s.io/knative-serving-aggregated-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/knative-serving-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-edit created
clusterrole.rbac.authorization.k8s.io/knative-serving-namespaced-view created
clusterrole.rbac.authorization.k8s.io/knative-serving-core created
clusterrole.rbac.authorization.k8s.io/knative-serving-podspecable-binding created
serviceaccount/controller created
clusterrole.rbac.authorization.k8s.io/knative-serving-admin created
clusterrolebinding.rbac.authorization.k8s.io/knative-serving-controller-admin created
clusterrolebinding.rbac.authorization.k8s.io/knative-serving-controller-addressable-resolver created
customresourcedefinition.apiextensions.k8s.io/images.caching.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/certificates.networking.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/configurations.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/clusterdomainclaims.networking.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/domainmappings.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/ingresses.networking.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/metrics.autoscaling.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/podautoscalers.autoscaling.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/revisions.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/routes.serving.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/serverlessservices.networking.internal.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/services.serving.knative.dev unchanged
image.caching.internal.knative.dev/queue-proxy created
configmap/config-autoscaler created
configmap/config-defaults created
configmap/config-deployment created
configmap/config-domain created
configmap/config-features created
configmap/config-gc created
configmap/config-leader-election created
configmap/config-logging created
configmap/config-network created
configmap/config-observability created
configmap/config-tracing created
horizontalpodautoscaler.autoscaling/activator created
poddisruptionbudget.policy/activator-pdb created
deployment.apps/activator created
service/activator-service created
deployment.apps/autoscaler created
service/autoscaler created
deployment.apps/controller created
service/controller created
deployment.apps/domain-mapping created
deployment.apps/domainmapping-webhook created
service/domainmapping-webhook created
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



Install a networking layer
The following commands install Istio and enable its Knative integration.
1,Install a properly configured Istio by following the Advanced Istio installation instructions or by running the command:
kubectl apply -l knative.dev/crd-install=true -f https://github.com/knative/net-istio/releases/download/knative-v1.1.0/istio.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.1.0/istio.yaml
2,Install the Knative Istio controller by running the command:
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.1.0/net-istio.yaml
3,Fetch the External IP address or CNAME by running the command:
kubectl --namespace istio-system get service istio-ingressgateway

Verify the installation
kubectl get pods -n knative-serving

 % kubectl get pods -n knative-serving
NAME                                    READY   STATUS             RESTARTS   AGE
activator-5dd67bff68-zwmbk              0/1     ImagePullBackOff   0          3m29s
autoscaler-67b5d4bfdf-zwg69             0/1     ErrImagePull       0          3m17s
controller-8474579d66-b48f7             0/1     ImagePullBackOff   0          3m16s
domain-mapping-6589cc8d65-klbdl         0/1     ImagePullBackOff   0          3m13s
domainmapping-webhook-5bb798954-txzdg   0/1     ImagePullBackOff   0          3m13s
webhook-759d957bbf-cdwdj                0/1     ImagePullBackOff   0          3m11s


% kubectl -n knative-serving get deploy
NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
activator               0/1     1            0           5m14s
autoscaler              0/1     1            0           5m2s
controller              0/1     1            0           5m1s
domain-mapping          0/1     1            0           4m59s
domainmapping-webhook   0/1     1            0           4m58s
webhook                 0/1     1            0           4m56s







Install Knative Eventing¶
To install Knative Eventing:

Install the required custom resource definitions (CRDs) by running the command:
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.1.0/eventing-crds.yaml
Install the core components of Eventing by running the command:
kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.1.0/eventing-core.yaml

Verify the installation
kubectl get pods -n knative-eventing

% kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.1.0/eventing-crds.yaml

customresourcedefinition.apiextensions.k8s.io/apiserversources.sources.knative.dev created
customresourcedefinition.apiextensions.k8s.io/brokers.eventing.knative.dev created
customresourcedefinition.apiextensions.k8s.io/channels.messaging.knative.dev created
customresourcedefinition.apiextensions.k8s.io/containersources.sources.knative.dev created
customresourcedefinition.apiextensions.k8s.io/eventtypes.eventing.knative.dev created
customresourcedefinition.apiextensions.k8s.io/parallels.flows.knative.dev created
customresourcedefinition.apiextensions.k8s.io/pingsources.sources.knative.dev created
customresourcedefinition.apiextensions.k8s.io/sequences.flows.knative.dev created
customresourcedefinition.apiextensions.k8s.io/sinkbindings.sources.knative.dev created
customresourcedefinition.apiextensions.k8s.io/subscriptions.messaging.knative.dev created
customresourcedefinition.apiextensions.k8s.io/triggers.eventing.knative.dev created




% kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v1.1.0/eventing-core.yaml
namespace/knative-eventing created
serviceaccount/eventing-controller created
clusterrolebinding.rbac.authorization.k8s.io/eventing-controller created
clusterrolebinding.rbac.authorization.k8s.io/eventing-controller-resolver created
clusterrolebinding.rbac.authorization.k8s.io/eventing-controller-source-observer created
clusterrolebinding.rbac.authorization.k8s.io/eventing-controller-sources-controller created
clusterrolebinding.rbac.authorization.k8s.io/eventing-controller-manipulator created
serviceaccount/pingsource-mt-adapter created
clusterrolebinding.rbac.authorization.k8s.io/knative-eventing-pingsource-mt-adapter created
serviceaccount/eventing-webhook created
clusterrolebinding.rbac.authorization.k8s.io/eventing-webhook created
rolebinding.rbac.authorization.k8s.io/eventing-webhook created
clusterrolebinding.rbac.authorization.k8s.io/eventing-webhook-resolver created
clusterrolebinding.rbac.authorization.k8s.io/eventing-webhook-podspecable-binding created
configmap/config-br-default-channel created
configmap/config-br-defaults created
configmap/default-ch-webhook created
configmap/config-ping-defaults created
configmap/config-features created
configmap/config-kreference-mapping created
configmap/config-leader-election created
configmap/config-logging created
configmap/config-observability created
configmap/config-tracing created
deployment.apps/eventing-controller created
deployment.apps/pingsource-mt-adapter created
horizontalpodautoscaler.autoscaling/eventing-webhook created
poddisruptionbudget.policy/eventing-webhook created
deployment.apps/eventing-webhook created
service/eventing-webhook created
customresourcedefinition.apiextensions.k8s.io/apiserversources.sources.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/brokers.eventing.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/channels.messaging.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/containersources.sources.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/eventtypes.eventing.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/parallels.flows.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/pingsources.sources.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/sequences.flows.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/sinkbindings.sources.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/subscriptions.messaging.knative.dev unchanged
customresourcedefinition.apiextensions.k8s.io/triggers.eventing.knative.dev unchanged
clusterrole.rbac.authorization.k8s.io/addressable-resolver created
clusterrole.rbac.authorization.k8s.io/service-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/serving-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/channel-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/broker-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/flows-addressable-resolver created
clusterrole.rbac.authorization.k8s.io/eventing-broker-filter created
clusterrole.rbac.authorization.k8s.io/eventing-broker-ingress created
clusterrole.rbac.authorization.k8s.io/eventing-config-reader created
clusterrole.rbac.authorization.k8s.io/channelable-manipulator created
clusterrole.rbac.authorization.k8s.io/meta-channelable-manipulator created
clusterrole.rbac.authorization.k8s.io/knative-eventing-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-messaging-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-flows-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-sources-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-bindings-namespaced-admin created
clusterrole.rbac.authorization.k8s.io/knative-eventing-namespaced-edit created
clusterrole.rbac.authorization.k8s.io/knative-eventing-namespaced-view created
clusterrole.rbac.authorization.k8s.io/knative-eventing-controller created
clusterrole.rbac.authorization.k8s.io/knative-eventing-pingsource-mt-adapter created
clusterrole.rbac.authorization.k8s.io/podspecable-binding created
clusterrole.rbac.authorization.k8s.io/builtin-podspecable-binding created
clusterrole.rbac.authorization.k8s.io/source-observer created
clusterrole.rbac.authorization.k8s.io/eventing-sources-source-observer created
clusterrole.rbac.authorization.k8s.io/knative-eventing-sources-controller created
clusterrole.rbac.authorization.k8s.io/knative-eventing-webhook created
role.rbac.authorization.k8s.io/knative-eventing-webhook created
validatingwebhookconfiguration.admissionregistration.k8s.io/config.webhook.eventing.knative.dev created
mutatingwebhookconfiguration.admissionregistration.k8s.io/webhook.eventing.knative.dev created
validatingwebhookconfiguration.admissionregistration.k8s.io/validation.webhook.eventing.knative.dev created
secret/eventing-webhook-certs created
mutatingwebhookconfiguration.admissionregistration.k8s.io/sinkbindings.webhook.sources.knative.dev created


 % kubectl get pods -n knative-eventing
NAME                                   READY   STATUS             RESTARTS   AGE
eventing-controller-5f7c968bf6-64s4f   0/1     ImagePullBackOff   0          58s
eventing-webhook-7576cf5d79-h55x8      0/1     ImagePullBackOff   0          55s



Optional: Install a default Channel (messaging) layer
Optional: Install a Broker layer:
Install the Kafka controller by running the following command:
kubectl apply -f https://github.com/knative-sandbox/eventing-kafka-broker/releases/download/knative-v1.1.0/eventing-kafka-controller.yaml

Install the Kafka Broker data plane by running the following command:
kubectl apply -f https://github.com/knative-sandbox/eventing-kafka-broker/releases/download/knative-v1.1.0/eventing-kafka-broker.yaml

Install optional Eventing extensions
kubectl apply -f https://github.com/knative-sandbox/eventing-kafka-broker/releases/download/knative-v1.1.0/eventing-kafka-controller.yaml
kubectl apply -f https://github.com/knative-sandbox/eventing-kafka-broker/releases/download/knative-v1.1.0/eventing-kafka-sink.yaml


Configure DNS¶
You can configure DNS to prevent the need to run curl commands with a host header.

The following tabs expand to show instructions for configuring DNS. Follow the procedure for the DNS of your choice
kubectl get ksvc


Install optional Serving extensions


