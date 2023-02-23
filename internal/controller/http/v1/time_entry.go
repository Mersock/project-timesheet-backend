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
		h.POST("", r.createTimeEntry)
	}

}

// createRole -.
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
