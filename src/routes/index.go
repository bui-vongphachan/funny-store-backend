package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	routeproduct "github.com/vongphachan/funny-store-backend/src/routes/products"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {

	router := gin.Default()

	CURRENT_PROXY := os.Getenv("CURRENT_PROXY")

	router.SetTrustedProxies([]string{CURRENT_PROXY})

	routeproduct.CreateDraft(db, router)

	return router
}
