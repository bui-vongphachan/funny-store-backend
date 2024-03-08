package serviceproductattributegroup

import (
	"context"
	"log"

	collectionname "github.com/vongphachan/funny-store-backend/src/constraints/table-names"
	modelproduct "github.com/vongphachan/funny-store-backend/src/models/products"
	"go.mongodb.org/mongo-driver/bson"
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

func FindById(db *mongo.Database, id *string) *modelproduct.AttributeGroup {
	context := context.TODO()

	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	filter := bson.M{"_id": objectID}

	result := db.Collection(collectionname.PRODUCT_ATTRIBUTE_GROUPS).FindOne(context, filter)
	if result.Err() != nil {
		log.Fatalln(result.Err().Error())
		return nil
	}

	var attributeGroup modelproduct.AttributeGroup
	err = result.Decode(&attributeGroup)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return &attributeGroup
}

// func ExcludeEmptyData(payload *modelproduct.AttributeGroup) *modelproduct.AttributeGroup {
// 	data := make(map[string]interface{})

// 	if payload.Title != "" {
// 		data["Title"] = payload.Title
// 	}

// 	output := &modelproduct.AttributeGroup{}
// }

// func Update(db *mongo.Database, id *string) {
// 	filter := bson.M{"_id": id}

// 	context := context.TODO()

// 	db.Collection(collectionname.PRODUCT_ATTRIBUTE_GROUPS).UpdateOne(context, filter, bson.M{"$set": update})

// }
