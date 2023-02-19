package repo

import (
	"database/sql"
	"fmt"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
)

// RolesRepo -.
type RolesRepo struct {
	*sql.DB
}

// NewRolesRepo -.
func NewRolesRepo(db *sql.DB) *RolesRepo {
	return &RolesRepo{db}
}

// GetRole -.
func (r *RolesRepo) GetRole() ([]entity.Roles, error) {
	var entities []entity.Roles
	results, err := r.DB.Query("SELECT id, name, created_at, updated_at FROM roles")
	if err != nil {
		return nil, fmt.Errorf("RolesRepo - GetRole - r.DB.Query: %w", err)
	}

	for results.Next() {
		var e entity.Roles
		err = results.Scan(&e.ID, &e.Name, &e.CreateAt, &e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}
