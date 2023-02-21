package v1

import (
	"errors"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// authRoutes -.
type authRoutes struct {
	au usecase.Auth
	l  logger.Interface
}

// newAuthRoutes -.
func newAuthRoutes(h *gin.RouterGroup, au usecase.Auth, l logger.Interface) {
	a := authRoutes{au, l}

	{
		h.POST("/signup", a.singUp)
		h.POST("/signin", a.singIn)
		h.POST("/renewAccess", a.renewAccess)
	}
}

// singUp -.
func (a authRoutes) singUp(c *gin.Context) {
	var req request.SignUpReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		a.l.Error(err, "http - v1 - Auth")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	userID, err := a.au.Signup(req)
	if err != nil {
		a.l.Error(err, "http - v1 - Roles")
		if errors.As(err, &ErrDuplicateRow) {
			errorResponse(c, http.StatusConflict, _defaultConflict)
			return
		}
		errorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusCreated, response.SignUpRes{
		ID: userID,
	})
}

// singIn -.
func (a authRoutes) singIn(c *gin.Context) {
	var req request.SignInReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		a.l.Error(err, "http - v1 - Auth")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	session, err := a.au.SignIn(req)
	if err != nil {
		a.l.Error(err, "http - v1 - Auth")
		errorResponse(c, http.StatusUnauthorized, _defaultUnauthorized)
		return
	}

	c.JSON(http.StatusOK, session)
}

// renewAccess -.
func (a authRoutes) renewAccess(c *gin.Context) {
	var req request.RenewTokenReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		a.l.Error(err, "http - v1 - Auth")
		if errors.As(err, &ve) {
			errorValidateRes(c, ve)
			return
		}
		errorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	session, err := a.au.RenewAccess(req)
	if err != nil {
		a.l.Error(err, "http - v1 - Auth")
		errorResponse(c, http.StatusUnauthorized, _defaultUnauthorized)
		return
	}

	c.JSON(http.StatusOK, session)
}
