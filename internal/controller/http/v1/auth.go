package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// authRoutes -.
type authRoutes struct {
	uu usecase.User
	l  logger.Interface
}

// newAuthRoutes -.
func newAuthRoutes(h *gin.RouterGroup, uu usecase.User, l logger.Interface) {
	a := authRoutes{uu, l}

	{
		h.POST("/signup", a.singUp)
		h.POST("/signin", a.singIn)
	}
}

// singUp -.
func (a authRoutes) singUp(c *gin.Context) {
	c.Status(http.StatusOK)
}

// singIn -.
func (a authRoutes) singIn(c *gin.Context) {
	c.Status(http.StatusOK)
}
