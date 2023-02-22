package entity

// Projects -.
type Projects struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Code     *string `json:"code"`
	CreateAt *string `json:"created_at"`
	UpdateAt *string `json:"updated_at,omitempty"`
}
