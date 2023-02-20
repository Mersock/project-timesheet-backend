package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// RolesRes -.
type RolesRes struct {
	Roles []entity.Roles `json:"data"`
	Total int            `json:"total"`
	utils.PaginationRes
}
