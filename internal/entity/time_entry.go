package entity

import "time"

// TimeEntry -.
type TimeEntry struct {
	ID        *int       `json:"id"`
	Status    *string    `json:"status"`
	WorkType  *string    `json:"work_type"`
	User      *string    `json:"user"`
	StartTime *time.Time `json:"start_time"`
	DateTime  *time.Time `json:"date_time"`
	CreateAt  *time.Time `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
}
