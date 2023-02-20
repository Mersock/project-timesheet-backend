package request

import "github.com/Mersock/project-timesheet-backend/internal/utils"

// RolesReq -.
type RolesReq struct {
	Name string `form:"name" binding:"omitempty"`
	utils.PaginationReq
}
