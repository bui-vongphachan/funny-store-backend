package routeproduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serviceproductattributegroup "github.com/vongphachan/funny-store-backend/src/services/product-attribute-group"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateType struct {
	ProductID primitive.ObjectID `json:"productId"`
}

func Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/attribute-group", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		var requestBody CreateType

		err := c.Bind(&requestBody)

		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		attributeGroup := serviceproductattributegroup.CreateEmpty(&requestBody.ProductID)

		serviceproductattributegroup.Save(db, attributeGroup)

		result["data"] = attributeGroup

		c.JSON(http.StatusOK, result)
	})

}
