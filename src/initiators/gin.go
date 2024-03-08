package initiators

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartGin(db *mongo.Database) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := setupRoutes(db)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Print("Server is running on port " + PORT + "...")

	r.Run("localhost:" + PORT)
}

func setupRoutes(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	CURRENT_PROXY := os.Getenv("CURRENT_PROXY")

	router.SetTrustedProxies([]string{CURRENT_PROXY})

	product_attribute_group.API_Create(db, router)
	product_attribute_group.API_Pagination(db, router)
	product_attribute_group.API_Update(db, router)

	return router
}