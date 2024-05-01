package handlers

import (
	"jira-for-peasants/services"
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

	handler.GET("", h.handleGetProjects)
	handler.GET("/:id", h.handleGetProject)
	handler.POST("", h.handleCreateProject)
	handler.PATCH("", h.handleUpdateProject)
	handler.DELETE("", h.handleDeleteProject)
}

func (h *ProjectHandler) handleGetProjects(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Projects")
}

func (h *ProjectHandler) handleGetProject(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, "Get Project by ID: "+id+"")
}

func (h *ProjectHandler) handleCreateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create Project")
}

func (h *ProjectHandler) handleUpdateProject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update Project")
}

func (h *ProjectHandler) handleDeleteProject(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete Project")
}
