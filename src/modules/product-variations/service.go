package product_variations

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"

	"github.com/vongphachan/funny-store-backend/src/modules/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeMatchPaginationPipeline(query url.Values) *bson.D {
	var matchStage bson.D

	utils.MakeObjectIdDocument(jsonID, query.Get(jsonID), &matchStage)

	utils.MakeObjectIdDocument(jsonProductID, query.Get(jsonProductID), &matchStage)

	if query.Has(jsonIsSingleVariant) && query.Get(jsonIsSingleVariant) != "" {

		if boolValue, err := strconv.ParseBool(query.Get(jsonIsSingleVariant)); err == nil {
			newDoc := bson.E{Key: jsonIsSingleVariant, Value: boolValue}
			matchStage = append(matchStage, newDoc)
		}
	}

	return &matchStage
}

func CreateEmpty(productId *primitive.ObjectID) (*ProductVariation, error) {
	output := ProductVariation{
		ID:              primitive.NewObjectID(),
		ProductID:       *productId,
		Delete:          false,
		IsSingleVariant: false,
		Stock:           0,
		Price:           0,
	}

	return &output, nil
}

func Save(db *mongo.Database, productVariation *ProductVariation) (*ProductVariation, error) {
	context := context.TODO()

	result, err := db.Collection(CollectionName).InsertOne(context, productVariation)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if result.InsertedID == nil {
		er := errors.New("inserted id is nil")
		log.Println(er.Error())
		return nil, er
	}

	return productVariation, nil
}

func BindNewData(input *ProductVariation, data *ProductVariation) (*ProductVariation, error) {
	if input == nil {
		er := errors.New("input is nil")
		log.Println(er.Error())
		return nil, er
	}

	if data == nil {
		er := errors.New("data is empty")
		log.Println(er.Error())
		return nil, er
	}

	data.Stock = input.Stock
	data.Price = input.Price

	return data, nil
}
