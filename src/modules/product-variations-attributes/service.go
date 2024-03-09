package product_variations_attributes

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateAttribute(db *mongo.Database, input *ProductVariationAttribute) (*ProductVariationAttribute, error) {

	filter := bson.D{{Key: "$and", Value: bson.A{
		bson.D{{Key: jsonAttributeGroupId, Value: input.AttributeGroupId}},
		bson.D{{Key: jsonVariationId, Value: input.VariationId}},
	}}}

	target := bson.E{Key: jsonAttributeId, Value: input.AttributeId}

	update := bson.D{{Key: "$set", Value: target}}

	options := options.Update().SetUpsert(true)

	result, err := db.Collection(CollectionName).UpdateOne(context.TODO(), filter, update, options)
	if err != nil {
		er := errors.New("unable to update product attribute")
		log.Println(er.Error())
		return nil, er
	}

	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		err := errors.New("attribute not updated")
		return nil, err
	}

	return nil, nil
}
