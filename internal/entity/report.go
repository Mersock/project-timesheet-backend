package entity

// ReportWorkType -.
type ReportWorkType struct {
	WorkTypeID   *int    `json:"work_type_id"`
	WorkTypeName *string `json:"work_type_name"`
	ProjectName  *string `json:"project_name"`
	Total        *int    `json:"total"`
}
