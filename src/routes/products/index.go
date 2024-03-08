package routeproduct

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	serviceproduct "github.com/vongphachan/funny-store-backend/src/services/product"
	serviceproductattribute "github.com/vongphachan/funny-store-backend/src/services/product-attribute"
	serviceproductattributegroup "github.com/vongphachan/funny-store-backend/src/services/product-attribute-group"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDraft(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/draft", func(c *gin.Context) {
		result := gin.H{
			"status":  400,
			"isError": true,
			"data":    nil,
			"message": "ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
		}

		product := serviceproduct.CreateEmpty()

		attributeGroup := serviceproductattributegroup.CreateEmpty(&product.ID)

		attribute := serviceproductattribute.CreateEmpty(&serviceproductattribute.CreateEmptyProps{
			Product:        product,
			AttributeGroup: attributeGroup,
		})

		serviceproduct.Save(db, product)

		serviceproductattributegroup.Save(db, attributeGroup)

		log.Println(attribute)

		result["data"] = product

		c.JSON(http.StatusOK, result)
	})

}
