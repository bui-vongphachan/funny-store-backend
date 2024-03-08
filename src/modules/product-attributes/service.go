package product_attribute

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty(productId *string, attributeGroupId *string) (*ProductAttribute, error) {
	productObjectId, err := primitive.ObjectIDFromHex(*productId)
	if err != nil {
		err := errors.New("invalid product id")
		log.Println(err.Error())
		return nil, err
	}

	productAttributeGroupObjectId, err := primitive.ObjectIDFromHex(*productId)
	if err != nil {
		err := errors.New("invalid attribute group id")
		log.Println(err.Error())
		return nil, err
	}

	output := ProductAttribute{
		ID:               primitive.NewObjectID().Hex(),
		Title:            "",
		Image:            "",
		ProductID:        productObjectId.String(),
		AttributeGroupID: productAttributeGroupObjectId.String(),
	}

	return &output, nil
}

func Save(db *mongo.Database, payload *ProductAttribute) *ProductAttribute {

	context := context.TODO()

	_, err := db.Collection(CollectionName).InsertOne(context, payload)

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
