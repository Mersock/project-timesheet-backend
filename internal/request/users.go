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
	SortBy    string `form:"sortBy" json:"-" binding:"omitempty,oneof=id email firstname lastname created_at updated_at role"`
	SortType  string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// GetUserByIDReq -.
type GetUserByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// DeleteUserReq -.
type DeleteUserReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
