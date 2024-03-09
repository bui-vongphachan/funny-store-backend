package product_attribute

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_Create(db *mongo.Database, r *gin.Engine) {
	r.POST("/products/attributes", func(c *gin.Context) {
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

func API_Pagination(db *mongo.Database, r *gin.Engine) {
	r.GET("/products/attributes", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		pipelines := mongo.Pipeline{}

		matchStage := MakeMatchPaginationPipeline(c.Request.URL.Query())

		if matchStage != nil {
			pipelines = append(pipelines, bson.D{{Key: "$match", Value: *matchStage}})
		}

		makePaginationQuery := utils.MakePaginationQueryType{
			DB:             db,
			UrlQuery:       c.Request.URL.Query(),
			CollectionName: CollectionName,
			MongoPipeline:  &pipelines,
		}

		items, err := utils.MakePaginationQuery(&makePaginationQuery)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		totalItems, err := utils.CountDocs(db, matchStage, CollectionName)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		output := gin.H{
			"totalItems": totalItems,
			"items":      items,
		}

		result["status"] = http.StatusOK
		result["isError"] = false
		result["data"] = output
		result["message"] = "ສໍາເລັດ"

		c.JSON(http.StatusOK, result)
	})
}
