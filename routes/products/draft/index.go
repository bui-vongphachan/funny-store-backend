package product_draft

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createdefaultproduct "github.com/vongphachan/funny-store-backend/routes/products/draft/create-default-product"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/draft", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		product := createdefaultproduct.Main()

		result["data"] = product

		c.JSON(http.StatusOK, result)
	})

}
