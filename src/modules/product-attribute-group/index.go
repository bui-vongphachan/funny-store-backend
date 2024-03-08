package product_attribute_group

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateType struct {
	ProductID primitive.ObjectID `json:"productId"`
}

func API_Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/attribute-group", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		var requestBody AttributeGroup

		err := c.Bind(&requestBody)

		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		attributeGroup, err := CreateEmpty(&requestBody.ProductID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		Save(db, attributeGroup)

		result["data"] = attributeGroup

		c.JSON(http.StatusOK, result)
	})

}

func API_Pagination(db *mongo.Database, r *gin.Engine) {
	r.GET("/product/attribute-group", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		pipeline := bson.A{}

		pipeline = *MakeMatchPaginationPipeline(c.Request.URL.Query(), &pipeline)

		pipeline = *utils.MakeSkipOffsetPipeLine(c.Request.URL.Query(), &pipeline)

		cursor, err := db.Collection(CollectionName).Aggregate(context.TODO(), pipeline)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		// Decode results
		var results []AttributeGroup
		if err := cursor.All(context.TODO(), &results); err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		c.JSON(http.StatusOK, result)
	})
}

func API_Update(db *mongo.Database, r *gin.Engine) {
	r.PATCH("/product/attribute-group/:id", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		attributeGroupId := c.Param("id")

		attributeGroup, err := FindById(db, &attributeGroupId)
		if err != nil || attributeGroup == nil {
			result["status"] = 404
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		var requestBody AttributeGroup
		if err := c.Bind(&requestBody); err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
		}

		newData, err := BindNewData(&requestBody, attributeGroup)
		if err != nil {
			result["status"] = 403
			result["message"] = "Invalid data"
			c.JSON(http.StatusOK, result)
			return
		}

		filter := bson.M{"_id": attributeGroup.ID}
		UpdateOne(db, &filter, newData)

		result["status"] = 200
		result["isError"] = false
		result["message"] = "Updated"
		result["data"] = newData

		c.JSON(http.StatusOK, result)
	})

}
