package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		h.GET("/count", u.getCountUser)
		h.GET("", u.getUsers)
		h.POST("", u.createUser)
		h.GET("/:id", u.getUserByID)
		h.PUT("/:id", u.updateUser)
		h.PUT("/password/:id", u.updateUserPassword)
		h.DELETE("/:id", u.deleteRole)
	}
}

// getUsers -.
func (r usersRoutes) getCountUser(c *gin.Context) {
	var req request.GetUsersReq

	//total rows
	total, err := r.uu.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}
	fmt.Println(total)
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// getUsers -.
func (r usersRoutes) getUsers(c *gin.Context) {
	var req request.GetUsersReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
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
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	users, err := r.uu.GetAllUsers(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
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

// getUserByID -.
func (r usersRoutes) getUserByID(c *gin.Context) {
	var req request.GetUserByIDReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	user, err := r.uu.GetUserByID(req.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetUserByIDRes{
		User: user,
	})
}

// createUser -.
func (r usersRoutes) createUser(c *gin.Context) {
	var req request.CreateUserReq

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

	userID, err := r.uu.CreateUser(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ErrDuplicateRow) {
			response.ErrorResponse(c, http.StatusConflict, _defaultConflict)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResByID(c, http.StatusCreated, userID)
}

// updateUser -.
func (r usersRoutes) updateUser(c *gin.Context) {
	var req request.UpdateUserReq
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

	rowAffected, err := r.uu.UpdateUser(req)
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

// updateUserPassword -.
func (r usersRoutes) updateUserPassword(c *gin.Context) {
	var req request.UpdateUserPasswordReq
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

	rowAffected, err := r.uu.UpdateUserPassword(req)
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

// deleteRole -.
func (r usersRoutes) deleteRole(c *gin.Context) {
	var req request.DeleteUserReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Users")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.uu.DeleteUser(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Users")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}
