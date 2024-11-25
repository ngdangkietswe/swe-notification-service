package main

import (
	"context"
	"encoding/json"
	"github.com/ngdangkietswe/swe-notification-service/kafka/constant"
	"github.com/ngdangkietswe/swe-notification-service/kafka/consumer"
	"github.com/ngdangkietswe/swe-notification-service/kafka/model"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	kConsumer := consumer.NewKConsumer(constant.TopicRegisterUser)
	kConsumer.Consume(context.Background(), func(msg kafka.Message) {
		var registerUser *model.RegisterUser
		err := json.Unmarshal(msg.Value, &registerUser)
		if err != nil {
			log.Printf("error while unmarshalling message: %v", err)
		}
	})
}
