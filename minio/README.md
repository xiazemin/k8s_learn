docker pull minio/minio

#docker.io/minio/minio:latest

https://www.jianshu.com/p/de835549ae96


https://www.jianshu.com/p/de835549ae96

docker run -p 9000:9000 --name my_minio \
  -e "MINIO_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_SECRET_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  -v  ~/Downloads:/data \
  -v  ~/Downloads:/root/.minio \
  minio/minio server /data

https://www.cnblogs.com/webenh/p/13334291.html

docker: Error response from daemon: Conflict. The container name "/my_minio" is already in use by container "37f2ed7b17ac1d1bd2ef73d519a157f5dfa7b28c0bebc868b92e6ba1892f2416". You have to remove (or rename) that container to be able to reuse that name.
See 'docker run --help'.


https://blog.csdn.net/weixin_41945228/article/details/92831173

 docker ps -a |grep my_minio


docker ps -a |grep myminio
558e85c2ec04   minio/minio              "/usr/bin/docker-ent…"   8 minutes ago       

 % docker rm 558e85c2ec04

 docker run -it -p 9000:9000 minio/minio server /data

 https://www.jianshu.com/p/52dbc679094a

 https://docs.min.io/



 docker run \
  -p 9000:9000 \
  -p 9001:9001 \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  minio/minio server /data --console-address ":9001"

  