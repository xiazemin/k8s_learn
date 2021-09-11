bin/zookeeper-server-start.sh config/zookeeper.properties

bin/kafka-server-start.sh config/server.properties

#https://www.jianshu.com/p/a743712beda5

bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 2 --topic test

bin/kafka-topics.sh --list --zookeeper localhost:2181

bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning

bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test

go run -tags dynamic main.go

go run -tags dynamic main.go localhost:9092 test test

 % bin/kafka-topics.sh --delete --topic test  --zookeeper localhost:2181
 % ps aux |grep kafka |grep -v grep|awk '{print $2}' |xargs kill -9
  % ps aux |grep zookeeper  |grep -v grep|awk '{print $2}' |xargs kill -9

 rm  /tmp/zookeeper/version-2/*
 rm /tmp/kafka-logs/*

.查看集群是否有broker没有运行
./bin/zookeeper-shell.sh 127.0.0.1:2181

ls /brokers/ids
[0]

