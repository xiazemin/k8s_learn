When we first added k8s inside Docker Desktop, we added a rule to promote all service accounts to be cluster admin. It helps people who install helm to start easily and to forget security. Maybe it's time to remove it (or at least make it optional).

Can you try to delete the ClusterRoleBinding named docker-for-desktop-binding and see if it works ?

https://github.com/docker/for-mac/issues/3694

