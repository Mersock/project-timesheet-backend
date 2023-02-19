package entity

// Roles -.
type Roles struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	CreateAt *string `json:"created_at"`
	UpdateAt *string `json:"updated_at,omitempty"`
}
