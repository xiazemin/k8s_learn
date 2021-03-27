https://www.codenong.com/cs106974917/

docker run -d --restart=unless-stopped --privileged=true -p 9443:443 -p 9080:80 \
-v ~/docker_volume/rancher_home/rancher:/var/lib/rancher \
-v ~/docker_volume/rancher_home/auditlog:/var/log/auditlog \
--name rancher rancher/rancher:v2.5.8-rc1

% docker inspect ee98e4339e2c669e02cde2b915068fa344f12db3c044a633d3546673f88da23c


https://localhost:9443/dashboard/c/local/explorer

 curl --insecure -sfL https://172.17.0.3:8443/v3/import/972gn5c6ppctx75p4rp4vs74ps2pxpdfqr2c8wzqqz9mtsvtzf8zvb_c-rkzlh.yaml | kubectl apply -f -  

 error: no objects passed to apply