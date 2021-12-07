https://www.cnblogs.com/yinzhengjie/p/10514206.html

https://www.cnblogs.com/huxi2b/p/6061110.html

offset 的单位是消息的条数，所以lag的单位也是消息的条数，证明


 ~ % cd software/apache-zookeeper-3.6.2-bin
 ./bin/zkServer.sh start-foreground

cd software/kafka_2.12-2.7.0
./bin/kafka-server-start.sh config/server.properties


 % bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning



 % bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
>1
>3
>test
>^@^@
>
>123456
>123456789011121314
abcdefg                                                ^ >^@hi
>abcedfhskgskhbdjbjjkjsjgskjbjbjdbjjdjbjkdkjjkdfjkhjkdjhgjbhhgfvgjkjgksjkjghdjkbkjgjkjdnbnjdjkbjdbndjbndkjbndkgkjgjgjkjsgkj


 % ./bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list  localhost:9092 -topic  test --time -2
test:0:0

% ./bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list  localhost:9092 -topic  test --time -1
test:0:13
% ./bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list  localhost:9092 -topic  test --time -1
test:0:14
% ./bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list  localhost:9092 -topic  test --time -1
test:0:15
% ./bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list  localhost:9092 -topic  test --time -1
test:0:16



% bin/kafka-console-consumer.sh --topic __consumer_offsets --bootstrap-server localhost:9092 --formatter "kafka.coordinator.group.GroupMetadataManager\$OffsetsMessageFormatter" --consumer.config config/consumer.properties  --from-beginning


