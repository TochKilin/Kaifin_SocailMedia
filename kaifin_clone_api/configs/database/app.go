package config

import (
	"log"
	"os"

	// "pg/pkg/utls"

	"github.com/joho/godotenv"

	utls "kaifin_clone_api/pkg/utls"
)

type AppConfig struct {
	AppHost string
	AppPort int
}

func NewConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file %v", err)
	}
	host := os.Getenv("API_HOST")
	port := utls.GetenvInt("API_PORT", 8888)

	return &AppConfig{
		AppHost: host,
		AppPort: port,
	}
}
