package product_variations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/products/variations", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
		}

		var requestBody ProductVariation

		err := c.Bind(&requestBody)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		productVariation, err := CreateEmpty(&requestBody.ProductID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		Save(db, productVariation)

		result["data"] = productVariation
		result["status"] = http.StatusCreated
		result["isError"] = false
		result["message"] = "Completed"

		c.JSON(http.StatusOK, result)
	})

}
