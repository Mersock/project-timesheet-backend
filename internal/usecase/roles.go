package usecase

import (
	"errors"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

var (
	errDuplicateRow = errors.New("duplicate")
)

// RolesUseCase -.
type RolesUseCase struct {
	repo RolesRepo
}

// NewRolesUseCase -.
func NewRolesUseCase(r RolesRepo) *RolesUseCase {
	return &RolesUseCase{repo: r}
}

// GetCount -.
func (uc *RolesUseCase) GetCount(req request.GetRolesReq) (int, error) {
	rows, err := uc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("RolesUseCase - GetCount - uc.repo.GetRows: %w", err)
	}
	return rows, nil
}

// GetAllRoles -.
func (uc *RolesUseCase) GetAllRoles(req request.GetRolesReq) ([]entity.Roles, error) {
	roles, err := uc.repo.Select(req)
	if err != nil {
		return nil, fmt.Errorf("RolesUseCase - GetAllRoles - uc.repo.GetRole: %w", err)
	}
	return roles, nil
}

// GetRoleByID -.
func (uc *RolesUseCase) GetRoleByID(roleID int) (entity.Roles, error) {
	var role entity.Roles
	role, err := uc.repo.SelectById(roleID)
	if err != nil {
		return role, fmt.Errorf("RolesUseCase - GetRoleByID - uc.repo.SelectById: %w", err)
	}
	return role, nil
}

// CreateRole -.
func (uc *RolesUseCase) CreateRole(req request.CreateRoleReq) (int64, error) {
	var roleID int64

	count, err := uc.repo.ChkUniqueInsert(req)
	if err != nil {
		return roleID, fmt.Errorf("RolesUseCase - CreateRole - uc.repo.ChkUniqueInsert: %w", err)
	}

	if count > 0 {
		return roleID, errDuplicateRow
	}

	roleID, err = uc.repo.Insert(req)
	if err != nil {
		return roleID, fmt.Errorf("RolesUseCase - CreateRole - uc.repo.Insert: %w", err)
	}
	return roleID, nil
}

// UpdateRole -.
func (uc *RolesUseCase) UpdateRole(req request.UpdateRoleReq) (int64, error) {
	var rowAffected int64

	count, err := uc.repo.ChkUniqueUpdate(req)
	if err != nil {
		return rowAffected, fmt.Errorf("RolesUseCase - UpdateRole - uc.repo.ChkUniqueUpdate: %w", err)
	}

	if count > 0 {
		return rowAffected, errDuplicateRow
	}

	rowAffected, err = uc.repo.Update(req)
	if err != nil {
		return rowAffected, fmt.Errorf("RolesUseCase - UpdateRole - uc.repo.Update: %w", err)
	}
	return rowAffected, nil
}
