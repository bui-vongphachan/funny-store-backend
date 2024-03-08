package product

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty() *Product {
	output := Product{
		ID:                  primitive.NewObjectID().Hex(),
		Title:               "",
		Description:         "",
		PreviewImages:       []string{},
		Gallery:             []string{},
		Delete:              false,
		HavingSingleVariant: false,
		Image:               "",
		IsDraft:             true,
	}

	return &output
}

func Save(db *mongo.Database, product *Product) *Product {

	_, err := db.Collection(CollectionName).InsertOne(context.TODO(), product)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return product
}
