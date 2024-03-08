package modelproduct

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttributeGroup struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	IsPrimary bool               `json:"isPrimary" bson:"isPrimary"`
	ProductID primitive.ObjectID `json:"productId" bson:"productId"`
	Delete    bool               `json:"delete" bson:"delete"`
}

type Attribute struct {
	ID                primitive.ObjectID  `json:"_id" bson:"_id"`
	Title             string              `json:"title" bson:"title"`
	Image             string              `json:"image,omitempty" bson:"image,omitempty"`
	ProductID         primitive.ObjectID  `json:"productId" bson:"productId"`
	AttributeGroupID  primitive.ObjectID  `json:"attributeGroupId" bson:"attributeGroupId"`
	Delete            bool                `json:"delete" bson:"delete"`
	CloudflareImageID string              `json:"cloudflareImageId,omitempty" bson:"cloudflareImageId,omitempty"`
	ReuseFrom         *primitive.ObjectID `json:"reuseFrom,omitempty" bson:"reuseFrom,omitempty"`
}

type Variation struct {
	ID              primitive.ObjectID  `json:"_id" bson:"_id"`
	ProductID       primitive.ObjectID  `json:"productId" bson:"productId"`
	Price           float64             `json:"price" bson:"price"`
	Delete          bool                `json:"delete" bson:"delete"`
	Stock           int                 `json:"stock" bson:"stock"`
	IsSingleVariant bool                `json:"isSingleVariant" bson:"isSingleVariant"`
	IsReady         bool                `json:"isReady" bson:"isReady"`
	ReuseForm       *primitive.ObjectID `json:"reuseForm,omitempty" bson:"reuseForm,omitempty"`
}

type VariationAttribute struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	VariationID      primitive.ObjectID `json:"variationId" bson:"variationId"`
	AttributeID      primitive.ObjectID `json:"attributeId" bson:"attributeId"`
	AttributeGroupID primitive.ObjectID `json:"attributeGroupId" bson:"attributeGroupId"`
}
