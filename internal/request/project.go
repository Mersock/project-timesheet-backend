package request

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/google/uuid"
)

// GetProjectsReq -.
type GetProjectsReq struct {
	Name     string `form:"name" binding:"omitempty"`
	Code     string `form:"code" binding:"omitempty"`
	SortBy   string `form:"sortBy" json:"-" binding:"omitempty,oneof=code name"`
	SortType string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// CreateProjectReq -.
type CreateProjectReq struct {
	Name        string    `form:"name" json:"name" binding:"required,max=255"`
	Code        uuid.UUID `form:"-" json:"-" binding:"required"`
	UserOwnerID int64     `form:"-" json:"-" binding:"required,numeric"`
}
