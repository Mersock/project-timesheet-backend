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

type TimeEntryList struct {
	ID          *int       `json:"id"`
	Status      *string    `json:"status"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	WorkType    *string    `json:"work_type"`
	ProjectCode *string    `json:"project_code"`
	ProjectName *string    `json:"project_name"`
	Email       *string    `json:"email"`
	Firstname   *string    `json:"firstname"`
	Lastname    *string    `json:"lastname"`
	CreateAt    *time.Time `json:"created_at"`
	UpdateAt    *time.Time `json:"updated_at,omitempty"`
}
