package entity

// Users -.
type Users struct {
	ID        *int    `json:"id"`
	Email     *string `json:"email"`
	Password  *string `json:"-"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	CreateAt  *string `json:"created_at"`
	UpdateAt  *string `json:"updated_at,omitempty"`
	Role      *string `json:"role"`
}
