package consumer

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-go-common-shared/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/constants"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
	"github.com/ngdangkietswe/swe-notification-service/kafka/payload"
	kafkago "github.com/segmentio/kafka-go"
	"log"
)

type RequestOvertimeConsumer struct {
	handler *handler.RequestOvertimeHandler
}

func NewRequestOvertimeConsumer(handler *handler.RequestOvertimeHandler) *RequestOvertimeConsumer {
	return &RequestOvertimeConsumer{
		handler: handler,
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (c *RequestOvertimeConsumer) Consume() {
	kConsumer := consumer.NewKConsumer(constants.TopicSendEmailRequestOvertime)
	kConsumer.Consume(context.Background(), func(msg kafkago.Message) {
		var requestOvertimePayload *payload.RequestOvertimePayload
		err := json.Unmarshal(msg.Value, &requestOvertimePayload)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
		c.handler.Handle(requestOvertimePayload)
	})
}
