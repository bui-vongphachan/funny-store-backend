package product_draft

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

		c.JSON(http.StatusOK, result)
	})

}
