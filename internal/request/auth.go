package request

// SignUpReq -.
type SignUpReq struct {
	Email     string `form:"email" json:"email" binding:"required,email,max=255"`
	Password  string `form:"password" json:"password" binding:"required,min=6"`
	Firstname string `form:"firstname" json:"firstname" binding:"required,max=255"`
	Lastname  string `form:"lastname" json:"lastname" binding:"required,max=255"`
	Role      int    `form:"-" json:"_" binding:"numeric"`
}

// SignInReq -.
type SignInReq struct {
	Email    string `form:"email" json:"email" binding:"required,email,max=255"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

// RenewTokenReq -.
type RenewTokenReq struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}
