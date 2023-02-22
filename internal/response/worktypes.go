package response

import "github.com/Mersock/project-timesheet-backend/internal/entity"

// GetWorkTypeByIDRes -.
type GetWorkTypeByIDRes struct {
	WorkType entity.WorkTypes `json:"data"`
}
