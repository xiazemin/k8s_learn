https://cloud.tencent.com/developer/article/1823421

http://www.catddm.com/archives/773

https://github.com/Shopify/sarama

https://blog.csdn.net/justlpf/article/details/107400534

https://www.cnblogs.com/zhaoxianxin/p/13432068.html
https://blog.csdn.net/luslin1711/article/details/105798571/


2021/09/11 22:21:54 kafka server: Replication-factor is invalid.
warn [producer clientid=console-producer] error while fetching metadata with correlation id 140 : {test=
warn [controller id=0, targetbrokerid=0] connection to node 0 (/202.112.238.170:9092) could not be established. broker may not be available.

 % vi config/server.properties
 listeners=PLAINTEXT://:9092
 advertised.listeners=PLAINTEXT://localhost:9092