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
	OwnerUserID int64     `form:"-" json:"-" binding:"required,numeric"`
	Members     []string  `form:"members" json:"members" binding:"omitempty,min=1"`
	WorkTypes   []string  `form:"work_types" json:"work_types" binding:"omitempty,min=1"`
}

// UpdateProjectReq -.
type UpdateProjectReq struct {
	ID              int                         `binding:"required,numeric"`
	Name            string                      `form:"name" json:"name" binding:"required"`
	AddWorkTypes    []string                    `form:"add_work_types" json:"add_work_types" binding:"omitempty,min=1"`
	DeleteWorkTypes []int                       `form:"delete_work_types" json:"delete_work_types" binding:"omitempty"`
	EditWorkTypes   []UpdateProjectWithWorkType `form:"edit_work_types" json:"edit_work_types" binding:"omitempty"`
}

type UpdateProjectWithWorkType struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

// UpdateProjectAddMoreMemberReq -.
type UpdateProjectAddMoreMemberReq struct {
	ID      int      `binding:"required,numeric"`
	Members []string `form:"members" json:"members" binding:"required,min=1"`
}

// DeleteProjectByReq -.
type DeleteProjectByReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// DeleteProjectMemberByReq -.
type DeleteProjectMemberByReq struct {
	ID     int `uri:"id" binding:"required,numeric"`
	UserID int `uri:"userID" binding:"required,numeric"`
}
