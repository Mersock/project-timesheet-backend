package v1

import (
	"errors"
	"github.com/Mersock/project-timesheet-backend/internal/entity"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Total int            `json:"total"`
	utils.Pagination
}

// getRoles -.
func (r rolesRoutes) getRoles(c *gin.Context) {
	var req request.RolesReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	//pagination
	paginate := utils.GeneratePaginationFromRequest(c)

	//total rows
	total, err := r.ru.GetRowsRoles(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		errorResponse(c, http.StatusInternalServerError, "Database error")
		return
	}

	roles, err := r.ru.GetAllRoles(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		errorResponse(c, http.StatusInternalServerError, "Database error")
		return
	}

	c.JSON(http.StatusOK, rolesResponse{
		Roles: roles,
		Total: total,
		Pagination: utils.Pagination{
			Limit: paginate.Limit,
			Page:  paginate.Page,
		},
	})
}
