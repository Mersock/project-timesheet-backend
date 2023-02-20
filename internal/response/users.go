package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetUsersRes -.
type GetUsersRes struct {
	Users []entity.Roles `json:"data"`
	Total int            `json:"total"`
	utils.PaginationRes
}
