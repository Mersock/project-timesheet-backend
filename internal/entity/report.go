package entity

// ReportWorkType -.
type ReportWorkType struct {
	WorkTypeID   *int    `json:"work_type_id"`
	WorkTypeName *string `json:"work_type_name"`
	ProjectName  *string `json:"project_name"`
	TotalSeconds *int    `json:"total_seconds"`
	TotalTime    *string `json:"total_time"`
}
