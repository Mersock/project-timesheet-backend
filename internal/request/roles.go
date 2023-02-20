package request

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetRolesReq -.
type GetRolesReq struct {
	Name     string `form:"name" binding:"omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty"`
	SortType string `form:"sortType" json:"-" binding:"omitempty"`
	utils.PaginationReq
}

// GetRoleByIDReq -.
type GetRoleByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// CreateRoleReq -.
type CreateRoleReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// UpdateRoleReq -.
type UpdateRoleReq struct {
	ID   int    `binding:"required,numeric"`
	Name string `form:"name" json:"name" binding:"required"`
}

// DeleteRoleReq -.
type DeleteRoleReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
