package v1

import (
	"database/sql"
	"errors"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// workTypesRoutes -.
type workTypesRoutes struct {
	wu usecase.WorkTypes
	l  logger.Interface
}

// newWorkTypesRoutes -.
func newWorkTypesRoutes(handler *gin.RouterGroup, wu usecase.WorkTypes, l logger.Interface) {
	u := workTypesRoutes{wu, l}

	h := handler.Group("/workTypes")
	{
		h.GET("/:id", u.getWorkTypeByID)
		h.PUT("/:id", u.updateWorkType)
	}
}

// getUserByID -.
func (r workTypesRoutes) getWorkTypeByID(c *gin.Context) {
	var req request.GetWorkTypeByIDReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - WorkTypes")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	worktype, err := r.wu.GetWorkTypeByID(req.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - WorkTypes")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetWorkTypeByIDRes{
		WorkType: worktype,
	})
}

// updateWorkType -.
func (r workTypesRoutes) updateWorkType(c *gin.Context) {
	var req request.UpdateWorkTypeReq
	req.ID, _ = strconv.Atoi(c.Param("id"))

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - WorkTypes")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.wu.UpdateWorkType(req)
	if err != nil {
		r.l.Error(err, "http - v1 - WorkTypes")
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
