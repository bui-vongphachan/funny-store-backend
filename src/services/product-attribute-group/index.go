package serviceproductattributegroup

import (
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEmpty(productId *primitive.ObjectID) *modelproduct.AttributeGroup {
	output := modelproduct.AttributeGroup{
		ID:        primitive.NewObjectID(),
		Title:     "",
		IsPrimary: true,
		ProductID: *productId,
		Delete:    false,
	}

	return &output
}
