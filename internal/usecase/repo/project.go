package repo

import (
	"database/sql"
	"fmt"
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

// Rollback -.
func (p *ProjectRepo) Rollback(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("ProjectRepo - Rollback - tx.Rollback(): %w", err)
	}
	return nil
}

// Commit -.
func (p *ProjectRepo) Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("ProjectRepo - Commit - tx.Commit(): %w", err)
	}
	return nil
}
