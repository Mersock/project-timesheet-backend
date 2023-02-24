package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reportRoutes -.
type reportRoutes struct {
	ru usecase.Report
	l  logger.Interface
}

// newReportRoutes -.
func newReportRoutes(handler *gin.RouterGroup, ru usecase.Roles, l logger.Interface) {
	r := reportRoutes{ru, l}

	h := handler.Group("/report", r.getReport)
	{
		h.GET("")
	}
}

// getReport -.
func (r reportRoutes) getReport(c *gin.Context) {
	c.Status(http.StatusOK)
}
