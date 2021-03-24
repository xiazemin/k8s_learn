curl -sfL http://rancher-mirror.cnrancher.com/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn sh -s - server \
--datastore-endpoint="mysql://username:password@tcp(hostname:3306)/database-name"

http://docs.rancher.cn/docs/rancher2/installation_new/resources/k8s-tutorials/ha-rke/_index/


brew install rke

http://docs.rancher.cn/docs/rke/installation/_index/#%E4%BD%BF%E7%94%A8-Homebrew-%E5%AE%89%E8%A3%85-RKE

 % cat .cluster.yml
nodes:
  - address: 127.0.0.1
    user: admin
    role:
      - controlplane
      - etcd
      - worker


rke up --config  .cluster.yml

FATA[0001] Cluster must have at least one etcd plane host: please specify one or more etcd in cluster config

https://blog.csdn.net/yezi1993/article/details/106430390/

 % kubectl get pods -n kube-system
NAME                                     READY   STATUS    RESTARTS   AGE
coredns-f9fd979d6-bmftr                  1/1     Running   7          2d18h
coredns-f9fd979d6-mj7t9                  1/1     Running   7          2d18h
etcd-docker-desktop                      1/1     Running   8          2d18h

 % kubectl describe pod etcd-docker-desktop   -n kube-system
https://192.168.65.4:2380

 % docker ps | grep etcd
 5d9fd422a058

 % docker exec -it 5d9fd422a058 sh
OCI runtime exec failed: exec failed: container_linux.go:367: starting container process caused: exec: "sh": executable file not found in $PATH: unknown



