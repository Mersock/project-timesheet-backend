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
	var req request.GetProjectsReq

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Projects")
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
	total, err := r.pu.GetCount(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Projects")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	projects, err := r.pu.GetAllProjects(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Projects")
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetProjectsRes{
		Projects: projects,
		Total:    total,
		PaginationRes: utils.PaginationRes{
			Limit:    paginate.Limit,
			Page:     paginate.Page,
			LastPage: utils.GetPageCount(total, paginate.Limit),
		},
	})
}

// getProjectByID -.
func (r projectsRoutes) getProjectByID(c *gin.Context) {
	var req request.GetProjectByIDReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Projects")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	project, err := r.pu.GetProjectsByIDWithUserWorkType(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Projects")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	c.JSON(http.StatusOK, response.GetProjectByIDWithUser{
		Data: project,
	})
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
	var req request.UpdateProjectReq
	req.ID, _ = strconv.Atoi(c.Param("id"))

	//validator
	if err := c.ShouldBind(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Projects")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.pu.UpdateProject(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Projects")
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

// deleteProject -.
func (r projectsRoutes) deleteProject(c *gin.Context) {
	var req request.DeleteProjectByReq

	//validator
	if err := c.ShouldBindUri(&req); err != nil {
		var ve validator.ValidationErrors
		r.l.Error(err, "http - v1 - Projects")
		if errors.As(err, &ve) {
			response.ErrorValidateRes(c, ve)
			return
		}
		response.ErrorResponse(c, http.StatusBadRequest, _defaultBadReq)
		return
	}

	rowAffected, err := r.pu.DeleteProject(req)
	if err != nil {
		r.l.Error(err, "http - v1 - Projects")
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, _defaultNotFound)
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, _defaultInternalServerErr)
		return
	}

	response.ResRowAffect(c, http.StatusOK, rowAffected)
}
