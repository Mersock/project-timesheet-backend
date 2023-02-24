package usecase

import (
	"database/sql"
	"errors"
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

// GetTimeEntryByID -.
func (uc *TimeEntriesUseCase) GetTimeEntryByID(timeEntryID int) (entity.TimeEntryList, error) {
	var timeEntry entity.TimeEntryList
	timeEntry, err := uc.repo.SelectByID(timeEntryID)
	if err != nil {
		return timeEntry, fmt.Errorf("TimeEntriesUseCase - GetTimeEntryByID - uc.repo.SelectById: %w", err)
	}
	return timeEntry, nil
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

// UpdateTimeEntry -.
func (uc *TimeEntriesUseCase) UpdateTimeEntry(req request.UpdateTimeEntryReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectByID(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("TimeEntriesUseCase - UpdateTimeEntry - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Update(req)
	if err != nil {
		return rowAffected, fmt.Errorf("TimeEntriesUseCase - UpdateTimeEntry - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}

// DeleteTimeEntry -.
func (uc *TimeEntriesUseCase) DeleteTimeEntry(req request.DeleteTimeEntryReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectByID(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("TimeEntriesUseCase - DeleteRole - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Delete(req)
	if err != nil {
		return rowAffected, fmt.Errorf("TimeEntriesUseCase - DeleteRole - uc.repo.Delete: %w", err)
	}

	return rowAffected, nil
}
