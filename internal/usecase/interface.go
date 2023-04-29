package usecase

import (
	"database/sql"

	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
)

type (
	// Report -.
	Report interface {
		GetAllWorkType(req request.GetWorkTypeReportReq) (response.GroupWorkTypeReport, error)
		GetWorkTypeCount(req request.GetWorkTypeReportReq) (int, error)
	}

	// ReportRepo -.
	ReportRepo interface {
		CountWorkType(req request.GetWorkTypeReportReq) (int, error)
		SelectWorkType(req request.GetWorkTypeReportReq) ([]entity.ReportWorkType, error)
	}

	// TimeEntry -.
	TimeEntry interface {
		GetCount(req request.GetTimeEntryReq) (int, error)
		GetAllTimeEntries(req request.GetTimeEntryReq) ([]entity.TimeEntryList, error)
		GetTimeEntryByID(timeEntryID int) (entity.TimeEntryList, error)
		CreateTimeEntry(req request.CreateTimeEntryReq) (int64, error)
		UpdateTimeEntry(req request.UpdateTimeEntryReq) (int64, error)
		DeleteTimeEntry(req request.DeleteTimeEntryReq) (int64, error)
	}

	// TimeEntryRepo -.
	TimeEntryRepo interface {
		Count(req request.GetTimeEntryReq) (int, error)
		Select(req request.GetTimeEntryReq) ([]entity.TimeEntryList, error)
		Insert(req request.CreateTimeEntryReq) (int64, error)
		SelectByID(timeEntryID int) (entity.TimeEntryList, error)
		Update(req request.UpdateTimeEntryReq) (int64, error)
		Delete(req request.DeleteTimeEntryReq) (int64, error)
	}

	//Status -.
	Status interface {
		GetCount(req request.GetStatusReq) (int, error)
		GetAllStatus(req request.GetStatusReq) ([]entity.Status, error)
		GetStatusByID(roleID int) (entity.Status, error)
		CreateStatus(req request.CreateStatusReq) (int64, error)
		UpdateStatus(req request.UpdateStatusReq) (int64, error)
		DeleteStatus(req request.DeleteStatusReq) (int64, error)
	}

	//StatusRepo -.
	StatusRepo interface {
		Count(req request.GetStatusReq) (int, error)
		Select(req request.GetStatusReq) ([]entity.Status, error)
		SelectById(statusID int) (entity.Status, error)
		Insert(req request.CreateStatusReq) (int64, error)
		Update(req request.UpdateStatusReq) (int64, error)
		Delete(req request.DeleteStatusReq) (int64, error)
		ChkUniqueInsert(req request.CreateStatusReq) (int, error)
		ChkUniqueUpdate(req request.UpdateStatusReq) (int, error)
	}

	//WorkTypes -.
	WorkTypes interface {
		GetWorkTypeByID(workTypeID int) (entity.WorkTypes, error)
		GetWorkTypeByProject(projectID int) ([]entity.WorkTypes, error)
		CreateWorkType(req request.CreateWorkTypeReq) (int64, error)
		UpdateWorkType(req request.UpdateWorkTypeReq) (int64, error)
		DeleteWorkType(req request.DeleteWorkTypeReq) (int64, error)
	}

	//WorkTypeRepo -.
	WorkTypeRepo interface {
		InsertWithProject(tx *sql.Tx, req request.CreateWorkTypeReq) (*sql.Tx, int64, error)
		UpdateWithProject(tx *sql.Tx, req request.UpdateWorkTypeReq) (*sql.Tx, int64, error)
		DeleteWithProject(tx *sql.Tx, req request.DeleteWorkTypeReq) (*sql.Tx, int64, error)
		Insert(req request.CreateWorkTypeReq) (int64, error)
		SelectById(workTypeID int) (entity.WorkTypes, error)
		SelectByProjectId(projectID int) ([]entity.WorkTypes, error)
		Update(req request.UpdateWorkTypeReq) (int64, error)
		Delete(req request.DeleteWorkTypeReq) (int64, error)
	}

	// DutyRepo -.
	DutyRepo interface {
		InsertOwner(tx *sql.Tx, projectID, ownerUserID int64) (*sql.Tx, error)
		InsertMember(tx *sql.Tx, projectID int64, member string) (*sql.Tx, error)
		Delete(projectID int64, userID int64) (int64, error)
	}

	//Project -.
	Project interface {
		GetCount(req request.GetProjectsReq) (int, error)
		GetAllProjects(req request.GetProjectsReq) ([]entity.Projects, error)
		GetProjectsByID(req request.GetProjectByIDReq) (entity.Projects, error)
		GetProjectsByIDWithUserWorkType(req request.GetProjectByIDReq) (entity.ProjectWithSliceUser, error)
		CreateProject(req request.CreateProjectReq) (int64, error)
		UpdateProject(req request.UpdateProjectReq) (int64, error)
		UpdateProjectAddMoreMember(req request.UpdateProjectAddMoreMemberReq) error
		DeleteProject(req request.DeleteProjectByReq) (int64, error)
		DeleteProjectMember(req request.DeleteProjectMemberByReq) (int64, error)
	}

	//ProjectRepo -.
	ProjectRepo interface {
		BeginTx() (*sql.Tx, error)
		Count(req request.GetProjectsReq) (int, error)
		Select(req request.GetProjectsReq) ([]entity.Projects, error)
		SelectById(projectID int) (entity.Projects, error)
		SelectByIdWithUser(projectID int) ([]entity.ProjectsWithUser, error)
		Insert(tx *sql.Tx, req request.CreateProjectReq) (*sql.Tx, int64, error)
		Update(tx *sql.Tx, req request.UpdateProjectReq) (*sql.Tx, int64, error)
		Delete(req request.DeleteProjectByReq) (int64, error)
	}

	//Auth -.
	Auth interface {
		Signup(req request.SignUpReq) (int64, error)
		SignIn(req request.SignInReq) (response.SignInRes, error)
		RenewAccess(req request.RenewTokenReq) (response.RenewTokenRes, error)
	}

	//User -.
	User interface {
		GetCount(req request.GetUsersReq) (int, error)
		GetAllUsers(req request.GetUsersReq) ([]entity.Users, error)
		GetUserByID(userID int) (entity.Users, error)
		CreateUser(req request.CreateUserReq) (int64, error)
		UpdateUser(req request.UpdateUserReq) (int64, error)
		UpdateUserPassword(req request.UpdateUserPasswordReq) (int64, error)
		DeleteUser(req request.DeleteUserReq) (int64, error)
	}

	//UserRepo -.
	UserRepo interface {
		Count(req request.GetUsersReq) (int, error)
		Select(req request.GetUsersReq) ([]entity.Users, error)
		SelectById(userID int) (entity.Users, error)
		SelectByEmail(email string) (entity.Users, error)
		Insert(req request.CreateUserReq) (int64, error)
		Delete(req request.DeleteUserReq) (int64, error)
		Update(req request.UpdateUserReq) (int64, error)
		UpdatePassword(req request.UpdateUserPasswordReq) (int64, error)
		ChkUniqueUpdate(req request.UpdateUserReq) (int, error)
		ChkUniqueInsert(req request.CreateUserReq) (int, error)
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
