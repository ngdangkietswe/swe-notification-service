package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-notification-service/data/repository"
	"github.com/ngdangkietswe/swe-notification-service/kafka"
	"github.com/ngdangkietswe/swe-notification-service/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/logger"
	"go.uber.org/fx"
)

func main() {
	// Initialize the config module
	config.Init()

	app := fx.New(
		logger.Module,
		repository.Module,
		kafka.Module,
		fx.Invoke(func(c *consumer.CdcAuthUsersConsumer) {
			go c.Consume()
		}),
		fx.Invoke(func(c *consumer.RegisterUserConsumer) {
			go c.Consume()
		}),
		fx.Invoke(func(c *consumer.ResetPasswordConsumer) {
			go c.Consume()
		}),
		fx.Invoke(func(c *consumer.RequestOvertimeConsumer) {
			go c.Consume()
		}),
		fx.Invoke(func(c *consumer.ReplyOvertimeConsumer) {
			go c.Consume()
		}),
	)

	app.Run()
}
