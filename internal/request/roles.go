package request

import "github.com/Mersock/project-timesheet-backend/internal/utils"

// RolesReq -.
type RolesReq struct {
	Name     string `form:"name" binding:"omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty"`
	SortType string `form:"sortType" json:"-" binding:"omitempty"`
	utils.PaginationReq
}
