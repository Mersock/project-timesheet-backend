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
		h.GET("/:id", r.getProjectByID)
		h.POST("", r.createProject)
		h.PUT("/:id", r.updateProject)
		h.DELETE("/:id", r.deleteProject)
	}
}

// getProject -.
func (r projectsRoutes) getProject(c *gin.Context) {
	c.Status(http.StatusOK)
}

// getProjectByID -.
func (r projectsRoutes) getProjectByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

// createProject -.
func (r projectsRoutes) createProject(c *gin.Context) {
	c.Status(http.StatusOK)
}

// updateProject -.
func (r projectsRoutes) updateProject(c *gin.Context) {
	c.Status(http.StatusOK)
}

// deleteProject -.
func (r projectsRoutes) deleteProject(c *gin.Context) {
	c.Status(http.StatusOK)
}
