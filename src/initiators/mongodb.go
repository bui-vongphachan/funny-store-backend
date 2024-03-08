package initiators

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

	fmt.Println("Starting MongoDB...")

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

	MONGODB_DBNAME := os.Getenv("MONGODB_DBNAME")

	if MONGODB_DBNAME == "" {
		log.Fatalln("MONGODB_DBNAME is not set")
		panic("MONGODB_DBNAME is not set")
	}

	database := client.Database(MONGODB_DBNAME)

	if database == nil {
		log.Fatalln("Database is not found")
		panic("Database is not found")
	}

	fmt.Println("MongoDB is connected")

	return database
}
