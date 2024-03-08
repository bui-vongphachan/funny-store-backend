package serviceproductattributegroup

import (
	"context"
	"log"

	collectionname "github.com/vongphachan/funny-store-backend/src/constraints/table-names"
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func Save(db *mongo.Database, attributeGroup *modelproduct.AttributeGroup) *modelproduct.AttributeGroup {

	context := context.TODO()

	_, err := db.Collection(collectionname.PRODUCT_ATTRIBUTE_GROUPS).InsertOne(context, attributeGroup)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return attributeGroup
}
