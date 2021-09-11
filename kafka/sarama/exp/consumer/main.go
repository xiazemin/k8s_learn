package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

// 消费者练习

func main() {
	// 生成消费者 实例
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Print(err)
		return
	}
	// 拿到 对应主题下所有分区
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	// 遍历所有分区
	for partition := range partitionList {
		//消费者 消费 对应主题的 具体 分区 指定 主题 分区 offset  return 对应分区的对象
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Println(err)
			return
		}

		// 运行完毕记得关闭
		defer pc.AsyncClose()

		// 去出对应的 消息
		// 通过异步 拿到 消息
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
