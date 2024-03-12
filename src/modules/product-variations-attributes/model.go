package product_variations_attributes

import (
	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	product_variations "github.com/vongphachan/funny-store-backend/src/modules/product-variations"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductVariationAttribute struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	AttributeGroupId primitive.ObjectID `json:"attributeGroupId" bson:"attributeGroupId"`
	VariationId      primitive.ObjectID `json:"variationId" bson:"variationId"`
	AttributeId      primitive.ObjectID `json:"attributeId" bson:"attributeId"`
	ProductId        primitive.ObjectID `json:"productId" bson:"productId"`
	OriginalId       primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type PopulatedProductVariationAttribute struct {
	Attribute      product_attribute.ProductAttribute
	Variation      product_variations.ProductVariation
	AttributeGroup product_attribute_group.AttributeGroup

	// Base struct
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	AttributeGroupId primitive.ObjectID `json:"attributeGroupId" bson:"attributeGroupId"`
	VariationId      primitive.ObjectID `json:"variationId" bson:"variationId"`
	AttributeId      primitive.ObjectID `json:"attributeId" bson:"attributeId"`
	ProductId        primitive.ObjectID `json:"productId" bson:"productId"`
	OriginalId       primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type Props_Replicate struct {
	ProductId       *primitive.ObjectID
	SourceList      *[]ProductVariationAttribute
	attributeGroups *map[string]*primitive.ObjectID
	attributes      *map[string]*primitive.ObjectID
	variations      *map[string]*primitive.ObjectID
}

type Props_FindAllByProductIdWithDataPopulation struct {
	DB             *mongo.Database
	ProductID      *primitive.ObjectID
	SessionContext *mongo.SessionContext
}
