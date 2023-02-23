package request

// GetTimeEntryReq -.
type GetTimeEntryReq struct {
	ProjectID string `form:"projectID" binding:"omitempty,numeric"`
}

// CreateTimeEntryReq -.
type CreateTimeEntryReq struct {
	StatusID      int    `form:"status_id" json:"status_id" binding:"required,numeric"`
	WorkTypeID    int    `form:"work_type_id" json:"work_type_id" binding:"required,numeric"`
	UserID        int    `form:"user_id" json:"user_id" binding:"required,numeric"`
	StartDateTime string `form:"start_date" json:"start_date" binding:"required,iso8601date"`
	EndDateTime   string `form:"end_date" json:"end_date" binding:"omitempty,iso8601date"`
}
