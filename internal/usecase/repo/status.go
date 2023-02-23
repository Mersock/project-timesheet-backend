package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// StatusRepo -.
type StatusRepo struct {
	*sql.DB
}

// NewStatusRepo -.
func NewStatusRepo(db *sql.DB) *RolesRepo {
	return &RolesRepo{db}
}

// Count -.
func (r *StatusRepo) Count(req request.GetStatusReq) (int, error) {
	var count int
	sqlRaw := "SELECT  COUNT(*) FROM statuses WHERE 1=1 "
	sqlCount := r.genRawSelectWithReq(sqlRaw, req)
	err := r.DB.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("StatusRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// Select -.
func (r *StatusRepo) Select(req request.GetStatusReq) ([]entity.Status, error) {
	var entities []entity.Status

	sqlRaw := "SELECT id, name, created_at, updated_at FROM statuses WHERE 1=1 "
	sqlSelect := r.genRawSelectWithReq(sqlRaw, req)
	mainQuery := r.genPaginateQuery(sqlSelect, req)
	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("StatusRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Status
		err = results.Scan(&e.ID, &e.Name, &e.CreateAt, &e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}

// SelectById -.
func (r *StatusRepo) SelectById(statusID int) (entity.Status, error) {
	var entity entity.Status

	sqlRaw := "SELECT id, name, created_at, updated_at FROM statuses WHERE id = ?"
	err := r.DB.QueryRow(sqlRaw, statusID).Scan(&entity.ID, &entity.Name, &entity.CreateAt, &entity.UpdateAt)
	if err != nil {
		return entity, fmt.Errorf("StatusRepo - SelectById - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// Insert -.
func (r *StatusRepo) Insert(req request.CreateStatusReq) (int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO statuses (name,created_at) values (?,NOW()) "
	result, err := r.DB.Exec(sqlRaw, req.Name)
	if err != nil {
		return insertId, fmt.Errorf("StatusRepo - Insert - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return insertId, fmt.Errorf("StatusRepo - Insert - result.LastInsertId: %w", err)
	}

	return insertId, nil
}

// Update -.
func (r *StatusRepo) Update(req request.UpdateStatusReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "UPDATE statuses SET name = ?, updated_at = NOW() WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.Name, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("StatusRepo - Update - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("StatusRepo - Update - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}

// Delete -.
func (r *StatusRepo) Delete(req request.DeleteStatusReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM statuses WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("StatusRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("StatusRepo - Delete - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}

// ChkUniqueInsert -.
func (r *StatusRepo) ChkUniqueInsert(req request.CreateStatusReq) (int, error) {
	var count int
	sqlRaw := fmt.Sprintf("SELECT  COUNT(*) FROM statuses WHERE name = '%s' ", req.Name)
	err := r.DB.QueryRow(sqlRaw).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("StatusRepo - ChkUniqueInsert - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// ChkUniqueUpdate -.
func (r *StatusRepo) ChkUniqueUpdate(req request.UpdateRoleReq) (int, error) {
	var count int
	sqlRaw := fmt.Sprintf("SELECT  COUNT(*) FROM statuses WHERE name = '%s' AND id != %d", req.Name, req.ID)
	fmt.Println(sqlRaw)
	err := r.DB.QueryRow(sqlRaw).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("StatusRepo - ChkUniqueUpdate - r.DB.QueryRow: %w", err)
	}
	return count, nil
}

// genRawSelectWithReq -.
func (r *StatusRepo) genRawSelectWithReq(sqlRaw string, req request.GetStatusReq) string {
	if req.Name != "" {
		sqlRaw = fmt.Sprintf("%s AND name LIKE '%%%s%%' ", sqlRaw, req.Name)
	}

	if req.SortBy != "" && req.SortType != "" {
		sqlRaw = fmt.Sprintf("%s ORDER BY %s %s", sqlRaw, req.SortBy, req.SortType)
	}

	return sqlRaw
}

// genPaginateQuery -.
func (r *StatusRepo) genPaginateQuery(sqlRaw string, req request.GetStatusReq) string {
	if req.Limit != nil && req.Page != nil {
		offset := (*req.Page - 1) * (*req.Limit)
		sqlRaw = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlRaw, *req.Limit, offset)
	}
	return sqlRaw
}
