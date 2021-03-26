% kubectl delete -f cert-manager.yaml
customresourcedefinition.apiextensions.k8s.io "certificaterequests.cert-manager.io" deleted
customresourcedefinition.apiextensions.k8s.io "certificates.cert-manager.io" deleted
customresourcedefinition.apiextensions.k8s.io "challenges.acme.cert-manager.io" deleted
customresourcedefinition.apiextensions.k8s.io "clusterissuers.cert-manager.io" deleted
customresourcedefinition.apiextensions.k8s.io "issuers.cert-manager.io" deleted
customresourcedefinition.apiextensions.k8s.io "orders.acme.cert-manager.io" deleted

 %  kubectl get pods --namespace cert-manager
No resources found in cert-manager namespace.

% helm list
NAME	NAMESPACE	REVISION	UPDATED	STATUS	CHART	APP VERSION

 % helm list  -n cattle-system
NAME           	NAMESPACE    	REVISION	UPDATED                                	STATUS  	CHART                        	APP VERSION
rancher-k8s    	cattle-system	1       	2021-03-24 17:56:49.843259 +0800 CST   	deployed	rancher-2.5.7                	v2.5.7
rancher-webhook	cattle-system	1       	2021-03-24 10:02:51.547798215 +0000 UTC	deployed	rancher-webhook-0.1.0-beta901	0.1.0-beta9

