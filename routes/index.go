package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	product_draft "github.com/vongphachan/funny-store-backend/routes/products/draft"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {

	router := gin.Default()

	CURRENT_PROXY := os.Getenv("CURRENT_PROXY")

	router.SetTrustedProxies([]string{CURRENT_PROXY})

	product_draft.Create(db, router)

	return router
}
