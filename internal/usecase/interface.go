package usecase

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

type (
	// Roles -.
	Roles interface {
		GetCount(req request.GetRolesReq) (int, error)
		GetAllRoles(req request.GetRolesReq) ([]entity.Roles, error)
		GetRoleByID(roleID int) (entity.Roles, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		Count(req request.GetRolesReq) (int, error)
		Select(req request.GetRolesReq) ([]entity.Roles, error)
		SelectById(roleID int) (entity.Roles, error)
	}
)
