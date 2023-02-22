package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetProjectsRes -.
type GetProjectsRes struct {
	Projects []entity.Projects `json:"data"`
	Total    int               `json:"total"`
	utils.PaginationRes
}
