https://kubernetes.github.io/ingress-nginx/examples/auth/client-certs/

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=nginxsvc/O=nginxsvc"

tls.crt         tls.key

kubectl create secret generic ca-secret --from-file=ca.crt=ca.crt
kubectl create secret generic tls-secret --from-file=tls.crt=server.crt --from-file=tls.key=server.key

curl -O https://kubernetes.github.io/ingress-nginx/examples/auth/client-certs/ingress.yaml

https://www.cnblogs.com/lan1x/p/5872915.html

1,生成私钥KEY
openssl genrsa -des3 -out server.key 2048
这一步执行完以后，cert目录下会生成server.key文件
2. 生成证书请求文件CSR
openssl req -new -key server.key -out server.csr
3. 生成CA的证书
前面提过X.509证书的认证者总是CA或由CA指定的人，所以得先生成一个CA的证书
openssl req -new -x509 -key server.key -out ca.crt -days 3650
4. 最后用第3步的CA证书给自己颁发一个证书玩玩
openssl x509 -req -days 3650 -in server.csr \
  -CA ca.crt -CAkey server.key \
  -CAcreateserial -out server.crt
执行完以后，cert目录下server.crt 就是我们需要的证书。当然，如果要在google等浏览器显示出安全的绿锁标志，自己颁发的证书肯定不好使，得花钱向第三方权威证书颁发机构申请(即：第4步是交给权威机构来做，我们只需要提交server.key、server.csr