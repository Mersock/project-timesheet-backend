package usecase

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

type (
	//User -.
	User interface {
		GetCount(req request.GetUsersReq) (int, error)
		GetAllUsers(req request.GetUsersReq) ([]entity.Users, error)
	}

	//UserRepo -.
	UserRepo interface {
		Count(req request.GetUsersReq) (int, error)
		Select(req request.GetUsersReq) ([]entity.Users, error)
	}

	// Roles -.
	Roles interface {
		GetCount(req request.GetRolesReq) (int, error)
		GetAllRoles(req request.GetRolesReq) ([]entity.Roles, error)
		GetRoleByID(roleID int) (entity.Roles, error)
		CreateRole(req request.CreateRoleReq) (int64, error)
		UpdateRole(req request.UpdateRoleReq) (int64, error)
		DeleteRole(req request.DeleteRoleReq) (int64, error)
	}

	//RolesRepo -.
	RolesRepo interface {
		Count(req request.GetRolesReq) (int, error)
		Select(req request.GetRolesReq) ([]entity.Roles, error)
		SelectById(roleID int) (entity.Roles, error)
		Insert(req request.CreateRoleReq) (int64, error)
		Update(req request.UpdateRoleReq) (int64, error)
		Delete(req request.DeleteRoleReq) (int64, error)
		ChkUniqueInsert(req request.CreateRoleReq) (int, error)
		ChkUniqueUpdate(req request.UpdateRoleReq) (int, error)
	}
)
