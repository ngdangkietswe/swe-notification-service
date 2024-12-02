package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	env := "local"
	if e := os.Getenv("ENV"); e != "" {
		env = e
	}

	if strings.ToLower(env) == "prod" {
		viper.AutomaticEnv()
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Can't read config file: %v", err)
			return
		}
	}

	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
	}

	log.Printf(
		"Application is running in %s environment\n"+
			"Config loaded: %v", env, *GlobalConfig)
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
