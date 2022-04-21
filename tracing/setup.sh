% docker pull jaegertracing/all-in-one

% docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
7737eeb4718fb172bec0d69f4cf4da44b2a2bc4214ad77f016508ae616cb9dab

#http://localhost:16686

#https://juejin.cn/post/6844903942309019661
#https://www.jaegertracing.io/docs/1.33/getting-started/
#https://blog.csdn.net/sniperking2008/article/details/103762543
#https://www.jianshu.com/p/ffc597bb4ce8
#https://www.jianshu.com/p/b5cd7b07a24e
#https://github.com/grpc-ecosystem/go-grpc-middleware

#https://zhuanlan.zhihu.com/p/427261147