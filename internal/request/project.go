package request

import "github.com/google/uuid"

// CreateProjectReq -.
type CreateProjectReq struct {
	Name   string    `form:"name" json:"name" binding:"required,max=255"`
	Code   uuid.UUID `form:"-" json:"-"`
	UserID int64     `form:"-" json:"-"`
}
