package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// ProjectsUseCase -.
type ProjectsUseCase struct {
	repo ProjectRepo
}

// NewProjectsUseCase -.
func NewProjectsUseCase(r ProjectRepo) *ProjectsUseCase {
	return &ProjectsUseCase{repo: r}
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
		return nil, fmt.Errorf("GetAllProjects - GetAllRoles - uc.repo.Select: %w", err)
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
	tx, err = pc.repo.InsertDuties(tx, projectID, req.UserOwnerID, true)
	if err != nil {
		tx.Rollback()
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties: %w", err)
	}

	tx.Commit()

	return projectID, nil
}
