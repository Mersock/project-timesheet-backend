package usecase

import (
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

// GetProjectsByIDWithUser -.
func (pc *ProjectsUseCase) GetProjectsByIDWithUser(req request.GetProjectByIDReq) ([]entity.ProjectsWithUser, error) {
	projects, err := pc.repo.SelectByIdWithUser(req.ID)
	if err != nil {
		return projects, fmt.Errorf("ProjectsUseCase - SelectByIdWithUser - uc.repo.SelectById: %w", err)
	}
	return projects, nil
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
	tx, err = pc.dutyRepo.Insert(tx, projectID, req.UserOwnerID, true)
	if err != nil {
		tx.Rollback()
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties - owner: %w", err)
	}

	//add member to project
	if req.Members != nil {
		for _, userID := range req.Members {
			tx, err = pc.dutyRepo.Insert(tx, projectID, *userID, false)
			if err != nil {
				tx.Rollback()
				return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties - member: %w", err)
			}
		}
	}

	//add work type to project
	if req.WorkTypes != nil {
		for _, name := range req.WorkTypes {
			reqWorkType := request.CreateWorkTypeReq{
				ProjectID: projectID,
				Name:      *name,
			}

			tx, _, err = pc.workTypeRepo.Insert(tx, reqWorkType)
			if err != nil {
				tx.Rollback()
				return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties - member: %w", err)
			}
		}
	}

	tx.Commit()

	return projectID, nil
}
