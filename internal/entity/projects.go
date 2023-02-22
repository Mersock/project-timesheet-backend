package entity

import "time"

// Projects -.
type Projects struct {
	ID       *int       `json:"id"`
	Name     *string    `json:"name"`
	Code     *string    `json:"code"`
	CreateAt *time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty"`
}
