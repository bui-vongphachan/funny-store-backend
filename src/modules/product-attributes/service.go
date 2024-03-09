package product_attribute

import (
	"context"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
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

func CountDocs(db *mongo.Database, filter *bson.D) (*int64, error) {
	if filter == nil {
		filter = &bson.D{}
	}

	cursor, err := db.Collection(CollectionName).CountDocuments(context.TODO(), *filter)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &cursor, err
}

func MakeMatchPaginationPipeline(query url.Values) *bson.D {
	var matchStage bson.D

	objectID, err := primitive.ObjectIDFromHex(query.Get(jsonID))
	if err == nil || query.Get(jsonID) != "" {
		newDoc := bson.E{Key: jsonID, Value: objectID}
		matchStage = append(matchStage, newDoc)
	}

	attributeGroupObjectId, err := primitive.ObjectIDFromHex(query.Get(jsonAttributeGroupId))
	if err == nil || query.Get(jsonAttributeGroupId) != "" {
		newDoc := bson.E{Key: jsonAttributeGroupId, Value: attributeGroupObjectId}
		matchStage = append(matchStage, newDoc)
	}

	productObjectId, err := primitive.ObjectIDFromHex(query.Get(jsonProductId))
	if err == nil || query.Get(jsonProductId) != "" {
		newDoc := bson.E{Key: jsonProductId, Value: productObjectId}
		matchStage = append(matchStage, newDoc)
	}

	if query.Has(jsonTitle) && query.Get(jsonTitle) != "" {
		regexOptions := bson.D{{
			Key:   "$regex",
			Value: query.Get(jsonTitle),
		}, {
			Key:   "$options",
			Value: "i",
		}}

		newDoc := bson.E{Key: jsonTitle, Value: regexOptions}

		matchStage = append(matchStage, newDoc)
	}

	if len(matchStage) == 0 {
		return nil
	}

	return &matchStage
}
