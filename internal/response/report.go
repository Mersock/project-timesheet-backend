package response

import "github.com/Mersock/project-timesheet-backend/internal/entity"

type GetReportWorkTypeRes struct {
	WorkTypeReport []entity.ReportWorkType `json:"data"`
}
