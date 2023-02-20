package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// response -.
type response struct {
	Error string `json:"error" example:"message"`
}

// validateRes -.
type validateRes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// errorResponse-.
func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func errorValidateRes(c *gin.Context, ve validator.ValidationErrors) {
	out := make([]validateRes, len(ve))
	for i, fe := range ve {
		out[i] = validateRes{fe.Field(), getValidateTag(fe)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
}

func getValidateTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "numeric":
		return "This field is required only numeric"
	}
	return "Unknown error"
}
