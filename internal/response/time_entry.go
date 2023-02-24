package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
)

// GetTimeEntriesRes -.
type GetTimeEntriesRes struct {
	TimeEntries []entity.TimeEntryList `json:"data"`
	Total       int                    `json:"total"`
	utils.PaginationRes
}

// GetTimeEntryByIDRes -.
type GetTimeEntryByIDRes struct {
	TimeEntry entity.TimeEntryList `json:"data"`
}
