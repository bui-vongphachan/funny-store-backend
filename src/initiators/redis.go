package initiators

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func NewRedis(chanel int) *redis.Client {

	log.Println("Starting Redis...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	REDIS_CONNECTION_STRING := os.Getenv("REDIS_CONNECTION_STRING")

	opt, err := redis.ParseURL(REDIS_CONNECTION_STRING)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return client
}
