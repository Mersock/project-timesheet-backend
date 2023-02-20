package request

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetUsersReq -.
type GetUsersReq struct {
	Firstname string `form:"firstname" binding:"omitempty"`
	Lastname  string `form:"lastname" binding:"omitempty"`
	Email     string `form:"email" binding:"omitempty"`
	Role      string `form:"role" binding:"omitempty"`
	SortBy    string `form:"sortBy" json:"-" binding:"omitempty"`
	SortType  string `form:"sortType" json:"-" binding:"omitempty"`
	utils.PaginationReq
}
