package product_attribute_group

const jsonID = "_id"
const jsonTitle = "title"
const jsonProductID = "productId"

type AttributeGroup struct {
	ID        string `json:"_id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	IsPrimary bool   `json:"isPrimary" bson:"isPrimary"`
	ProductID string `json:"productId" bson:"productId"`
	Delete    bool   `json:"delete" bson:"delete"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID        string `json:"_id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	IsPrimary bool   `json:"isPrimary" bson:"isPrimary"`
	ProductID string `json:"productId" bson:"productId"`
	Delete    bool   `json:"delete" bson:"delete"`
}
