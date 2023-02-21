package v1

import (
	"database/sql"
	"errors"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// usersRoutes -.
type usersRoutes struct {
	uu usecase.User
	l  logger.Interface
}

// newUsersRoutes -.
func newUsersRoutes(handler *gin.RouterGroup, uu usecase.User, l logger.Interface) {
	u := usersRoutes{uu, l}

	h := handler.Group("/user")
	{
		h.GET("", u.getAllUsers)
		h.GET("/:id", u.getUser)
		h.DELETE("/:id", u.deleteRole)
	}
}

// getUsers -.
func (r usersRoutes) getAllUsers(c *gin.Context) {
	var req request.GetAllUsersReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	//pagination
	paginate := utils.GeneratePaginationFromRequest(c)
	req.Limit = &paginate.Limit
	req.Page = &paginate.Page

	//sortBy
	if strings.ToLower(req.SortBy) == "id" {
		req.SortBy = "users.id"
	}
	if strings.ToLower(req.SortBy) == "role" {
		req.SortBy = "roles.name"
	}

	//total rows
	total, err := r.uu.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		errorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	users, err := r.uu.GetAllUsers(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		errorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetUsersRes{
		Users: users,
		Total: total,
		PaginationRes: utils.PaginationRes{
			Limit:    paginate.Limit,
			Page:     paginate.Page,
			LastPage: utils.GetPageCount(total, paginate.Limit),
		},
	})
}

// getUser -.
func (r usersRoutes) getUser(c *gin.Context) {
	var req request.GetUserReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	user, err := r.uu.GetUser(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		if errors.Is(err, sql.ErrNoRows) {
			errorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		errorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetUserByIDRes{
		User: user,
	})
}

// deleteRole -.
func (r usersRoutes) deleteRole(c *gin.Context) {
	var req request.DeleteUserReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.uu.DeleteUser(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		if errors.Is(err, sql.ErrNoRows) {
			errorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		errorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.UpdateRoleRes{
		RowAffected: rowAffected,
	})
}
