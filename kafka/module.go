package kafka

import (
	"github.com/ngdangkietswe/swe-notification-service/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	handler.NewRegisterUserHandler,
	handler.NewResetPasswordHandler,
	handler.NewRequestOvertimeHandler,
	handler.NewReplyOvertimeHandler,

	consumer.NewRegisterUserConsumer,
	consumer.NewResetPasswordConsumer,
	consumer.NewRequestOvertimeConsumer,
	consumer.NewReplyOvertimeConsumer,
)
