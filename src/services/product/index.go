package serviceproduct

import (
	"context"
	"encoding/json"
	"fmt"

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

type SaveToDatabaseProps struct {
	DB      *mongo.Database
	Product *model.Product
}

func SaveToDatabase(props *SaveToDatabaseProps) *model.Product {

	context := context.TODO()

	cursur, err := props.DB.Collection(collectionname.PRODUCT).InsertOne(context, props.Product)

	if err != nil {
		return nil
	}

	jsonData, err := json.MarshalIndent(cursur, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)

	return props.Product
}
