package entity

import "time"

// WorkTypes -.
type WorkTypes struct {
	ID       *int       `json:"id"`
	Name     *string    `json:"name"`
	CreateAt *time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty"`
	Project  *string    `json:"project"`
}
