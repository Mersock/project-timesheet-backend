package request

import "github.com/google/uuid"

// CreateProjectReq -.
type CreateProjectReq struct {
	Name        string    `form:"name" json:"name" binding:"required,max=255"`
	Code        uuid.UUID `form:"-" json:"-" binding:"required"`
	UserOwnerID int64     `form:"-" json:"-" binding:"required,numeric"`
}
