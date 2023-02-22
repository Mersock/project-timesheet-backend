package repo

import (
	"database/sql"
	"fmt"
)

// DutiesRepo -.
type DutiesRepo struct {
	*sql.DB
}

// NewDutiesRepo -.
func NewDutiesRepo(db *sql.DB) *DutiesRepo {
	return &DutiesRepo{db}
}

// Insert -.
func (r *DutiesRepo) Insert(tx *sql.Tx, projectID int64, userID int64, isOwner bool) (*sql.Tx, error) {

	sqlRaw := "INSERT INTO duties (project_id,user_id,is_owner) values (?,?,?) "
	_, err := tx.Exec(sqlRaw, projectID, userID, isOwner)

	if err != nil {
		return tx, fmt.Errorf("ProjectRepo - InsertDuties - r.DB.Exec: %w", err)
	}

	return tx, nil
}
