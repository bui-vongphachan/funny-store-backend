package swapProductVariation_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

// AttributeGroupLookUp represents the attribute group lookup structure.
type AttributeGroupLookUp struct {
	ID        string `json:"_id"`
	Title     string `json:"title"`
	IsPrimary bool   `json:"isPrimary"`
	ProductId string `json:"productId"`
	Delete    bool   `json:"delete"`
	ReuseFrom string `json:"reuseFrom"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

// AttributeLookUp represents the attribute lookup structure.
type AttributeLookUp struct {
	ID                string `json:"_id"`
	Title             string `json:"title"`
	Image             string `json:"image"`
	ProductId         string `json:"productId"`
	AttributeGroupId  string `json:"attributeGroupId"`
	Delete            bool   `json:"delete"`
	CloudflareImageId string `json:"cloudflareImageId"`
	ReuseFrom         string `json:"reuseFrom"`
}

// VariationAttribute represents the variation attribute structure.
type VariationAttribute struct {
	ID                   string               `json:"_id"`
	AttributeGroupId     string               `json:"attributeGroupId"`
	VariationId          string               `json:"variationId"`
	AttributeId          string               `json:"attributeId"`
	AttributeGroupLookUp AttributeGroupLookUp `json:"attributeGroupLookUp"`
	AttributeLookUp      AttributeLookUp      `json:"attributeLookUp"`
}

// Variant represents a single variant structure.
type Variant struct {
	IsReady             bool                 `json:"isReady"`
	VariationAttributes []VariationAttribute `json:"variationAttributes"`
	ID                  string               `json:"_id"`
	ProductId           string               `json:"productId"`
	Price               int                  `json:"price"`
	Stock               int                  `json:"stock"`
	Delete              bool                 `json:"delete"`
	IsSingleVariant     bool                 `json:"isSingleVariant"`
}

func TestExampleSwapProductVariation(t *testing.T) {
	files, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
		defer files.Close()
	}

	fileContent, err := io.ReadAll(files)
	if err != nil {
		panic(err)
	}

	var variations []Variant

	// Decode the JSON data into the struct
	err = json.Unmarshal(fileContent, &variations)
	if err != nil {
		log.Fatal(err)
	}

	for _, variation := range variations {
		for _, variationAttribute := range variation.VariationAttributes {
			fmt.Println(variationAttribute)
		}
	}
}

func TrackRemainingElements() {

}
