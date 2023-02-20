package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// RolesRepo -.
type RolesRepo struct {
	*sql.DB
}

// NewRolesRepo -.
func NewRolesRepo(db *sql.DB) *RolesRepo {
	return &RolesRepo{db}
}

// Count -.
func (r *RolesRepo) Count(req request.GetRolesReq) (int, error) {
	var count int
	sqlRaw := "SELECT  COUNT(*) FROM `roles` "
	sqlCount := genRawSelectWithReq(sqlRaw, req)
	err := r.DB.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("RolesRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// Select -.
func (r *RolesRepo) Select(req request.GetRolesReq) ([]entity.Roles, error) {
	var entities []entity.Roles

	sqlRaw := "SELECT `id`, `name`, `created_at`, `updated_at` FROM `roles` WHERE 1=1 "
	sqlSelect := genRawSelectWithReq(sqlRaw, req)
	mainQuery := genPaginateQuery(sqlSelect, req)
	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("RolesRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Roles
		err = results.Scan(&e.ID, &e.Name, &e.CreateAt, &e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}

// SelectById -.
func (r *RolesRepo) SelectById(roleID int) (entity.Roles, error) {
	var entity entity.Roles

	sqlRaw := "SELECT `id`, `name`, `created_at`, `updated_at` FROM `roles` WHERE `id` = ?"
	err := r.DB.QueryRow(sqlRaw, roleID).Scan(&entity.ID, &entity.Name, &entity.CreateAt, &entity.UpdateAt)
	if err != nil {
		return entity, fmt.Errorf("RolesRepo - SelectById - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// Insert -.
func (r *RolesRepo) Insert(req request.CreateRoleReq) (int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO roles (name,created_at) values (?,NOW()) "
	result, err := r.DB.ExecContext(context.Background(), sqlRaw, req.Name)
	if err != nil {
		return insertId, fmt.Errorf("RolesRepo - Insert - r.DB.ExecContext: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return insertId, fmt.Errorf("RolesRepo - Insert - result.LastInsertId: %w", err)
	}

	return insertId, nil
}

// ChkUniqueInsert -.
func (r *RolesRepo) ChkUniqueInsert(req request.CreateRoleReq) (int, error) {
	var count int
	sqlRaw := fmt.Sprintf("SELECT  COUNT(*) FROM `roles` WHERE name = '%s' ", req.Name)
	err := r.DB.QueryRow(sqlRaw).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("RolesRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// genRawSelectWithReq -.
func genRawSelectWithReq(sqlRaw string, req request.GetRolesReq) string {
	if req.Name != "" {
		sqlRaw = fmt.Sprintf("%s AND `name` LIKE '%%%s%%' ", sqlRaw, req.Name)
	}

	if req.SortBy != "" && req.SortType != "" {
		sqlRaw = fmt.Sprintf("%s ORDER BY %s %s", sqlRaw, req.SortBy, req.SortType)
	}

	return sqlRaw
}

// genPaginateQuery -.
func genPaginateQuery(sqlRaw string, req request.GetRolesReq) string {
	if req.Limit != nil && req.Page != nil {
		offset := (*req.Page - 1) * (*req.Limit)
		sqlRaw = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlRaw, *req.Limit, offset)
	}
	return sqlRaw
}
