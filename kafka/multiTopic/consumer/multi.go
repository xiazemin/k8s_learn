package consumer

import (
	"fmt"
	"time"

	"multiTopic/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MultiConsumer interface {
	Listen()
}

type multiConsumer struct {
	Topics []string
	*kafka.Consumer
	batchPool []*kafka.Message
	batchMap  map[string]struct{}
	batchSize int
	timeout   time.Duration
}

func NewMultiConsumer() MultiConsumer {
	_batchSize := 10
	return &multiConsumer{
		Topics:    []string{config.TopicTest, config.TopicTest1},
		Consumer:  NewKafkaConsumerSingle(),
		batchSize: _batchSize,
		timeout:   10 * time.Second,
		batchPool: make([]*kafka.Message, 0, _batchSize),
		batchMap:  make(map[string]struct{}, _batchSize),
	}
}

func (m *multiConsumer) Listen() {
	if err := m.Consumer.SubscribeTopics(m.Topics, nil); err != nil {
		panic(err)
	}
	for true {
		event := m.Consumer.Poll(1000)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			fmt.Printf("len of pool %d \n", len(m.batchPool))
			if _, ok := m.batchMap[string(e.Value)]; !ok {
				m.batchPool = append(m.batchPool, e)
				m.batchMap[string(e.Value)] = struct{}{}
			}
			if len(m.batchPool) >= m.batchSize || (len(m.batchPool) > 0 && time.Since(m.batchPool[0].Timestamp) > m.timeout) {
				for _, msg := range m.batchPool {
					fmt.Printf("batch topic:%#v", *(msg.TopicPartition.Topic))
				}
				m.batchPool = m.batchPool[0:0]
				m.batchMap = make(map[string]struct{}, m.batchSize)
			}
		default:
			fmt.Println(e)
		}
	}
}
