package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// 构建 生产者
	// 生成 生产者配置文件
	config := sarama.NewConfig()
	// 设置生产者 消息 回复等级 0 1 all
	config.Producer.RequiredAcks = sarama.NoResponse //sarama.WaitForAll
	//kafka server: Replication-factor is invalid.
	// 设置生产者 成功 发送消息 将在什么 通道返回
	config.Producer.Return.Successes = true
	// 设置生产者 发送的分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 构建 消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("123")

	// 连接 kafka
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Print(err)
		return
	}
	defer producer.Close()
	// 发送消息
	message, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(message, " ", offset)

}
