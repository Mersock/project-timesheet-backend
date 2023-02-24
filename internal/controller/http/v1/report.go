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
)

// reportRoutes -.
type reportRoutes struct {
	ru usecase.Report
	l  logger.Interface
}

// newReportRoutes -.
func newReportRoutes(handler *gin.RouterGroup, ru usecase.Report, l logger.Interface) {
	r := reportRoutes{ru, l}

	h := handler.Group("/report")
	{
		h.POST("/workType", r.getWorkTypeReport)
	}
}

// getReport -.
func (r reportRoutes) getWorkTypeReport(c *gin.Context) {
	var req request.GetWorkTypeReportReq

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

	_, err := r.ru.GetWorkTypeCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Report")
		r.l.Error(err, "http - v1 - Roles")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	workTypeReport, err := r.ru.GetAllWorkType(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Report")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetReportWorkTypeRes{
		WorkTypeReport: workTypeReport,
	})
}
