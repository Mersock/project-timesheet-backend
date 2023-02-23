package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// WorkTypesUseCase -.
type WorkTypesUseCase struct {
	repo WorkTypeRepo
}

// NewWorkTypesUseCase -.
func NewWorkTypesUseCase(r WorkTypeRepo) *WorkTypesUseCase {
	return &WorkTypesUseCase{repo: r}
}

// GetWorkTypeByID -.
func (uc *WorkTypesUseCase) GetWorkTypeByID(workTypeID int) (entity.WorkTypes, error) {
	var workType entity.WorkTypes
	workType, err := uc.repo.SelectById(workTypeID)
	if err != nil {
		return workType, fmt.Errorf("WorkTypesUseCase - GetWorkTypeByID - uc.repo.SelectById: %w", err)
	}
	return workType, nil
}

// GetWorkTypeByProject -.
func (uc *WorkTypesUseCase) GetWorkTypeByProject(projectID int) ([]entity.WorkTypes, error) {
	var workTypes []entity.WorkTypes
	workTypes, err := uc.repo.SelectByProjectId(projectID)
	if err != nil {
		return workTypes, fmt.Errorf("WorkTypesUseCase - GetRoleByID - uc.repo.SelectByProjectId: %w", err)
	}
	return workTypes, nil
}

// UpdateWorkType -.
func (uc *WorkTypesUseCase) UpdateWorkType(req request.UpdateWorkTypeReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("WorkTypesUseCase - UpdateWorkType - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Update(req)
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesUseCase - UpdateWorkType - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}

// DeleteWorkType -.
func (uc *WorkTypesUseCase) DeleteWorkType(req request.DeleteWorkTypeReq) (int64, error) {
	var rowAffected int64

	_, err := uc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("WorkTypesUseCase - DeleteWorkType - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = uc.repo.Delete(req)
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesUseCase - DeleteWorkType - uc.repo.Delete: %w", err)
	}

	return rowAffected, nil
}
