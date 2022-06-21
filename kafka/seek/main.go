package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	groupID := "test"
	topic := "test"
	offset := 2
	startTime := time.Now().Unix() - 600

	hostname, _ := os.Hostname()
	kconsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"metadata.broker.list":     "192.168.1.65:9091, 192.168.1.65:9092, 192.168.1.65:9093",
		"group.id":                 groupID,
		"client.id":                hostname,
		"enable.auto.commit":       false,
		"enable.auto.offset.store": false,
		"auto.offset.reset":        "earliest",
		"enable.partition.eof":     true,
	})
	fmt.Println(err)
	if kconsumer == nil {
		fmt.Println(kconsumer)
		return
	}

	if err := kconsumer.Subscribe(topic, nil); err != nil {
		fmt.Println(err)
	}

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

	fmt.Println(kconsumer.Assignment())
	for i := 0; i < 10; i++ {
		ev := kconsumer.Poll(int(400))
		fmt.Println(kconsumer.Assignment())
		fmt.Println(ev)
	}

	offsets, err := kconsumer.OffsetsForTimes([]kafka.TopicPartition{{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(startTime * 1000),
	}}, 400)
	fmt.Println(offsets, err)
	fmt.Println(kconsumer.Assignment())
	for i := 0; i < 10; i++ {
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
