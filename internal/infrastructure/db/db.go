package persistence

import (
	"fmt"

	configuration "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/infrastructure/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(config configuration.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName, config.DBSSL)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\u2714 Connected to database%s\n", "\033[32m", "\033[0m")

	return db, nil
}
