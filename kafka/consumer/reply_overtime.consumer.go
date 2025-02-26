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

type ReplyOvertimeConsumer struct {
	handler *handler.ReplyOvertimeHandler
}

func NewReplyOvertimeConsumer(handler *handler.ReplyOvertimeHandler) *ReplyOvertimeConsumer {
	return &ReplyOvertimeConsumer{
		handler: handler,
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (c *ReplyOvertimeConsumer) Consume() {
	kConsumer := consumer.NewKConsumer(constants.TopicSendEmailReplyOvertime)
	kConsumer.Consume(context.Background(), func(msg kafkago.Message) {
		var replyOvertime *payload.ReplyOvertimePayload
		err := json.Unmarshal(msg.Value, &replyOvertime)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
		c.handler.Handle(replyOvertime)
	})
}
