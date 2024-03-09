package utils

import (
	"context"
	"log"
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeSkipStage(query *url.Values) *primitive.D {

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

func MakeLimitStage(query *url.Values) *primitive.D {
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

func MakePaginationQuery(props *MakePaginationQueryType) (*[]bson.M, error) {

	skipStage := MakeSkipStage(&props.UrlQuery)
	*props.MongoPipeline = append(*props.MongoPipeline, *skipStage)

	limitStage := MakeLimitStage(&props.UrlQuery)
	*props.MongoPipeline = append(*props.MongoPipeline, *limitStage)

	cursor, err := props.DB.Collection(props.CollectionName).Aggregate(context.TODO(), *props.MongoPipeline)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	items := []bson.M{}
	if err := cursor.All(context.TODO(), &items); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &items, nil
}

func CountDocs(db *mongo.Database, filter *bson.D, collectionName string) (*int64, error) {
	if filter == nil {
		filter = &bson.D{}
	}

	cursor, err := db.Collection(collectionName).CountDocuments(context.TODO(), *filter)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &cursor, err
}
