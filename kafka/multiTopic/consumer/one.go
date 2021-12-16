package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer interface {
	Listen()
}

type consumer struct {
	Topic string
	*kafka.Consumer
}

func NewConsumer(topic string) Consumer {
	return &consumer{
		Topic:    topic,
		Consumer: NewKafkaConsumerSingle(),
	}
}

func (m *consumer) Listen() {
	if err := m.Consumer.Subscribe(m.Topic, nil); err != nil {
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
