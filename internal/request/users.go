package request

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetAllUsersReq -.
type GetAllUsersReq struct {
	Firstname string `form:"firstname" binding:"omitempty"`
	Lastname  string `form:"lastname" binding:"omitempty"`
	Email     string `form:"email" binding:"omitempty"`
	Role      string `form:"role" binding:"omitempty"`
	SortBy    string `form:"sortBy" json:"-" binding:"omitempty,oneof=id email firstname lastname created_at updated_at role"`
	SortType  string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// CreateUserReq -.
type CreateUserReq struct {
	Email     string `form:"email" json:"email" binding:"required,email,max=255"`
	Password  string `form:"password" json:"password" binding:"required,min=6"`
	Firstname string `form:"firstname" json:"firstname" binding:"required,max=255"`
	Lastname  string `form:"lastname" json:"lastname" binding:"required,max=255"`
	RoleID    int
}

// GetUserReq -.
type GetUserReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// DeleteUserReq -.
type DeleteUserReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}