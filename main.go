package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vongphachan/funny-store-backend/helpers"
	"github.com/vongphachan/funny-store-backend/routes"
)

func main() {
	db := helpers.StartMongoDB()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := routes.SetupRouter(db)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Print("Server is running on port " + PORT + "...")

	r.Run(":" + PORT)
}
