package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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
}
