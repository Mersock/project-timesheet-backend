package repo

import (
	"database/sql"
	"fmt"
	"strings"
)

// DutiesRepo -.
type DutiesRepo struct {
	*sql.DB
}

// NewDutiesRepo -.
func NewDutiesRepo(db *sql.DB) *DutiesRepo {
	return &DutiesRepo{db}
}

// InsertOwner -.
func (r *DutiesRepo) InsertOwner(tx *sql.Tx, projectID int64, OwnerUserID int64) (*sql.Tx, error) {

	sqlRaw := "INSERT INTO duties (project_id,user_id,is_owner) values (?,?,?) "
	_, err := tx.Exec(sqlRaw, projectID, OwnerUserID, true)

	if err != nil {
		return tx, fmt.Errorf("DutiesRepo - InsertOwner - r.DB.Exec: %w", err)
	}

	return tx, nil
}

// InsertMember -.
func (r *DutiesRepo) InsertMember(tx *sql.Tx, projectID int64, members []string) (*sql.Tx, error) {

	sqlRaw := "INSERT INTO duties (project_id,user_id,is_owner) SELECT ? as project_id, users.id, false as is_owner from users where users.email in (?)"
	_, err := tx.Exec(sqlRaw, projectID, strings.Join(members, ","))

	if err != nil {
		return tx, fmt.Errorf("DutiesRepo - InsertMember - r.DB.Exec: %w", err)
	}

	return tx, nil
}

// Delete -.
func (r *DutiesRepo) Delete(projectID int64, userID int64) (int64, error) {
	var rowAffected int64

	sqlRaw := "DELETE FROM duties WHERE project_id = ? AND user_id = ?"
	result, err := r.DB.Exec(sqlRaw, projectID, userID)
	if err != nil {
		return rowAffected, fmt.Errorf("DutiesRepo - Delete - r.DB.Exec: %w", err)
	}
	rowAffected, err = result.RowsAffected()
	if err != nil {
		return rowAffected, fmt.Errorf("DutiesRepo - Delete - result.rowAffected: %w", err)
	}
	return rowAffected, nil
}
