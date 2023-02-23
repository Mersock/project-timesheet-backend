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

// GetProjectByIDReq -.
type GetProjectByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// CreateProjectReq -.
type CreateProjectReq struct {
	Name        string    `form:"name" json:"name" binding:"required,max=255"`
	Code        uuid.UUID `form:"-" json:"-" binding:"required"`
	UserOwnerID int64     `form:"-" json:"-" binding:"required,numeric"`
	Members     []*int64  `form:"members" json:"members" binding:"omitempty,min=1"`
	WorkTypes   []*string `form:"work_types" json:"work_types" binding:"omitempty,min=1"`
}

// UpdateProjectReq -.
type UpdateProjectReq struct {
	ID   int    `binding:"required,numeric"`
	Name string `form:"name" json:"name" binding:"required"`
}

// DeleteProjectReq -.
type DeleteProjectByReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
