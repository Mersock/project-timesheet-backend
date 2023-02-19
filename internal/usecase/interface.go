package usecase

import "github.com/Mersock/project-timesheet-backend/internal/entity"

type (
	// Roles -.
	Roles interface {
		GetAllRoles() ([]entity.Roles, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		GetRole() ([]entity.Roles, error)
	}
)
