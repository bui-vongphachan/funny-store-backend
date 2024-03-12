package product_attribute_group

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AttributeGroup struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	IsPrimary  bool               `json:"isPrimary" bson:"isPrimary"`
	ProductID  primitive.ObjectID `json:"productId" bson:"productId"`
	Delete     bool               `json:"delete" bson:"delete"`
	OriginalID primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	IsPrimary  bool               `json:"isPrimary" bson:"isPrimary"`
	ProductID  primitive.ObjectID `json:"productId" bson:"productId"`
	Delete     bool               `json:"delete" bson:"delete"`
	OriginalID primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type Props_Relicate struct {
	NewProductID   *primitive.ObjectID
	SourceList     *[]AttributeGroup
	DB             *mongo.Database
	SessionContext *mongo.SessionContext
}
