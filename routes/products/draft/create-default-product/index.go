package createdefaultproduct

import (
	typeproduct "github.com/vongphachan/funny-store-backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Main() typeproduct.Product {
	output := typeproduct.Product{
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
