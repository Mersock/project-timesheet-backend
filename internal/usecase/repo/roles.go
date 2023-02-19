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

// GetRows -.
func (r *RolesRepo) GetRows() (int, error) {
	sqlRaw := "SELECT  COUNT(*) FROM `roles` "
	var count int
	err := r.DB.QueryRow(sqlRaw).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("RolesRepo - GetRows - r.DB.QueryRow: %w", err)
	}
	defer r.DB.Close()

	return count, nil
}

// GetRole -.
func (r *RolesRepo) GetRole() ([]entity.Roles, error) {
	var entities []entity.Roles
	sqlRaw := "SELECT `id`, `name`, `created_at`, `updated_at` FROM `roles` ORDER BY `id`"
	results, err := r.DB.Query(sqlRaw)
	if err != nil {
		return nil, fmt.Errorf("RolesRepo - GetRole - r.DB.Query: %w", err)
	}
	defer results.Close()

	for results.Next() {
		var e entity.Roles
		err = results.Scan(&e.ID, &e.Name, &e.CreateAt, &e.UpdateAt)
		entities = append(entities, e)
	}

	return entities, nil
}
