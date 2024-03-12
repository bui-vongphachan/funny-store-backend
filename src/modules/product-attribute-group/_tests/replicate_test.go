package product_attribute_group_test

import (
	"testing"

	product_attribute_group "github.com/vongphachan/funny-store-backend/src/modules/product-attribute-group"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReplicate(t *testing.T) {
	newProductID := primitive.NewObjectID()

	sourceProductId := primitive.NewObjectID()
	sourceList := []product_attribute_group.AttributeGroup{}

	for i := 0; i < 1; i++ {
		attributeGroup := product_attribute_group.AttributeGroup{
			ID:        primitive.NewObjectID(),
			Title:     "title",
			IsPrimary: true,
			ProductID: sourceProductId,
			Delete:    false,
		}

		sourceList = append(sourceList, attributeGroup)
	}

	props := product_attribute_group.Props_Relicate{
		DB:             nil,
		NewProductID:   &newProductID,
		SourceList:     &sourceList,
		SessionContext: nil,
	}

	newList := product_attribute_group.Replicate(&props)

	if len(*newList) != len(sourceList) {
		t.Errorf("Replicate() = %v, want %v", len(*newList), len(sourceList))
	}

	for i := 0; i < len(*newList); i++ {
		if (*newList)[i].ID == sourceList[i].ID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ID, sourceList[i].ID)
		}

		if (*newList)[i].ProductID != newProductID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ProductID, newProductID)
		}

		if (*newList)[i].OriginalID != sourceList[i].ID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].OriginalID, sourceList[i].ID)
		}

		if (*newList)[i].ProductID == sourceProductId {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ProductID, sourceProductId)
		}
	}
}
