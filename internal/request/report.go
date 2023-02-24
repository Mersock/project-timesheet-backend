package request

// GetWorkTypeReportReq -.
type GetWorkTypeReportReq struct {
	ProjectCode string `form:"project_code" json:"project_code" binding:"required"`
}
