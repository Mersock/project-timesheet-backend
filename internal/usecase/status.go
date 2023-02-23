package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// StatusUseCase -.
type StatusUseCase struct {
	repo StatusRepo
}

// NewStatusUseCase -.
func NewStatusUseCase(r StatusRepo) *StatusUseCase {
	return &StatusUseCase{repo: r}
}

// GetCount -.
func (uc *StatusUseCase) GetCount(req request.GetStatusReq) (int, error) {
	rows, err := uc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("RolesUseCase - GetCount - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllStatus -.
func (uc *StatusUseCase) GetAllStatus(req request.GetStatusReq) ([]entity.Status, error) {
	status, err := uc.repo.Select(req)
	if err != nil {
		return nil, fmt.Errorf("StatusUseCase - GetAllStatus - uc.repo.Select: %w", err)
	}
	return status, nil
}

// GetStatusByID -.
func (uc *StatusUseCase) GetStatusByID(roleID int) (entity.Status, error) {
	var status entity.Status
	status, err := uc.repo.SelectById(roleID)
	if err != nil {
		return status, fmt.Errorf("StatusUseCase - GetStatusByID - uc.repo.SelectById: %w", err)
	}
	return status, nil
}

// CreateStatus -.
func (uc *StatusUseCase) CreateStatus(req request.CreateStatusReq) (int64, error) {
	var statusID int64

	count, err := uc.repo.ChkUniqueInsert(req)
	if err != nil {
		return statusID, fmt.Errorf("StatusUseCase - CreateStatus - uc.repo.ChkUniqueInsert: %w", err)
	}

	if count > 0 {
		return statusID, ErrDuplicateRow
	}

	statusID, err = uc.repo.Insert(req)
	if err != nil {
		return statusID, fmt.Errorf("StatusUseCase - CreateStatus - uc.repo.Insert: %w", err)
	}
	return statusID, nil
}

// UpdateStatus -.
func (uc *StatusUseCase) UpdateStatus(req request.UpdateStatusReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("StatusUseCase - UpdateStatus - uc.repo.SelectById: %w", err)
	}

	count, err := uc.repo.ChkUniqueUpdate(req)
	if err != nil {
		return rowAffected, fmt.Errorf("StatusUseCase - UpdateStatus - uc.repo.ChkUniqueUpdate: %w", err)
	}

	if count > 0 {
		return rowAffected, ErrDuplicateRow
	}

	rowAffected, err = uc.repo.Update(req)
	if err != nil {
		return rowAffected, fmt.Errorf("StatusUseCase - UpdateStatus - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}

// DeleteStatus -.
func (uc *StatusUseCase) DeleteStatus(req request.DeleteStatusReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("StatusUseCase - DeleteStatus - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Delete(req)
	if err != nil {
		return rowAffected, fmt.Errorf("StatusUseCase - DeleteStatus - uc.repo.Delete: %w", err)
	}

	return rowAffected, nil
}
