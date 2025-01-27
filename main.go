package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
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
		kafka.Module,
		fx.Invoke(func(c *consumer.RegisterUserConsumer) {
			go c.Consume()
		}),
		fx.Invoke(func(c *consumer.ResetPasswordConsumer) {
			go c.Consume()
		}),
	)

	app.Run()
}
