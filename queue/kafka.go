package queue

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaService struct {
	Producer *Producer
	Consumer *Consumer
}

func NewKafkaService(topic string) *KafkaService {
	_, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, 0)
	if err != nil {
		panic(err.Error())
	}

	c := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       topic,
		GroupID:     "image-group",
		Partition:   0,
		StartOffset: kafka.LastOffset,
	},
	)

	p := &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: topic,
	}

	return &KafkaService{
		Consumer: NewConsumer(c),
		Producer: NewProducer(p),
	}
}
