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
	"strconv"
)

// timeEntryRoutes -.
type timeEntryRoutes struct {
	tu usecase.TimeEntry
	l  logger.Interface
}

// newTimeEntryRoutes -.
func newTimeEntryRoutes(handler *gin.RouterGroup, tu usecase.TimeEntry, l logger.Interface) {
	r := timeEntryRoutes{tu, l}

	h := handler.Group("/timeEntry")
	{
		h.GET("", r.getTimeEntries)
		h.PUT("/:id", r.updateTimeEntry)
		h.POST("", r.createTimeEntry)
	}

}

// getTimeEntries -.
func (r timeEntryRoutes) getTimeEntries(c *gin.Context) {
	var req request.GetTimeEntryReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Time Entry")
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
	total, err := r.tu.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Time Entry")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	timeEntries, err := r.tu.GetAllTimeEntries(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Time Entry")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetTimeEntriesRes{
		TimeEntries: timeEntries,
		Total:       total,
		PaginationRes: utils.PaginationRes{
			Limit:    paginate.Limit,
			Page:     paginate.Page,
			LastPage: utils.GetPageCount(total, paginate.Limit),
		},
	})
}

// createTimeEntry -.
func (r timeEntryRoutes) createTimeEntry(c *gin.Context) {
	var req request.CreateTimeEntryReq
	userId, _ := strconv.ParseInt(c.Request.Header.Get("x-user-id"), 10, 64)
	req.UserID = userId

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Time Entry")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	timeEntryID, err := r.tu.CreateTimeEntry(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Time Entry")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResByID(c, http.StatusCreated, timeEntryID)
}

// updateTimeEntry -.
func (r timeEntryRoutes) updateTimeEntry(c *gin.Context) {
	var req request.UpdateTimeEntryReq
	req.ID, _ = strconv.Atoi(c.Param("id"))
	userId, _ := strconv.ParseInt(c.Request.Header.Get("x-user-id"), 10, 64)

	req.UserID = userId

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Time Entry")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.tu.UpdateTimeEntry(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Time Entry")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}

		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}
