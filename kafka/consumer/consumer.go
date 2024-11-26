package consumer

import (
	"context"
	"github.com/ngdangkietswe/swe-notification-service/configs"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KConsumer struct {
	Reader *kafka.Reader
}

func NewKConsumer(topic string) *KConsumer {
	return &KConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{configs.GlobalConfig.KafkaBrokers},
			Topic:    topic,
			GroupID:  configs.GlobalConfig.KafkaConsumerGroup,
			MinBytes: 10e3,                   // 10KB
			MaxBytes: 10e6,                   // 10MB
			MaxWait:  500 * time.Millisecond, // 500ms
		}),
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (k *KConsumer) Consume(ctx context.Context, handler func(msg kafka.Message)) {
	log.Printf("Starting consumer for topic: %s", k.Reader.Config().Topic)

	for {
		msg, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error while consuming message: %v", err)
			break
		}
		log.Printf("Received message: topic=%s partition=%d offset=%d key=%s value=%s",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		// Call the handler to process the message
		handler(msg)
	}
}
