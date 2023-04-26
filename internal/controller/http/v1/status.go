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

// statusRoutes -.
type statusRoutes struct {
	su usecase.Status
	l  logger.Interface
}

// newStatusRoutes -.
func newStatusRoutes(handler *gin.RouterGroup, su usecase.Status, l logger.Interface) {
	r := statusRoutes{su, l}

	h := handler.Group("/status")
	{
		h.GET("", r.getStatus)
		h.GET("/:id", r.getStatusByID)
		h.POST("", r.createStatus)
		h.PUT("/:id", r.updateStatus)
		h.DELETE("/:id", r.deleteStatus)
	}

}

// getStatus -.
func (r statusRoutes) getStatus(c *gin.Context) {
	var req request.GetStatusReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Status")
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
	total, err := r.su.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	status, err := r.su.GetAllStatus(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetStatusRes{
		Status: status,
		Total:  total,
		PaginationRes: utils.PaginationRes{
			Limit:    paginate.Limit,
			Page:     paginate.Page,
			LastPage: utils.GetPageCount(total, paginate.Limit),
		},
	})
}

// getStatusByID -.
func (r statusRoutes) getStatusByID(c *gin.Context) {
	var req request.GetStatusByIDReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Status")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	status, err := r.su.GetStatusByID(req.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetStatusByIDRes{
		Status: status,
	})
}

// createStatus -.
func (r statusRoutes) createStatus(c *gin.Context) {
	var req request.CreateStatusReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Status")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	statusID, err := r.su.CreateStatus(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
		if errors.As(err, &ErrDuplicateRow) {
			response.ErrorResponse(c, http.StatusConflict, _defaultConflict)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResByID(c, http.StatusCreated, statusID)
}

// updateStatus -.
func (r statusRoutes) updateStatus(c *gin.Context) {
	var req request.UpdateStatusReq
	req.ID, _ = strconv.Atoi(c.Param("id"))

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Status")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.su.UpdateStatus(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
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

// deleteStatus -.
func (r statusRoutes) deleteStatus(c *gin.Context) {
	var req request.DeleteStatusReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Status")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.su.DeleteStatus(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Status")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}
