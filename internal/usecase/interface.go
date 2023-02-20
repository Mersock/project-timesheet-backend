package usecase

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

type (
	// Roles -.
	Roles interface {
		GetRowsRoles(req request.RolesReq) (int, error)
		GetAllRoles(req request.RolesReq) ([]entity.Roles, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		GetRole(req request.RolesReq) ([]entity.Roles, error)
		GetRows(req request.RolesReq) (int, error)
	}
)
