package consumer

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-go-common-shared/constants"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-go-common-shared/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
	"github.com/segmentio/kafka-go"
	"log"
)

type ResetPasswordConsumer struct {
	handler *handler.ResetPasswordHandler
}

func NewResetPasswordConsumer(handler *handler.ResetPasswordHandler) *ResetPasswordConsumer {
	return &ResetPasswordConsumer{
		handler: handler,
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (c *ResetPasswordConsumer) Consume() {
	kConsumer := consumer.NewKConsumer(constants.TopicResetPassword)
	kConsumer.Consume(context.Background(), func(msg kafka.Message) {
		var resetPassword *domain.ResetPassword
		err := json.Unmarshal(msg.Value, &resetPassword)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
		c.handler.Handle(resetPassword)
	})
}
