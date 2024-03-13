package admins

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty() {}

func Save() {}

func FindOneById() {}

func FindAll() {}

func FindOneByEmail(db *mongo.Database, email *string) (*Admin, error) {
	filter := primitive.D{{Key: Field_Email, Value: email}}

	admin := Admin{}
	err := db.Collection(CollectionName).FindOne(context.TODO(), filter).Decode(&admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}
