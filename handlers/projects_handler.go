package handlers

import (
	err_pkg "jira-for-peasants/errors"
	"jira-for-peasants/requests"
	"jira-for-peasants/responses"
	"jira-for-peasants/services"
	"jira-for-peasants/utils"
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
		return err_pkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err_pkg.UnprocessableEntity(err.Error())
	}

	status, e := h.projectService.GetProjectKeyPrefixUsed(ctx, u.KeyPrefix)

	if e != nil {
		return e
	}

	if status {
		return err_pkg.BadRequest(err_pkg.ProjectKeyPrefixUsed)
	}

	newProject, err := h.projectService.CreateProject(ctx, services.CreateProjectParams{
		UserId:         utils.GetUser(c),
		OrganizationId: u.OrganizationId,
		Name:           u.Name,
		KeyPrefix:      u.KeyPrefix,
		Type:           u.Type,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.NewProjectResponse(
		newProject.ID,
		newProject.Name,
		newProject.KeyPrefix,
		newProject.Type,
		newProject.OrganizationID,
		newProject.UserID,
	))
}

func (h *ProjectHandler) handleUpdateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update Project")
}

func (h *ProjectHandler) handleDeleteProject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete Project")
}

func (h *ProjectHandler) handleGetProjects(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Projects")
}

func (h *ProjectHandler) handleGetProject(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, "Get Project by ID: "+id+"")
}
