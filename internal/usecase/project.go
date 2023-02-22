package usecase

import (
	"fmt"
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

	tx, err = pc.repo.InsertDuties(tx, projectID, 123, true)
	if err != nil {
		tx.Rollback()
		return projectID, fmt.Errorf("ProjectsUseCase - CreateProject - pc.repo.InsertDuties: %w", err)
	}

	tx.Commit()

	return projectID, nil
}
