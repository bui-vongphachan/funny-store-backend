package product_attribute

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductAttribute struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	Title             string             `json:"title" bson:"title"`
	Image             string             `json:"image,omitempty" bson:"image,omitempty"`
	ProductID         primitive.ObjectID `json:"productId" bson:"productId"`
	AttributeGroupID  primitive.ObjectID `json:"attributeGroupId" bson:"attributeGroupId"`
	Delete            bool               `json:"delete" bson:"delete"`
	CloudflareImageID string             `json:"cloudflareImageId,omitempty" bson:"cloudflareImageId,omitempty"`
	OriginalID        primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID                primitive.ObjectID `json:"_id" bson:"_id"`
	Title             string             `json:"title" bson:"title"`
	Image             string             `json:"image,omitempty" bson:"image,omitempty"`
	ProductID         primitive.ObjectID `json:"productId" bson:"productId"`
	AttributeGroupID  primitive.ObjectID `json:"attributeGroupId" bson:"attributeGroupId"`
	Delete            bool               `json:"delete" bson:"delete"`
	CloudflareImageID string             `json:"cloudflareImageId,omitempty" bson:"cloudflareImageId,omitempty"`
	OriginalID        primitive.ObjectID `json:"originalId" bson:"originalId"`
}
