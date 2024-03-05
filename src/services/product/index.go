package serviceproduct

import (
	model "github.com/vongphachan/funny-store-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEmpty() model.Product {
	output := model.Product{
		ID:                  primitive.NewObjectID().Hex(),
		Title:               "",
		Description:         "",
		PreviewImages:       []string{},
		Gallery:             []string{},
		Delete:              false,
		HavingSingleVariant: false,
		Image:               "",
		IsDraft:             true,
	}

	return output
}
