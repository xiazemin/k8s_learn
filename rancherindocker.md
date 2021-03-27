

https://www.codenong.com/cs106974917/

docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:v2.5.8-rc1

% docker inspect ee98e4339e2c669e02cde2b915068fa344f12db3c044a633d3546673f88da23c


https://localhost:9443/dashboard/c/local/explorer

 curl --insecure -sfL https://172.17.0.3:8443/v3/import/972gn5c6ppctx75p4rp4vs74ps2pxpdfqr2c8wzqqz9mtsvtzf8zvb_c-rkzlh.yaml | kubectl apply -f -  

 error: no objects passed to apply



 docker pull rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
 Digest: sha256:829c90bb77913d3b48bb714829a2c05ed20e1de79daa21ac19bc747d1ae7a2ff
Status: Downloaded newer image for rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
docker.io/rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64

docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64


https://localhost:9443/

 % docker inspect --format '{{ .NetworkSettings.IPAddress }}' rancher
172.17.0.2
docker logs -f rancher

https://www.jianshu.com/p/5fb3e1a998d6

172.17.0.2:9443

save url


add a cluster

Create a new Kubernetes cluster

Existing nodes

sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-head --server https://localhost:9443 --token 75g2wzfx9x8jm65sgkxn7fh7kpt5s66bkxvd82kv76767xhwdkgzwd --ca-checksum 6297744ecff40c7d107d52d2f24adfe7e3f93e91b91d63894c5afddb839fb28d --worker

Password:
Unable to find image 'rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-head' locally
docker: Error response from daemon: manifest for rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-head not found: manifest unknown: manifest unknown.
See 'docker run --help'.



docker pull rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
Digest: sha256:e4b328852406b70482cc129aa459e70ac279b3f6efc8419dc265a15aae493250
Status: Downloaded newer image for rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
docker.io/rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64



sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  rancher/rancher-agent:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64 --server https://localhost:9443 --token 75g2wzfx9x8jm65sgkxn7fh7kpt5s66bkxvd82kv76767xhwdkgzwd --ca-checksum 6297744ecff40c7d107d52d2f24adfe7e3f93e91b91d63894c5afddb839fb28d --worker

Password:
979e7e8a605c8707421bf90d64fc342dea3722eab911e9b5ad0193c321e9f8dd


 curl --insecure -sfL https://172.17.0.2:9443/v3/import/f78q6hbx5tr4gqqh4sp95km2fzxzmh6k6p28p65wgwg7w5vkdrhtbg_c-vwndp.yaml | kubectl apply -f -   


combined from similar events): Failed to create pod sandbox: rpc error: code = Unknown desc = failed to create containerd task: failed to mount rootfs component &{overlay overlay [index=off workdir=/var/lib/rancher/k3s/agent/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/3303/work

https://github.com/k3s-io/k3s/issues/19


% mkdir -p  ~/docker_volume/rancher_home/k3s

docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
-v ~/docker_volume/rancher_home/k3s:/var/lib/rancher/k3s \
--name rancher rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64


ode = Unknown desc = failed to get sandbox image "docker.io/rancher/pause:3.1": failed to pull image "docker.io/rancher/pause:3.1": failed to pull and unpack image "docker.io/rancher/pause:3.1": failed to resolve reference "docker.io/rancher/pause:3.1": failed to do request: Head "https://registry-1.docker.io/v2/rancher/pause/manifests/3.1": dial tcp: lookup registry-1.docker.io: no such host

https://stackoverflow.com/questions/61325361/k3d-failed-to-pull-image-docker-io-rancher-pause3-1


{
  "registry-mirrors": ["https://2gotd6wc.mirror.aliyuncs.com"]
}

% docker pull docker.io/rancher/pause:3.1
3.1: Pulling from rancher/pause
7675586df687: Pull complete
Digest: sha256:d22591b61e9c2b52aecbf07106d5db313c4f178e404d660b32517b18fcbf0144
Status: Downloaded newer image for rancher/pause:3.1
docker.io/rancher/pause:3.1

% docker rename rancher/pause:3.1 docker.io/rancher/pause:3.1
Error response from daemon: No such container: rancher/pause:3.1
Error: failed to rename container named rancher/pause:3.1

 % docker tag da86e6ba6ca1  docker.io/rancher/pause:3.1
% docker images |grep rancher
rancher/rancher-agent                                             master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64   069623e57568   10 hours ago    504MB
rancher/rancher                                                   master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64   6cc449849ff1   10 hours ago    981MB
rancher/pause                                                     3.1                                                           da86e6ba6ca1   3 years ago     742kB


docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
-v ~/docker_volume/rancher_home/k3s:/var/lib/rancher/k3s \
--name rancher rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
fc644600e386d2702e20b647a32d273d2a14c5f51dfa397dbe18508a1f75669c

https://github.com/rancher/k3d/issues/144
https://stackoverflow.com/questions/61325361/k3d-failed-to-pull-image-docker-io-rancher-pause3-1

 % docker exec -it fc644600e386d2702e20b647a32d273d2a14c5f51dfa397dbe18508a1f75669c sh

 sh-4.4# ls /etc/systemd/
system


# cat /etc/resolv.conf
# DNS requests are forwarded to the host. DHCP DNS options are ignored.
nameserver 192.168.65.5

 % mkdir -p docker_volume/rancher_home/etc

docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
-v ~/docker_volume/rancher_home/k3s:/var/lib/rancher/k3s \
-v ~/docker_volume/rancher_home/etc:/etc \
--name rancher rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64
baa230450dc8664994f246e0e6202775d10a078e6578809b00743eaa0a8db986

vi  ~/docker_volume/rancher_home/etc/resolv.conf
nameserver 1.1.1.1

https://www.jianshu.com/p/dbc8d9a8374e

% docker ps |grep rancher
baa230450dc8   rancher/rancher:master-baa461856167e28afea67c12667e92b3f242f467-linux-amd64   "entrypoint.sh"          17 minutes ago      Up 5 minutes       0.0.0.0:9080->80/tcp, 0.0.0.0:9443->443/tcp   rancher

 % docker exec -it baa230450dc8 sh

 #!/bin/bash
set -e

exec tini -- rancher --http-listen-port=80 --https-listen-port=443 --audit-log-path=${AUDIT_LOG_PATH} --audit-level=${AUDIT_LEVEL} --audit-log-maxage=${AUDIT_LOG_MAXAGE} --audit-log-maxbackup=${AUDIT_LOG_MAXBACKUP} --audit-log-maxsize=${AUDIT_LOG_MAXSIZE} "${@}"


#!/bin/bash
set -e

exec tini -- rancher --http-listen-port=9080 --https-listen-port=9443 --audit-log-path=${AUDIT_LOG_PATH} --audit-level=${AUDIT_LEVEL} --audit-log-maxage=${AUDIT_LOG_MAXAGE} --audit-log-maxbackup=${AUDIT_LOG_MAXBACKUP} --audit-log-maxsize=${AUDIT_LOG_MAXSIZE} "${@}""

% docker restart  baa230450dc8
baa230450dc8

 % docker images | grep "k8s.gcr.io/pause"
k8s.gcr.io/pause                                                  3.2                                                           80d28bedfe5d   13 months ago   683kB

https://blog.ilemonrain.com/docker/rancher-with-k3s.html
http://hutao.tech/k8s-source-code-analysis/prepare/debug-environment.html