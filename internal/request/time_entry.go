package request

// GetTimeEntryReq -.
type GetTimeEntryReq struct {
	ProjectID string `form:"projectID" binding:"omitempty,numeric"`
}

// CreateTimeEntryReq -.
type CreateTimeEntryReq struct {
	StatusID   int    `form:"status_id" json:"status_id" binding:"required,numeric"`
	WorkTypeID int    `form:"work_type_id" json:"work_type_id" binding:"required,numeric"`
	UserID     int64  `form:"-" json:"-" binding:"required,numeric"`
	StartDate  string `form:"start_date" json:"start_date" binding:"required,iso8601date"`
	EndDate    string `form:"end_date" json:"end_date" binding:"omitempty,iso8601date"`
}
