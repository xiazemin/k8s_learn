package consumer

import (
	"multiTopic/config"
	"os"
	"strings"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaConsumerOnce sync.Once
var kafkaConsumer *kafka.Consumer

func NewKafkaConsumerSingle() *kafka.Consumer {
	kafkaConsumerOnce.Do(func() {
		hostname, _ := os.Hostname()
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"metadata.broker.list": strings.Join(config.BrokerList, ","),
			"group.id":             config.Group,
			"client.id":            hostname,
		})
		if err != nil {
			panic(err)
		}
		kafkaConsumer = consumer
	})

	return kafkaConsumer
}

func NewKafkaConsumer(offset string) *kafka.Consumer {
	hostname, _ := os.Hostname()
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"metadata.broker.list":            strings.Join(config.BrokerList, ","),
		"group.id":                        config.Group,
		"client.id":                       hostname,
		"auto.offset.reset":               offset, // "latest",
		"enable.auto.commit":              true,
		"enable.auto.offset.store":        false,
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		panic(err)
	}

	return consumer
}
