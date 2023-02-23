package repo

import (
	"database/sql"
	"fmt"
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
