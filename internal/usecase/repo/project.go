package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// ProjectRepo -.
type ProjectRepo struct {
	*sql.DB
}

// NewProjectRepo -.
func NewProjectRepo(db *sql.DB) *ProjectRepo {
	return &ProjectRepo{db}
}

// BeginTx -.
func (r *ProjectRepo) BeginTx() (*sql.Tx, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("ProjectRepo - BeginTx - p.DB.Begin: %w", err)
	}

	return tx, nil
}

// Count -.
func (r *ProjectRepo) Count(req request.GetProjectsReq) (int, error) {
	var count int

	sqlRaw := "SELECT COUNT(*) "
	sqlRaw += "FROM projects "
	sqlRaw += "INNER JOIN duties ON duties.project_id = projects.id "
	sqlRaw += "WHERE 1=1 "
	sqlCount := r.genRawSelectWithReq(sqlRaw, req)

	err := r.DB.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("ProjectRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// Select -.
func (r *ProjectRepo) Select(req request.GetProjectsReq) ([]entity.Projects, error) {
	var entities []entity.Projects

	sqlRaw := "SELECT id,name,code,created_at,updated_at "
	sqlRaw += "FROM projects "
	sqlRaw += "INNER JOIN duties ON duties.project_id = projects.id "
	sqlRaw += "WHERE 1=1 "
	sqlSelect := r.genRawSelectWithReq(sqlRaw, req)
	mainQuery := r.genPaginateQuery(sqlSelect, req)
	fmt.Println(mainQuery)
	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Projects
		err = results.Scan(&e.ID, &e.Name, &e.Code, &e.CreateAt, &e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}

// Insert -.
func (r *ProjectRepo) Insert(tx *sql.Tx, req request.CreateProjectReq) (*sql.Tx, int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO projects (code,name,created_at) values (?,?,NOW()) "
	result, err := tx.Exec(sqlRaw, req.Code, req.Name)

	if err != nil {
		return tx, insertId, fmt.Errorf("ProjectRepo - Insert - r.DB.Exec: %w", err)
	}

	insertId, err = result.LastInsertId()
	if err != nil {
		return tx, insertId, fmt.Errorf("ProjectRepo - Insert - result.LastInsertId: %w", err)
	}

	return tx, insertId, nil
}

// genRawSelectWithReq -.
func (r *ProjectRepo) genRawSelectWithReq(sqlRaw string, req request.GetProjectsReq) string {
	if req.Name != "" {
		sqlRaw = fmt.Sprintf("%s AND name LIKE '%%%s%%' ", sqlRaw, req.Name)
	}

	if req.Code != "" {
		sqlRaw = fmt.Sprintf("%s AND code LIKE '%%%s%%' ", sqlRaw, req.Code)
	}

	return sqlRaw
}

// genPaginateQuery -.
func (r *ProjectRepo) genPaginateQuery(sqlRaw string, req request.GetProjectsReq) string {
	if req.Limit != nil && req.Page != nil {
		offset := (*req.Page - 1) * (*req.Limit)
		sqlRaw = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlRaw, *req.Limit, offset)
	}
	return sqlRaw
}
