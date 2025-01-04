package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-notification-service/kafka"
	"github.com/ngdangkietswe/swe-notification-service/kafka/consumer"
	"go.uber.org/fx"
)

func main() {
	// Initialize the config module
	config.Init()

	app := fx.New(
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
