package utils

import (
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
)

type MakePaginationQueryType struct {
	DB             *mongo.Database
	CollectionName string
	UrlQuery       url.Values
	MongoPipeline  *mongo.Pipeline
}
