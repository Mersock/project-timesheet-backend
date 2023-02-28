package entity

// ReportWorkType -.
type ReportWorkType struct {
	ProjectID    *int    `json:"project_id"`
	ProjectName  *string `json:"project_name"`
	WorkTypeID   *int    `json:"work_type_id"`
	WorkTypeName *string `json:"work_type_name"`
	TotalSeconds *int    `json:"total_seconds"`
	TotalTime    *string `json:"total_time"`
}
