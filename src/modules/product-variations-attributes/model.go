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
}

type PopulatedProductVariationAttribute struct {
	Attribute      product_attribute.ProductAttribute
	Variation      product_variations.ProductVariation
	AttributeGroup product_attribute_group.AttributeGroup
}

type ReplicateProps struct {
	DB              *mongo.Database
	TargetProductID *primitive.ObjectID
	ProductId       *primitive.ObjectID
}
