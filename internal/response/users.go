package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetUsersRes -.
type GetUsersRes struct {
	Users []entity.Users `json:"data"`
	Total int            `json:"total"`
	utils.PaginationRes
}

// GetUserByIDRes -.
type GetUserByIDRes struct {
	User entity.Users `json:"data"`
}
