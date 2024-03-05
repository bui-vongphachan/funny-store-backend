package modelproduct

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttributeGroup struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	IsPrimary bool               `json:"isPrimary"`
	ProductID primitive.ObjectID `json:"productId"`
	Delete    bool               `json:"delete"`
}

type Attribute struct {
	ID                primitive.ObjectID  `json:"_id"`
	Title             string              `json:"title"`
	Image             string              `json:"image,omitempty"`
	ProductID         primitive.ObjectID  `json:"productId"`
	AttributeGroupID  primitive.ObjectID  `json:"attributeGroupId"`
	Delete            bool                `json:"delete"`
	CloudflareImageID string              `json:"cloudflareImageId,omitempty"`
	ReuseFrom         *primitive.ObjectID `json:"reuseFrom,omitempty"`
}

type Variation struct {
	ID              primitive.ObjectID  `json:"_id"`
	ProductID       primitive.ObjectID  `json:"productId"`
	Price           float64             `json:"price"`
	Delete          bool                `json:"delete"`
	Stock           int                 `json:"stock"`
	IsSingleVariant bool                `json:"isSingleVariant"`
	IsReady         bool                `json:"isReady"`
	ReuseForm       *primitive.ObjectID `json:"reuseForm,omitempty"`
}

type VariationAttribute struct {
	ID               primitive.ObjectID `json:"_id"`
	VariationID      primitive.ObjectID `json:"variationId"`
	AttributeID      primitive.ObjectID `json:"attributeId"`
	AttributeGroupID primitive.ObjectID `json:"attributeGroupId"`
}
