package request

import "github.com/Mersock/project-timesheet-backend/internal/utils"

// GetTimeEntryReq -.
type GetTimeEntryReq struct {
	ProjectName string `form:"projectName" binding:"omitempty"`
	Email       string `form:"email" binding:"omitempty"`
	Firstname   string `form:"firstname" binding:"omitempty"`
	Lastname    string `form:"lastname" binding:"omitempty"`
	Status      string `form:"status" binding:"omitempty"`
	SortBy      string `form:"sortBy" json:"-" binding:"omitempty,oneof=id created_at updated_at"`
	SortType    string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// CreateTimeEntryReq -.
type CreateTimeEntryReq struct {
	StatusID   int    `form:"status_id" json:"status_id" binding:"required,numeric"`
	WorkTypeID int    `form:"work_type_id" json:"work_type_id" binding:"required,numeric"`
	UserID     int64  `form:"-" json:"-" binding:"required,numeric"`
	StartDate  string `form:"start_date" json:"start_date" binding:"required,iso8601date"`
	EndDate    string `form:"end_date" json:"end_date" binding:"omitempty,iso8601date"`
}
