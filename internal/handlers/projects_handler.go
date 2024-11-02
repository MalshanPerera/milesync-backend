package handlers

import (
	"jira-for-peasants/internal/api/requests"
	"jira-for-peasants/internal/api/responses"
	"jira-for-peasants/internal/services"
	errpkg "jira-for-peasants/pkg/errors"
	"jira-for-peasants/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService *services.ProjectService
}

func NewProjectHandler(projectService *services.ProjectService) ProjectHandler {
	return ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) RegisterRoutes(handler *echo.Group) {

	handler.POST("", h.handleCreateProject)
	handler.PATCH("", h.handleUpdateProject)
	handler.DELETE("", h.handleDeleteProject)
	handler.GET("", h.handleGetProjects)
	handler.GET("/:id", h.handleGetProject)
}

func (h *ProjectHandler) handleCreateProject(c echo.Context) error {
	u := new(requests.CreateProjectRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return errpkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	status, e := h.projectService.GetProjectKeyPrefixUsed(ctx, u.KeyPrefix)

	if e != nil {
		return e
	}

	if status {
		return errpkg.BadRequest(errpkg.ProjectKeyPrefixUsed)
	}

	newProject, err := h.projectService.CreateProject(ctx, services.CreateProjectParams{
		UserId:         utils.GetUser(c),
		OrganizationId: utils.GetOrganization(c),
		Name:           u.Name,
		KeyPrefix:      u.KeyPrefix,
		Type:           u.Type,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.ProjectResponse{
		ID:             newProject.ID,
		Name:           newProject.Name,
		KeyPrefix:      newProject.KeyPrefix,
		Type:           newProject.Type,
		OrganizationId: newProject.OrganizationID,
		CreatedBy:      newProject.UserID,
	})
}

func (h *ProjectHandler) handleUpdateProject(c echo.Context) error {
	u := new(requests.UpdateProjectRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return errpkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	status, e := h.projectService.GetProjectKeyPrefixUsed(ctx, u.KeyPrefix)

	if e != nil {
		return e
	}

	if status {
		return errpkg.BadRequest(errpkg.ProjectKeyPrefixUsed)
	}

	updatedProject, err := h.projectService.UpdateProject(ctx, services.UpdateProjectParams{
		UserId:    utils.GetUser(c),
		ID:        u.ID,
		Name:      u.Name,
		KeyPrefix: u.KeyPrefix,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.ProjectResponse{
		ID:             updatedProject.ID,
		Name:           updatedProject.Name,
		KeyPrefix:      updatedProject.KeyPrefix,
		Type:           updatedProject.Type,
		OrganizationId: updatedProject.OrganizationID,
		CreatedBy:      updatedProject.UserID,
	})
}

func (h *ProjectHandler) handleDeleteProject(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()

	err := h.projectService.DeleteProject(ctx, id, utils.GetUser(c))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, id)
}

func (h *ProjectHandler) handleGetProjects(c echo.Context) error {
	ctx := c.Request().Context()

	userId := utils.GetUser(c)
	organizationId := utils.GetOrganization(c)

	allProjects, err := h.projectService.GetProjects(ctx, userId, organizationId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, allProjects)
}

func (h *ProjectHandler) handleGetProject(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, "Get Project by ID: "+id+"")
}
