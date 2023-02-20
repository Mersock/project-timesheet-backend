package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GeneratePaginationFromRequest -.
func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	limit := 10
	page := 1
	sortBy := ""
	sortType := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
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
