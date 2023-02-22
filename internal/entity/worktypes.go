package entity

// WorkTypes -.
type WorkTypes struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	CreateAt *string `json:"created_at"`
	UpdateAt *string `json:"updated_at,omitempty"`
	Project  *string `json:"project"`
}
