bin/zkServer.sh --config conf start-foreground
bin/kafka-server-start.sh config/server.properties
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 2 --topic test
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 2 --topic test1
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test1


go get -tags dynamic github.com/confluentinc/confluent-kafka-go/kafka 
go run  -tags dynamic main.go multiConsumer 
