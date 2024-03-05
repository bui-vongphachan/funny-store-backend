package model

type Product struct {
	ID                  string   `json:"_id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	PreviewImages       []string `json:"previewImages"`
	Gallery             []string `json:"gallery"`
	Delete              bool     `json:"delete"`
	HavingSingleVariant bool     `json:"havingSingleVariant"`
	Image               string   `json:"image,omitempty"`
	IsDraft             bool     `json:"isDraft"`
}
