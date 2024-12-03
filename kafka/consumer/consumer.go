package consumer

import (
	"context"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

type KConsumer struct {
	Reader *kafka.Reader
}

func NewKConsumer(topic string) *KConsumer {
	return &KConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  strings.Split(config.GetString("KAFKA_BROKERS", "localhost:9092"), ","),
			Topic:    topic,
			GroupID:  config.GetString("KAFKA_CONSUMER_GROUP", "swe-consumer-group"),
			MinBytes: 10e3,                   // 10KB
			MaxBytes: 10e6,                   // 10MB
			MaxWait:  500 * time.Millisecond, // 500ms
		}),
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (k *KConsumer) Consume(ctx context.Context, handler func(msg kafka.Message)) {
	log.Printf("Starting consumer for topic: %s with groupID: %s, brokers: %s",
		k.Reader.Config().Topic, k.Reader.Config().GroupID, k.Reader.Config().Brokers)

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
