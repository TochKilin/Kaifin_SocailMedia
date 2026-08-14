package radis

import (
	// standard library
	"log"
	"os"

	// community package

	"github.com/joho/godotenv"

	utls "kaifin_clone_api/pkg/utls"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
	RedisExpire   int
}

// load redis from env
func InitRedis() *RedisConfig {
	// load env
	err := godotenv.Load()
	// check env file
	if err != nil {
		log.Printf("No .env file found, using system environment variables")
	}

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")
	redis_db := utls.GetenvInt("REDIS_DB", 2)
	redis_exprie := utls.GetenvInt("REDIS_EXPIRE", 60)
	// new redis object
	return &RedisConfig{
		RedisHost:     redis_host,
		RedisPort:     redis_port,
		RedisPassword: redis_password,
		RedisDB:       redis_db,
		RedisExpire:   redis_exprie,
	}
}
