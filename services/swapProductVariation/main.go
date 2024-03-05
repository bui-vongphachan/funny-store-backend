package swapProductVariation

import (
	"context"
	"log"

	collectionname "github.com/vongphachan/funny-store-backend/constraints"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SwapProductVariation(db *mongo.Database, session *mongo.Session, fromProductObjectId *primitive.ObjectID, toProductObjectId *primitive.ObjectID) {

	if fromProductObjectId == nil || toProductObjectId == nil {
		return
	}

	context := context.TODO()

	searchStage := bson.M{"$match": bson.M{
		"productId": fromProductObjectId,
	}}

	cursur, err := db.Collection(collectionname.PRODUCT_VARIATION).Aggregate(
		context,
		[]bson.M{
			searchStage,
		},
	)

	if err != nil {
		panic(err)
	}

	log.Println(cursur)

}
