package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	_defaultInternalServerErr = "Internal server error"
	_defaultNotFoundErr       = "Not found"
	_defaultBadReq            = "Bad request"
)

// response -.
type errRes struct {
	Error string `json:"error" example:"message"`
}

// validateRes -.
type validateRes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// errorResponse-.
func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errRes{msg})
}

// errorValidateRes -.
func errorValidateRes(c *gin.Context, ve validator.ValidationErrors) {
	out := make([]validateRes, len(ve))
	for i, fe := range ve {
		out[i] = validateRes{fe.Field(), utils.GetValidateTag(fe)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
}
