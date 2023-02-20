package utils

// Pagination -.
type Pagination struct {
	Limit    string `form:"limit" json:"limit" binding:"numeric,omitempty"`
	Page     string `form:"page" json:"page" binding:"numeric,omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty"`
	SortType string `form:"sortType" json:"-" binding:"omitempty"`
	Total    string `json:"total"`
}
