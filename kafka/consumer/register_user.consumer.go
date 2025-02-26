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

type RegisterUserConsumer struct {
	handler *handler.RegisterUserHandler
}

func NewRegisterUserConsumer(handler *handler.RegisterUserHandler) *RegisterUserConsumer {
	return &RegisterUserConsumer{
		handler: handler,
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (c *RegisterUserConsumer) Consume() {
	kConsumer := consumer.NewKConsumer(constants.TopicRegisterUser)
	kConsumer.Consume(context.Background(), func(msg kafka.Message) {
		var registerUser *domain.RegisterUser
		err := json.Unmarshal(msg.Value, &registerUser)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
		c.handler.Handle(registerUser)
	})
}
