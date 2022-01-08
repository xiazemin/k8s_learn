
% kubectl get deploy ratings-v1  -o yaml

 - lastTransitionTime: "2022-01-08T06:27:55Z"
    lastUpdateTime: "2022-01-08T06:27:55Z"
    message: 'admission webhook "namespace.sidecar-injector.istio.io" denied the request:
      template: inject:1: function "Template_Version_And_Istio_Version_Mismatched_Check_Installation"
      not defined'
    reason: FailedCreate


https://stackoverflow.com/questions/69737931/istio-version-template-version-and-istio-version-mismatched


If your Bookinfo deployment is stuck in a pending state, you might see the following error:

admission webhook "sidecar-injector.istio.io" denied the request: template:
      inject:1: function "Template_Version_And_Istio_Version_Mismatched_Check_Installation"
      not defined
Your istioctl version does not match the IstioOperator version that was used during Istio installation. Ensure that you download the same version of istioctl, which is 1.12.1 in this example.

https://docs.solo.io/gloo-mesh-enterprise/main/getting_started/managed_kubernetes/


% kubectl get deploy -o wide -n istio-system
NAME                   READY   UP-TO-DATE   AVAILABLE   AGE    CONTAINERS    IMAGES                                   SELECTOR
istio-egressgateway    1/1     1            1           4h9m   istio-proxy   docker.io/querycapistio/proxyv2:1.12.1   app=istio-egressgateway,istio=egressgateway
istio-ingressgateway   1/1     1            1           4h9m   istio-proxy   docker.io/querycapistio/proxyv2:1.12.1   app=istio-ingressgateway,istio=ingressgateway
istiod                 3/3     3            3           54m    discovery     docker.io/istiojfh/pilot:1.8.1-arm64     istio=pilot


https://aiqianji.com/openoker/kubeflow-manifests/raw/master/patch/istiod.yaml

https://raw.githubusercontent.com/DickChesterwood/istio-fleetman/master/_course_files/warmup-exercise/2-istio-minikube.yaml

https://git.zero-downtime.net/ZeroDownTime/kubezero/pulls/38.patch


https://hub.docker.com/r/querycapistio/proxyv2/tags?page=1&name=1.8.1

docker pull querycapistio/proxyv2:1.8.1
1.8.1: Pulling from querycapistio/proxyv2
1.8.1: Pulling from querycapistio/proxyv2
04da93b342eb: Already exists 
b235194751de: Already exists 
606a67bb8db9: Already exists 
f8d3bd8f9548: Pull complete 
4d95f829ec51: Pull complete 
a1f6af12b3c4: Pull complete 
08aa46a139f3: Pull complete 
66f9bfc34fbc: Pull complete 
cf1fe789dbb8: Pull complete 
Digest: sha256:df20066122a5bf1cd5835465dc36eedaa7fd8507ed3275731e522ab6ef82deca
Status: Downloaded newer image for querycapistio/proxyv2:1.8.1
docker.io/querycapistio/proxyv2:1.8.1


Istio/istio-egressgateway.yaml
Istio/istio-ingressgateway.yaml

% kubectl delete -f ../istio-egressgateway.yaml 
deployment.apps "istio-egressgateway" deleted
% kubectl delete -f ../istio-ingressgateway.yaml
deployment.apps "istio-ingressgateway" deleted

   operator.istio.io/version: 1.8.1
   image: docker.io/querycapistio/proxyv2:1.8.1


% kubectl apply -f ../istio-egressgateway.yaml
deployment.apps/istio-egressgateway created
% kubectl apply -f ../istio-ingressgateway.yaml
deployment.apps/istio-ingressgateway created



% kubectl delete -f samples/bookinfo/platform/kube/bookinfo.yaml
service "details" deleted
serviceaccount "bookinfo-details" deleted
deployment.apps "details-v1" deleted
service "ratings" deleted
serviceaccount "bookinfo-ratings" deleted
deployment.apps "ratings-v1" deleted
service "reviews" deleted
serviceaccount "bookinfo-reviews" deleted
deployment.apps "reviews-v1" deleted
deployment.apps "reviews-v2" deleted
deployment.apps "reviews-v3" deleted
service "productpage" deleted
serviceaccount "bookinfo-productpage" deleted
deployment.apps "productpage-v1" deleted


 % kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
service/details created
serviceaccount/bookinfo-details created
deployment.apps/details-v1 created
service/ratings created
serviceaccount/bookinfo-ratings created
deployment.apps/ratings-v1 created
service/reviews created
serviceaccount/bookinfo-reviews created
deployment.apps/reviews-v1 created
deployment.apps/reviews-v2 created
deployment.apps/reviews-v3 created
service/productpage created
serviceaccount/bookinfo-productpage created
deployment.apps/productpage-v1 created


 % git checkout 1.8.1
Note: switching to '1.8.1'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c <new-branch-name>

Or undo this operation with:

  git switch -



% kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
service/details unchanged
serviceaccount/bookinfo-details unchanged
deployment.apps/details-v1 configured
service/ratings unchanged
serviceaccount/bookinfo-ratings unchanged
deployment.apps/ratings-v1 configured
service/reviews unchanged
serviceaccount/bookinfo-reviews unchanged
deployment.apps/reviews-v1 configured
deployment.apps/reviews-v2 configured
deployment.apps/reviews-v3 configured
service/productpage unchanged
serviceaccount/bookinfo-productpage unchanged
deployment.apps/productpage-v1 configured


 % kubectl get deploy details-v1 -o yaml  


     message: 'admission webhook "namespace.sidecar-injector.istio.io" denied the request:
      template: inject:1: function "Template_Version_And_Istio_Version_Mismatched_Check_Installation"
      not defined'
    reason: FailedCreate


% istioctl version
client version: 1.12.1
control plane version: 1.8-dev-806fb24bc121bf93ea06f6a38b7ccb3d78d1f326
data plane version: 1.8.1 (2 proxies)



 % kubectl get cm -n istio-system istio-sidecar-injector -o yaml
 https://github.com/istio/istio/issues/15152
 kubectl -n istio-system get cm istio-sidecar-injector -o jsonpath="{.data.config}"

 https://cloud.google.com/service-mesh/docs/troubleshooting/troubleshoot-webhook

 
