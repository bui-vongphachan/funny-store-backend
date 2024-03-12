package product_attribute

import (
	"context"
	"errors"
	"log"
	"net/url"

	"github.com/vongphachan/funny-store-backend/src/modules/utils"
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
		log.Println(err.Error())
		return nil
	}

	return payload
}

func MakeMatchPaginationPipeline(query url.Values) *bson.D {
	var matchStage bson.D

	utils.MakeObjectIdDocument(jsonID, query.Get(jsonID), &matchStage)

	utils.MakeObjectIdDocument(jsonAttributeGroupId, query.Get(jsonAttributeGroupId), &matchStage)

	utils.MakeObjectIdDocument(jsonProductId, query.Get(jsonProductId), &matchStage)

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

func FindById(db *mongo.Database, id *string) (*ProductAttribute, error) {
	context := context.TODO()

	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	cursor := db.Collection(CollectionName).FindOne(context, filter)
	if cursor.Err() != nil {
		log.Println(cursor.Err().Error())
		return nil, cursor.Err()
	}

	var result ProductAttribute
	if err := cursor.Decode(&result); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &result, nil
}

func BindNewData(input *ProductAttribute, data *ProductAttribute) (*ProductAttribute, error) {
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

func UpdateOne(db *mongo.Database, filter *bson.M, payload *ProductAttribute) (*ProductAttribute, error) {

	context := context.TODO()

	update := bson.M{"$set": payload}

	result, err := db.Collection(CollectionName).UpdateOne(context, filter, update)
	if err != nil {
		er := errors.New("unable to update product attribute")
		log.Println(er.Error())
		return nil, er
	}

	if result.MatchedCount == 0 {
		log.Println("attribute not product matched")
	}

	if result.ModifiedCount == 0 {
		log.Println("attribute not product updated")
	}

	return payload, nil
}

func FindAllByProductId(db *mongo.Database, productId *string, sessionContext *mongo.SessionContext) ([]*ProductAttribute, error) {

	objectID, err := primitive.ObjectIDFromHex(*productId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	pipelines := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.D{{Key: jsonProductId, Value: objectID}, {Key: jsonDelete, Value: false}}}},
	}

	cursor, err := db.Collection(CollectionName).Aggregate(*sessionContext, pipelines)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result []*ProductAttribute
	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func Replicate(productId *primitive.ObjectID, input *[]ProductAttribute) *[]ProductAttribute {

	newList := *input

	for index, item := range *input {

		newList[index] = ProductAttribute{
			ID:               primitive.NewObjectID(),
			Title:            item.Title,
			Image:            item.Image,
			ProductID:        *productId,
			AttributeGroupID: item.AttributeGroupID,
			Delete:           item.Delete,
			OrgirinalID:      item.ID,
		}
	}

	return &newList
}
