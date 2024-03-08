package utils

import (
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeSkipOffsetPipeLine(query url.Values, pipeline *primitive.A) *primitive.A {

	limit := 10
	skip := 0

	if query.Has(jsonLimit) {
		intValue, err := strconv.Atoi(query.Get(jsonLimit))

		if err != nil {
			limit = 10
		} else if intValue < 0 {
			limit = 10
		} else if intValue > 100 {
			limit = 100
		}

		newPipeline := bson.D{{Key: "$limit", Value: limit}}

		*pipeline = append(*pipeline, newPipeline)
	}

	if query.Has(jsonSkip) {
		intValue, err := strconv.Atoi(query.Get(jsonSkip))
		if err != nil {
			skip = 0
		} else if intValue < 0 {
			skip = 0
		}

		newPipeline := bson.D{{Key: "$skip", Value: skip}}

		*pipeline = append(*pipeline, newPipeline)
	}

	return pipeline
}
