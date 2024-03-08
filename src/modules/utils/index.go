package utils

import (
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func MakeSkipOffsetPipeLine(query *url.Values) *bson.D {
	var pipeline bson.D

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

		pipeline = append(pipeline, bson.E{Key: "$limit", Value: limit})
	}

	if query.Has(jsonSkip) {
		intValue, err := strconv.Atoi(query.Get(jsonSkip))
		if err != nil {
			skip = 0
		} else if intValue < 0 {
			skip = 0
		}

		pipeline = append(pipeline, bson.E{Key: "$skip", Value: skip})

	}

	pipeline = append(pipeline, bson.E{Key: "$skip", Value: 0})

	return &pipeline
}
