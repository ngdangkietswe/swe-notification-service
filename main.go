package main

import (
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-notification-service/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/kafka/handler"
)

func main() {
	// Initialize the config module
	config.Init()
	// RegisterUserConsumer consumes messages from the kafka topic "auth.register_user.v1"
	registerUserHandler := handler.NewRegisterUserHandler()
	registerUserConsumer := consumer.NewRegisterUserConsumer(registerUserHandler)
	registerUserConsumer.Consume()
}
