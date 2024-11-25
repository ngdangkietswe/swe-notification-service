package consumer

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-notification-service/kafka/constant"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
	"github.com/ngdangkietswe/swe-notification-service/kafka/model"
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
	kConsumer := NewKConsumer(constant.TopicRegisterUser)
	kConsumer.Consume(context.Background(), func(msg kafka.Message) {
		var registerUser *model.RegisterUser
		err := json.Unmarshal(msg.Value, &registerUser)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
		c.handler.Handle(registerUser)
	})
}
