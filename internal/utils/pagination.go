package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// PaginationReq -.
type PaginationReq struct {
	Limit *int `form:"limit" binding:"omitempty,numeric,min=1"`
	Page  *int `form:"page" binding:"omitempty,numeric,min=1"`
}

// PaginationRes -.
type PaginationRes struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	LastPage int `json:"last_page"`
}

// GeneratePaginationFromRequest -.
func GeneratePaginationFromRequest(c *gin.Context) PaginationRes {
	limit := 10
	page := 1
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
		}
	}
	return PaginationRes{
		Limit: limit,
		Page:  page,
	}

}

// GetPageCount -.
func GetPageCount(total, limit int) int {
	return (total + limit - 1) / limit
}
