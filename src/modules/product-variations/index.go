package product_variations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
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

func API_Pagination(db *mongo.Database, r *gin.Engine) {
	r.GET("/products/variations", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
		}

		pipelines := mongo.Pipeline{}

		matchStage := MakeMatchPaginationPipeline(c.Request.URL.Query())

		if matchStage != nil {
			if len(*matchStage) > 0 {
				pipelines = append(pipelines, bson.D{{Key: "$match", Value: *matchStage}})
			}
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
		result["message"] = "Success"

		c.JSON(http.StatusOK, result)
	})
}
