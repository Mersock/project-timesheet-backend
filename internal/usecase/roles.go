package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// RolesUseCase -.
type RolesUseCase struct {
	repo RolesRepo
}

// NewRolesUseCase -.
func NewRolesUseCase(r RolesRepo) *RolesUseCase {
	return &RolesUseCase{repo: r}
}

// GetRowsRoles -.
func (uc *RolesUseCase) GetRowsRoles(req request.RolesReq) (int, error) {
	rows, err := uc.repo.GetRows(req)
	if err != nil {
		return rows, fmt.Errorf("RolesUseCase - GetRowsRoles - uc.repo.GetRows: %w", err)
	}
	return rows, nil
}

// GetAllRoles -.
func (uc *RolesUseCase) GetAllRoles(req request.RolesReq) ([]entity.Roles, error) {
	roles, err := uc.repo.GetRole(req)
	if err != nil {
		return nil, fmt.Errorf("RolesUseCase - GetAllRoles - uc.repo.GetRole: %w", err)
	}
	return roles, nil
}
