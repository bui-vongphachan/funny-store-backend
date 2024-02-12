package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func StartMongoDB() *mongo.Database {

	fmt.Println("Hello, World!")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	MONGO_URI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return client.Database("testing")
}
