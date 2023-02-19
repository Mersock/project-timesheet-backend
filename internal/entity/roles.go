package entity

import "time"

// Roles -.
type Roles struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at,omitempty"`
}
