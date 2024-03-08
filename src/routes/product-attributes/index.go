package routeproduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
	serviceproductattribute "github.com/vongphachan/funny-store-backend/src/services/product-attribute"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateType struct {
	ProductID        primitive.ObjectID `json:"productId"`
	AttributeGroupID primitive.ObjectID `json:"attributeGroupId"`
}

func Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/attribute", func(c *gin.Context) {
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

		attribute := serviceproductattribute.CreateEmpty(&requestBody.ProductID, &requestBody.AttributeGroupID)

		serviceproductattribute.Save(db, attribute)

		result["data"] = attribute

		c.JSON(http.StatusOK, result)
	})

}
