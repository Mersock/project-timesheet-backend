package request

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetStatusReq -.
type GetStatusReq struct {
	Name     string `form:"name" binding:"omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty,oneof=id name created_at updated_at"`
	SortType string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// GetStatusByIDReq -.
type GetStatusByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// CreateStatusReq -.
type CreateStatusReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

// UpdateStatusReq -.
type UpdateStatusReq struct {
	ID   int    `binding:"required,numeric"`
	Name string `form:"name" json:"name" binding:"required"`
}

// DeleteStatusReq -.
type DeleteStatusReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
