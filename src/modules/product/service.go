package product

import (
	"context"
	"errors"
	"log"

	"github.com/vongphachan/funny-store-backend/src/modules/utils"
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

func FindByObjectID(props *FindByObjectIDProps) (*Product, error) {
	filter := primitive.D{{Key: Field_ID, Value: props.ID}}

	var product Product
	err := props.DB.Collection(CollectionName).FindOne(*props.SessionContext, filter).Decode(&product)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &product, nil
}

func FindByID(db *mongo.Database, id *string, sessionContext *mongo.SessionContext) (*Product, error) {
	objectId, err := utils.MakeObjectId(*id)
	if err != nil {
		return nil, err
	}

	product, err := FindByObjectID(&FindByObjectIDProps{
		DB:             db,
		ID:             objectId,
		SessionContext: sessionContext,
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return product, nil
}

func Replicate(props *ReplicateProps) (*Product, error) {
	originalProduct, err := FindByObjectID(&FindByObjectIDProps{
		DB:             props.DB,
		ID:             props.SourceProductID,
		SessionContext: props.SessionContext,
	})
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
