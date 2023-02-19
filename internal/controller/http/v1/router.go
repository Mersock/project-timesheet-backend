package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(handler *gin.Engine, l logger.Interface, ru usecase.Roles) {
	//options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//healthcheck
	handler.GET("/api/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	//routers
	h := handler.Group("/api/v1")
	{
		newRolesRoutes(h, ru, l)
	}
}
