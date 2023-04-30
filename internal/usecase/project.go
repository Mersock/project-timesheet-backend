package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// ProjectsUseCase -.
type ProjectsUseCase struct {
	repo         ProjectRepo
	dutyRepo     DutyRepo
	workTypeRepo WorkTypeRepo
}

// NewProjectsUseCase -.
func NewProjectsUseCase(r ProjectRepo, dutyRepo DutyRepo, workTypeRepo WorkTypeRepo) *ProjectsUseCase {
	return &ProjectsUseCase{
		repo:         r,
		dutyRepo:     dutyRepo,
		workTypeRepo: workTypeRepo,
	}
}

// GetCount -.
func (pc *ProjectsUseCase) GetCount(req request.GetProjectsReq) (int, error) {
	rows, err := pc.repo.Count(req)
	if err != nil {
		return rows, fmt.Errorf("ProjectsUseCase - GetCount - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllProjects -.
func (pc *ProjectsUseCase) GetAllProjects(req request.GetProjectsReq) ([]entity.Projects, error) {
	projects, err := pc.repo.Select(req)
	if err != nil {
		return nil, fmt.Errorf("ProjectsUseCase - GetAllRoles - uc.repo.Select: %w", err)
	}
	return projects, nil
}

// GetProjectsByID -.
func (pc *ProjectsUseCase) GetProjectsByID(req request.GetProjectByIDReq) (entity.Projects, error) {
	projects, err := pc.repo.SelectById(req.ID)
	if err != nil {
		return projects, fmt.Errorf("ProjectsUseCase - GetProjectsByID - uc.repo.SelectById: %w", err)
	}
	return projects, nil
}

// GetProjectsByIDWithUserWorkType -.
func (pc *ProjectsUseCase) GetProjectsByIDWithUserWorkType(req request.GetProjectByIDReq) (entity.ProjectWithSliceUser, error) {
	var projectWithUsers entity.ProjectWithSliceUser

	_, err := pc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return projectWithUsers, sql.ErrNoRows
		}
		return projectWithUsers, fmt.Errorf("ProjectsUseCase - UpdateProject - uc.repo.SelectById: %w", err)
	}

	selectProject, err := pc.repo.SelectByIdWithUser(req.ID)
	if err != nil {
		return projectWithUsers, fmt.Errorf("ProjectsUseCase - SelectByIdWithUser - uc.repo.SelectById: %w", err)
	}
	selectWorktypes, err := pc.workTypeRepo.SelectByProjectId(req.ID)
	if err != nil {
		return projectWithUsers, fmt.Errorf("ProjectsUseCase - SelectByIdWithUser - uc.repo.SelectById: %w", err)
	}

	projectWithUsers = pc.mappingProjectsByIDWithUser(selectProject)
	projectWithUsers.WorkTypes = selectWorktypes

	return projectWithUsers, nil
}

// GetProjectsByCodeWithUserWorkType -.
func (pc *ProjectsUseCase) GetProjectsByCode(req request.GetProjectByCodeReq) (entity.ProjectWithSliceUser, error) {
	var projectWithUsers entity.ProjectWithSliceUser

	_, err := pc.repo.SelectByCode(req.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return projectWithUsers, sql.ErrNoRows
		}
		return projectWithUsers, fmt.Errorf("ProjectsUseCase - UpdateProject - uc.repo.SelectById: %w", err)
	}

	selectProject, err := pc.repo.SelectByCodeWithUser(req.Code)
	if err != nil {
		return projectWithUsers, fmt.Errorf("ProjectsUseCase - SelectByIdWithUser - uc.repo.SelectById: %w", err)
	}

	projectWithUsers = pc.mappingProjectsByIDWithUser(selectProject)

	return projectWithUsers, nil
}

// CreateProject -.
func (pc *ProjectsUseCase) CreateProject(req request.CreateProjectReq) (int64, error) {
	var projectID int64

	tx, err := pc.repo.BeginTx()
	if err != nil {
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.BeginTx: %w", err)
	}

	tx, projectID, err = pc.repo.Insert(tx, req)
	if err != nil {
		tx.Rollback()
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.Insert: %w", err)
	}

	//create by owner
	tx, err = pc.dutyRepo.InsertOwner(tx, projectID, req.OwnerUserID)
	if err != nil {
		tx.Rollback()
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertOwner - owner: %w", err)
	}

	//add member to project
	if len(req.Members) != 0 {
		for _, member := range req.Members {
			tx, err = pc.dutyRepo.InsertMember(tx, projectID, member)
			if err != nil {
				tx.Rollback()
				return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties - member: %w", err)
			}
		}
	}

	//add work type to project
	if len(req.WorkTypes) != 0 {
		for _, name := range req.WorkTypes {
			reqWorkType := request.CreateWorkTypeReq{
				ProjectID: projectID,
				Name:      name,
			}

			tx, _, err = pc.workTypeRepo.InsertWithProject(tx, reqWorkType)
			if err != nil {
				tx.Rollback()
				return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties - member: %w", err)
			}
		}
	}

	tx.Commit()

	return projectID, nil
}

// UpdateProject -.
func (pc *ProjectsUseCase) UpdateProject(req request.UpdateProjectReq) (int64, error) {
	var rowAffected int64

	_, err := pc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("ProjectsUseCase - UpdateProject - uc.repo.SelectById: %w", err)
	}

	tx, err := pc.repo.BeginTx()
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectsUseCase - UpdateProject - pc.repo.BeginTx: %w", err)
	}

	tx, rowAffected, err = pc.repo.Update(tx, req)
	if err != nil {
		tx.Rollback()
		return rowAffected, fmt.Errorf("ProjectsUseCase - UpdateProject - uc.repo.Update: %w", err)
	}

	//add member to project
	if len(req.Members) != 0 {
		for _, member := range req.Members {
			tx, err = pc.dutyRepo.InsertMember(tx, int64(req.ID), member)
			if err != nil {
				tx.Rollback()
				return rowAffected, fmt.Errorf("ProjectsUseCase - UpdateProject - pc.repo.InsertMember - member: %w", err)
			}
		}
	}

	//add work type to project
	if len(req.AddWorkTypes) != 0 {
		for _, name := range req.AddWorkTypes {
			reqWorkType := request.CreateWorkTypeReq{
				Name:      name,
				ProjectID: int64(req.ID),
			}

			tx, _, err = pc.workTypeRepo.InsertWithProject(tx, reqWorkType)
			if err != nil {
				tx.Rollback()
				return rowAffected, fmt.Errorf("ProjectsUseCase - CreateProject - pc.workTypeRepo.InsertWithProject - member: %w", err)
			}
		}
	}

	//update work to project
	if len(req.EditWorkTypes) != 0 {
		for _, obj := range req.EditWorkTypes {
			reqWorkType := request.UpdateWorkTypeReq{
				Name: obj.Name,
				ID:   obj.ID,
			}

			tx, _, err = pc.workTypeRepo.UpdateWithProject(tx, reqWorkType)
			if err != nil {
				tx.Rollback()
				return rowAffected, fmt.Errorf("ProjectsUseCase - CreateProject - pc.workTypeRepo.UpdateWithProject - member: %w", err)
			}
		}
	}

	//delete worktype
	if len(req.DeleteWorkTypes) != 0 {
		for _, id := range req.DeleteWorkTypes {
			reqWorkType := request.DeleteWorkTypeReq{
				ID: id,
			}
			tx, _, err = pc.workTypeRepo.DeleteWithProject(tx, reqWorkType)
			if err != nil {
				tx.Rollback()
				return rowAffected, fmt.Errorf("ProjectsUseCase - CreateProject - pc.workTypeRepo.UpdateWithProject - member: %w", err)
			}
		}
	}

	tx.Commit()

	return rowAffected, nil
}

// UpdateProjectAddMoreMember -.
func (pc *ProjectsUseCase) UpdateProjectAddMoreMember(req request.UpdateProjectAddMoreMemberReq) error {

	_, err := pc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		return fmt.Errorf("ProjectsUseCase - UpdateProjectAddMoreMember - uc.repo.SelectById: %w", err)
	}

	tx, err := pc.repo.BeginTx()
	if err != nil {
		return fmt.Errorf("ProjectsUseCase - UpdateProjectAddMoreMember - pc.repo.BeginTx: %w", err)
	}

	//add member to project
	if req.Members != nil {
		for _, member := range req.Members {
			tx, err = pc.dutyRepo.InsertMember(tx, int64(req.ID), member)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("ProjectsUseCase - UpdateProjectAddMoreMember - pc.repo.Insert - member: %w", err)
			}
		}
	}

	tx.Commit()

	return nil
}

// DeleteProjectMember -.
func (pc *ProjectsUseCase) DeleteProjectMember(req request.DeleteProjectMemberByReq) (int64, error) {
	var rowAffected int64

	rowAffected, err := pc.dutyRepo.Delete(int64(req.ID), int64(req.UserID))
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectsUseCase - DeleteProjectMember - pc.dutyRepo.Delete: %w", err)
	}

	return rowAffected, nil
}

// DeleteProject -.
func (pc *ProjectsUseCase) DeleteProject(req request.DeleteProjectByReq) (int64, error) {
	var rowAffected int64

	_, err := pc.repo.SelectById(req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return rowAffected, sql.ErrNoRows
		}
		return rowAffected, fmt.Errorf("ProjectsUseCase - DeleteRole - uc.repo.SelectById: %w", err)
	}

	rowAffected, err = pc.repo.Delete(req)
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectsUseCase - DeleteRole - uc.repo.Delete: %w", err)
	}

	return rowAffected, nil
}

// mappingProjectsByIDWithUser -.
func (pc *ProjectsUseCase) mappingProjectsByIDWithUser(selectProject []entity.ProjectsWithUser) entity.ProjectWithSliceUser {
	var projectWithUsers entity.ProjectWithSliceUser
	var project entity.Projects
	var users []entity.UsersInProject

	for _, v := range selectProject {
		project = entity.Projects{
			ID:       v.ID,
			Name:     v.Name,
			Code:     v.Code,
			CreateAt: v.CreateAt,
			UpdateAt: v.UpdateAt,
		}
		user := entity.UsersInProject{
			UserID:    *v.UserID,
			Email:     *v.Email,
			Firstname: *v.Firstname,
			Lastname:  *v.Lastname,
			Role:      *v.Role,
		}
		users = append(users, user)
	}

	projectWithUsers = entity.ProjectWithSliceUser{
		Projects: project,
		Users:    users,
	}

	return projectWithUsers
}
