package serviceproduct

import (
	"context"
	"log"

	collectionname "github.com/vongphachan/funny-store-backend/src/constraints/table-names"
	model "github.com/vongphachan/funny-store-backend/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEmpty() *model.Product {
	output := model.Product{
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

func Save(db *mongo.Database, product *model.Product) *model.Product {

	context := context.TODO()

	_, err := db.Collection(collectionname.PRODUCT).InsertOne(context, product)

	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return product
}
