package datasource

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-notification-service/data/ent"
	"log"
)

func NewEntClient() *ent.Client {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s",
			config.GetString("DB_HOST", "localhost"),
			config.GetInt("DB_PORT", 5432),
			config.GetString("DB_USER", "postgres"),
			config.GetString("DB_NAME", "SweNotification"),
			config.GetString("DB_PASSWORD", "123456"),
			config.GetString("DB_SSL_MODE", "disable"),
			config.GetString("DB_SEARCH_PATH", "swenotification"),
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
