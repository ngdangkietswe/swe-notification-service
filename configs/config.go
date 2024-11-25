package configs

import (
	"github.com/spf13/viper"
	"log"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't read config file: %v", err)
		return
	}

	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
		return
	}

	GlobalConfig = config

	return
}

type Configuration struct {
	AppName            string `mapstructure:"APP_NAME"`
	KafkaBrokers       string `mapstructure:"KAFKA_BROKER"`
	KafkaConsumerGroup string `mapstructure:"KAFKA_CONSUMER_GROUP"`
	SmtpHost           string `mapstructure:"SMTP_HOST"`
	SmtpPort           string `mapstructure:"SMTP_PORT"`
	SmtpUsername       string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword       string `mapstructure:"SMTP_PASSWORD"`
}
