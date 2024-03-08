package product_attribute_group

import (
	"context"
	"errors"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeMatchPaginationPipeline(query url.Values, pipeline *primitive.A) *primitive.A {
	var filterDoc bson.D

	objectID, err := primitive.ObjectIDFromHex(query.Get(jsonID))
	if err == nil {
		// add new filter document
		newDoc := bson.D{{Key: jsonID, Value: objectID}}
		filterDoc = append(filterDoc, newDoc...)
	}

	productObjectId, err := primitive.ObjectIDFromHex(query.Get(jsonID))
	if err != nil {
		newDoc := bson.D{{Key: jsonProductID, Value: productObjectId}}
		filterDoc = append(filterDoc, newDoc...)
	}

	if query.Has(jsonTitle) && query.Get(jsonTitle) != "" {
		// find any doc that contains the title
		regexOptions := bson.D{{
			Key:   "$regex",
			Value: query.Get(jsonTitle),
		}, {
			Key:   "$options",
			Value: "i",
		}}

		newDoc := bson.D{{Key: "title", Value: regexOptions}}

		filterDoc = append(filterDoc, newDoc...)
	}

	*pipeline = append(*pipeline, filterDoc)

	return pipeline
}

func CreateEmpty(productId *string) (*AttributeGroup, error) {
	productObjectId, err := primitive.ObjectIDFromHex(*productId)
	log.Println("productObjectId")
	log.Println(productObjectId)
	if err != nil {
		er := errors.New("invalid product id")
		log.Println(er.Error())
		return nil, er
	}

	output := AttributeGroup{
		ID:        primitive.NewObjectID().Hex(),
		Title:     "",
		IsPrimary: true,
		ProductID: productObjectId.Hex(),
		Delete:    false,
	}

	return &output, nil
}

func Save(db *mongo.Database, attributeGroup *AttributeGroup) *AttributeGroup {

	context := context.TODO()

	_, err := db.Collection(CollectionName).InsertOne(context, attributeGroup)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return attributeGroup
}

func FindById(db *mongo.Database, id *string) *AttributeGroup {
	context := context.TODO()

	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	filter := bson.M{"_id": objectID}

	result := db.Collection(CollectionName).FindOne(context, filter)
	if result.Err() != nil {
		log.Fatalln(result.Err().Error())
		return nil
	}

	var attributeGroup AttributeGroup
	err = result.Decode(&attributeGroup)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return &attributeGroup
}

func BindNewData(input *AttributeGroup, data *AttributeGroup) (*AttributeGroup, error) {
	if input == nil {
		er := errors.New("input is nil")
		log.Println(er.Error())
		return nil, er
	}

	if data == nil {
		er := errors.New("data is empty")
		log.Println(er.Error())
		return nil, er
	}

	if input.Title != "" {
		data.Title = input.Title
	}

	data.Delete = input.Delete

	return data, nil
}

func UpdateOne(db *mongo.Database, filter *bson.M, payload *AttributeGroup) (*AttributeGroup, error) {

	context := context.TODO()

	update := bson.M{"$set": payload}

	result, err := db.Collection(CollectionName).UpdateOne(context, filter, update)
	if err != nil {
		er := errors.New("unable to update attribute group")
		log.Println(er.Error())
		return nil, er
	}

	if result.MatchedCount == 0 {
		log.Println("attribute group not matched")
	}

	if result.ModifiedCount == 0 {
		log.Println("attribute group not updated")
	}

	return payload, nil
}
