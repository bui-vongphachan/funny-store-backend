package utils

type PaginationQuery struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}
