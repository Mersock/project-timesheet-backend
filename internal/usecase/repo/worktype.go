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
func (r *WorkTypesRepo) Insert(req request.CreateWorkTypeReq) (int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO work_types (name,project_id,created_at) values (?,?,NOW()) "
	result, err := r.DB.Exec(sqlRaw, req.Name, req.ProjectID)
	if err != nil {
		return insertId, fmt.Errorf("WorkTypesRepo - Insert - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return insertId, fmt.Errorf("WorkTypesRepo - Insert - result.LastInsertId: %w", err)
	}

	return insertId, nil
}

// InsertWithProject -.
func (r *WorkTypesRepo) InsertWithProject(tx *sql.Tx, req request.CreateWorkTypeReq) (*sql.Tx, int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO work_types (name,project_id,created_at) values (?,?,NOW()) "
	result, err := tx.Exec(sqlRaw, req.Name, req.ProjectID)
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - InsertWithProject - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - InsertWithProject - result.LastInsertId: %w", err)
	}

	return tx, insertId, nil
}

// UpdateWithProject -.
func (r *WorkTypesRepo) UpdateWithProject(tx *sql.Tx, req request.UpdateWorkTypeReq) (*sql.Tx, int64, error) {
	var insertId int64

	sqlRaw := "UPDATE work_types SET name = ?, updated_at = NOW() WHERE id = ?"
	result, err := tx.Exec(sqlRaw, req.Name, req.ID)
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - UpdateWithProject - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return tx, insertId, fmt.Errorf("WorkTypesRepo - UpdateWithProject - result.LastInsertId: %w", err)
	}

	return tx, insertId, nil
}

// Delete -.
func (r *WorkTypesRepo) DeleteWithProject(tx *sql.Tx, req request.DeleteWorkTypeReq) (*sql.Tx, int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM work_types WHERE id = ?"
	result, err := tx.Exec(sqlRaw, req.ID)
	if err != nil {
		return tx, rowAffected, fmt.Errorf("WorkTypesRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return tx, rowAffected, fmt.Errorf("WorkTypesRepo - Delete - result.rowAffected: %w", err)
	}
	return tx, rowAffected, nil
}

// SelectById -.
func (r *WorkTypesRepo) SelectById(workTypeID int) (entity.WorkTypes, error) {
	var entity entity.WorkTypes

	sqlRaw := "SELECT work_types.id, work_types.name, work_types.created_at, work_types.updated_at, projects.name as project "
	sqlRaw += "FROM work_types "
	sqlRaw += "INNER JOIN projects ON work_types.project_id = projects.id "
	sqlRaw += "WHERE work_types.id = ? "
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

// SelectByProjectId -.
func (r *WorkTypesRepo) SelectByProjectId(projectID int) ([]entity.WorkTypes, error) {
	var entities []entity.WorkTypes

	sqlRaw := "SELECT work_types.id, work_types.name, work_types.created_at, work_types.updated_at, projects.name as project "
	sqlRaw += "FROM work_types "
	sqlRaw += "INNER JOIN projects ON work_types.project_id = projects.id "
	sqlRaw += "WHERE work_types.project_id = ? "
	results, err := r.DB.Query(sqlRaw, projectID)
	if err != nil {
		return entities, fmt.Errorf("WorkTypesRepo - SelectByProjectId - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.WorkTypes
		err = results.Scan(&e.ID, &e.Name, &e.CreateAt, &e.UpdateAt, &e.Project)
		entities = append(entities, e)
	}

	return entities, nil
}

// Update -.
func (r *WorkTypesRepo) Update(req request.UpdateWorkTypeReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "UPDATE work_types SET name = ?, updated_at = NOW() WHERE id = ?"
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
func (r *WorkTypesRepo) Delete(req request.DeleteWorkTypeReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM work_types WHERE id = ?"
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
