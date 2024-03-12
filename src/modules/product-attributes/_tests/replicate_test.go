package product_attribute_test

import (
	"testing"

	product_attribute "github.com/vongphachan/funny-store-backend/src/modules/product-attributes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReplicate(t *testing.T) {
	newProductID := primitive.NewObjectID()

	sourceProductId := primitive.NewObjectID()
	sourceList := []product_attribute.ProductAttribute{}

	for i := 0; i < 1; i++ {
		attributeGroup := product_attribute.ProductAttribute{
			ID:        primitive.NewObjectID(),
			Title:     "title",
			Image:     "image",
			ProductID: sourceProductId,
		}

		sourceList = append(sourceList, attributeGroup)
	}

	newList := product_attribute.Replicate(&newProductID, &sourceList)

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

		if (*newList)[i].OrgirinalID != sourceList[i].ID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].OrgirinalID, sourceList[i].ID)
		}

		if (*newList)[i].ProductID == sourceProductId {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ProductID, sourceProductId)
		}

		if (*newList)[i].ProductID != newProductID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ProductID, newProductID)
		}
	}
}
