package response

// GetReportWorkTypeRes -.
type GetReportWorkTypeRes struct {
	Data GroupWorkTypeReport `json:"data"`
}

// GroupWorkTypeReport -.
type GroupWorkTypeReport struct {
	ProjectID   int               `json:"project_id"`
	ProjectName string            `json:"project_name"`
	WorkTypes   []WorkTypesReport `json:"worktypes"`
}

// WorkTypes -.
type WorkTypesReport struct {
	WorkTypeID   int    `json:"work_type_id"`
	WorkTypeName string `json:"work_type_name"`
	TotalSeconds int    `json:"total_seconds"`
	TotalTime    string `json:"total_time"`
}
