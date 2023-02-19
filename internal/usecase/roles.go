package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
)

// RolesUseCase -.
type RolesUseCase struct {
	repo RolesRepo
}

// NewRolesUseCase -.
func NewRolesUseCase(r RolesRepo) *RolesUseCase {
	return &RolesUseCase{repo: r}
}

// GetAllRoles -.
func (uc *RolesUseCase) GetAllRoles() ([]entity.Roles, error) {
	roles, err := uc.repo.GetRole()
	if err != nil {
		return nil, fmt.Errorf("RolesUseCase - GetAllRoles - uc.repo.GetRole: %w", err)
	}
	return roles, nil
}
