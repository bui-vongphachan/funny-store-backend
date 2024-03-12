package product

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	product_variations "github.com/vongphachan/funny-store-backend/src/modules/product-variations"
	product_variations_attributes "github.com/vongphachan/funny-store-backend/src/modules/product-variations-attributes"
	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func API_CreateDraft(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/draft", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
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

		productVariations, err := product_variations.CreateEmpty(&product.ID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		saveResult, err := Save(db, product)
		if err != nil || saveResult == nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		product_attribute_group.Save(db, attributeGroup)

		product_attribute.Save(db, attribute)

		product_variations.Save(db, productVariations)

		result["data"] = product
		result["status"] = http.StatusCreated
		result["isError"] = false
		result["message"] = "Success"

		c.JSON(http.StatusOK, result)
	})

}

func API_Replicate(db *mongo.Database, r *gin.Engine) {
	r.POST("/product/replicate", func(c *gin.Context) {
		result := gin.H{
			"status":  http.StatusBadRequest,
			"isError": true,
			"data":    nil,
			"message": "Invalid data",
		}

		var requestBody PropsReplicateAPI

		err := c.Bind(&requestBody)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		sourceProductId, err := utils.MakeObjectId(requestBody.SourceProductID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		targetProductId, err := utils.MakeObjectId(requestBody.TargetProductID)
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		session, err := db.Client().StartSession()
		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}

		opts := options.Transaction()

		err = mongo.WithSession(context.Background(), session, func(sessionCtx mongo.SessionContext) error {
			err := session.StartTransaction(opts)
			if err != nil {
				return err
			}

			var product *Product
			var attributeGroups *[]product_attribute_group.AttributeGroup
			var attributes *[]product_attribute.ProductAttribute
			var variationAttributes *[]product_variations_attributes.ProductVariationAttribute
			var variations *[]product_variations.ProductVariation

			attributeMap := map[string]*primitive.ObjectID{}
			variationMap := map[string]*primitive.ObjectID{}
			attributeGroupMap := map[string]*primitive.ObjectID{}

			// Retrieve data
			{
				// get product data
				{
					replicateProps := ReplicateProps{
						DB:              db,
						TargetProductID: targetProductId,
						SourceProductID: sourceProductId,
						SessionContext:  &sessionCtx,
					}
					result, err := Replicate(&replicateProps)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}

					product = result
				}
				// get product attribute group data
				{
					result, err := product_attribute_group.FindAllByProductId(db, &requestBody.SourceProductID, &sessionCtx)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					attributeGroups = result
				}
				// get product attribute data
				{
					result, err := product_attribute.FindAllByProductId(db, &requestBody.SourceProductID, &sessionCtx)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					attributes = result
				}
				// get product variation data
				{
					result, err := product_variations.FindAllByProductId(db, &requestBody.SourceProductID)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					variations = result
				}
				// get product variation attribute data
				{
					result, err := product_variations_attributes.FindByProductId(
						db,
						sourceProductId,
					)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					variationAttributes = result
				}
			}

			// Replicate data
			{
				// replicate product attribute group
				{
					updatedItems, err := product_attribute_group.RelicateAndSave(&product_attribute_group.Props_Relicate{
						DB:             db,
						NewProductID:   targetProductId,
						SourceList:     attributeGroups,
						SessionContext: &sessionCtx,
					})
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					for _, item := range *updatedItems {
						attributeGroupMap[item.OriginalID.String()] = &item.ID
					}
				}
				// replicate product attributes
				{
					updatedItems, err := product_attribute.ReplicateAndSave(
						targetProductId,
						attributes,
						db,
						&sessionCtx,
					)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					for _, item := range *updatedItems {
						attributeMap[item.OriginalID.String()] = &item.ID
					}
				}
				// replicate product variations
				{
					updatedItems, err := product_variations.ReplicateAndSave(
						targetProductId,
						variations,
						db,
						&sessionCtx,
					)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
					for _, item := range *updatedItems {
						variationMap[item.OriginalID.String()] = &item.ID
					}
				}
				// replicate product variation attributes
				{
					_, err := product_variations_attributes.RelicateAndSave(&product_variations_attributes.Props_Replicate{
						ProductId:       targetProductId,
						SourceList:      variationAttributes,
						AttributeGroups: &attributeGroupMap,
						Attributes:      &attributeMap,
						Variations:      &variationMap,
					}, db, &sessionCtx)
					if err != nil {
						session.AbortTransaction(sessionCtx)
						return err
					}
				}
			}

			result["data"] = product
			result["status"] = http.StatusCreated
			result["isError"] = false
			result["message"] = "Success"

			session.CommitTransaction(sessionCtx)

			return nil
		})

		if err != nil {
			result["message"] = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		c.JSON(http.StatusOK, result)

	})
}
