package product_variation_swap

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SwapProductVariation(db *mongo.Database, r *gin.Engine) {
	r.PATCH("/product/variation/swap", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		c.JSON(http.StatusOK, result)
	})
}
