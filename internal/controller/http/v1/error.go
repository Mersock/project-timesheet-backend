package v1

import "github.com/gin-gonic/gin"

// response -.
type response struct {
	Error string `json:"error" example:"message"`
}

// errorResponse-.
func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
