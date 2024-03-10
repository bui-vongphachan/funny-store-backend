package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID                  primitive.ObjectID `json:"_id" bson:"_id"`
	Title               string             `json:"title" bson:"title"`
	Description         string             `json:"description" bson:"description"`
	PreviewImages       []string           `json:"previewImages" bson:"previewImages"`
	Gallery             []string           `json:"gallery" bson:"gallery"`
	Delete              bool               `json:"delete" bson:"delete"`
	HavingSingleVariant bool               `json:"havingSingleVariant" bson:"havingSingleVariant"`
	Image               string             `json:"image,omitempty" bson:"image,omitempty"`
	IsDraft             bool               `json:"isDraft" bson:"isDraft"`
	OriginalID          primitive.ObjectID `json:"originalId" bson:"originalId"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID                  primitive.ObjectID `json:"_id" bson:"_id"`
	Title               string             `json:"title" bson:"title"`
	Description         string             `json:"description" bson:"description"`
	PreviewImages       []string           `json:"previewImages" bson:"previewImages"`
	Gallery             []string           `json:"gallery" bson:"gallery"`
	Delete              bool               `json:"delete" bson:"delete"`
	HavingSingleVariant bool               `json:"havingSingleVariant" bson:"havingSingleVariant"`
	Image               string             `json:"image,omitempty" bson:"image,omitempty"`
	IsDraft             bool               `json:"isDraft" bson:"isDraft"`
}

type ReplicateProps struct {
	DB              *mongo.Database
	TargetProductID *primitive.ObjectID
	SourceProductID *primitive.ObjectID
}

type PropsReplicateAPI struct {
	TargetProductID string `json:"targetProductId"`
	SourceProductID string `json:"sourceProductId"`
}
