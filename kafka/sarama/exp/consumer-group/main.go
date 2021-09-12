package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

type consumerGroupHandler struct {
	name string
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		fmt.Println("ERROR", err)
	}
}

func consume(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) {
	fmt.Println(name + "start")
	wg.Done()
	ctx := context.Background()
	for {
		topics := []string{"test"}
		handler := consumerGroupHandler{name: name}
		err := (*group).Consume(ctx, topics, handler)
		fmt.Println("consume group end")
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	group1, err := sarama.NewConsumerGroupFromClient("c1", client)
	if err != nil {
		panic(err)
	}
	group2, err := sarama.NewConsumerGroupFromClient("c2", client)
	if err != nil {
		panic(err)
	}
	group3, err := sarama.NewConsumerGroupFromClient("c3", client)
	if err != nil {
		panic(err)
	}
	defer group1.Close()
	defer group2.Close()
	defer group3.Close()
	wg.Add(3)
	go consume(&group1, &wg, "c1")
	go consume(&group2, &wg, "c2")
	go consume(&group3, &wg, "c3")
	wg.Wait()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}
