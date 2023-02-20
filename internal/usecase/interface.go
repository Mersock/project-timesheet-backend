package usecase

import "github.com/Mersock/project-timesheet-backend/internal/entity"

type (
	// Roles -.
	Roles interface {
		GetRowsRoles() (int, error)
		GetAllRoles() ([]entity.Roles, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		GetRole() ([]entity.Roles, error)
		GetRows() (int, error)
	}
)