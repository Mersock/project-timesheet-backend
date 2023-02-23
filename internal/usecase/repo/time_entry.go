package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// TimeEntryRepo -.
type TimeEntryRepo struct {
	*sql.DB
}

// NewTimeEntryRepoRepo -.
func NewTimeEntryRepoRepo(db *sql.DB) *TimeEntryRepo {
	return &TimeEntryRepo{db}
}

// Count -.
func (r *TimeEntryRepo) Count(req request.GetTimeEntryReq) (int, error) {
	var count int
	sqlRaw := "SELECT  COUNT(*) FROM time_entries WHERE 1=1 "
	sqlCount := r.genRawSelectWithReq(sqlRaw, req)
	err := r.DB.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("TimeEntryRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// Select -.
func (r *TimeEntryRepo) Select(req request.GetTimeEntryReq) ([]entity.TimeEntryList, error) {
	var entities []entity.TimeEntryList

	sqlRaw := "SELECT "
	sqlRaw += "time_entries.id,"
	sqlRaw += "statuses.name status, "
	sqlRaw += "work_types.name work_type, "
	sqlRaw += "time_entries.start_time, "
	sqlRaw += "time_entries.end_time, "
	sqlRaw += "projects.name project_name, "
	sqlRaw += "users.email,"
	sqlRaw += "users.firstname, "
	sqlRaw += "users.lastname, "
	sqlRaw += "time_entries.created_at,"
	sqlRaw += "time_entries.updated_at "
	sqlRaw += "FROM time_entries "
	sqlRaw += "INNER JOIN statuses ON time_entries.status_id = statuses.id "
	sqlRaw += "INNER JOIN work_types ON time_entries.work_type_id = work_types.id "
	sqlRaw += "INNER JOIN projects ON work_types.project_id = projects.id "
	sqlRaw += "INNER JOIN users ON time_entries.user_id = users.id "
	sqlRaw += "WHERE 1=1 "
	sqlSelect := r.genRawSelectWithReq(sqlRaw, req)
	mainQuery := r.genPaginateQuery(sqlSelect, req)
	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("TimeEntryRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.TimeEntryList
		err = results.Scan(&e.ID, &e.Status, &e.WorkType, &e.StartTime, &e.EndTime, &e.ProjectName,
			&e.Email, e.Firstname, e.Lastname, e.CreateAt, e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}

// Insert -.
func (r *TimeEntryRepo) Insert(req request.CreateTimeEntryReq) (int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO time_entries (status_id,work_type_id,user_id,start_time,end_time,created_at) values (?,?,?,?,?,NOW()) "
	result, err := r.DB.Exec(sqlRaw, req.StatusID, req.WorkTypeID, req.UserID, req.StartDate, req.EndDate)
	if err != nil {
		return insertId, fmt.Errorf("TimeEntryRepo - Insert - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return insertId, fmt.Errorf("TimeEntryRepo - Insert - result.LastInsertId: %w", err)
	}

	return insertId, nil
}

// genRawSelectWithReq -.
func (r *TimeEntryRepo) genRawSelectWithReq(sqlRaw string, req request.GetTimeEntryReq) string {
	//if req.Name != "" {
	//	sqlRaw = fmt.Sprintf("%s AND name LIKE '%%%s%%' ", sqlRaw, req.Name)
	//}
	//
	if req.SortBy != "" && req.SortType != "" {
		sqlRaw = fmt.Sprintf("%s ORDER BY %s %s", sqlRaw, req.SortBy, req.SortType)
	}

	return sqlRaw
}

// genPaginateQuery -.
func (r *TimeEntryRepo) genPaginateQuery(sqlRaw string, req request.GetTimeEntryReq) string {
	if req.Limit != nil && req.Page != nil {
		offset := (*req.Page - 1) * (*req.Limit)
		sqlRaw = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlRaw, *req.Limit, offset)
	}
	return sqlRaw
}
