package serviceproductattribute

import (
	"context"
	"log"

	collectionname "github.com/vongphachan/funny-store-backend/src/constraints/table-names"
	model "github.com/vongphachan/funny-store-backend/src/models"
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func Save(db *mongo.Database, payload *modelproduct.Attribute) *modelproduct.Attribute {

	context := context.TODO()

	_, err := db.Collection(collectionname.PRODUCT_ATTRIBUTE).InsertOne(context, payload)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return payload
}

func SaveBulk(db *mongo.Database, payload *[]interface{}) *[]interface{} {
	context := context.TODO()

	_, err := db.Collection(collectionname.PRODUCT_ATTRIBUTE).InsertMany(context, *payload)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return payload
}
