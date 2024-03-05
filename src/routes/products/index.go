package routeproduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serviceproduct "github.com/vongphachan/funny-store-backend/src/services/product"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDraft(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/draft", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		product := serviceproduct.CreateEmpty()

		result["data"] = product

		c.JSON(http.StatusOK, result)
	})

}
