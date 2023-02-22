package v1

import (
	"errors"
	"github.com/Mersock/project-timesheet-backend/internal/request"
	"github.com/Mersock/project-timesheet-backend/internal/response"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// projectsRoutes -.
type projectsRoutes struct {
	pu usecase.Project
	l  logger.Interface
}

// newProjectsRoutes -.
func newProjectsRoutes(handler *gin.RouterGroup, pu usecase.Project, l logger.Interface) {
	r := projectsRoutes{pu, l}

	h := handler.Group("/project")
	{
		h.GET("", r.getProject)
		h.GET("/:id", r.getProjectByID)
		h.POST("", r.createProject)
		h.PUT("/:id", r.updateProject)
		h.DELETE("/:id", r.deleteProject)
	}
}

// getProject -.
func (r projectsRoutes) getProject(c *gin.Context) {
	c.Status(http.StatusOK)
}

// getProjectByID -.
func (r projectsRoutes) getProjectByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

// createProject -.
func (r projectsRoutes) createProject(c *gin.Context) {
	var req request.CreateProjectReq

	code, _ := uuid.NewRandom()
	userId, _ := strconv.ParseInt(c.Request.Header.Get("x-user-id"), 10, 64)
	req.Code = code
	req.UserOwnerID = userId

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Project")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	projectID, err := r.pu.CreateProject(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Project")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResByID(c, http.StatusCreated, projectID)
}

// updateProject -.
func (r projectsRoutes) updateProject(c *gin.Context) {
	c.Status(http.StatusOK)
}

// deleteProject -.
func (r projectsRoutes) deleteProject(c *gin.Context) {
	c.Status(http.StatusOK)
}
