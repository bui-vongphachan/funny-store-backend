package product

type Product struct {
	ID                  string   `json:"_id" bson:"_id"`
	Title               string   `json:"title" bson:"title"`
	Description         string   `json:"description" bson:"description"`
	PreviewImages       []string `json:"previewImages" bson:"previewImages"`
	Gallery             []string `json:"gallery" bson:"gallery"`
	Delete              bool     `json:"delete" bson:"delete"`
	HavingSingleVariant bool     `json:"havingSingleVariant" bson:"havingSingleVariant"`
	Image               string   `json:"image,omitempty" bson:"image,omitempty"`
	IsDraft             bool     `json:"isDraft" bson:"isDraft"`
}

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`

	// extends base type
	ID                  string   `json:"_id" bson:"_id"`
	Title               string   `json:"title" bson:"title"`
	Description         string   `json:"description" bson:"description"`
	PreviewImages       []string `json:"previewImages" bson:"previewImages"`
	Gallery             []string `json:"gallery" bson:"gallery"`
	Delete              bool     `json:"delete" bson:"delete"`
	HavingSingleVariant bool     `json:"havingSingleVariant" bson:"havingSingleVariant"`
	Image               string   `json:"image,omitempty" bson:"image,omitempty"`
	IsDraft             bool     `json:"isDraft" bson:"isDraft"`
}
