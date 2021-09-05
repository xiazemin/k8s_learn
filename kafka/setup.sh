bin/zookeeper-server-start.sh config/zookeeper.properties

bin/kafka-server-start.sh config/server.properties

#https://www.jianshu.com/p/a743712beda5

bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

bin/kafka-topics.sh --list --zookeeper localhost:2181

bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning

bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test

go run -tags dynamic main.go

go run -tags dynamic main.go localhost:9092 test test