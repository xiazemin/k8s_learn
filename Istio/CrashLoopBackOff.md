Insufficient resources—lack of resources prevents the container from loading
Locked file—a file was already locked by another container
Locked database—the database is being used and locked by other pods
Failed reference—reference to scripts or binaries that are not present on the container
Setup error—an issue with the init-container setup in Kubernetes
Config loading error—a server cannot load the configuration file
Misconfigurations—a general file system misconfiguration
Connection issues—DNS or kube-DNS is not able to connect to a third-party service
Deploying failed services—an attempt to deploy services/applications that have already failed (e.g. due to a lack of access to other services)


https://komodor.com/learn/how-to-fix-crashloopbackoff-kubernetes-error/


https://istio.io/latest/docs/ops/configuration/mesh/app-health-check/

% kubectl get all -n istio-io-health
No resources found in istio-io-health namespace.



 % kubectl label namespace default istio-injection=enabled
error: 'istio-injection' already has a value (enabled), and --overwrite is false


% kubectl apply -f ../istio/samples/bookinfo/platform/kube/bookinfo.yaml
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




https://istio.io/latest/blog/2019/data-plane-setup/


