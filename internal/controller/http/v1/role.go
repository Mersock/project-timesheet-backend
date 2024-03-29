package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// rolesRoutes -.
type rolesRoutes struct {
	ru usecase.Roles
	l  logger.Interface
}

// newRolesRoutes -.
func newRolesRoutes(handler *gin.RouterGroup, ru usecase.Roles, l logger.Interface) {
	r := rolesRoutes{ru, l}

	h := handler.Group("/role")
	{
		h.GET("", r.getRoles)
		h.GET("/:id", r.getRoleByID)
		h.POST("", r.createRole)
		h.PUT("/:id", r.updateRole)
		h.DELETE("/:id", r.deleteRole)
	}

}

// getRoles -.
func (r rolesRoutes) getRoles(c *gin.Context) {
	var req request.GetRolesReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	//pagination
	paginate := utils.GeneratePaginationFromRequest(c)
	req.Limit = &paginate.Limit
	req.Page = &paginate.Page

	//total rows
	total, err := r.ru.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	roles, err := r.ru.GetAllRoles(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetRolesRes{
		Roles: roles,
		Total: total,
		PaginationRes: utils.PaginationRes{
			Limit:    paginate.Limit,
			Page:     paginate.Page,
			LastPage: utils.GetPageCount(total, paginate.Limit),
		},
	})
}

// getRoleByID -.
func (r rolesRoutes) getRoleByID(c *gin.Context) {
	var req request.GetRoleByIDReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	role, err := r.ru.GetRoleByID(req.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetRoleByIDRes{
		Role: role,
	})
}

// createRole -.
func (r rolesRoutes) createRole(c *gin.Context) {
	var req request.CreateRoleReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	roleID, err := r.ru.CreateRole(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ErrDuplicateRow) {
			response.ErrorResponse(c, http.StatusConflict, _defaultConflict)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResByID(c, http.StatusCreated, roleID)
}

// updateRole -.
func (r rolesRoutes) updateRole(c *gin.Context) {
	var req request.UpdateRoleReq
	req.ID, _ = strconv.Atoi(c.Param("id"))

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.ru.UpdateRole(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}

		if errors.As(err, &ErrDuplicateRow) {
			response.ErrorResponse(c, http.StatusConflict, _defaultConflict)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}

// deleteRole -.
func (r rolesRoutes) deleteRole(c *gin.Context) {
	var req request.DeleteRoleReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.ru.DeleteRole(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}
