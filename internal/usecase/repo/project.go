package repo

import (
	"database/sql"
	"fmt"
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
func (p *ProjectRepo) BeginTx() (*sql.Tx, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("ProjectRepo - BeginTx - p.DB.Begin: %w", err)
	}

	return tx, nil
}

// Insert -.
func (p *ProjectRepo) Insert(tx *sql.Tx, req request.CreateProjectReq) (*sql.Tx, int64, error) {
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

// InsertDuties -.
func (p *ProjectRepo) InsertDuties(tx *sql.Tx, projectID int64, userID int64, isOwner bool) (*sql.Tx, error) {

	sqlRaw := "INSERT INTO duties (project_id,user_id,is_owner) values (?,?,NOW()) "
	_, err := tx.Exec(sqlRaw, projectID, userID, isOwner)

	if err != nil {
		return tx, fmt.Errorf("ProjectRepo - InsertDuties - r.DB.Exec: %w", err)
	}

	return tx, nil
}
