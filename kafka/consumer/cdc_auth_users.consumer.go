package consumer

import (
	"context"
	"encoding/json"
	kafkacommon "github.com/ngdangkietswe/swe-go-common-shared/kafka"
	"github.com/ngdangkietswe/swe-go-common-shared/kafka/consumer"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/ngdangkietswe/swe-notification-service/constants"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type CdcAuthUsersConsumer struct {
	logger  *logger.Logger
	handler *handler.CdcAuthUsersHandler
}

func NewCdcAuthUsersConsumer(logger *logger.Logger, handler *handler.CdcAuthUsersHandler) *CdcAuthUsersConsumer {
	return &CdcAuthUsersConsumer{
		logger:  logger,
		handler: handler,
	}
}

// Consume continuously listens for messages from the Kafka topic.
func (c *CdcAuthUsersConsumer) Consume() {
	kConsumer := consumer.NewKConsumer(constants.TopicCdcAuthUsers)
	kConsumer.Consume(context.Background(), func(msg kafka.Message) {
		var cdcEventMsg *kafkacommon.CdcEventMsg
		err := json.Unmarshal(msg.Value, &cdcEventMsg)
		if err != nil {
			c.logger.Error("Error while unmarshalling message", zap.String("topic", constants.TopicCdcAuthUsers), zap.Error(err))
		}
		c.handler.Handle(cdcEventMsg)
	})
}
