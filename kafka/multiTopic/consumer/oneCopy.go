package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerCopy interface {
	Listen()
}

type consumerCopy struct {
	Topic string
	*kafka.Consumer
}

func NewConsumerCopy(topic, offset string) ConsumerCopy {
	return &consumerCopy{
		Topic:    topic,
		Consumer: NewKafkaConsumer(offset),
	}
}

func (m *consumerCopy) Listen() {
	// if err := m.Consumer.Subscribe(m.Topic, nil); err != nil {
	// 	panic(err)
	// }

	if err := m.Consumer.SubscribeTopics([]string{m.Topic}, nil); err != nil {
		panic(err)
	}
	for true {
		event := m.Consumer.Poll(1000)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			fmt.Printf("topic:%#v", *(e.TopicPartition.Topic))
		default:
			fmt.Println(e)
		}
	}
}
