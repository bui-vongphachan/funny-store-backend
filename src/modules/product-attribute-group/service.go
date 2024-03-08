package product_attribute_group

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
