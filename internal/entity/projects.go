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

// ProjectsWithUser -.
type ProjectsWithUser struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	Code      *string    `json:"code"`
	CreateAt  *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
	UserID    *int       `json:"user_id"`
	Email     *string    `json:"email"`
	Firstname *string    `json:"firstname"`
	Lastname  *string    `json:"lastname"`
	Role      *string    `json:"role"`
}

// ProjectWithSliceUser -.
type ProjectWithSliceUser struct {
	Projects
	Users     []UsersInProject `json:"users"`
	WorkTypes []WorkTypes      `json:"work_types"`
}

// UsersInProject -.
type UsersInProject struct {
	UserID    int    `json:"user_id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}
