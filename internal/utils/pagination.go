package utils

import (
	"github.com/gin-gonic/gin"
)

// Pagination -.
type Pagination struct {
	Limit    string `form:"limit" json:"limit" binding:"numeric,omitempty"`
	Page     string `form:"page" json:"page" binding:"numeric,omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty"`
	SortType string `form:"sortType" json:"-" binding:"omitempty"`
	Total    string `json:"total"`
}

// GeneratePaginationFromRequest -.
func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	limit := "10"
	page := "1"
	sortBy := ""
	sortType := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit = queryValue
			break
		case "page":
			page = queryValue
			break
		case "sortBy":
			sortBy = queryValue
			break
		case "sortType":
			sortType = queryValue
			break
		}
	}
	return Pagination{
		Limit:    limit,
		Page:     page,
		SortBy:   sortBy,
		SortType: sortType,
	}

}
