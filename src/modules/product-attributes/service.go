package product_attribute

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty(productId *primitive.ObjectID, attributeGroupId *primitive.ObjectID) (*ProductAttribute, error) {

	output := ProductAttribute{
		ID:               primitive.NewObjectID(),
		Title:            "",
		Image:            "",
		ProductID:        *productId,
		AttributeGroupID: *attributeGroupId,
	}

	return &output, nil
}

func Save(db *mongo.Database, payload *ProductAttribute) *ProductAttribute {

	_, err := db.Collection(CollectionName).InsertOne(context.TODO(), payload)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return payload
}

func SaveBulk(db *mongo.Database, payload *[]interface{}) *[]interface{} {
	context := context.TODO()

	_, err := db.Collection(CollectionName).InsertMany(context, *payload)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return payload
}
