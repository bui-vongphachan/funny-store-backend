package product_attribute_group

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

		productAttribute, err := product_attribute.CreateEmpty(&requestBody.ProductID, &attributeGroup.ID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusBadRequest, result)
			return
		}

		product_attribute.Save(db, productAttribute)

		result["data"] = attributeGroup
		result["status"] = http.StatusCreated
		result["isError"] = false
		result["message"] = "ສໍາເລັດ"

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

func API_Update(db *mongo.Database, r *gin.Engine) {
	r.PATCH("/product/attribute-group/:id", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "Unable to update",
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
