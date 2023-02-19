package v1

import (
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// rolesRoutes -.
type rolesRoutes struct {
	ru usecase.Roles
	l  logger.Interface
}

// newRolesRoutes -.
func newRolesRoutes(handler *gin.RouterGroup, ru usecase.Roles, l logger.Interface) {
	r := rolesRoutes{ru, l}

	h := handler.Group("/roles")
	{
		h.GET("", r.getRoles)
	}

}

// rolesResponse -.
type rolesResponse struct {
	Roles []entity.Roles `json:"data"`
}

// getRoles -.
func (r rolesRoutes) getRoles(c *gin.Context) {
	roles, err := r.ru.GetAllRoles()
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		errorResponse(c, http.StatusInternalServerError, "database error")
		return
	}
	c.JSON(http.StatusOK, rolesResponse{roles})
}
