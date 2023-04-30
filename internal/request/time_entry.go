package request

import "github.com/Mersock/project-timesheet-backend/internal/utils"

// GetTimeEntryReq -.
type GetTimeEntryReq struct {
	ProjectCode string `form:"projectName" binding:"omitempty"`
	Email       string `form:"email" binding:"omitempty"`
	Firstname   string `form:"firstname" binding:"omitempty"`
	Lastname    string `form:"lastname" binding:"omitempty"`
	Status      string `form:"status" binding:"omitempty"`
	SortBy      string `form:"sortBy" json:"-" binding:"omitempty,oneof=id created_at updated_at"`
	SortType    string `form:"sortType" json:"-" binding:"omitempty,oneof=asc desc"`
	utils.PaginationReq
}

// GetTimeEntryByIDReq -.
type GetTimeEntryByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// CreateTimeEntryReq -.
type CreateTimeEntryReq struct {
	StatusID   int    `form:"status_id" json:"status_id" binding:"required,numeric"`
	WorkTypeID int    `form:"work_type_id" json:"work_type_id" binding:"required,numeric"`
	UserID     int64  `form:"-" json:"-" binding:"required,numeric"`
	StartTime  string `form:"start_time" json:"start_time" binding:"required,iso8601date"`
	EndTime    string `form:"end_time" json:"end_time" binding:"omitempty,iso8601date"`
}

// UpdateTimeEntryReq -.
type UpdateTimeEntryReq struct {
	ID         int    `binding:"required,numeric"`
	StatusID   int    `form:"status_id" json:"status_id" binding:"required,numeric"`
	WorkTypeID int    `form:"work_type_id" json:"work_type_id" binding:"required,numeric"`
	UserID     int64  `form:"-" json:"-" binding:"required,numeric"`
	StartTime  string `form:"start_time" json:"start_time" binding:"required,iso8601date"`
	EndTime    string `form:"end_time" json:"end_time" binding:"omitempty,iso8601date"`
}

// DeleteTimeEntryReq -.
type DeleteTimeEntryReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
