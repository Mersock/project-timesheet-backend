package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetStatusRes -.
type GetStatusRes struct {
	Status []entity.Status `json:"data"`
	Total  int             `json:"total"`
	utils.PaginationRes
}

// GetStatusByIDRes -.
type GetStatusByIDRes struct {
	Status entity.Status `json:"data"`
}
