package entity

import "time"

// Users -.
type Users struct {
	ID        *int       `json:"id"`
	Email     *string    `json:"email"`
	Password  *string    `json:"-"`
	Firstname *string    `json:"firstname"`
	Lastname  *string    `json:"lastname"`
	CreateAt  *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
	Role      *string    `json:"role"`
}
