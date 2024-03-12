package product_variations_attributes_test

import (
	"log"
	"testing"

	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	product_variations "github.com/vongphachan/funny-store-backend/src/modules/product-variations"
	product_variations_attributes "github.com/vongphachan/funny-store-backend/src/modules/product-variations-attributes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReplicate(t *testing.T) {
	sourceProductId := primitive.NewObjectID()
	sourceAttributes := []product_attribute.ProductAttribute{}
	sourceAttributeGroups := []product_attribute_group.AttributeGroup{}
	sourceVariations := []product_variations.ProductVariation{}
	sourceVariationAttributes := []product_variations_attributes.ProductVariationAttribute{}

	{
		for i := 0; i < 2; i++ {
			attributeGroup := product_attribute_group.AttributeGroup{
				ID:        primitive.NewObjectID(),
				ProductID: sourceProductId,
			}

			sourceAttributeGroups = append(sourceAttributeGroups, attributeGroup)
		}

		for i := 0; i < 2; i++ {
			attribute := product_attribute.ProductAttribute{
				ID:               primitive.NewObjectID(),
				ProductID:        sourceProductId,
				AttributeGroupID: sourceAttributeGroups[i].ID,
			}

			sourceAttributes = append(sourceAttributes, attribute)
		}

		for i := 0; i < 2; i++ {
			variation := product_variations.ProductVariation{
				ID:        primitive.NewObjectID(),
				ProductID: sourceProductId,
			}

			sourceVariations = append(sourceVariations, variation)
		}

		for i := 0; i < 2; i++ {
			variationAttribute := product_variations_attributes.ProductVariationAttribute{
				ID:               primitive.NewObjectID(),
				AttributeGroupId: sourceAttributeGroups[i].ID,
				VariationId:      sourceVariations[i].ID,
				AttributeId:      sourceAttributes[i].ID,
				ProductId:        sourceProductId,
			}

			sourceVariationAttributes = append(sourceVariationAttributes, variationAttribute)
		}
	}

	newProductID := primitive.NewObjectID()
	newAttributes := product_attribute.Replicate(&newProductID, &sourceAttributes)
	newVariations := product_variations.Replicate(&newProductID, &sourceVariations)
	newAttributeGroups := product_attribute_group.Replicate(&newProductID, &sourceAttributeGroups)

	attributeMap := map[string]*primitive.ObjectID{}
	variationMap := map[string]*primitive.ObjectID{}
	attributeGroupMap := map[string]*primitive.ObjectID{}

	{
		for _, item := range *newAttributes {
			log.Println("newAttributes", item)
			attributeMap[item.OriginalID.String()] = &item.ID
		}

		for _, item := range *newVariations {
			variationMap[item.OriginalID.String()] = &item.ID
		}

		for _, item := range *newAttributeGroups {
			attributeGroupMap[item.OriginalID.String()] = &item.ID
		}
	}

	props := product_variations_attributes.Props_Replicate{
		ProductId:       &newProductID,
		SourceList:      &sourceVariationAttributes,
		Attributes:      &attributeMap,
		Variations:      &variationMap,
		AttributeGroups: &attributeGroupMap,
	}

	newVariationAttributes := product_variations_attributes.Replicate(&props)

	if len(*newVariationAttributes) != len(sourceVariationAttributes) {
		t.Errorf("Replicate() = %v, want %v", len(*newVariationAttributes), len(sourceVariationAttributes))
	}

	for i, item := range *newVariationAttributes {
		if item.OriginalId != sourceVariationAttributes[i].ID {
			t.Errorf("Replicate() = %v, want %v", item.OriginalId, sourceVariationAttributes[i].ID)
		}

		if item.ProductId != newProductID {
			t.Errorf("Replicate() = %v, want %v", item.ProductId, newProductID)
		} else if item.ProductId != sourceProductId {
			t.Errorf("Replicate() = %v, want %v", item.ProductId, sourceProductId)
		}

		if item.AttributeId != *attributeMap[sourceVariationAttributes[i].AttributeId.String()] {
			t.Errorf("Replicate() = %v, want %v", item.AttributeId, *attributeMap[sourceVariationAttributes[i].AttributeId.String()])
		}

		if item.VariationId != *variationMap[sourceVariationAttributes[i].VariationId.String()] {
			t.Errorf("Replicate() = %v, want %v", item.VariationId, *variationMap[sourceVariationAttributes[i].VariationId.String()])
		}

		if item.AttributeGroupId != *attributeGroupMap[sourceVariationAttributes[i].AttributeGroupId.String()] {
			t.Errorf("Replicate() = %v, want %v", item.AttributeGroupId, *attributeGroupMap[sourceVariationAttributes[i].AttributeGroupId.String()])
		}

		log.Println("newVariationAttributes", item)
	}
}
