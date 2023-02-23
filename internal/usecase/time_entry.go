package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
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

// GetCount -.
func (uc *TimeEntriesUseCase) GetCount(req request.GetTimeEntryReq) (int, error) {
	rows, err := uc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("TimeEntriesUseCase - GetCount - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllTimeEntries -.
func (uc *TimeEntriesUseCase) GetAllTimeEntries(req request.GetTimeEntryReq) ([]entity.TimeEntryList, error) {
	roles, err := uc.repo.Select(req)
	if err != nil {
		return nil, fmt.Errorf("TimeEntriesUseCase - GetAllRoles - uc.repo.Select: %w", err)
	}
	return roles, nil
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
