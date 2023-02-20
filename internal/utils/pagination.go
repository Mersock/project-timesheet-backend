package utils

import (
	"github.com/gin-gonic/gin"
)

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
