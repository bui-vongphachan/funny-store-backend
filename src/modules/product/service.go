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
		log.Println(err.Error())
		return nil, err
	}

	if result.InsertedID == nil {
		err := errors.New("inserted id is nil")
		log.Println(err.Error())
		return nil, err
	}

	return product, nil
}

func FindByID(db *mongo.Database, id *primitive.ObjectID) (*Product, error) {
	filter := primitive.D{{Key: Field_ID, Value: id}}

	var product Product
	err := db.Collection(CollectionName).FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &product, nil
}

func Replicate(props *ReplicateProps) (*Product, error) {
	originalProduct, err := FindByID(props.DB, props.SourceProductID)
	if err != nil {
		return nil, err
	}

	originalProduct.ID = primitive.NewObjectID()
	originalProduct.OriginalID = *props.TargetProductID
	originalProduct.IsDraft = true
	originalProduct.Delete = false
	originalProduct.Title = originalProduct.Title + " (Copy)"

	_, err = Save(props.DB, originalProduct)
	if err != nil {
		return nil, err
	}

	return originalProduct, nil
}
