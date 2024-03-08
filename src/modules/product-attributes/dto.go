package product_attribute

type ProductAttribute struct {
	ID                string `json:"_id" bson:"_id"`
	Title             string `json:"title" bson:"title"`
	Image             string `json:"image,omitempty" bson:"image,omitempty"`
	ProductID         string `json:"productId" bson:"productId"`
	AttributeGroupID  string `json:"attributeGroupId" bson:"attributeGroupId"`
	Delete            bool   `json:"delete" bson:"delete"`
	CloudflareImageID string `json:"cloudflareImageId,omitempty" bson:"cloudflareImageId,omitempty"`
	ReuseFrom         string `json:"reuseFrom,omitempty" bson:"reuseFrom,omitempty"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID                string `json:"_id" bson:"_id"`
	Title             string `json:"title" bson:"title"`
	Image             string `json:"image,omitempty" bson:"image,omitempty"`
	ProductID         string `json:"productId" bson:"productId"`
	AttributeGroupID  string `json:"attributeGroupId" bson:"attributeGroupId"`
	Delete            bool   `json:"delete" bson:"delete"`
	CloudflareImageID string `json:"cloudflareImageId,omitempty" bson:"cloudflareImageId,omitempty"`
	ReuseFrom         string `json:"reuseFrom,omitempty" bson:"reuseFrom,omitempty"`
}
