package product_variations_attributes

import (
	"context"
	"errors"
	"log"

	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	product_variations "github.com/vongphachan/funny-store-backend/src/modules/product-variations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateAttribute(db *mongo.Database, input *ProductVariationAttribute) (*ProductVariationAttribute, error) {

	filter := bson.D{{Key: "$and", Value: bson.A{
		bson.D{{Key: jsonAttributeGroupId, Value: input.AttributeGroupId}},
		bson.D{{Key: jsonVariationId, Value: input.VariationId}},
		bson.D{{Key: jsonProductId, Value: input.ProductId}},
	}}}

	target := bson.E{Key: jsonAttributeId, Value: input.AttributeId}

	update := bson.D{{Key: "$set", Value: target}}

	options := options.Update().SetUpsert(true)

	result, err := db.Collection(CollectionName).UpdateOne(context.TODO(), filter, update, options)
	if err != nil {
		er := errors.New("unable to update product attribute")
		log.Println(er.Error())
		return nil, er
	}

	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		err := errors.New("attribute not updated")
		return nil, err
	}

	return nil, nil
}

func FindAllByProductIdWithDataPopulation(props *Props_FindAllByProductIdWithDataPopulation) (*[]PopulatedProductVariationAttribute, error) {
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: Field_ProductId, Value: props.ProductID}}}}

	productAttributeLookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: product_attribute.CollectionName},
		{Key: "localField", Value: Field_AttributeId},
		{Key: "foreignField", Value: product_attribute.Field_ID},
		{Key: "as", Value: "attribute"},
	}}}

	productAttributeUnwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$attribute"}}}}

	productVariationLookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: product_variations.CollectionName},
		{Key: "localField", Value: Field_VariationId},
		{Key: "foreignField", Value: product_variations.Field_OriginalID},
		{Key: "as", Value: "variation"},
	}}}

	productVariationUnwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$variation"}}}}

	productAttributeGroupLookupStage := bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: product_attribute_group.CollectionName},
		{Key: "localField", Value: Field_AttibuteGroupId},
		{Key: "foreignField", Value: product_attribute_group.Field_OriginalID},
		{Key: "as", Value: "attributeGroup"},
	}}}

	productAttributeGroupUnwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$attributeGroup"}}}}

	cursor, err := props.DB.Collection(CollectionName).Aggregate(*props.SessionContext, mongo.Pipeline{
		matchStage,
		productAttributeLookupStage,
		productAttributeUnwindStage,
		productVariationLookupStage,
		productVariationUnwindStage,
		productAttributeGroupLookupStage,
		productAttributeGroupUnwindStage,
	})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result []PopulatedProductVariationAttribute
	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &result, nil
}

func Replicate(props *Props_Replicate) (*[]ProductVariationAttribute, error) {

	items, err := FindAllByProductIdWithDataPopulation(&Props_FindAllByProductIdWithDataPopulation{
		DB:             props.DB,
		ProductID:      props.ProductId,
		SessionContext: props.SessionContext,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Convert items to ProductVariationAttribute with new ObjectID and its ancestor's ObjectID
	newList := []ProductVariationAttribute{}
	for _, item := range *items {
		newItem := ProductVariationAttribute{
			ID:               primitive.NewObjectID(),
			AttributeGroupId: item.AttributeGroup.ID,
			VariationId:      item.Variation.ID,
			AttributeId:      item.Attribute.ID,
			OriginalId:       item.ID,
		}

		newList = append(newList, newItem)
	}

	return &newList, nil
}

func RelicateAndSave(props *Props_Replicate) (*[]ProductVariationAttribute, error) {
	replicatedItems, err := Replicate(props)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Convert replicatedItems to []interface{}
	var items []interface{}
	for _, item := range *replicatedItems {
		items = append(items, item)
	}

	result, err := props.DB.Collection(CollectionName).InsertMany(*props.SessionContext, items)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if len(result.InsertedIDs) == 0 {
		err := errors.New("unable to insert replicated items")
		log.Println(err.Error())
		return nil, err
	}

	return replicatedItems, nil
}
