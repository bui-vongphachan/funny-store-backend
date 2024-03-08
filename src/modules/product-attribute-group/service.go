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

func MakeMatchPaginationPipeline(query url.Values) *primitive.D {
	var matchStage bson.D

	log.Println("matchStage")
	log.Println(matchStage)

	objectID, err := primitive.ObjectIDFromHex(query.Get(jsonID))
	if err == nil || query.Get(jsonID) != "" {
		// add new filter document
		newDoc := bson.E{Key: jsonID, Value: objectID}
		matchStage = append(matchStage, newDoc)
	}

	productObjectId, err := primitive.ObjectIDFromHex(query.Get(jsonProductID))
	if err == nil || query.Get(jsonProductID) != "" {
		newDoc := bson.E{Key: jsonProductID, Value: productObjectId}
		matchStage = append(matchStage, newDoc)
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

		matchStage = append(matchStage, newDoc...)
	}

	if len(matchStage) == 0 {
		return nil
	}

	return &matchStage
}

func CreateEmpty(productId *string) (*AttributeGroup, error) {
	productObjectId, err := primitive.ObjectIDFromHex(*productId)
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

func Save(db *mongo.Database, attributeGroup *AttributeGroup) (*AttributeGroup, error) {

	context := context.TODO()

	result, err := db.Collection(CollectionName).InsertOne(context, attributeGroup)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	if result.InsertedID == nil {
		er := errors.New("inserted id is nil")
		log.Println(er.Error())
		return nil, er
	}

	return attributeGroup, nil
}

func FindById(db *mongo.Database, id *string) (*AttributeGroup, error) {
	context := context.TODO()

	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result := db.Collection(CollectionName).FindOne(context, filter)
	if result.Err() != nil {
		log.Fatalln(result.Err().Error())
		return nil, result.Err()
	}

	var attributeGroup AttributeGroup
	if err := result.Decode(&attributeGroup); err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &attributeGroup, nil
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
