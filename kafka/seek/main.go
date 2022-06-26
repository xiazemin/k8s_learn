package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	groupID   string = "test"
	topic     string = "test"
	groupList string = "localhost:9092"

	offset    int   = 4
	startTime int64 = time.Now().Unix() - 600 - 8*3600
	count     int   = 6
)

func getConsumer() *kafka.Consumer {

	hostname, _ := os.Hostname()

	/*
	   % ./bin/zkServer.sh  restart
	   % ./bin/kafka-server-start.sh config/server.properties
	   % ./kafka-console-producer.sh --topic test --bootstrap-server localhost:9092
	   //https://blog.csdn.net/weixin_41631266/article/details/110261829

	   % ./kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 3 --topic test

	   % sh kafka-topics.sh --list --bootstrap-server localhost:9092
	   __consumer_offsets
	   test
	   test1
	   test2
	   test3

	   https://blog.csdn.net/weixin_42324471/article/details/120523733

	   % sh kafka-console-consumer.sh --topic __consumer_offsets  --bootstrap-server localhost:9092 --from-beginning


	   % sh kafka-consumer-groups.sh -bootstrap-server localhost:9092 --list
	   console-consumer-53390

	   % sh kafka-consumer-groups.sh -bootstrap-server localhost:9092 --group console-consumer-53390  --describe



	   % sh kafka-console-consumer.sh --bootstrap-server localhost:9092 --group test --topic test --from-beginning
	   1
	   2
	   3
	   4
	   5


	   n % sh kafka-consumer-groups.sh -bootstrap-server localhost:9092 --group test  --describe

	   GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                          HOST            CLIENT-ID
	   test            test            0          5               5               0               consumer-test-1-05706b82-35ff-47a5-a101-5300f870ab56 /192.168.31.231 consumer-test-1

	   % sh kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test
	   >6
	   >7
	   >8
	   >9

	   % sh kafka-consumer-groups.sh -bootstrap-server localhost:9092 --group test  --describe

	   Consumer group 'test' has no active members.

	   GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
	   test            test            0          5               9               4               -               -               -
	*/
	kconsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"metadata.broker.list":     groupList,
		"group.id":                 groupID,
		"client.id":                hostname,
		"enable.auto.commit":       false,
		"enable.auto.offset.store": false,
		"auto.offset.reset":        "earliest",
		"enable.partition.eof":     true,
		"debug":                    "msg", //  "generic, broker, topic, metadata, feature, queue, msg, protocol, cgrp, security, fetch, interceptor, plugin, consumer, admin, eos, mock, assignor, conf, all",
	})
	//https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	//https://github.com/edenhill/librdkafka/blob/master/INTRODUCTION.md
	fmt.Println(err)
	if kconsumer == nil {
		fmt.Println(kconsumer)
		return nil
	}
	return kconsumer
}

func sub(kconsumer *kafka.Consumer) {
	// if err := kconsumer.Subscribe(topic, nil); err != nil {
	// 	fmt.Println(err)
	// }

	//1，每次请求都初始化consumer没有问题
	//2，全局subscribeTopic没有问题
	//3，全局初始化consumer，请求里订阅就会有问题了

	if err := kconsumer.SubscribeTopics([]string{topic}, nil); err != nil {
		fmt.Println(err)
	}
}

func assign(kconsumer *kafka.Consumer) {

	if err := kconsumer.Assign([]kafka.TopicPartition{{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(offset),
	}}); err != nil {
		fmt.Println(err)
	}

	if err := kconsumer.Seek(kafka.TopicPartition{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(offset),
	}, 400); err != nil {
		fmt.Println(err)
	}

	as, err := kconsumer.Assignment()
	fmt.Println(as, err)
	if err := kconsumer.Assign([]kafka.TopicPartition{{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(as[0].Offset),
	}}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(kconsumer.Assignment())
}

func main() {
	kconsumer := getConsumer()
	//sub(kconsumer)
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		sub(kconsumer)
		time.Sleep(100 * time.Millisecond)
		//时序问题，导致seek在旧的结构体上进行，但是旧的结构体已经被删除了
		assign(kconsumer)
		for i := 0; i < count; i++ {
			ev := kconsumer.Poll(int(400))
			as, err := kconsumer.Assignment()
			if len(as) > 0 {
				fmt.Println(as, err, int64(as[0].Offset))
			}
			fmt.Println("\033[31m", ev, "\033[m")
			if ev != nil {
				if _, ok := ev.(*kafka.Message); ok {
					fmt.Println("\033[31moffsets&\033[m:", ev.(*kafka.Message).TopicPartition)
				}
			}
		}
		//seekByTime(kconsumer)
	})
	fmt.Println(http.ListenAndServe(":8088", nil))
}

func seekByTime(kconsumer *kafka.Consumer) {
	offsets, err := kconsumer.OffsetsForTimes([]kafka.TopicPartition{{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(startTime * 1000),
	}}, 400)
	fmt.Println(offsets, err)
	fmt.Println(kconsumer.Assignment())
	fmt.Println("---------")
	for i := 0; i < count; i++ {
		ev := kconsumer.Poll(int(400))
		if ev == nil {
			offset++
			if err := kconsumer.Assign([]kafka.TopicPartition{{
				Topic:     &topic,
				Partition: 0,
				Offset:    kafka.Offset(offset),
			}}); err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(kconsumer.Assignment())
		fmt.Println(ev)
		switch e := ev.(type) {
		case *kafka.Message:
			if err := kconsumer.Assign([]kafka.TopicPartition{e.TopicPartition}); err != nil {
				fmt.Println(err)
			}
		}
	}

	low, high, err := kconsumer.GetWatermarkOffsets(topic, 0)
	fmt.Println(low, high, err)
}

// % curl http://127.0.0.1:8088/

/*
分区没有变
[test[0]@stored] <nil> -1000
<nil>
*/
