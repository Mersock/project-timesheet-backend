package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetRolesRes -.
type GetRolesRes struct {
	Roles []entity.Roles `json:"data"`
	Total int            `json:"total"`
	utils.PaginationRes
}

// GetRoleByIDRes -.
type GetRoleByIDRes struct {
	Role entity.Roles `json:"data"`
}
