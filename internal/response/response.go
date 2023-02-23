package response

import (
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// IDRes -.
type IDRes struct {
	ID int64 `json:"id"`
}

// RowAffectRes -.
type RowAffectRes struct {
	RowAffected int64 `json:"row_affected"`
}

// errRes -.
type errRes struct {
	Error string `json:"error" example:"message"`
}

// validateRes -.
type validateRes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ResByID -.
func ResByID(c *gin.Context, code int, id int64) {
	c.JSON(code, IDRes{ID: id})
}

// ResRowAffect -.
func ResRowAffect(c *gin.Context, code int, rowAffected int64) {
	c.JSON(code, RowAffectRes{RowAffected: rowAffected})
}

// ErrorResponse -.
func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, errRes{msg})
}

// ErrorValidateRes -.
func ErrorValidateRes(c *gin.Context, ve validator.ValidationErrors) {
	out := make([]validateRes, len(ve))
	for i, fe := range ve {
		out[i] = validateRes{fe.Field(), utils.GetValidateTag(fe)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
}
