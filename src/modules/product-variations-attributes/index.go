package product_variations_attributes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Update(db *mongo.Database, r *gin.Engine) {
	r.PUT("/products/variations/attributes", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
		}

		var requestBody ProductVariationAttribute

		err := c.Bind(&requestBody)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		if _, err := UpdateAttribute(db, &requestBody); err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		result["status"] = http.StatusOK
		result["isError"] = false
		result["data"] = nil
		result["message"] = "Success"

		c.JSON(http.StatusOK, result)
	})
}
