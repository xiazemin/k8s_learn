% helm repo add stable https://kubernetes-charts.storage.googleapis.com/
Error: repo "https://kubernetes-charts.storage.googleapis.com/" is no longer available; try "https://charts.helm.sh/stable" instead

 % helm repo add bitnami https://charts.bitnami.com/bitnami
"bitnami" has been added to your repositories

 % helm repo add incubator https://kubernetes-charts-incubator.storage.googleapis.com/
Error: repo "https://kubernetes-charts-incubator.storage.googleapis.com/" is no longer available; try "https://charts.helm.sh/incubator" instead

helm repo update # Make sure we get the latest list of charts

% helm repo list
NAME    URL                               
bitnami https://charts.bitnami.com/bitnami

% helm search repo stable
No results found

$ helm install stable/mysql --generate-name