package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// ReportUseCase -.
type ReportUseCase struct {
	repo ReportRepo
}

// NewReportUseCase -.
func NewReportUseCase(r ReportRepo) *ReportUseCase {
	return &ReportUseCase{
		repo: r,
	}
}

// GetWorkTypeCount -.
func (pc *ReportUseCase) GetWorkTypeCount(req request.GetWorkTypeReportReq) (int, error) {
	rows, err := pc.repo.CountWorkType(req)
	if err != nil {
		return rows, fmt.Errorf("ReportUseCase - CountWorkType - uc.repo.Count: %w", err)
	}
	return rows, nil
}

// GetAllWorkType -.
func (pc *ReportUseCase) GetAllWorkType(req request.GetWorkTypeReportReq) ([]entity.ReportWorkType, error) {
	workTypes, err := pc.repo.SelectWorkType(req)
	if err != nil {
		return nil, fmt.Errorf("ProjectsUseCase - GetAllRoles - uc.repo.Select: %w", err)
	}
	return workTypes, nil
}
