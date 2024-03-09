package product_variations

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductVariation struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ProductID       primitive.ObjectID `json:"productId" bson:"productId"`
	IsSingleVariant bool               `json:"isSingleVariant" bson:"isSingleVariant"`
	Stock           int                `json:"stock" bson:"stock"`
	Price           float64            `json:"price" bson:"price"`
	Delete          bool               `json:"delete" bson:"delete"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	ProductID       primitive.ObjectID `json:"productId" bson:"productId"`
	IsSingleVariant bool               `json:"isSingleVariant" bson:"isSingleVariant"`
	Stock           int                `json:"stock" bson:"stock"`
	Price           float64            `json:"price" bson:"price"`
	Delete          bool               `json:"delete" bson:"delete"`
}
