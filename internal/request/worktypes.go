package request

// CreateWorkTypeReq -.
type CreateWorkTypeReq struct {
	Name      string `form:"name" json:"name" binding:"required,max=255"`
	ProjectID int64  `form:"project_id" json:"project_id" binding:"required,numeric"`
}

// GetWorkTypeByIDReq -.
type GetWorkTypeByIDReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}

// UpdateWorkTypeReq -.
type UpdateWorkTypeReq struct {
	ID   int    `binding:"required,numeric"`
	Name string `form:"name" json:"name" binding:"required"`
}

// DeleteWorkTypeReq -.
type DeleteWorkTypeReq struct {
	ID int `uri:"id" binding:"required,numeric"`
}
