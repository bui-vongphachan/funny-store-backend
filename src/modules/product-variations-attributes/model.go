package product_variations_attributes

type ProductVariationAttribute struct {
	ID               string `json:"_id" bson:"_id"`
	AttributeGroupId string `json:"attributeGroupId" bson:"attributeGroupId"`
	VariationId      string `json:"variationId" bson:"variationId"`
	AttributeId      string `json:"attributeId" bson:"attributeId"`
}
