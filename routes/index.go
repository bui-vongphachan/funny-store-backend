package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	product_variation_swap "github.com/vongphachan/funny-store-backend/routes/products/variations/swap"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {

	router := gin.Default()

	CURRENT_PROXY := os.Getenv("CURRENT_PROXY")

	router.SetTrustedProxies([]string{CURRENT_PROXY})

	product_variation_swap.SwapProductVariation(db, router)

	return router
}
