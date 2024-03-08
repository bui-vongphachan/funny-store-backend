package routeproductattributegroup

import (
	"net/http"

	"github.com/gin-gonic/gin"
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
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

func Update(db *mongo.Database, r *gin.Engine) {
	r.PATCH("/product/attribute-group/:id", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		attributeGroupId := c.Param("id")

		attributeGroup := serviceproductattributegroup.FindById(db, &attributeGroupId)
		if attributeGroup == nil {
			result["status"] = 404
			result["message"] = "Data not found"
			c.JSON(http.StatusOK, result)
			return
		}

		var requestBody modelproduct.AttributeGroup

		err := c.Bind(&requestBody)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		newData, err := serviceproductattributegroup.BindNewData(&requestBody, attributeGroup)
		if err != nil {
			result["status"] = 403
			result["message"] = "Invalid data"
			c.JSON(http.StatusOK, result)
			return
		}

		// update := bson.M{"$set": bson.M{"avg_rating": 4}}

		// context := context.TODO()

		// db.Collection(collectionname.PRODUCT_ATTRIBUTE_GROUPS).UpdateOne(context, filter, bson.M{"$set": update})

		result["status"] = 200
		result["isError"] = false
		result["message"] = "Updated"
		result["data"] = newData

		c.JSON(http.StatusOK, result)
	})

}
