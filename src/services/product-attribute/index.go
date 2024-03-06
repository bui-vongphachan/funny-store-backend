package serviceproductattribute

import (
	model "github.com/vongphachan/funny-store-backend/src/models"
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateEmptyProps struct {
	Product        *model.Product
	AttributeGroup *modelproduct.AttributeGroup
}

func CreateEmpty(props *CreateEmptyProps) *modelproduct.Attribute {

	output := modelproduct.Attribute{
		ID:               primitive.NewObjectID(),
		Title:            "",
		Image:            "",
		ProductID:        props.Product.ID,
		AttributeGroupID: props.AttributeGroup.ID,
	}

	return &output
}
