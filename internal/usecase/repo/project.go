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
	sqlRaw += "WHERE 1=1 "
	sqlSelect := r.genRawSelectWithReq(sqlRaw, req)
	mainQuery := r.genPaginateQuery(sqlSelect, req)

	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Projects
		err = results.Scan(&e.ID, &e.Name, &e.Code, &e.CreateAt, &e.UpdateAt)
		if err != nil {
			return nil, fmt.Errorf("ProjectRepo - Select - r.DB.Query: %w", err)
		}
		entities = append(entities, e)
	}

	return entities, nil
}

// SelectById -.
func (r *ProjectRepo) SelectById(projectID int) (entity.Projects, error) {
	var entity entity.Projects

	sqlRaw := "SELECT id,name,code,created_at,updated_at "
	sqlRaw += "FROM projects "
	sqlRaw += "WHERE id = ? "
	err := r.DB.QueryRow(sqlRaw, projectID).Scan(&entity.ID,
		&entity.Name,
		&entity.Code,
		&entity.CreateAt,
		&entity.UpdateAt,
	)
	if err != nil {
		return entity, fmt.Errorf("ProjectRepo - SelectById - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// SelectByIdWithUser -.
func (r *ProjectRepo) SelectByIdWithUser(projectID int) ([]entity.ProjectsWithUser, error) {
	var entities []entity.ProjectsWithUser

	sqlRaw := "SELECT projects.id,projects.code,projects.name,projects.created_at,projects.updated_at,"
	sqlRaw += "users.id as user_id,users.email,users.firstname,users.lastname,roles.name as role_name "
	sqlRaw += "FROM projects "
	sqlRaw += "INNER JOIN duties ON duties.project_id = projects.id "
	sqlRaw += "INNER JOIN users ON duties.user_id = users.id "
	sqlRaw += "INNER JOIN roles ON users.role_id = roles.id "
	sqlRaw += "WHERE projects.id = ? "
	results, err := r.DB.Query(sqlRaw, projectID)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepo - SelectByIdWithUser - r.DB.QueryRow: %w", err)
	}

	for results.Next() {
		var e entity.ProjectsWithUser
		err = results.Scan(&e.ID, &e.Code, &e.Name, &e.CreateAt, &e.UpdateAt, &e.UserID, &e.Email, &e.Firstname, &e.Lastname, &e.Role)
		if err != nil {
			return nil, fmt.Errorf("ProjectRepo - SelectByIdWithUser - r.DB.Query: %w", err)
		}
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

// Update -.
func (r *ProjectRepo) Update(req request.UpdateProjectReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "UPDATE projects SET name = ?, updated_at = NOW() WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.Name, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectRepo - Update - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectRepo - Update - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}

// Delete -.
func (r *ProjectRepo) Delete(req request.DeleteProjectByReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM projects WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("ProjectRepo - Delete - result.rowAffected: %w", err)
	}
	return rowAffected, nil
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
