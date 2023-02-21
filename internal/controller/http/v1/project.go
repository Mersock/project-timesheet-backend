package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// projectsRoutes -.
type projectsRoutes struct {
	pu usecase.Project
	l  logger.Interface
}

// newProjectsRoutes -.
func newProjectsRoutes(handler *gin.RouterGroup, pu usecase.Project, l logger.Interface) {
	r := projectsRoutes{pu, l}

	h := handler.Group("/project")
	{
		h.GET("", r.getProject)
	}
}

func (r projectsRoutes) getProject(c *gin.Context) {
	c.Status(http.StatusOK)
}
