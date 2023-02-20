package utils

type Pagination struct {
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Total    string `json:"total"`
	SortBy   string `json:"-"`
	SortType string `json:"-"`
}
