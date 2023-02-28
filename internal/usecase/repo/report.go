package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// ReportRepo -.
type ReportRepo struct {
	*sql.DB
}

// NewReportRepo -.
func NewReportRepo(db *sql.DB) *ReportRepo {
	return &ReportRepo{db}
}

// CountWorkType -.
func (r *ReportRepo) CountWorkType(req request.GetWorkTypeReportReq) (int, error) {
	var count int

	sqlRaw := "SELECT  COUNT(*) FROM time_entries  "
	sqlRaw += "INNER JOIN work_types ON time_entries.work_type_id = work_types.id "
	sqlRaw += "INNER JOIN projects ON work_types.project_id = projects.id "
	sqlRaw += "WHERE projects.code = ?"
	err := r.DB.QueryRow(sqlRaw, req.ProjectCode).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("ReportRepo - Count - r.DB.QueryRow: %w", err)
	}
	return count, nil
}

// SelectWorkType -.
func (r *ReportRepo) SelectWorkType(req request.GetWorkTypeReportReq) ([]entity.ReportWorkType, error) {
	var entities []entity.ReportWorkType

	sqlRaw := "SELECT  "
	sqlRaw += "projects.id AS project_id, "
	sqlRaw += "projects.name AS project_name, "
	sqlRaw += "work_types.id AS work_type_id, "
	sqlRaw += "work_types.name AS work_type_name,"
	sqlRaw += "SUM(TIMESTAMPDIFF(SECOND, start_time , end_time)) AS total_seconds "
	sqlRaw += "FROM time_entries "
	sqlRaw += "INNER JOIN work_types ON time_entries.work_type_id = work_types.id "
	sqlRaw += "INNER JOIN projects ON work_types.project_id = projects.id "
	sqlRaw += "WHERE projects.code = ? "
	sqlRaw += "GROUP BY work_types.id, work_types.name, projects.name "
	sqlRaw += "ORDER BY projects.id ASC, work_types.id ASC"
	results, err := r.DB.Query(sqlRaw, req.ProjectCode)
	if err != nil {
		return entities, fmt.Errorf("ReportRepo - Count - r.DB.Query: %w", err)
	}
	for results.Next() {
		var e entity.ReportWorkType
		err = results.Scan(&e.ProjectID, &e.ProjectName, &e.WorkTypeID, &e.WorkTypeName, &e.TotalSeconds)
		entities = append(entities, e)
	}

	return entities, nil
}
