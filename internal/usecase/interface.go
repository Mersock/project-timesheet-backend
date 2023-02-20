package usecase

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

type (
	// Roles -.
	Roles interface {
		GetCount(req request.RolesReq) (int, error)
		GetAllRoles(req request.RolesReq) ([]entity.Roles, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		Count(req request.RolesReq) (int, error)
		Select(req request.RolesReq) ([]entity.Roles, error)
	}
)
