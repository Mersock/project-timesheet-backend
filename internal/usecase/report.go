package usecase

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"time"
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
func (pc *ReportUseCase) GetAllWorkType(req request.GetWorkTypeReportReq) (response.GroupWorkTypeReport, error) {
	var res response.GroupWorkTypeReport

	workTypes, err := pc.repo.SelectWorkType(req)
	if err != nil {
		return res, fmt.Errorf("ProjectsUseCase - GetAllRoles - uc.repo.Select: %w", err)
	}
	workTypes = pc.calWorkTypeTotalTime(workTypes)

	res = pc.groupWorkTypeReport(workTypes)

	return res, nil
}

// calWorkTypeTotalTime
func (pc *ReportUseCase) calWorkTypeTotalTime(workTypes []entity.ReportWorkType) []entity.ReportWorkType {
	//mapping time
	for i, workType := range workTypes {
		t := time.Unix(int64(*workType.TotalSeconds), 10)
		totalTime := t.UTC().Format("15:04:05")
		workTypes[i].TotalTime = &totalTime
	}
	return workTypes
}

// groupWorkTypeReport
func (pc *ReportUseCase) groupWorkTypeReport(workTypes []entity.ReportWorkType) response.GroupWorkTypeReport {
	var res response.GroupWorkTypeReport

	for _, workType := range workTypes {
		var report response.WorkTypesReport
		res.ProjectID = *workType.ProjectID
		res.ProjectName = *workType.ProjectName
		report.WorkTypeID = *workType.WorkTypeID
		report.WorkTypeName = *workType.WorkTypeName
		report.TotalSeconds = *workType.TotalSeconds
		report.TotalTime = *workType.TotalTime
		res.WorkTypes = append(res.WorkTypes, report)
	}

	return res
}
