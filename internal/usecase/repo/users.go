package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
)

// UsersRepo -.
type UsersRepo struct {
	*sql.DB
}

// NewUsersRepo -.
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

// Count -.
func (r *UsersRepo) Count(req request.GetUsersReq) (int, error) {
	var count int

	sqlRaw := "SELECT COUNT(*) "
	sqlRaw += "FROM users "
	sqlRaw += "INNER JOIN roles ON roles.id = users.id "
	sqlRaw += "WHERE 1=1 "
	sqlCount := r.genRawSelectWithReq(sqlRaw, req)

	err := r.DB.QueryRow(sqlCount).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("UsersRepo - Count - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// Select -.
func (r *UsersRepo) Select(req request.GetUsersReq) ([]entity.Users, error) {
	var entities []entity.Users

	sqlRaw := "SELECT users.id, email, firstname, lastname, users.created_at, users.updated_at,roles.name as role "
	sqlRaw += "FROM users "
	sqlRaw += "INNER JOIN roles ON roles.id = users.id "
	sqlRaw += "WHERE 1=1 "
	sqlSelect := r.genRawSelectWithReq(sqlRaw, req)
	mainQuery := r.genPaginateQuery(sqlSelect, req)

	results, err := r.DB.Query(mainQuery)
	if err != nil {
		return nil, fmt.Errorf("UsersRepo - Select - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Users
		err = results.Scan(&e.ID, &e.Email, &e.Firstname, &e.Lastname, &e.CreateAt, &e.UpdateAt, &e.Role)
		entities = append(entities, e)
	}

	return entities, nil
}

// SelectById -.
func (r *UsersRepo) SelectById(userID int) (entity.Users, error) {
	var entity entity.Users

	sqlRaw := "SELECT users.id, email, firstname, lastname, users.created_at, users.updated_at,roles.name as role "
	sqlRaw += "FROM users "
	sqlRaw += "INNER JOIN roles ON roles.id = users.id "
	sqlRaw += "WHERE users.id = ? "
	err := r.DB.QueryRow(sqlRaw, userID).Scan(&entity.ID,
		&entity.Email,
		&entity.Firstname,
		&entity.Lastname,
		&entity.CreateAt,
		&entity.UpdateAt,
		&entity.Role,
	)
	if err != nil {
		return entity, fmt.Errorf("UsersRepo - SelectById - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// SelectByEmail -.
func (r *UsersRepo) SelectByEmail(email string) (entity.Users, error) {
	var entity entity.Users

	sqlRaw := "SELECT users.id, email, firstname, lastname, users.created_at, users.updated_at,roles.name as role "
	sqlRaw += "FROM users "
	sqlRaw += "INNER JOIN roles ON roles.id = users.id "
	sqlRaw += "WHERE users.email = ? "
	err := r.DB.QueryRow(sqlRaw, email).Scan(&entity.ID,
		&entity.Email,
		&entity.Firstname,
		&entity.Lastname,
		&entity.CreateAt,
		&entity.UpdateAt,
		&entity.Role,
	)
	if err != nil {
		return entity, fmt.Errorf("UsersRepo - SelectByEmail - r.DB.QueryRow: %w", err)
	}

	return entity, nil
}

// Insert -.
func (r *UsersRepo) Insert(req request.CreateUserReq) (int64, error) {
	var insertId int64

	sqlRaw := "INSERT INTO users (email,password,firstname,lastname,role_id,created_at) values (?,?,?,?,?,NOW()) "
	result, err := r.DB.Exec(sqlRaw, req.Email, req.Password, req.Firstname, req.Lastname, req.RoleID)
	if err != nil {
		return insertId, fmt.Errorf("UsersRepo - Insert - r.DB.Exec: %w", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		return insertId, fmt.Errorf("UsersRepo - Insert - result.LastInsertId: %w", err)
	}

	return insertId, nil
}

// Delete -.
func (r *UsersRepo) Delete(req request.DeleteUserReq) (int64, error) {
	var rowAffected int64
	sqlRaw := "DELETE FROM users WHERE id = ?"
	result, err := r.DB.Exec(sqlRaw, req.ID)
	if err != nil {
		return rowAffected, fmt.Errorf("UsersRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("UsersRepo - Delete - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}

// ChkUniqueInsert -.
func (r *UsersRepo) ChkUniqueInsert(req request.CreateUserReq) (int, error) {
	var count int
	sqlRaw := fmt.Sprintf("SELECT  COUNT(*) FROM users WHERE email = '%s' ", req.Email)
	err := r.DB.QueryRow(sqlRaw).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("UsersRepo - ChkUniqueInsert - r.DB.QueryRow: %w", err)
	}

	return count, nil
}

// genRawSelectWithReq -.
func (r *UsersRepo) genRawSelectWithReq(sqlRaw string, req request.GetUsersReq) string {
	if req.Email != "" {
		sqlRaw = fmt.Sprintf("%s AND email LIKE '%%%s%%' ", sqlRaw, req.Email)
	}

	if req.Firstname != "" {
		sqlRaw = fmt.Sprintf("%s AND firstname LIKE '%%%s%%' ", sqlRaw, req.Firstname)
	}

	if req.Lastname != "" {
		sqlRaw = fmt.Sprintf("%s AND lastname LIKE '%%%s%%' ", sqlRaw, req.Lastname)
	}

	if req.Role != "" {
		sqlRaw = fmt.Sprintf("%s AND roles.name LIKE '%%%s%%' ", sqlRaw, req.Role)
	}

	if req.SortBy != "" && req.SortType != "" {
		sqlRaw = fmt.Sprintf("%s ORDER BY %s %s", sqlRaw, req.SortBy, req.SortType)
	}

	return sqlRaw
}

// genPaginateQuery -.
func (r *UsersRepo) genPaginateQuery(sqlRaw string, req request.GetUsersReq) string {
	if req.Limit != nil && req.Page != nil {
		offset := (*req.Page - 1) * (*req.Limit)
		sqlRaw = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlRaw, *req.Limit, offset)
	}
	return sqlRaw
}
