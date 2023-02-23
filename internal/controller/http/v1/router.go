package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/Mersock/project-timesheet-backend/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(
	handler *gin.Engine,
	l logger.Interface,
	tokenMaker token.Maker,
	ru usecase.Roles,
	uu usecase.User,
	au usecase.Auth,
	pu usecase.Project,
	wu usecase.WorkTypes,
	su usecase.Status,
) {
	//options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//healthcheck
	handler.GET("/api/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	auth := handler.Group("/api/v1/auth")
	newAuthRoutes(auth, au, l)

	//routers require auth
	h := handler.Group("/api/v1")
	//h.Use(middleware.AuthMiddleware(tokenMaker))
	{
		newRolesRoutes(h, ru, l)
		newUsersRoutes(h, uu, l)
		newProjectsRoutes(h, pu, l)
		newWorkTypesRoutes(h, wu, l)
		newStatusRoutes(h, su, l)
	}
}
