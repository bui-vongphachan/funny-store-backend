package product_attribute_test

import (
	"testing"

	product_variations "github.com/vongphachan/funny-store-backend/src/modules/product-variations"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReplicate(t *testing.T) {
	newProductID := primitive.NewObjectID()

	sourceProductId := primitive.NewObjectID()
	sourceList := []product_variations.ProductVariation{}

	for i := 0; i < 1; i++ {
		attributeGroup := product_variations.ProductVariation{
			ID:        primitive.NewObjectID(),
			ProductID: sourceProductId,
			Stock:     0,
			Price:     0,
		}

		sourceList = append(sourceList, attributeGroup)
	}

	newList := product_variations.Replicate(&newProductID, &sourceList)

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

		if (*newList)[i].ProductID != newProductID {
			t.Errorf("Replicate() = %v, want %v", (*newList)[i].ProductID, newProductID)
		}
	}
}
