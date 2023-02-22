package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// WorkTypesRepo -.
type WorkTypesRepo struct {
	*sql.DB
}

// NewWorkTypesRepo -.
func NewWorkTypesRepo(db *sql.DB) *WorkTypesRepo {
	return &WorkTypesRepo{db}
}

// Insert -.
func (r *WorkTypesRepo) Insert(tx *sql.Tx, req request.CreateWorkTypeReq) (*sql.Tx, int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO worktypes (name,project_id,created_at) values (?,?,NOW()) "
	result, err := tx.Exec(sqlRaw, req.Name, req.ProjectID)
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - Insert - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - Insert - result.LastInsertId: %w", err)
	}

	return tx, insertId, nil
}
