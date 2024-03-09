package product

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty() *Product {
	output := Product{
		ID:                  primitive.NewObjectID(),
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

func Save(db *mongo.Database, product *Product) (*Product, error) {

	result, err := db.Collection(CollectionName).InsertOne(context.TODO(), product)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	if result.InsertedID == nil {
		err := errors.New("inserted id is nil")
		log.Println(err.Error())
		return nil, err
	}

	return product, nil
}
