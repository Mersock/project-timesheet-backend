package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
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

// SelectById -.
func (r *WorkTypesRepo) SelectById(workTypeID int) (entity.WorkTypes, error) {
	var entity entity.WorkTypes

	sqlRaw := "SELECT id, name, created_at, updated_at,projects.name as project"
	sqlRaw += "FROM worktypes "
	sqlRaw += "INNER JOIN projects ON worktypes.id = projects.id"
	sqlRaw += "WHERE id = ? "
	err := r.DB.QueryRow(sqlRaw, workTypeID).Scan(&entity.ID,
		&entity.Name,
		&entity.CreateAt,
		&entity.UpdateAt,
		&entity.Project,
	)
	if err != nil {
		return entity, fmt.Errorf("WorkTypesRepo - SelectById - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// Update -.
func (r *WorkTypesRepo) Update(req request.UpdateWorkTypeReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "UPDATE worktypes SET name = ?, updated_at = NOW() WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.Name, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesRepo - Update - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesRepo - Update - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}

// Delete -.
func (r *WorkTypesRepo) Delete(req request.DeleteUserReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM worktypes WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("WorkTypesRepo - Delete - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}
