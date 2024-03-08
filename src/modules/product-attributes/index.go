package product_attribute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/attribute", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		var requestBody ProductAttribute

		err := c.Bind(&requestBody)

		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		attribute, err := CreateEmpty(&requestBody.ProductID, &requestBody.AttributeGroupID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		Save(db, attribute)

		result["data"] = attribute

		c.JSON(http.StatusOK, result)
	})

}
