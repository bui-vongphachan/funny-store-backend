package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	"go.mongodb.org/mongo-driver/mongo"
)

func API_CreateDraft(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/draft", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		product := CreateEmpty()

		attributeGroup, err := product_attribute_group.CreateEmpty(&product.ID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		attribute, err := product_attribute.CreateEmpty(&product.ID, &attributeGroup.ID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		Save(db, product)

		product_attribute_group.Save(db, attributeGroup)

		product_attribute.Save(db, attribute)

		result["data"] = product

		c.JSON(http.StatusOK, result)
	})

}
