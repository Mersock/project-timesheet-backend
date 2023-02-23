package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// TimeEntriesUseCase -.
type TimeEntriesUseCase struct {
	repo TimeEntryRepo
}

// NewTimeEntriesUseCase -.
func NewTimeEntriesUseCase(r TimeEntryRepo) *TimeEntriesUseCase {
	return &TimeEntriesUseCase{repo: r}
}

// CreateTimeEntry -.
func (uc *TimeEntriesUseCase) CreateTimeEntry(req request.CreateTimeEntryReq) (int64, error) {
	var timeEntryID int64

	timeEntryID, err := uc.repo.Insert(req)
	if err != nil {
		return timeEntryID, fmt.Errorf("TimeEntriesUseCase - CreateTimeEntry - uc.repo.Insert: %w", err)
	}
	return timeEntryID, nil
}
