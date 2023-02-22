package entity

// Projects -.
type Projects struct {
	ID       *int    `json:"id"`
	Name     *string `json:"email"`
	Code     *string `json:"-"`
	CreateAt *string `json:"created_at"`
	UpdateAt *string `json:"updated_at,omitempty"`
}
