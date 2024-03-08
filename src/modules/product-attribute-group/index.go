package product_attribute_group

import (
	"context"
	"log"
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

		pipelines := mongo.Pipeline{}

		matchStage := MakeMatchPaginationPipeline(c.Request.URL.Query())

		if matchStage != nil {
			pipelines = append(pipelines, *matchStage)
		}

		skipStage := utils.MakeSkipStage(c.Request.URL.Query())
		pipelines = append(pipelines, *skipStage)

		limitStage := utils.MakeLimitStage(c.Request.URL.Query())
		pipelines = append(pipelines, *limitStage)

		cursor, err := db.Collection(CollectionName).Aggregate(context.TODO(), pipelines)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		items := []bson.M{}
		if err := cursor.All(context.TODO(), &items); err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		totalItems, err := CountDocs(db, matchStage)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		log.Println("totalItems", *totalItems)

		output := gin.H{
			"totalItems": totalItems,
			"items":      items,
		}

		log.Println("output", output)

		result["status"] = http.StatusOK
		result["isError"] = false
		result["data"] = output
		result["message"] = "ສໍາເລັດ"

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
