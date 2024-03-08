package utils

import (
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeSkipStage(query url.Values) *primitive.D {

	limit := 10

	if query.Has(jsonLimit) {
		intValue, err := strconv.Atoi(query.Get(jsonLimit))

		if err != nil {
			limit = 10
		} else if intValue < 0 {
			limit = 10
		} else if intValue > 100 {
			limit = 100
		}

	}

	return &bson.D{{
		Key:   "$limit",
		Value: limit,
	}}
}

func MakeLimitStage(query url.Values) *primitive.D {
	skip := 0

	if query.Has(jsonSkip) {
		intValue, err := strconv.Atoi(query.Get(jsonSkip))
		if err != nil {
			skip = 0
		} else if intValue < 0 {
			skip = 0
		}

	}

	return &bson.D{{
		Key:   "$skip",
		Value: skip,
	}}
}
